// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

// Define the tnt2engine type and it's methods

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/bgallie/jc1"
	"golang.org/x/crypto/sha3"
)

var (
	EngineLayout   = "rrprrprr"
	proFormaRotors = []*Rotor{
		// Define the proforma rotors used to create the actual rotors to use.
		new(Rotor).New(1783, 863, 1033, []byte{
			184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
			154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
			59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232,
			87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57,
			131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26,
			247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14,
			129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245,
			41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40,
			78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37,
			232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250,
			81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66,
			50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94,
			204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123,
			111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220,
			12, 95, 253, 145, 133, 55, 237, 183, 0, 150, 157, 68, 6, 92, 11, 77,
			241, 50, 172, 211, 182, 22, 174, 9, 82, 194, 116, 145, 66, 69, 111, 0}),
		new(Rotor).New(1753, 1494, 1039, []byte{
			100, 120, 105, 253, 78, 6, 70, 91, 136, 33, 73, 16, 15, 13, 174, 206,
			97, 207, 186, 14, 141, 185, 228, 85, 161, 253, 190, 198, 234, 193, 63, 20,
			63, 229, 90, 58, 254, 193, 63, 69, 156, 75, 113, 145, 167, 124, 26, 38,
			94, 117, 42, 25, 81, 251, 172, 67, 175, 138, 159, 85, 66, 180, 187, 101,
			204, 45, 222, 90, 143, 217, 32, 9, 109, 71, 24, 223, 43, 196, 181, 175,
			67, 118, 69, 154, 201, 178, 228, 137, 216, 184, 102, 29, 148, 77, 27, 139,
			90, 20, 115, 102, 91, 37, 244, 44, 9, 254, 144, 216, 214, 201, 70, 160,
			127, 154, 161, 160, 125, 210, 16, 141, 151, 211, 117, 153, 153, 75, 141, 252,
			109, 76, 251, 215, 116, 31, 224, 156, 56, 112, 40, 36, 180, 156, 214, 190,
			122, 206, 11, 172, 52, 68, 167, 87, 53, 234, 125, 167, 21, 100, 193, 166,
			26, 9, 237, 249, 101, 142, 141, 49, 210, 254, 139, 72, 88, 148, 223, 216,
			251, 70, 63, 0, 182, 75, 137, 218, 178, 155, 101, 102, 195, 226, 193, 26,
			9, 12, 147, 186, 248, 43, 5, 117, 133, 78, 14, 201, 165, 155, 206, 57,
			120, 35, 117, 215, 16, 129, 104, 133, 173, 50, 38, 200, 240, 210, 250, 157,
			12, 140, 182, 16, 67, 146, 32, 30, 26, 92, 157, 195, 158, 117, 29, 26,
			115, 201, 171, 66, 251, 125, 141, 213, 131, 127, 40, 102}),
		new(Rotor).New(1721, 1250, 660, []byte{
			25, 134, 2, 219, 108, 110, 170, 11, 12, 129, 29, 172, 198, 2, 14, 255,
			158, 7, 103, 114, 63, 69, 173, 156, 249, 147, 235, 203, 90, 200, 233, 73,
			38, 137, 10, 93, 176, 253, 64, 85, 46, 136, 21, 220, 37, 109, 149, 169,
			165, 153, 37, 42, 63, 35, 65, 196, 237, 215, 100, 226, 151, 53, 172, 215,
			240, 111, 136, 4, 47, 134, 80, 10, 165, 192, 212, 158, 48, 116, 89, 211,
			76, 120, 62, 226, 174, 97, 105, 33, 118, 245, 247, 162, 179, 90, 207, 178,
			69, 114, 201, 206, 93, 130, 79, 199, 223, 120, 233, 66, 86, 178, 59, 104,
			16, 217, 189, 78, 8, 249, 139, 156, 141, 222, 143, 8, 155, 96, 216, 156,
			210, 214, 108, 1, 80, 147, 10, 50, 53, 32, 78, 176, 6, 183, 11, 251,
			130, 192, 204, 184, 131, 159, 142, 127, 170, 183, 238, 60, 6, 47, 77, 30,
			125, 91, 170, 213, 209, 57, 250, 143, 252, 174, 54, 177, 55, 216, 220, 17,
			194, 54, 199, 66, 201, 194, 117, 226, 223, 146, 194, 177, 11, 93, 66, 182,
			46, 122, 253, 161, 204, 40, 167, 40, 92, 37, 134, 155, 0, 231, 21, 105,
			73, 171, 159, 246, 182, 91, 87, 50, 12, 5, 182, 217, 220, 84, 23, 24,
			2, 59, 88, 141, 5, 28, 254, 61, 15, 206, 228, 126, 138, 90, 57, 243,
			39, 215, 151, 181, 144, 211, 147, 198}),
		new(Rotor).New(1741, 1009, 1513, []byte{
			59, 155, 29, 153, 190, 106, 54, 89, 63, 156, 123, 112, 152, 24, 237, 200,
			85, 31, 249, 221, 7, 186, 76, 48, 229, 63, 232, 43, 60, 224, 108, 113,
			71, 154, 254, 136, 83, 102, 6, 108, 108, 138, 65, 104, 190, 98, 197, 120,
			244, 159, 191, 154, 224, 194, 37, 255, 51, 135, 123, 162, 17, 170, 199, 216,
			247, 94, 186, 218, 204, 48, 242, 65, 203, 30, 22, 226, 242, 57, 40, 32,
			22, 231, 138, 222, 125, 10, 125, 108, 24, 59, 221, 99, 156, 96, 214, 129,
			20, 227, 252, 198, 205, 71, 208, 99, 94, 247, 115, 76, 198, 106, 70, 134,
			143, 223, 158, 226, 204, 99, 210, 71, 139, 87, 33, 236, 30, 244, 49, 223,
			228, 215, 142, 236, 68, 74, 166, 97, 216, 67, 14, 41, 128, 40, 55, 70,
			235, 130, 50, 118, 198, 96, 87, 26, 134, 122, 174, 119, 237, 6, 239, 91,
			84, 144, 211, 239, 252, 172, 143, 151, 5, 249, 200, 38, 149, 31, 224, 68,
			100, 250, 25, 173, 38, 74, 133, 18, 244, 7, 138, 0, 85, 143, 137, 140,
			38, 95, 191, 129, 109, 227, 224, 28, 66, 39, 80, 45, 49, 78, 63, 245,
			4, 42, 118, 84, 72, 204, 145, 70, 139, 113, 103, 179, 35, 211, 87, 205,
			38, 235, 135, 115, 15, 14, 19, 163, 29, 185, 234, 35, 191, 251, 64, 151,
			9, 166, 252, 7, 125, 133, 7, 156, 45, 14}),
		new(Rotor).New(1723, 1293, 1046, []byte{
			59, 137, 3, 62, 80, 176, 170, 169, 12, 135, 154, 73, 218, 169, 34, 130,
			71, 240, 156, 66, 122, 214, 138, 174, 35, 15, 210, 20, 0, 17, 47, 172,
			227, 243, 160, 166, 101, 87, 0, 83, 16, 204, 69, 56, 249, 1, 107, 129,
			30, 236, 248, 46, 59, 16, 136, 240, 7, 68, 175, 181, 102, 24, 221, 34,
			206, 73, 37, 100, 74, 5, 82, 49, 42, 77, 33, 219, 30, 140, 122, 201,
			173, 86, 171, 7, 139, 239, 119, 224, 83, 33, 167, 38, 38, 252, 238, 109,
			173, 151, 153, 182, 170, 199, 109, 174, 85, 177, 165, 37, 171, 94, 247, 29,
			178, 32, 54, 252, 180, 240, 170, 188, 119, 168, 101, 220, 147, 32, 153, 5,
			15, 239, 180, 141, 232, 143, 14, 49, 98, 69, 224, 22, 134, 220, 139, 165,
			26, 189, 188, 120, 113, 196, 95, 124, 238, 91, 217, 213, 114, 32, 177, 200,
			216, 95, 142, 54, 252, 162, 46, 35, 191, 106, 48, 42, 71, 37, 16, 157,
			79, 66, 33, 12, 120, 31, 247, 54, 48, 189, 177, 142, 183, 152, 122, 252,
			139, 150, 164, 251, 77, 9, 128, 220, 145, 27, 85, 162, 42, 154, 151, 87,
			176, 158, 233, 135, 198, 224, 14, 216, 73, 28, 240, 129, 130, 85, 77, 101,
			56, 212, 76, 210, 78, 21, 17, 60, 130, 231, 20, 210, 179, 86, 116, 29,
			121, 144, 166, 0, 136, 120, 97, 5}),
		new(Rotor).New(1733, 1313, 1414, []byte{
			141, 233, 47, 225, 230, 220, 229, 226, 34, 136, 160, 200, 162, 159, 148, 163,
			157, 133, 38, 86, 25, 23, 18, 48, 5, 98, 112, 20, 37, 159, 82, 163,
			209, 135, 40, 197, 152, 8, 255, 234, 149, 22, 158, 19, 235, 186, 173, 247,
			109, 77, 243, 223, 143, 165, 33, 110, 122, 181, 130, 242, 116, 132, 205, 43,
			4, 81, 85, 99, 152, 109, 9, 180, 190, 100, 204, 226, 97, 214, 214, 200,
			169, 61, 53, 107, 128, 231, 15, 42, 162, 156, 119, 166, 223, 143, 234, 16,
			220, 234, 132, 0, 200, 20, 164, 12, 216, 165, 86, 49, 149, 83, 200, 208,
			151, 80, 65, 60, 102, 69, 55, 248, 199, 233, 6, 239, 204, 212, 244, 89,
			255, 240, 54, 232, 189, 143, 233, 51, 44, 167, 97, 2, 71, 233, 154, 155,
			213, 203, 55, 110, 48, 187, 130, 84, 87, 71, 158, 91, 42, 21, 229, 161,
			2, 176, 152, 186, 16, 99, 185, 200, 245, 89, 186, 173, 54, 78, 101, 242,
			169, 224, 83, 242, 78, 39, 93, 123, 86, 196, 13, 82, 104, 92, 139, 230,
			35, 84, 182, 162, 19, 119, 20, 62, 214, 197, 134, 75, 57, 52, 91, 37,
			225, 167, 86, 81, 159, 46, 98, 38, 166, 49, 253, 37, 220, 156, 187, 92,
			92, 4, 17, 20, 89, 244, 147, 114, 180, 179, 208, 196, 42, 227, 66, 2,
			166, 64, 12, 142, 162, 228, 83, 106, 244}),
	}
	// Define the proforma permutators used to create the actual permutators to use.
	proFormPermutators = []*Permutator{
		new(Permutator).New([]int{43, 57, 73, 83}, []byte{
			207, 252, 142, 205, 239, 35, 230, 62, 69, 94, 166, 89, 184, 81, 144, 120,
			27, 167, 39, 224, 75, 243, 87, 99, 47, 105, 163, 123, 129, 225, 2, 242,
			65, 43, 12, 113, 30, 102, 240, 78, 137, 109, 112, 210, 214, 118, 106, 22,
			232, 181, 164, 255, 70, 198, 160, 44, 231, 20, 228, 53, 85, 238, 178, 133,
			95, 194, 245, 234, 13, 147, 134, 25, 244, 91, 176, 38, 46, 1, 217, 249,
			250, 52, 182, 73, 206, 140, 216, 145, 60, 218, 213, 8, 151, 101, 156, 5,
			241, 67, 49, 42, 212, 180, 92, 21, 16, 130, 128, 126, 98, 199, 162, 188,
			117, 191, 66, 84, 57, 208, 158, 247, 41, 131, 227, 155, 61, 165, 253, 51,
			119, 103, 179, 93, 122, 83, 183, 116, 79, 222, 50, 59, 80, 110, 186, 141,
			90, 152, 127, 107, 54, 71, 185, 161, 169, 34, 148, 146, 157, 138, 24, 237,
			76, 196, 192, 251, 189, 201, 219, 86, 68, 37, 33, 82, 11, 170, 246, 72,
			229, 28, 32, 132, 23, 197, 108, 236, 220, 17, 150, 190, 171, 96, 26, 204,
			209, 31, 211, 4, 14, 136, 195, 45, 172, 111, 154, 36, 149, 226, 202, 187,
			193, 223, 139, 175, 124, 9, 3, 58, 125, 88, 15, 6, 121, 235, 221, 200,
			114, 254, 135, 168, 7, 29, 159, 48, 40, 115, 143, 203, 215, 77, 18, 55,
			56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10}),
		new(Permutator).New([]int{49, 51, 73, 83}, []byte{
			248, 250, 32, 91, 122, 166, 115, 61, 178, 111, 37, 35, 82, 167, 157, 66,
			22, 65, 47, 1, 195, 182, 190, 73, 19, 218, 237, 76, 140, 155, 18, 11,
			30, 207, 105, 49, 230, 83, 10, 251, 52, 136, 99, 212, 108, 154, 113, 41,
			185, 44, 102, 226, 135, 165, 94, 27, 6, 177, 162, 161, 209, 200, 33, 23,
			197, 120, 71, 249, 125, 244, 217, 38, 0, 128, 95, 80, 214, 254, 163, 203,
			180, 137, 100, 235, 16, 58, 78, 173, 3, 118, 148, 191, 15, 7, 149, 219,
			39, 129, 75, 158, 224, 92, 147, 144, 236, 60, 29, 9, 252, 51, 139, 97,
			43, 87, 193, 222, 85, 223, 127, 153, 192, 13, 143, 70, 151, 123, 211, 72,
			93, 194, 229, 42, 17, 146, 196, 107, 215, 112, 231, 21, 124, 86, 132, 238,
			26, 189, 98, 172, 201, 175, 188, 88, 114, 5, 25, 64, 103, 246, 45, 57,
			109, 63, 81, 62, 204, 106, 179, 199, 116, 141, 186, 121, 84, 210, 79, 156,
			216, 14, 253, 233, 46, 55, 138, 34, 74, 20, 245, 89, 198, 133, 239, 142,
			234, 24, 176, 213, 169, 241, 90, 232, 28, 240, 183, 227, 56, 247, 160, 152,
			202, 4, 159, 104, 187, 31, 174, 48, 168, 67, 40, 50, 134, 228, 181, 170,
			225, 126, 54, 36, 220, 208, 150, 117, 255, 221, 101, 69, 77, 110, 243, 206,
			130, 59, 205, 242, 184, 164, 131, 12, 2, 119, 96, 171, 53, 68, 8, 145}),
	}
	counter *Counter = new(Counter)
	jc1Key  *jc1.UberJc1
)

