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

// New fills the (empty) rotor r with the given size, start, step and rotor data.
func (r *Rotor) New(size, start, step int, rotor []byte) *Rotor {
	r.Start, r.Current = start, start
	r.Size = size
	r.Step = step
	r.Rotor = append([]byte(nil), rotor...)
	r.sliceRotor()
	return r
}

// Update - updates the given Rotor with a new size, start, step and (psudo)
//   - random rotor data.
func (r *Rotor) Update(random *Rand) {
	// Get size, start and step of the new rotor
	rotorSize := RotorSizes[rotorSizesIndex]
	rotorSizesIndex--
	start := random.Intn(rotorSize)
	step := random.Intn(rotorSize-1) + 1
	// byteCnt is the total number of bytes needed to hold rotorSize bits + a slice of 256 bits
	byteCnt := ((rotorSize + CipherBlockSize + 7) / 8)
	// blkBytes is the number of bytes rotor r needs to increase to hold the new rotor.
	blkBytes := byteCnt - len(r.Rotor)
	// Adjust the size of r.Rotor to match the new rotor size.
	adjRotor := make([]byte, blkBytes)
	r.Rotor = append(r.Rotor, adjRotor...)
	// Fill the rotor with random data using tntengine Rand function to generate the
	// random data to fill the rotor.  As an optimization, set the random.idx equal to
	// CypherBlockSize.  This will ensure that the random.Read() will only call the
	// encryption function one time per block instead of up to two times per block.
	random.idx = CipherBlockBytes
	blk := make(CipherBlock, CipherBlockBytes)
	CipherBlocksToRead := (rotorSize + 7) / CipherBlockSize
	j := 0
	for i := 0; i < CipherBlocksToRead; i++ {
		random.Read(blk)
		copy(r.Rotor[j:], blk[:])
		j += CipherBlockBytes
	}
	r.Size = rotorSize
	r.Step = step
	r.Start, r.Current = start, start
	r.sliceRotor() // Append the first 256 bits of the rotor to the end of the rotor
}

// sliceRotor appends the first 256 bits of the rotor to the end of the rotor.
func (r *Rotor) sliceRotor() {
	var size, sBlk, sBit, Rshift, Lshift uint
	var i int
	size = uint(r.Size)
	sBlk = size >> 3
	sBit = size & 7
	Rshift = 8 - sBit
	Lshift = sBit
	if sBit == 0 {
		copy(r.Rotor[sBlk:], r.Rotor[0:CipherBlockBytes])
	} else {
		// The copy appending will be done at the byte level instead of the bit level
		// so that we only loop 32 times instead of 256 times.
		for i = 0; i < CipherBlockBytes; i++ {
			r.Rotor[sBlk] &= (0xff >> Rshift)       // Clear out the bits that will be replaced
			r.Rotor[sBlk] |= (r.Rotor[i] << Lshift) // and add in the bits from the beginning of the rotor
			sBlk++
			r.Rotor[sBlk] = (r.Rotor[i] >> Rshift) // Seed the next byte at the end with the remaining bits from the beginning byte.
		}
	}
}

// SetIndex positions the rotor to the position it would be in after
// processing idx number of blocks.
func (r *Rotor) SetIndex(idx *big.Int) {
	// Special case if idx == 0
	if idx.Sign() == 0 {
		r.Current = r.Start
	} else {
		// Calculate the new r.Current:
		// r.Current = mod(((idx * r.Step) + r.Start), r.Size) + r.Start
		p := new(big.Int)
		q := new(big.Int)
		rem := new(big.Int)
		p = p.Mul(idx, new(big.Int).SetInt64(int64(r.Step)))
		p = p.Add(p, new(big.Int).SetInt64(int64(r.Start)))
		_, rem = q.DivMod(p, new(big.Int).SetInt64(int64(r.Size)), rem)
		r.Current = int(rem.Int64())
	}
}

// Always return nil since the block count is not tracked for rotors.
func (r *Rotor) Index() *big.Int {
	return nil
}

// Get the number of bits in the CipherBlock from rotor r.
func (r *Rotor) getRotorBlock(bytes int) CipherBlock {
	// This code handles short blocks to accomadate file lenghts
	// that are not multiples of "CipherBlockBytes"
	ress := make([]byte, bytes)
	rotor := r.Rotor
	sBit := r.Current & 7
	bIdx := r.Current >> 3
	// The copy operates at the byte level instead of the bit level
	// so we only loop 32 times instead of 256 times.
	if sBit == 0 {
		copy(ress, rotor[bIdx:])
	} else {
		sLeft := 8 - sBit
		for bCnt := 0; bCnt < bytes; bCnt++ {
			ress[bCnt] = rotor[bIdx]>>sBit |
				(rotor[bIdx+1] << sLeft)
			bIdx++
		}
	}
	// Step the rotor to its new position.
	r.Current = (r.Current + r.Step) % r.Size
	return ress
}

// ApplyF encrypts the given block of data using the rotor r.
func (r *Rotor) ApplyF(blk CipherBlock) CipherBlock {
	// Add (not XOR) the rotor bits to the bits in the input block.
	return AddBlock(blk, r.getRotorBlock(len(blk)))
}

// ApplyG decrypts the given block of data using the rotor r.
func (r *Rotor) ApplyG(blk CipherBlock) CipherBlock {
	// Subtract (not XOR) the rotor bits from the bits in the input block.
	return SubBlock(blk, r.getRotorBlock(len(blk)))
}

// String converts a Rotor to a string representation of the Rotor.
func (r *Rotor) String() string {
	var output bytes.Buffer
	rotorLen := len(r.Rotor)
	output.WriteString(fmt.Sprintf("new(Rotor).New(%d, %d, %d, []byte{\n",
		r.Size, r.Start, r.Step))
	for i := 0; i < rotorLen; i += 16 {
		output.WriteString("\t")
		if i+16 < rotorLen {
			for _, k := range r.Rotor[i : i+15] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d,", r.Rotor[i+15]))
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
