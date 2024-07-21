package uuidconv

import (
	"github.com/google/uuid"
)

// UUID2DoubleInt converts a UUID to two 64-bit integers.
//
// The function takes a UUID as input and returns two integers, high and low.
// The high integer is constructed by concatenating the first 8 bytes of the UUID,
// and the low integer is constructed by concatenating the last 8 bytes of the UUID.
//
// Parameters:
// - v: the UUID to be converted.
//
// Returns:
// - high: the high part of the UUID as a 64-bit integer.
// - low: the low part of the UUID as a 64-bit integer.
//
//nolint:mnd,varnamelen
func UUID2DoubleInt(v uuid.UUID) (int64, int64) {
	// Extract the bytes of the UUID and concatenate them to form the high and low integers.
	high := int64(v[0]) | int64(v[1])<<8 | int64(v[2])<<16 | int64(v[3])<<24 |
		int64(v[4])<<32 | int64(v[5])<<40 | int64(v[6])<<48 | int64(v[7])<<56

	low := int64(v[8]) | int64(v[9])<<8 | int64(v[10])<<16 | int64(v[11])<<24 |
		int64(v[12])<<32 | int64(v[13])<<40 | int64(v[14])<<48 | int64(v[15])<<56

	return high, low
}

// DoubleInt2UUID converts two 64-bit integers to a UUID.
//
// The function takes two integers, high and low, as input and returns a UUID.
// The UUID is constructed by concatenating the bytes of the high and low integers.
// The first 8 bytes of the UUID are constructed from the high integer,
// and the last 8 bytes are constructed from the low integer.
//
// Parameters:
// - high: the high part of the UUID as a 64-bit integer.
// - low: the low part of the UUID as a 64-bit integer.
//
// Returns:
// - uuid: the UUID constructed from the high and low integers.
//
// Example:
// The function converts the integers 6211610197754262546 and 6086219141769098904 to a UUID.
// The UUID is constructed by concatenating the bytes of the high and low integers.
// The first 8 bytes of the UUID are constructed from the high integer,
// and the last 8 bytes are constructed from the low integer.
// The resulting UUID is 12345678-9012-3456-9876-543210987654.
//
//nolint:mnd
func DoubleInt2UUID(high int64, low int64) uuid.UUID {
	return uuid.UUID{
		// Extract the bytes of the high integer and concatenate them to form the first 8 bytes of the UUID.
		byte(high),
		byte(high >> 8),
		byte(high >> 16),
		byte(high >> 24),
		byte(high >> 32),
		byte(high >> 40),
		byte(high >> 48),
		byte(high >> 56),
		// Extract the bytes of the low integer and concatenate them to form the last 8 bytes of the UUID.
		byte(low),
		byte(low >> 8),
		byte(low >> 16),
		byte(low >> 24),
		byte(low >> 32),
		byte(low >> 40),
		byte(low >> 48),
		byte(low >> 56),
	}
}
