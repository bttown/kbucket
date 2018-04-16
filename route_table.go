package kbucket

import (
	"bytes"
	"io"
	"log"
	"os"
)

type RouteTable struct {
	BitSize             int
	ownerID             []byte
	Buckets             []*Bucket
	lastChangedBucketID int
}

func NewFromDumpFile(filepath string) (*RouteTable, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var table = new(RouteTable)
	err = load(table, file)
	return table, err
}

func New(ownerID []byte) *RouteTable {
	table := RouteTable{
		ownerID: ownerID,
		BitSize: len(ownerID) * 8,
	}

	var k = 8
	buckets := make([]*Bucket, table.BitSize)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = NewBucket(k)
	}
	table.Buckets = buckets

	return &table
}

func (table *RouteTable) OwnerID() []byte {
	return table.ownerID
}

func (table *RouteTable) Neighbors(nodeID []byte) []interface{} {
	bucketID := table.BucketID(nodeID)
	bucket := table.Buckets[bucketID]

	log.Println(bucketID)

	nodes := bucket.Nodes()
	if len(nodes) == 0 {
		nodes = table.Buckets[table.lastChangedBucketID].Nodes()
	}

	return nodes
}

func (table *RouteTable) Add(v Contact) bool {
	bucketID := table.BucketID(v.ID)
	bucket := table.Buckets[bucketID]

	ok := bucket.Add(v)
	if ok {
		// log.Printf("add %s to bucket %d", v.GetStringID(), bucketID)
	}

	table.lastChangedBucketID = bucketID
	return ok
}

func (table *RouteTable) Dump(writer io.Writer) error {
	return dump(table, writer)
}

func (table *RouteTable) BucketID(nodeID []byte) int {
	if bytes.Equal(nodeID, table.ownerID) {
		return 0
	}
	var i int
	var bite byte
	var bitDiff int
	var v byte
	for i, bite = range table.ownerID {
		v = bite ^ nodeID[i]
		switch {
		case v > 0x70:
			bitDiff = 8
			goto calc
		case v > 0x40:
			bitDiff = 7
			goto calc
		case v > 0x20:
			bitDiff = 6
			goto calc
		case v > 0x10:
			bitDiff = 5
			goto calc
		case v > 0x08:
			bitDiff = 4
			goto calc
		case v > 0x04:
			bitDiff = 3
			goto calc
		case v > 0x02:
			bitDiff = 2
			goto calc
		}
	}

calc:

	return i*8 + (8 - bitDiff)
}
