package encoder

import (
	"bytes"
	"log"

	"github.com/HewlettPackard/structex"
)

// Struct representing the 5GS update type IE in 3GPP TS 24501 (Version 16.9.0) chapter 9.11.3.9A
// The order of fields is import so that structex.EncodeByteBuffer() puts the correct bytes at the correct places
type Nas5GSUpdateType struct {
	// Identifier of the IE (Information Element)
	IEI uint8 // Byte 0
	// Length of the IE
	Length uint8 // Byte 1
	// Signifies whether SMS over NAS is supported or not
	SMS_Requested uint8 `bitfield:"1"` // Byte 2
	// Signifies if the UE's radio capabilities needs an update
	NG_RAN_RCU uint8 `bitfield:"1"`
	// Optimizes small data transfers by routing through the control plane
	PNB_CIoT_5GS uint8 `bitfield:"2"`
	// Whether data is routed through User Plane or Control Plane
	PNB_CIoT_EPS uint8 `bitfield:"2"`
	// Two spare bits that should always be zero
	// This is achieved using reserved
	Spare uint8 `bitfield:"2,reserved"`
}

// This function encodes data in a Nas5GSUpdateType into a passed bytestream
func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer) {
	// Encodes struct with bitfields into a byte array
	new_buffer, ok := structex.EncodeByteBuffer(ie)

	if ok != nil {
		// Display custom prefix for the log message
		log.SetPrefix("Encode")
		// Remove timestamp, line numbers and so from the log message
		log.SetFlags(0)
		// Output the log message
		// Alternatively use .Fatal or .Panic to terminate the program
		log.Print(ok)
		return
	}

	// Write the encoded byte array to the passed buffer
	buffer.Write(new_buffer)
	// Error handling for buffer.Write not needed because this function will panic on an error
}
