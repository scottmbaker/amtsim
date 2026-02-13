package consts

// Provisioning Modes
const (
	ProvisioningModePre = 0
	ProvisioningModeACM = 1 // Admin Control Mode
	ProvisioningModeCCM = 4 // Client Control Mode
)

// Provisioning States
const (
	ProvisioningStatePre  = 0 // Pre-provisioning
	ProvisioningStateIn   = 1 // In-provisioning
	ProvisioningStatePost = 2 // Post-provisioning
)

// Remote Access Status
const (
	RemoteAccessStatusNotConnected = 0
	RemoteAccessStatusConnecting   = 1
	RemoteAccessStatusConnected    = 2
)

// Power States
const (
	PowerStateOn = 2
	// Add others if needed
)
