// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.
package tnt2engine

// Define a permutator used in tnt2engine

import (
	"bytes"
	"fmt"
	"math/big"
	"sync"
)

var (
	// CycleSizes is an array of cycles to use when cycling the permutation table.
	// There are 4 cycles in each entry and they meet the following criteria:
	//      1.  The sum of the cycles is equal to 256.
	//      2.  The cycles are relatively prime to each other. (This maximizes
	//          the number of unique states the permutation can be in for the
	//          given cycles).
	CycleSizes = [...][NumberPermutationCycles]int{
		{61, 63, 65, 67}, // Number of unique states: 16,736,265
		{53, 65, 67, 71}, // Number of unique states: 16,387,685
		{55, 57, 71, 73}, // Number of unique states: 16,248,705
		{53, 61, 63, 79}, // Number of unique states: 16,090,641
		{43, 57, 73, 83}, // Number of unique states: 14,850,609
		{49, 51, 73, 83}, // Number of unique states: 15,141,441
		{47, 53, 73, 83}, // Number of unique states: 15,092,969
		{47, 53, 71, 85}} // Number of unique states: 15,033,185

	// CyclePermutations is an array of possible orderings that a particular
	// set of four (4) cycle sizes can take.  This is used to increase the number
	// of bitperms that can be generated from the randp table, increasing the
	// complexity that the cryptoanalysis faces.
	CyclePermutations = [...][NumberPermutationCycles]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 2, 1}, {0, 3, 1, 2},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 2, 0}, {1, 3, 0, 2},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 1, 0}, {2, 3, 0, 1},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 1, 0}, {3, 2, 0, 1}}
	cycleSizes      []int
	cycleSizesIndex int
)

// Cycle describes a cycle for the permutator so it can adjust the permutation
// table used to permutate the block.  TNT2 currently uses 4 cycles to rearrange
// Randp into bitPerm
type Cycle struct {
	Start   int // The starting point (into randp) for this cycle.
	Length  int // The length of the cycle.
	Current int // The point in the cycle [0 .. cycle.length-1] to start
}

// Permutator is a type that defines a permutation cryptor in TNT2.
type Permutator struct {
	CurrentState  int                   // Current number of cycles for this permutator.
	MaximalStates int                   // Maximum number of cycles this permutator can have before repeating.
	Cycles        []Cycle               // Cycles ordered by the current permutation.
	Randp         []byte                // Values 0 - 255 in a random order.
	bitPerm       [CipherBlockSize]byte // Permutation table created from Randp.
}

// New creates a permutator and initializes it
func (p *Permutator) New(cycleSizes []int, randp []byte) *Permutator {
	p.CurrentState = 0
	p.MaximalStates = 1
	p.Randp = append(p.Randp[:0], randp...)
	// Create the Cycles structure based on the given cycle sizes
	p.Cycles = make([]Cycle, len(cycleSizes))
	for i := range cycleSizes {
		p.Cycles[i].Length = cycleSizes[i]
		p.Cycles[i].Current = 0
		// Adjust the start to reflect the lenght of the previous cycles
		if i == 0 { // no previous cycle so start at 0
			p.Cycles[i].Start = 0
		} else {
			p.Cycles[i].Start = p.Cycles[i-1].Start + p.Cycles[i-1].Length
		}
		// Calculate the maximum number of states the permutator can take.
		p.MaximalStates *= p.Cycles[i].Length
	}
	p.cycle()
	return p
}

// Update will update the given (proForma) permutator in place using
// (psudo-random) data generated by the TNT2 encrytption algorithm
// using the proForma rotors and permutators.
func (p *Permutator) Update(random *Rand) {
	// updated permutators start with a current state of 0
	p.CurrentState = 0
	p.MaximalStates = 1
	// Create a table of byte values [0...255] in a random order
	for i, val := range random.Perm(CipherBlockSize) {
		p.Randp[i] = byte(val)
	}
	// Chose a size of the cycles to use and randomize order of the values
	length := len(CycleSizes[cycleSizesIndex])
	cycles := make([]int, length)
	randi := random.Perm(length)
	for idx, val := range randi {
		cycles[idx] = CycleSizes[cycleSizes[cycleSizesIndex]][val]
	}
	cycleSizesIndex = (cycleSizesIndex + 1) % len(CycleSizes)
	// update p.Cycles based on the new cycle sizes
	for i := range cycles {
		p.Cycles[i].Length = cycles[i]
		p.Cycles[i].Current = 0
		// Adjust the start to reflect the lenght of the previous cycles
		if i == 0 { // no previous cycle so start at 0
			p.Cycles[i].Start = 0
		} else {
			p.Cycles[i].Start = p.Cycles[i-1].Start + p.Cycles[i-1].Length
		}
		p.MaximalStates *= p.Cycles[i].Length
	}
	p.cycle()
}

