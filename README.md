# tnt2engine

The encryption engine used by the TNT2 - Infinite Key Encryption System.

This project was created to allow the original tntengine project to be follow more closely the TNT Infinite Key Encryption System decribed in the article:

>__An Infinite Key Encryption System.__    
[Dr. Dobbs Journal Volume 9, Number 94, 1984](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up)

___v1.6.1___   
This release removes the `CyclePermutations` array, which is no longer needed.

___v1.6.0___   
- The code no longer randomizes the order of the machine, but instead leaves the placement of rotors and permutators the same as given in engineLayout (defaults to `rrprrprr`) but randomizes the order of the rotors and permutators.  This prevents things like `pprrrrrr` from occuring.
- Changed how the rotor sizes are selected when updating the rotors.  It now selects the (number of rotors) largest sizes to maximize the number of bytes that can be encrypted before the psudo-random data repeats.  For the default `engineLayout`, 8.152525 * 10<sup>37</sup> bytes that can be encrypted before the psudo-random data repeats.  The **tnt2engine** (currently) supports up to 40 rotors, which would increase the peroid to 3.714144 * 10<sup>170</sup>.
- Changed the `sliceRotor()` and `getRotorBlock()` to operate at the block level instead of at the bit level.  This gives a good speed increase of those functions.  The `sliceRotor()` function is called 6 times when initializing the tntengine and `getRotorBlock()` many time when initializing the tnt2engine and is called 6 times for each 32 byte block in the file to be encrypted.   Here are the benchmark results:
```
goos: linux
goarch: amd64
cpu: AMD Ryzen 9 3900X 12-Core Processor
=============================================================
   OrigSliceRotor-8    375.2  ns/op     0 B/op    0 allocs/op
    NewSliceRotor-8     65.04 ns/op     0 B/op    0 allocs/op
OrigGetRotorBlock-8    717.9  ns/op    32 B/op    1 allocs/op
 NewGetRotorBlock-8     55.55 ns/op    32 B/op    1 allocs/op
```
___v1.5.1___   
This release replaces the depreciated `NewUberJc1()` call with `jc1.UberJc1.New()` call.

___v1.5.0___   
This release:
- Changes how a short last block is handled.  This enables files to decrypted without needing to know the number of encrypted bytes when decrypting.  **NOTE:** *Files encrypted with earlier versions will not decrypt correctly with this version.*
- Various code cleanup and refactoring.

___v1.4.2___  
This release updated/corrected comments and fixed the following:
- Removed unused constants.
- Removed the now unused `CypherBlock.Marshall()` and `CypherBlock.Unmarshall()` methods.
- Fixed bug when creating the proforma machine from a proforma data file that caused a panic.
- Fixed bug where the amount to step the rotor could be set to zero or Rotor.Size, causing the rotor to reuse the same bytes all the time.
- Corrected the calculation of the adjustment size for the rotor when updating an existing rotor.
- Set the `rotor.Current` to `rotor.Start` instead of `rotor.Size` when updating an existing rotor.

___v1.4.1___  
Change `tnt2engine.Rand.Read()`  to only initialize the `Rand.blk` with key used to store the next block count encrypted by UberJC1 on the first call to `Rand.Read()`._  Subsequent encryption of `Rand.blk` uses the results of the previous encryption of `Rand.blk`.

___v1.4.0___  
1. This release of **tnt2engine** fixes an issue where `permutator.SetIndex()` did not update the current value of the cycles.
2. The JC1 package is now only used to generate the key to store the next block count,
3. The `jc1Key` is no longer stored as part the the TntEngine..
4. All test in _rand_test.go_ will now execute correctly when running all tests at once or when running each test individually.
5. `Rand.NewRand()` is now depreciated.  The `New()` method of `*Rand` instance replaces it.
6. Added new `Update(*Rand)` method to the `Cryptor` interface.
7. Added `rotor.New()` (replacing `NewRotor()`) and `permutator.New()` (replacing `NewPermutator()`).
8. Implemented `rotor.Update()`, `permutator.Update()`, and `counter.Update()`.

___v1.3.0___  
1. This release of **tnt2engine** fixes additional issues with how `*big.Int` variables are assigned values.
2. Added tests for `tnt2engine.Rand`.

___v1.2.0___  
This release of **tnt2engine** changes the parameter of `createProFormaMachine(proFormaFilename string)` to `createProFormaMachine(pfmMachine io.Reader)`.
