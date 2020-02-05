package config

// RPCConfig defines RPC endpoint and mode ["raw" and "web"]
type RPCConfig struct {
	Endpoint string
	Mode     string
	CertPath string
	KeyPath  string
}
