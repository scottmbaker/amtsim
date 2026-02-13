package amt

import (
	"amtsim/internal/amt/consts"
	"encoding/json"
	"log"
	"os"
	"sync"
)

const StateFile = "device_state.json"

type DeviceState struct {
	PowerState         int    `json:"powerState"`
	ProvisioningMode   int    `json:"provisioningMode"`
	ProvisioningState  int    `json:"provisioningState"`
	RemoteAccessStatus int    `json:"remoteAccessStatus"` // 0=Not Connected, 1=Connecting, 2=Connected
	UUID               string `json:"uuid"`
}

var (
	currentState DeviceState
	stateMu      sync.RWMutex
	once         sync.Once
)

func ensureInitialized() {
	once.Do(func() {
		// Set defaults
		currentState = DeviceState{
			PowerState:        2, // On
			ProvisioningMode:  consts.ProvisioningModePre,
			ProvisioningState: consts.ProvisioningStatePre,
		}
		loadState()
	})
}

func loadState() {
	data, err := os.ReadFile(StateFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to read state file: %v", err)
		}
		return
	}
	if err := json.Unmarshal(data, &currentState); err != nil {
		log.Printf("Failed to unmarshal state: %v", err)
	}
}

func saveState() {
	data, err := json.MarshalIndent(currentState, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal state: %v", err)
		return
	}
	if err := os.WriteFile(StateFile, data, 0644); err != nil {
		log.Printf("Failed to write state file: %v", err)
	}
}

// GetPowerState returns the current power state safely
func GetPowerState() int {
	ensureInitialized()
	stateMu.RLock()
	defer stateMu.RUnlock()
	return currentState.PowerState
}

// SetPowerState updates the power state safely
func SetPowerState(state int) {
	ensureInitialized()
	stateMu.Lock()
	defer stateMu.Unlock()
	currentState.PowerState = state
	saveState()
}

// GetProvisioningMode returns the current provisioning mode
func GetProvisioningMode() int {
	ensureInitialized()
	stateMu.RLock()
	defer stateMu.RUnlock()
	return currentState.ProvisioningMode
}

// SetProvisioningMode updates the provisioning mode
func SetProvisioningMode(mode int) {
	ensureInitialized()
	stateMu.Lock()
	defer stateMu.Unlock()
	currentState.ProvisioningMode = mode
	saveState()
}

// GetProvisioningState returns the current provisioning state
func GetProvisioningState() int {
	ensureInitialized()
	stateMu.RLock()
	defer stateMu.RUnlock()
	return currentState.ProvisioningState
}

// SetProvisioningState updates the provisioning state
func SetProvisioningState(state int) {
	ensureInitialized()
	stateMu.Lock()
	currentState.ProvisioningState = state
	saveState()
	stateMu.Unlock()
}

// GetRemoteAccessStatus returns the current remote access status
func GetRemoteAccessStatus() int {
	ensureInitialized()
	stateMu.RLock()
	defer stateMu.RUnlock()
	return currentState.RemoteAccessStatus
}

// SetRemoteAccessStatus updates the remote access status
func SetRemoteAccessStatus(status int) {
	ensureInitialized()
	stateMu.Lock()
	defer stateMu.Unlock()
	currentState.RemoteAccessStatus = status
	saveState()
}

// GetUUID returns the current UUID
func GetUUID() string {
	ensureInitialized()
	stateMu.RLock()
	defer stateMu.RUnlock()
	return currentState.UUID
}

// SetUUID updates the UUID. Does not save to disk to avoid overwrite on restart unless needed?
// Actually we should save it if it's persistent, but main.go sets it from flag every run.
// Let's save it to be safe.
func SetUUID(uuid string) {
	ensureInitialized()
	stateMu.Lock()
	defer stateMu.Unlock()
	currentState.UUID = uuid
	saveState()
}
