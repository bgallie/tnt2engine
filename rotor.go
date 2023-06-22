// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

// Define the rotor used by tnt2engine.

import (
	"bytes"
	"fmt"
	"math/big"
)

var (
	// RotorSizes is an array of possible rotor sizes.  It consists of prime
	// numbers less than 8192 to ensure that the rotor sizes are realitivly prime.
	// The rotor sizes selected from this list will maximizes the number of unique
	// states the rotors can take.  The number of unique states range from
	// 183,599,058,301,611,293,854,881 to 297,245,983,088,018,794,170,091
	RotorSizes = [...]int{
		7823, 7829, 7841, 7853, 7867, 7873, 7877, 7879, 7883, 7901,
		7907, 7919, 7927, 7933, 7937, 7949, 7951, 7963, 7993, 8009,
		8011, 8017, 8039, 8053, 8059, 8069, 8081, 8087, 8089, 8093,
		8101, 8111, 8117, 8123, 8147, 8161, 8167, 8171, 8179, 8191}
	rotorSizes      []int
	rotorSizesIndex int
)

// Rotor - the type of the TNT2 rotor
type Rotor struct {
	Size    int    // the size in bits for this rotor
	Start   int    // the initial starting position of the rotor
	Step    int    // the step size in bits for this rotor
	Current int    // the current position of this rotor
	Rotor   []byte // the rotor
}

// New - initializes a new Rotor with the given size, start, step and rotor data.
func (r *Rotor) New(size, start, step int, rotor []byte) *Rotor {
	r.Start, r.Current = start, start
	r.Size = size
	r.Step = step
	r.Rotor = rotor
	r.sliceRotor()
	return r
}

// Update - updates the given Rotor with a new size, start, step and (psudo)
//   - random rotor data.
func (r *Rotor) Update(random *Rand) {
	// Get size, start and step of the new rotor
	rotorSize := RotorSizes[rotorSizes[rotorSizesIndex]]
	rotorSizesIndex = (rotorSizesIndex + 1) % len(RotorSizes)
	start := random.Intn(rotorSize)
	step := random.Intn(rotorSize-1) + 1
	// byteCnt is the total number of bytes needed to hold rotorSize bits + a slice of 256 bits
	byteCnt := ((rotorSize + CipherBlockSize + 7) / 8)
	// blkBytes is the number of bytes rotor r needs to increase to hold the new rotor.
	blkBytes := byteCnt - len(r.Rotor)
	// Adjust the size of r.Rotor to match the new rotor size.
	adjRotor := make([]byte, blkBytes)
	r.Rotor = append(r.Rotor, adjRotor...)
	// Fill the rotor with random data using tnt2engine Rand function to generate the
	// random data to fill the rotor.
	random.Read(r.Rotor)
	r.Size = rotorSize
	r.Step = step
	r.Start, r.Current = start, start
	r.sliceRotor()
}

// sliceRotor - appends the first 256 bits of the rotor to the end of the rotor.
func (r *Rotor) sliceRotor() {
	var i, j uint
	j = uint(r.Size)
	for i = 0; i < 256; i++ {
		if GetBit(r.Rotor, i) {
			SetBit(r.Rotor, j)
		} else {
			ClrBit(r.Rotor, j)
		}
		j++
	}
}

// SetIndex - set the current rotor position based on the given index
func (r *Rotor) SetIndex(idx *big.Int) {
	// Special case if idx == 0
	if idx.Sign() == 0 {
		r.Current = r.Start
	} else {
		p := new(big.Int)
		q := new(big.Int)
		rem := new(big.Int)
		p = p.Mul(idx, new(big.Int).SetInt64(int64(r.Step)))
		p = p.Add(p, new(big.Int).SetInt64(int64(r.Start)))
		_, rem = q.DivMod(p, new(big.Int).SetInt64(int64(r.Size)), rem)
		r.Current = int(rem.Int64())
	}
}

// Index - Rotor does no track the index.
func (r *Rotor) Index() *big.Int {
	return nil
}

// Get the number of bytes in "blk" from the given rotor.
func (r *Rotor) getRotorBlock(blk CipherBlock) CipherBlock {
	ress := make([]byte, len(blk))
	rotor := r.Rotor
	idx := r.Current
	blockSize := len(blk) * BitsPerByte

	for cnt := 0; cnt < blockSize; cnt++ {
		if GetBit(rotor, uint(idx)) {
			SetBit(ress, uint(cnt))
		}

		idx++
	}

	r.Current = (r.Current + r.Step) % r.Size
	return ress
}

// ApplyF - encrypts the given block of data.
func (r *Rotor) ApplyF(blk CipherBlock) CipherBlock {
	return AddBlock(blk, r.getRotorBlock(blk))
}

// ApplyG - decrypts the given block of data
func (r *Rotor) ApplyG(blk CipherBlock) CipherBlock {
	return SubBlock(blk, r.getRotorBlock(blk))
}

// String - converts a Rotor to a string representation of the Rotor.
func (r *Rotor) String() string {
	var output bytes.Buffer
	rotorLen := len(r.Rotor)
	output.WriteString(fmt.Sprintf("rotor.New(%d, %d, %d, []byte{\n",
		r.Size, r.Start, r.Step))
	for i := 0; i < rotorLen; i += 16 {
		output.WriteString("\t")
		if i+16 < rotorLen {
			for _, k := range r.Rotor[i : i+16] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
		} else {
			l := len(r.Rotor[i:])
			for _, k := range r.Rotor[i : i+l-1] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d})", r.Rotor[i+l-1]))
		}
		output.WriteString("\n")
	}

	return output.String()
}
