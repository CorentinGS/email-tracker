package config

const (
	GCThresholdPercent = 0.7 // GCThresholdPercent is the threshold for garbage collection

	GCLimit = 1024 * 1024 * 1024 // GCLimit is the limit for garbage collection

	NonceLength = 16 // NonceLength is the length of the nonce : 16 bytes * 8 bits/byte = 128 bits
)