// Tnt2Engine type defines the encryption/decryption machine (rotors and
// permutators).
type Tnt2Engine struct {
	engineType    string // "E)ncrypt" or "D)ecrypt"
	engine        []Crypter
	left, right   chan CipherBlock
	cntrKey       string
	maximalStates *big.Int
}

// Left is a getter that returns the input channel for the Tnt2Engine.
func (e *Tnt2Engine) Left() chan CipherBlock {
	return e.left
}

// Right is a getter that returns the output channel for the Tnt2Engine.
func (e *Tnt2Engine) Right() chan CipherBlock {
	return e.right
}

// CounterKey is a getter that returns the SHAKE256 hash for the secret key.
// This is used to set/retrieve that next block to use in encrypting data
// from the file used to save the next block to use..
func (e *Tnt2Engine) CounterKey() string {
	return e.cntrKey
}

// Index is a getter that returns the block number of the next block to be
// encrypted.
func (e *Tnt2Engine) Index() (cntr *big.Int) {
	if len(e.engine) != 0 {
		machine := e.engine[len(e.engine)-1]
		switch machine.(type) {
		default:
			cntr = BigZero
		case *Counter:
			cntr = machine.Index()
		}
	}

	return
}

// SetIndex is a setter function that sets the rotors and permutators so that
// the TntEngine will be ready start encrypting/decrypting at the correct block.
func (e *Tnt2Engine) SetIndex(iCnt *big.Int) {
	for _, machine := range e.engine {
		machine.SetIndex(new(big.Int).Set(iCnt))
	}
}

