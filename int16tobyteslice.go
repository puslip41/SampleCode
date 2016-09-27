package main

import (
					"fmt"
	"encoding/binary"
)

func main() {
	var i int16 = 258
	
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))

	fmt.Println("LittleEndian : %d -> %x", i, b)

	binary.BigEndian.PutUint16(b, uint16(i))
	fmt.Println("BigEndian : %d -> %x", i, b)
}
