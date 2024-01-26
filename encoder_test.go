package encoder

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	// Create test struct
	ie := Nas5GSUpdateType{
		IEI:           1,
		Length:        2,
		SMS_Requested: 1,
		NG_RAN_RCU:    1,
		PNB_CIoT_5GS:  0,
		PNB_CIoT_EPS:  0,
	}

	// Create buffer and fill it with the expected result
	expected_buffer := bytes.Buffer{}
	expected_buffer.Write([]byte{1, 2, 3})
	// Create buffer to hold the actual result
	actual_buffer := bytes.Buffer{}

	// Run the Encode function on the test struct
	ie.Encode(&actual_buffer)

	// Compare the contents (.Bytes()) of expected_buffer and actual_buffer
	if !bytes.Equal(expected_buffer.Bytes(), actual_buffer.Bytes()) {
		t.Errorf("expected -> %#x\n  actual -> %#x\n", expected_buffer.Bytes(), actual_buffer.Bytes())
	}
}