// SetEngineType is a setter function that sets the engineType [D)ecrypt or E)crypt]
// of the Tnt2Engine.
func (e *Tnt2Engine) SetEngineType(engineType string) {
	switch string(strings.TrimSpace(engineType)[0]) {
	case "d", "D":
		e.engineType = "D"
	case "e", "E":
		e.engineType = "E"
	default:
		log.Fatalf("Missing or incorrect Tnt2Engine engineType: [%s]", engineType)
	}
}

// Engine is a getter function that returns a slice containing the rotors and
// permutators for the Tnt2Engine.
func (e *Tnt2Engine) Engine() []Crypter {
	return e.engine
}

// EngineType is a getter function that returns the engine type of the TntMachine.
func (e *Tnt2Engine) EngineType() string {
	return e.engineType
}

// MaximalStates is a getter function that returns maximum number of states that the
// engine can be in before repeating.
func (e *Tnt2Engine) MaximalStates() *big.Int {
	return e.maximalStates
}

func countLayoutType(cType rune) int {
	var cnt int
	for _, v := range EngineLayout {
		if cType == v {
			cnt++
		}
	}
	return cnt
}

// Init will initialize the Tnt2Engine generating new Rotors and Permutators using
// the proForma rotors and permutators in complex way, updating the rotors and
// permutators in place.
func (e *Tnt2Engine) Init(secret []byte, proFormaFileName string) {
	rCnt := countLayoutType('r')
	pCnt := countLayoutType('p')
	jc1Key = new(jc1.UberJc1).New(secret)
	// Create an ecryption machine based on the proForma rotors and permutators.
	var pfmReader io.Reader = nil
	if len(proFormaFileName) != 0 {
		in, err := os.Open(proFormaFileName)
		checkFatal(err)
		defer in.Close()
		pfmReader = bufio.NewReader(in)
	}
	e.engine = *createProFormaMachine(pfmReader)
	e.left, e.right = createEncryptMachine(e.engine...)
	// Get a SHA-3 hash of the encryption key.  This is used as a key to store
	// the count of blocks already encrypted to use as a starting point for the
	// encryption of the next message.
	k := make([]byte, 1024)
	blk := make(CipherBlock, CipherBlockBytes)
	h := blk[:]
	d := sha3.NewShake256()
	d.Write(jc1Key.XORKeyStream(k))
	d.Read(h)
	// Encrypt the hash starting at block 1234567890 (no good reason for this number)
	// to make it specific to the proForma machine used.
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	e.SetIndex(iCnt)
	e.left <- blk
	blk = <-e.right
	e.cntrKey = hex.EncodeToString(blk[:])
	e.SetIndex(BigZero)
	// Create a random number function [func(max int) int] that uses psudo-
	// random data generated the proforma encryption machine.
	random := new(Rand).New(e)
	// Get the last _rCnt_ rotor sizes (to maximize the period of the generator).
	rotorSizesIndex = len(RotorSizes) - 1
	// Create a permutaion of cycle sizes indices to allow picking the cycle
	// sizes in a random order based on the key.
	cycleSizes = random.Perm(len(CycleSizes))[0:2]
	cycleSizesIndex = 0
	// get the number of rotors and permutators
	rIdx := 0
	pIdx := 0
	// Update the rotors and permutators in a very non-linear fashion.
	rotors := make([]Crypter, rCnt)
	permutators := make([]Crypter, pCnt)
	e.maximalStates = new(big.Int).Set(BigOne)
	for _, machine := range e.engine {
		machine.Update(random)
		switch v := machine.(type) {
		default:
			fmt.Fprintf(os.Stderr, "Unknown machine: %v\n", v)
		case *Rotor:
			e.maximalStates = e.maximalStates.Mul(e.maximalStates, big.NewInt(int64(machine.(*Rotor).Size)))
			rotors[rIdx] = machine
			rIdx++
		case *Permutator:
			e.maximalStates = e.maximalStates.Mul(e.maximalStates, big.NewInt(int64(machine.(*Permutator).MaximalStates)))
			permutators[pIdx] = machine
			pIdx++
		case *Counter:
			machine.SetIndex(BigZero)
		}
	}
	// Update any addition rotors.
	for rIdx < rCnt {
		rotors[rIdx] = new(Rotor)
		rotors[rIdx].Update(random)
		rIdx++
	}
	// Update any additional permutators
	for pIdx < pCnt {
		permutators[pIdx] = new(Permutator)
		permutators[pIdx].Update(random)
		pIdx++
	}
	// Now that we have created the new rotors and permutators from the proform
	// machine, populate the Tnt2Engine with them using a random order for the
	// rotors and the permutators (without changing the layout in engineLayout).
	newMachine := make([]Crypter, len(EngineLayout)+1)
	rotorOrder := random.Perm(rCnt)
	permOrder := random.Perm(pCnt)
	rIdx, pIdx = 0, 0
	for idx, val := range EngineLayout {
		if val == 'r' {
			newMachine[idx] = rotors[rotorOrder[rIdx]]
			rIdx++
		} else if val == 'p' {
			newMachine[idx] = permutators[permOrder[pIdx]]
			pIdx++
		}
	}
	counter.SetIndex(BigZero)
	newMachine[len(newMachine)-1] = counter
	e.engine = newMachine
}

