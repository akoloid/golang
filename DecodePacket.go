package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type DecodedStruct struct {
	Short1 int16
	Chars1 string
	Byte1  uint8
	Chars2 string
	Short2 int16
	Chars3 string
	Long1  uint32
}

func decodePacket(packet []byte) (DecodedStruct, error) {
	if len(packet) != 44 {
		return DecodedStruct{}, fmt.Errorf("Invalid packet size. Expected 44 bytes.")
	}

	var decoded DecodedStruct

	reader := binary.LittleEndian
	offset := 0

	if err := binary.Read(bytes.NewReader(packet[offset:offset+2]), reader, &decoded.Short1); err != nil {
		return DecodedStruct{}, err
	}
	offset += 2

	decoded.Chars1 = string(packet[offset : offset+12])
	offset += 12

	if err := binary.Read(bytes.NewReader(packet[offset:offset+1]), reader, &decoded.Byte1); err != nil {
		return DecodedStruct{}, err
	}
	offset += 1

	decoded.Chars2 = string(packet[offset : offset+8])
	offset += 8

	if err := binary.Read(bytes.NewReader(packet[offset:offset+2]), reader, &decoded.Short2); err != nil {
		return DecodedStruct{}, err
	}
	offset += 2

	decoded.Chars3 = string(packet[offset : offset+15])
	offset += 15

	if err := binary.Read(bytes.NewReader(packet[offset:offset+4]), reader, &decoded.Long1); err != nil {
		return DecodedStruct{}, err
	}

	return decoded, nil
}

func main() {
	packet := []byte{0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64,
		0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75,
		0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15}

	decoded, err := decodePacket(packet)
	if err != nil {
		fmt.Println("Error decoding packet:", err)
		return
	}

	fmt.Printf("Decoded struct: %+v\n", decoded)
}
