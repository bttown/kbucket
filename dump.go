package kbucket

import (
	"encoding/json"
	"io"
	"time"
)

type snapshotTable struct {
	BitSize int
	OwnerID []byte
	Buckets []*snapshotBucket
}

type snapshotBucket struct {
	K           int
	LastChanged time.Time

	Items map[string]*element
}

func dump(table *RouteTable, writer io.Writer) error {
	buckets := make([]*snapshotBucket, 0, len(table.Buckets))
	for i := 0; i < len(table.Buckets); i++ {
		buckets = append(buckets, &snapshotBucket{
			K:           table.Buckets[i].K,
			LastChanged: table.Buckets[i].lastChanged,
			Items:       table.Buckets[i].items,
		})
	}
	snapTable := &snapshotTable{
		BitSize: table.BitSize,
		OwnerID: table.ownerID,
		Buckets: buckets,
	}

	return json.NewEncoder(writer).Encode(snapTable)
}

func load(table *RouteTable, reader io.Reader) error {
	var snapshot = new(snapshotTable)
	err := json.NewDecoder(reader).Decode(snapshot)
	if err != nil {
		panic(err)
	}

	buckets := make([]*Bucket, 0, len(snapshot.Buckets))
	for i := 0; i < len(snapshot.Buckets); i++ {
		buckets = append(buckets, &Bucket{
			K:           snapshot.Buckets[i].K,
			lastChanged: snapshot.Buckets[i].LastChanged,
			items:       snapshot.Buckets[i].Items,
		})
	}
	table.ownerID = snapshot.OwnerID
	table.BitSize = snapshot.BitSize
	table.Buckets = buckets
	return nil
}