// BuildCiperMachine will create a "machine" to encrypt or decrypt data sent to the
// left channel and outputed on the right channel for the Tnt2Engine.  The engineType
// determines weither a encrypt machine or a decrypt machine will be created.
func (e *Tnt2Engine) BuildCipherMachine() {
	switch e.engineType {
	case "D":
		e.left, e.right = createDecryptMachine(e.engine...)
	case "E":
		e.left, e.right = createEncryptMachine(e.engine...)
	default:
		log.Fatalf("Missing or incorrect Tnt2Engine engineType: [%s]", e.engineType)

	}
}

// CloseCipherMachine will close down the cipher machine by exiting the go function
// that performs the encryption/decryption using the individual rotors/permutators.
// This is done by passing the CipherMachine a CypherBlock with a length of zero (0).
func (e *Tnt2Engine) CloseCipherMachine() {
	blk := new(CipherBlock)
	e.Left() <- *blk
	<-e.Right()
}

// createProFormaMachine initializes the proForma machine used to create the
// TNT2 encryption machine.  If the machineFileName is not empty then the
// proForma machine is loaded from that file, else the hardcoded rotors and
// permutators are used to initialize the proForma machine.
func createProFormaMachine(pfmReader io.Reader) *[]Crypter {
	newMachine := make([]Crypter, 8)
	// getCyclesSizes will extract the lengths of the given permutation cycles
	// and return them as a slice of ints.
	getCycleSizes := func(cycles []Cycle) []int {
		cycleSizes := make([]int, len(cycles))
		for i, v := range cycles {
			cycleSizes[i] = v.Length
		}
		return cycleSizes
	}
	if pfmReader == nil {
		// Create the proforma encryption machine.  The layout of the machine is:
		// 		rotor, rotor, permutator, rotor, rotor, permutator, rotor, rotor
		// ----------------------------------------------------------------------
		// The ProFormaMachine is created by making a copy of the hardcoded proforma
		// rotors and permutators.  This resolves an issue running tests where
		// Tnt2Engine.Init() is called multiple times which caused a failure on
		// the second call.
		newMachine[0] = new(Rotor).New(proFormaRotors[0].Size, proFormaRotors[0].Start, proFormaRotors[0].Step, append([]byte(nil), proFormaRotors[0].Rotor...))
		newMachine[1] = new(Rotor).New(proFormaRotors[1].Size, proFormaRotors[1].Start, proFormaRotors[1].Step, append([]byte(nil), proFormaRotors[1].Rotor...))
		newMachine[2] = new(Permutator).New(getCycleSizes(proFormPermutators[0].Cycles), append([]byte(nil), proFormPermutators[0].Randp...))
		newMachine[3] = new(Rotor).New(proFormaRotors[2].Size, proFormaRotors[2].Start, proFormaRotors[2].Step, append([]byte(nil), proFormaRotors[2].Rotor...))
		newMachine[4] = new(Rotor).New(proFormaRotors[3].Size, proFormaRotors[3].Start, proFormaRotors[3].Step, append([]byte(nil), proFormaRotors[3].Rotor...))
		newMachine[5] = new(Permutator).New(getCycleSizes(proFormPermutators[1].Cycles), append([]byte(nil), proFormPermutators[1].Randp...))
		newMachine[6] = new(Rotor).New(proFormaRotors[4].Size, proFormaRotors[4].Start, proFormaRotors[4].Step, append([]byte(nil), proFormaRotors[4].Rotor...))
		newMachine[7] = new(Rotor).New(proFormaRotors[5].Size, proFormaRotors[5].Start, proFormaRotors[5].Step, append([]byte(nil), proFormaRotors[5].Rotor...))
	} else {
		jDecoder := json.NewDecoder(pfmReader)
		// Create the proforma encryption machine from the given proforma machine file.
		// The layout of the machine is:
		// 		rotor, rotor, permutator, rotor, rotor, permutator, rotor, rotor
		newMachine[0] = new(Rotor)
		newMachine[1] = new(Rotor)
		newMachine[2] = new(Permutator)
		newMachine[3] = new(Rotor)
		newMachine[4] = new(Rotor)
		newMachine[5] = new(Permutator)
		newMachine[6] = new(Rotor)
		newMachine[7] = new(Rotor)

		for _, machine := range newMachine {
			switch v := machine.(type) {
			default:
				fmt.Fprintf(os.Stderr, "Unknown machine: %v\n", v)
			case *Rotor:
				err := jDecoder.Decode(&machine)
				checkFatal(err)
			case *Permutator:
				err := jDecoder.Decode(&machine)
				checkFatal(err)
			}
		}
	}

	return &newMachine
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
