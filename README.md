# sort
quick sort

# Installation
 
Install quick sort with go get command.
 
```bash
go get github.com/y4asse/sort
```
 
# Usage
 
Please create go code named "main.go".

```go:main.go
package main

import (
	"fmt"

	"github.com/y4asse/sort"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	sortedArr := sort.QuickSort(arr)
	fmt.Println(sortedArr)
}
```
 
Run "go run main.go"
 
```bash
go run main.go
```
 
# Author
 
* Hayase Taiki
* Twitter : https://twitter.com/y4asse
 
# License
 
this library is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 

Thank you!