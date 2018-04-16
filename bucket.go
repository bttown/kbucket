package kbucket

import (
	"sync"
	"time"
)

type Bucket struct {
	K           int
	lastChanged time.Time

	items map[string]*element

	rw sync.RWMutex
}

func NewBucket(k int) *Bucket {
	bucket := Bucket{
		K:           k,
		lastChanged: time.Now(),
	}
	bucket.items = make(map[string]*element)

	return &bucket
}

func (bucket *Bucket) Len() int {
	bucket.rw.RLock()
	defer bucket.rw.RUnlock()

	return len(bucket.items)
}

func (bucket *Bucket) Add(v Contact) bool {
	bucket.rw.Lock()
	defer bucket.rw.Unlock()

	key := v.GetStringID()
	_, ok := bucket.items[key]
	if ok {
		return false
	}

	if len(bucket.items) == bucket.K {
		for key, _ := range bucket.items {
			delete(bucket.items, key)
			break
		}
	}

	element := newElement(v)
	bucket.items[key] = element

	return true
}

func (bucket *Bucket) Nodes() []interface{} {
	bucket.rw.RLock()
	defer bucket.rw.RUnlock()

	var l = make([]interface{}, 0, len(bucket.items))
	for _, e := range bucket.items {
		l = append(l, e.Val)
	}
	return l
}
