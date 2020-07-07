# tutorial
golang examples.

## 1. go mod
 gomod repository has provided a test env that you done `go mod init` (generate `go.mod` file) under your local project repository and `import` the package. Just remember turning the `GO111MODULE` of golang env variable ON (on|off|auto, no upper case on every letter).
 Tip: If you turn on the `Go111MODULE`, your coding of `import` packages in the project would find package from `pkg` repository not usually  `src` repository you known that means your project would report error like **_can not find your package from `GOROOT`_** which the `go.mod` was not included.

  go.mod:
  ```
  module test/gomod
  
  go 1.14
  
  require github.com/chenqgp/tutorial [v0.0.0-20200702084203-bbe18f8b3a9b or tags]
  ```
## 2. Array Slice and Map
###### Array
Array is one of `data type` in GO. It has a **centainly known length** and each element was a **same type** in array, when you delivered a parameter from array that it always **passed the value** of the array not its pointer. The values in array have continually stored by one next to one in memory, so the operation of the array is very fast.
###### Slice
Slice is a `data structures`. It's easy to manage sets of the datas, Slice was different from Array that Slice has a **modifiable length**. It's completely hard to tell clearly that Slice has numerous features, thouth, but We listed many examples in arraySliceMap folder. [_Just check it out!_](https://github.com/chenqgp/tutorial/blob/master/arraySliceMap/slice.go#L51) when you pass Slice as paramiter to a function or method, this situation consumes less memory than passing Array. The pointer, length, capacity of a Slice totally deals with 24 bytes(every one of them takes 8 byte) in your memory, it just delivered a copy of the Slice rather than underlying datas.
###### Map 
Map is one of `data structures`. It used to store plenty of unordered paire key-value. In a sense Map similar to slice, so we would list something special cases about Map， [_see!_](https://github.com/chenqgp/tutorial/blob/master/arraySliceMap/map.go#L33)
