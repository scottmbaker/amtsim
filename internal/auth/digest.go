package auth

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type DigestMiddleware struct {
	Realm     string
	UserPass  map[string]string
	Nonces    map[string]time.Time
	NonceLock sync.Mutex
}

func NewDigestMiddleware(realm string, userPass map[string]string) *DigestMiddleware {
	return &DigestMiddleware{
		Realm:    realm,
		UserPass: userPass,
		Nonces:   make(map[string]time.Time),
	}
}

func (m *DigestMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			m.requireAuth(w)
			return
		}

		if !strings.HasPrefix(authHeader, "Digest ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authMap := parseAuthHeader(authHeader[7:])

		// Validate Nonce
		if !m.validateNonce(authMap["nonce"]) {
			m.requireAuth(w)
			return
		}

		// Calculate Expected Response
		ha1 := md5Sum(authMap["username"] + ":" + m.Realm + ":" + m.UserPass[authMap["username"]])
		ha2 := md5Sum(r.Method + ":" + authMap["uri"])
		expectedr := md5Sum(ha1 + ":" + authMap["nonce"] + ":" + authMap["nc"] + ":" + authMap["cnonce"] + ":" + authMap["qop"] + ":" + ha2)

		if expectedr != authMap["response"] {
			m.requireAuth(w)
			return
		}

		next(w, r)
	}
}

func (m *DigestMiddleware) requireAuth(w http.ResponseWriter) {
	nonce := m.generateNonce()
	w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Digest realm="%s", qop="auth", nonce="%s", opaque=""`, m.Realm, nonce))
	w.WriteHeader(http.StatusUnauthorized)
}

func (m *DigestMiddleware) generateNonce() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Printf("Failed to generate nonce: %v", err)
		return "nonce-gen-failed" // Should not happen usually, avoiding panic
	}
	nonce := hex.EncodeToString(b)

	m.NonceLock.Lock()
	defer m.NonceLock.Unlock()
	m.Nonces[nonce] = time.Now()

	return nonce
}

func (m *DigestMiddleware) validateNonce(nonce string) bool {
	m.NonceLock.Lock()
	defer m.NonceLock.Unlock()

	t, ok := m.Nonces[nonce]
	if !ok {
		return false
	}

	// Expire after 5 minutes
	if time.Since(t) > 5*time.Minute {
		delete(m.Nonces, nonce)
		return false
	}

	return true
}

func md5Sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func parseAuthHeader(header string) map[string]string {
	result := make(map[string]string)
	parts := strings.Split(header, ",")
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			key := kv[0]
			val := strings.Trim(kv[1], `"`)
			result[key] = val
		}
	}
	return result
}
