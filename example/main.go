package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bttown/kbucket"
)

func main() {
	ownerID, _ := hex.DecodeString("f2404441fce23c585ff0170c387c24e859a7704a")
	table := kbucket.New(ownerID)

	nID := []byte{242, 64, 68, 65, 252, 226, 60, 88, 95, 240, 23, 12, 56, 124, 36, 232, 89, 167, 112, 74}
	fmt.Println(hex.EncodeToString(nID))
	fmt.Println(table.BucketID(nID))
}
