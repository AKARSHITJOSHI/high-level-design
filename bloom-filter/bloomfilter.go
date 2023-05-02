package filter

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash"
)

const (
	filterSize = 1000
)

type BloomFilter struct {
	bitfield [filterSize]bool
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		// bitfield: [filterSize]bool,
	}
}

func getHash(h hash.Hash, value string) int {
	result := h.Sum([]byte(value))
	data := binary.BigEndian.Uint64(result)
	return int(data) % filterSize // cast the int64
}

func (b *BloomFilter) Add(input string) {
	var md5 = md5.New()
	var sha256 = sha256.New()
	var sha512 = sha512.New()
	b.bitfield[getHash(md5, input)] = true
	b.bitfield[getHash(sha256, input)] = true
	b.bitfield[getHash(sha512, input)] = true
}

func (b *BloomFilter) IsPresent(input string) bool {
	var md5 = md5.New()
	var sha256 = sha256.New()
	var sha512 = sha512.New()
	return b.bitfield[getHash(md5, input)] && b.bitfield[getHash(sha256, input)] && b.bitfield[getHash(sha512, input)]
}

func (b *BloomFilter) Clear() {
	for i := 0; i < len(b.bitfield); i++ {
		b.bitfield[i] = false
	}
}
