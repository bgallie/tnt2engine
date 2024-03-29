// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

// Define the random number generator using tnt2engine as the source.

import (
	"fmt"
	"os"
)

var emptyBlk CipherBlock

const (
	intSize = 32 << (^uint(0) >> 63)
	rngMax  = 1 << 63
	rngMask = rngMax - 1
)

type Rand struct {
	tnt2Machine *Tnt2Engine
	idx         int
	blk         CipherBlock
}

func NewRand(src *Tnt2Engine) *Rand {
	fmt.Fprintln(os.Stderr, "WARNING: rand.NewRand() is deprecated.  Use Rand.New() instead")
	return new(Rand).New(src)
}

func (rnd *Rand) New(src *Tnt2Engine) *Rand {
	rnd.tnt2Machine = src
	rnd.idx = CipherBlockBytes
	return rnd
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if intSize == 32 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 {
	return int32(r.Int63() >> 32)
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 {
	return uint32(r.Int63() >> 31)
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (rnd *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return rnd.Int63() & (n - 1)
	}
	max := int64(rngMask - (1<<63)%uint64(n))
	v := rnd.Int63()
	for v > max {
		v = rnd.Int63()
	}
	return v % n
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rnd *Rand) Int63() int64 {
	return rnd.Int64() & rngMask
}

// Int64 returns a pseudo-random 64-bit value as a int64.
func (rnd *Rand) Int64() int64 {
	bytes := make([]byte, 8)
	_, _ = rnd.Read(bytes)
	var n int64
	for _, val := range bytes {
		n = (n << 8) | int64(val)
	}
	return n
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (rnd *Rand) Uint64() uint64 {
	bytes := make([]byte, 8)
	_, _ = rnd.Read(bytes)
	var n uint64
	for _, val := range bytes {
		n = (n << 8) | uint64(val)
	}
	return n
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n).
func (rnd *Rand) Perm(n int) (res []int) {
	if n < 0 {
		panic(fmt.Sprintf("Perm called with a negative argument [%d]", n))
	}
	res = make([]int, n)
	if n > 0 {
		if n == 1 {
			res[0] = 0
		} else {
			for i := range res {
				res[i] = i
			}
		}
	}
	for i := (n - 1); i > 0; i-- {
		j := rnd.Intn(i)
		res[i], res[j] = res[j], res[i]
	}
	return
}

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
// Read should not be called concurrently with any other Rand method.
func (rnd *Rand) Read(p []byte) (n int, err error) {
	// On the first call to Read(), initialize rnd.blk with 32 pseudo-random
	// bytes that are based on the secret key so that the same sequence
	// will be generated.
	if string(rnd.blk) == string(emptyBlk) {
		cntrKeyBytes := rnd.tnt2Machine.cntrKey[:]
		cntrKeyBytes = jc1Key.XORKeyStream(cntrKeyBytes)
		rnd.blk = make(CipherBlock, CipherBlockBytes)
		_ = copy(rnd.blk[:], cntrKeyBytes)
	}
	// Clear out p to receive the pseudo-random data.
	p = p[:0]
	// Read pseudo-random data generated by the tnt2engine (32 bytes at a time)
	// and appending them into p.  It reads as many 32 bytes blocks as needed
	// to fill p.  If not all bytes in the 32 byte block are needed, then the
	// extra bytes are saved for the next call to Read()
	left := rnd.tnt2Machine.Left()
	right := rnd.tnt2Machine.Right()
	for {
		if rnd.idx >= CipherBlockBytes { // No more data in rnd.blk, so get more
			left <- rnd.blk
			rnd.blk = <-right
			rnd.idx = 0
		}
		leftInBlk := len(rnd.blk) - rnd.idx // calculate how many bytes are left in rnd.blk
		remaining := cap(p) - len(p)        // calculate how many bytes are still needed to fill p.
		if remaining >= leftInBlk {         // there is enough room to all the bytes in rnd.blk
			p = append(p, rnd.blk[rnd.idx:]...) // so append all of the bytes into p.
			rnd.idx += leftInBlk
		} else { // append only the needed bytes into p.
			p = append(p, rnd.blk[rnd.idx:rnd.idx+remaining]...)
			rnd.idx += remaining
			break
		}
		// and repeat until p is filled.
	}

	return len(p), nil
}
