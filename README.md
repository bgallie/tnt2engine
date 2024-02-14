# tnt2engine

The encryption engine used by the TNT2 - Infinite Key Encryption System.

This project was created to allow the original tntengine project to be follow more closely the TNT Infinite Key Encryption System decribed in the article:

>__An Infinite Key Encryption System.__    
[Dr. Dobbs Journal Volume 9, Number 94, 1984](https://archive.org/details/1984-08-dr-dobbs-journal/page/44/mode/2up)

The version numbering for this project (v1.4.2) matches the version number of the original __tntengine__ project.  The list of changes in the original __tntengine__ rolled into this version are:

___v.1.4.2___  
This release updated/corrected comments and fixed the following:
- Removed unused constants.
- Removed the now unused _CypherBlock.Marshall_ and _CypherBlock.Unmarshall_ methods.
- Fixed bug when creating the proforma machine from a proforma data file that caused a panic.
- Fixed bug where the amount to step the rotor could be set to zero or Rotor.Size, causing the rotor to reuse the same bytes all the time.
- Corrected the calculation of the adjustment size for the rotor when updating an existing rotor.
- Set the rotor.Current to _rotor.Start_ instead of _rotor.Size_ when updating an existing rotor.

___v1.4.1___  
Change **tntengine.Rand.Read**  to only initialize the _Rand.blk_ with key used to store the next block count encrypted by UberJC1 on the first call to _Rand.Read()._  Subsequent encryption of _Rand.blk_ uses the results of the previous encryption of _Rand.blk_.

___v1.4.0___  
1. This release of **tntengine** fixes an issue where permutator.SetIndex() did not update the current value of the cycles.
2. TJC1 package is now only used to generate the key to store the next block count,
3. The jc1Key is no longer stored as part the the TntEngine..
4. Change **tntengine.Rand.Read**  to only initialize the _Rand.blk_ with key used to store the next block count on the first call to _Rand.Read()._  Subsequent encryption of _Rand.blk_ uses the results of the previous encryption of _Rand.blk_.
5. All test in _rand_test.go_ will now execute correctly when running all tests at once or when running each test individually.
6. _Rand.NewRand()_ is now depreciated.  The _New()_ method of *Rand instance replaces it.
7. Added new _Update(*Rand)_ method to the **Cryptor** interface.
8. Added _rotor.New()_ (replacing _NewRotor()_) and _permutator.New()_ (replacing _NewPermutator()_).
9. Implemented _rotor.Update()_, _permutator.Update()_, and _counter.Update()_.

___v1.3.0___  
1. This release of **tntengine** fixes additional issues with how _*big.Int_ variables are assigned values.
2. Change **tntengine.Rand.Read**  to only initialize the _Rand.blk_ with UberJC1 on the first call to _Rand.Read()._  Subsequent encryption of _Rand.blk_ uses the results of the previous encryption of _Rand.blk_.
3. Added tests for **tntengine.Rand**.

___v1.2.0___  
This release of **tntengine** changes the parameter of `createProFormaMachine(proFormaFilename string)` to `createProFormaMachine(pfmMachine io.Reader)`.
