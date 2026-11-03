package config

type Deployment struct {
	StableDecimals       uint8    `json:"stableDecimals"`
	ChainID              uint64   `json:"chainId"`
	Pool                 string   `json:"pool"`
	Oracle               string   `json:"oracle"`
	Multisig             string   `json:"multisig"`
	Stable               string   `json:"stable"`
	WETH                 string   `json:"weth"`
	Router               string   `json:"router"`
	Factory              string   `json:"factory"`
	SP                   string   `json:"sp"`
	JP                   string   `json:"jp"`
	Operator             string   `json:"operator"`
	Admins               []string `json:"admins"`
	Threshold            uint64   `json:"threshold"`
	DeploymentStartBlock uint64   `json:"deploymentStartBlock"`
}
