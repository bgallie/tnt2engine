# tnt2engine

The encryption engine used by the TNT2 - Infinite Key Encryption System.

This project was created to allow the original tntengine project to be follow more closely the TNT Infinite Key Encryption System decribed in the article:

>__An Infinite Key Encryption System.__    
[Dr. Dobbs Journal Volume 9, Number 94, 1984](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up)

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
