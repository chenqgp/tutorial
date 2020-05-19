# tutorial
golang examples.

### 1. go mod
* gomod repository has provided a test env that you done `go mod init` (generate `go.mod` file) under your local project repository and `import` the package. Just remember turning the `GO111MODULE` of golang env variable ON (on|off|auto, no upper case on every letter).
Tip: If you turn on the `Go111MODULE`, your coding of `import` packages in the project would find package from `pkg` repository not usually  `src` repository you known that means your project would report error like **_can not find your package from `GOROOT`_** which the `go.mod` was not included.

  go.mod:
  ```
  module test/gomod
  
  go 1.14
  
  require github.com/chenqgp/tutorial [v1.0.1-0.20200518143801-d085fb9b89a5 or tags]
  ```
### 2. Array Slice and Map
* Array
It has a **centainly known length** and each element was a **same type** in array, when you delivered a parameter from array that it always **passed the value** of the array not its pointer.
