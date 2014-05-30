package util

// util.RamUsageEstimator.java

// amd64 system
const (
	NUM_BYTES_CHAR = 2 // UTF8 uses 1-4 bytes to represent each rune
	NUM_BYTES_INT  = 8

	/* Number of bytes to represent an object reference */
	NUM_BYTES_OBJECT_REF = 8
)
