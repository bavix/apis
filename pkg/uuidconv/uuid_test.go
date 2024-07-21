package uuidconv_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/bavix/apis/pkg/uuidconv"
)

// UUIDTestSuite is a test suite for testing the UUID conversion functions.
//
// It is a struct that embeds the Suite struct from the testify package.
type UUIDTestSuite struct {
	suite.Suite
}

// TestUuidTestSuite runs the UUIDTestSuite tests.
//
// The function calls the Run method of the Suite struct from the testify package
// to run the tests for the UUIDTestSuite.
func TestUuidTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(UUIDTestSuite))
}

// TestUUID2DoubleInt tests the conversion of a UUID to two 64-bit integers.
//
// The function tests the UUID2DoubleInt function from the uuidconv package.
// It defines a struct for expected values and a list of test cases. Each test
// case consists of a name, expected values, and a UID (UUID) to be tested. The function then
// iterates over the test cases and uses the Run method of the Suite struct
// from the testify package to run a subtest for each case. Inside the subtest,
// the function calls the UUID2DoubleInt function and checks if the returned
// values are equal to the expected values.
//
//nolint:funlen
func (s *UUIDTestSuite) TestUUID2DoubleInt() {
	// Define struct for expected values
	//
	// The struct contains the expected values for the high and low parts of
	// the UID.
	type expectedValues struct {
		high int64
		low  int64
	}

	// Define the test cases
	//
	// The list of test cases contains the name of the test case, the expected
	// values, and the UID to be tested.
	tests := []struct {
		name     string
		expected expectedValues
		uid      uuid.UUID
	}{
		{
			name:     "Zero UID",
			expected: expectedValues{high: 0, low: 0},
			uid:      uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		},
		{
			name:     "UID with low part",
			expected: expectedValues{high: 0, low: 72057594037927936},
			uid:      uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			name:     "UID with high part",
			expected: expectedValues{high: 72057594037927936, low: 0},
			uid:      uuid.MustParse("00000000-0000-0001-0000-000000000000"),
		},
		{
			name:     "UID with negative low part",
			expected: expectedValues{high: 0, low: -1},
			uid:      uuid.MustParse("00000000-0000-0000-FFFF-FFFFFFFFFFFF"),
		},
		{
			name:     "UID with negative high part",
			expected: expectedValues{high: -1, low: 0},
			uid:      uuid.MustParse("FFFFFFFF-FFFF-FFFF-0000-000000000000"),
		},
		{
			name:     "UID with negative parts",
			expected: expectedValues{high: -1, low: -1},
			uid:      uuid.MustParse("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF"),
		},
		{
			name:     "UID with negative high part and low part",
			expected: expectedValues{high: -858980353, low: 0},
			uid:      uuid.MustParse("FFFFCCCC-FFFF-FFFF-0000-000000000000"),
		},
		{
			name:     "Complex UID",
			expected: expectedValues{high: 3623603779236289292, low: -5403687274121261928},
			uid:      uuid.MustParse("0c5b2444-70a0-4932-980c-b4dc0d3f02b5"),
		},
	}

	// Test each case
	//
	// The function iterates over the list of test cases, runs a subtest for
	// each case, and checks if the returned values are equal to the expected
	// values.
	for _, test := range tests {
		s.Run(test.name, func() {
			high, low := uuidconv.UUID2DoubleInt(test.uid)
			s.Equal(test.expected.high, high, "Unexpected high part")
			s.Equal(test.expected.low, low, "Unexpected low part")
		})
	}
}

// TestDoubleInt2UUID tests the conversion of two 64-bit integers to a UUID.
//
// The function tests the DoubleInt2UUID function from the uuidconv package.
// It defines a struct for test cases and a list of test cases. Each test
// case consists of a name, two 64-bit integers, and a UUID. The function then
// iterates over the test cases and uses the Run method of the Suite struct
// from the testify package to run a subtest for each case. Inside the subtest,
// the function calls the DoubleInt2UUID function and checks if the returned
// UUID is equal to the expected UUID.
//
//nolint:funlen
func (s *UUIDTestSuite) TestDoubleInt2UUID() {
	tests := []struct {
		name string
		high int64
		low  int64
		guid uuid.UUID
	}{
		{
			name: "Zero UUID",
			high: 0,
			low:  0,
			guid: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		},
		{
			name: "UUID with low part",
			high: 0,
			low:  72057594037927936,
			guid: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			name: "UUID with high part",
			high: 72057594037927936,
			low:  0,
			guid: uuid.MustParse("00000000-0000-0001-0000-000000000000"),
		},
		{
			name: "UUID with negative low part",
			high: 0,
			low:  -1,
			guid: uuid.MustParse("00000000-0000-0000-FFFF-FFFFFFFFFFFF"),
		},
		{
			name: "UUID with negative high part",
			high: -1,
			low:  0,
			guid: uuid.MustParse("FFFFFFFF-FFFF-FFFF-0000-000000000000"),
		},
		{
			name: "UUID with negative parts",
			high: -1,
			low:  -1,
			guid: uuid.MustParse("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF"),
		},
		{
			name: "UUID with negative high part and low part",
			high: -858980353,
			low:  0,
			guid: uuid.MustParse("FFFFCCCC-FFFF-FFFF-0000-000000000000"),
		},
		{
			name: "Complex UUID",
			high: 3623603779236289292,
			low:  -5403687274121261928,
			guid: uuid.MustParse("0c5b2444-70a0-4932-980c-b4dc0d3f02b5"),
		},
		{
			name: "UUID with negative high part and positive low part",
			high: -858980353,
			low:  72057594037927936,
			guid: uuid.MustParse("FFFFCCCC-FFFF-FFFF-0000-000000000001"),
		},
		{
			name: "UUID with positive high part and negative low part",
			high: 72057594037927936,
			low:  -1,
			guid: uuid.MustParse("00000000-0000-0001-FFFF-FFFFFFFFFFFF"),
		},
		{
			name: "UUID with positive high part and positive low part",
			high: 72057594037927936,
			low:  72057594037927936,
			guid: uuid.MustParse("00000000-0000-0001-0000-000000000001"),
		},
		{
			name: "UUID with negative high part and negative low part",
			high: -858980353,
			low:  -858980353,
			guid: uuid.MustParse("ffffcccc-ffff-ffff-ffff-ccccffffffff"),
		},
	}

	// Test each case
	//
	// The function iterates over the list of test cases, runs a subtest for
	// each case, and checks if the returned UUID is equal to the expected UUID.
	for _, test := range tests {
		s.Run(test.name, func() {
			guid := uuidconv.DoubleInt2UUID(test.high, test.low)
			s.Equal(test.guid, guid, "Unexpected UUID")
		})
	}
}
