
### Golang tutorial
```bash
go run script.go
go build script.go

go get github.com/v-braun/go2p

# run UNIT TEST
go test tests/hasher_test.go

go env GOPATH 

home/user/go # base envioenment
home/user/go/src # where project lives
 
go install # exe are in ../../bin folder

go.sum file (yarn.lock)

go mod init github.com/user/golang-bot 
# crea go.mod file che e' simile al requirements.txt

require (
	rsc.io/quote v1.5.2
)
```

go.mod file (requirements.txt di python)

``` go
require (
	rsc.io/quote v1.5.2
)
```

imports
```
import (
	"fmt"
	"github.com/Edmartt/go-password-hasher/hasher"
	)
func main(){}
```

``` go
import "math" 
import "fmt" # input / output

import (
	"fmt"
	"math"

)

func main(){ ... }  //all here

fmt.Println("Hello world")

// ASSEGNAZIONE
var x int
var y string
var z bool
x := 5
y := 'test'
z := 0
var a int = 5
var sum int = x + a 

// IF
if x > 6 {
fmt.Println(sum)
} else if x < 0 {

} else {
}

// ARRAY
var a [5] int
a := [5]int{1,2,3,4,5} // ARRAY (immutabile di lunghezza)
a[1] = 8 // ARRAY ASSIGNMENT

b := [ ]int{1,2,3,4,5} // SLICE (tipo lista di python)
b = append(a,13)

// DIZIONARO
dizionario := make(map[string]int)
dizionario['a']=1
dizionario['b']=2
dizionario['c']=3
