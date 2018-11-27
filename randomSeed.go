package main

import (
	"crypto/rand"
	"encoding/binary"
)

func randomInt64Seed() int64 {
	key := [8]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	return int64(binary.BigEndian.Uint64(key[:]))
}
