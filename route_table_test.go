package kbucket

import (
	"fmt"
	"testing"
)

func TestBucketID(t *testing.T) {
	ownerID := []byte{0xf0}
	router := New(ownerID) // bit size is 4
	bucketID := router.BucketID([]byte{0x80})
	fmt.Println(bucketID)
}
