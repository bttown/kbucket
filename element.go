package kbucket

import (
	"encoding/hex"
	"net"
	"time"
)

type Contact struct {
	ID []byte
	net.UDPAddr
}

func (c *Contact) GetID() []byte {
	return c.ID
}

func (c *Contact) GetStringID() string {
	return hex.EncodeToString(c.ID)
}

type element struct {
	Val Contact

	LastChanged time.Time
}

func newElement(v Contact) *element {
	elm := element{
		Val:         v,
		LastChanged: time.Now(),
	}

	return &elm
}