// cycle bitPerm to it's next state.
func (p *Permutator) nextState() {
	for idx := 0; idx < len(p.Cycles); idx++ {
		p.Cycles[idx].Current = (p.Cycles[idx].Current + 1) % p.Cycles[idx].Length
	}
	p.CurrentState = (p.CurrentState + 1) % p.MaximalStates
	p.cycle()
}

// cycle will create a new bitPerm from Randp based on the current cycle.
func (p *Permutator) cycle() {
	var wg sync.WaitGroup
	for cycleIdx := range p.Cycles {
		wg.Add(1)
		go func(cIdx int) {
			defer wg.Done()
			curCycle := p.Cycles[cIdx]
			cycle := p.Randp[curCycle.Start : curCycle.Start+curCycle.Length]
			sIdx := curCycle.Current
			length := curCycle.Length
			for _, val := range cycle {
				p.bitPerm[val] = p.Randp[cycle[sIdx]]
				sIdx = (sIdx + 1) % length
			}
		}(cycleIdx)
	}
	wg.Wait()
}

// SetIndex - set the Permutator to the state it would be in after encoding 'idx - 1' blocks
// of data.
func (p *Permutator) SetIndex(idx *big.Int) {
	q := new(big.Int)
	r := new(big.Int)
	_, r = q.DivMod(idx, big.NewInt(int64(p.MaximalStates)), r)
	p.CurrentState = int(r.Int64())
	for i := 0; i < NumberPermutationCycles; i++ {
		p.Cycles[i].Current = p.CurrentState % p.Cycles[i].Length
	}
	p.cycle()
}

// Index returns the current index of the cryptor.  For permeutators, this
// returns nil.
func (p *Permutator) Index() *big.Int {
	return nil
}

// ApplyF performs forward permutation on the 32 byte block of data.
func (p *Permutator) ApplyF(blk CipherBlock) CipherBlock {
	if len(blk) == CipherBlockBytes {
		ress := make([]byte, CipherBlockBytes)
		for i, v := range p.bitPerm {
			if GetBit(blk, uint(i)) {
				SetBit(ress, uint(v))
			}
		}
		p.nextState()
		blk = ress
	}
	return blk
}

// ApplyG performs the reverse permutation on the 32 byte block of data.
func (p *Permutator) ApplyG(blk CipherBlock) CipherBlock {
	if len(blk) == CipherBlockBytes {
		ress := make([]byte, CipherBlockBytes)
		for i, v := range p.bitPerm {
			if GetBit(blk, uint(v)) {
				SetBit(ress, uint(i))
			}
		}
		p.nextState()
		blk = ress
	}
	return blk
}

// String formats a string representing the permutator (as Go source code).
func (p *Permutator) String() string {
	var output bytes.Buffer
	output.WriteString("permutator.New([]int{")
	for _, v := range p.Cycles[0 : NumberPermutationCycles-1] {
		output.WriteString(fmt.Sprintf("%d, ", v.Length))
	}
	output.WriteString(fmt.Sprintf("%d}, []byte{\n", p.Cycles[NumberPermutationCycles-1].Length))
	for i := 0; i < CipherBlockSize; i += 16 {
		output.WriteString("\t")
		if i != (CipherBlockSize - 16) {
			for _, k := range p.Randp[i : i+16] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
		} else {
			for _, k := range p.Randp[i : i+15] {
				output.WriteString(fmt.Sprintf("%d, ", k))
			}
			output.WriteString(fmt.Sprintf("%d})", p.Randp[i+15]))
		}
		output.WriteString("\n")
	}
	return output.String()
}
