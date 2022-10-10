# Golang tutorial


# activate go env
source /etc/profile

# Shell

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

go mod init github.com/user/golang-bot 
# crea go.mod file che e' simile al requirements.txt

require (
	rsc.io/quote v1.5.2
)
```


# go.mod
crea il file per il tracking delle dipendenze
``` bash
go mod init
```
go.mod file (requirements.txt di python)

``` go
require (
	rsc.io/quote v1.5.2
)
```

# go.sum
go.sum file (yarn.lock)
```
 
```

# imports
``` go
import (
	"fmt"
	"github.com/Edmartt/go-password-hasher/hasher"
	)

func main(){

}
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


```

# variabili

esempio
``` go
var numero int
numero := 5
var numero int = 5
```

tipi
``` go
var  //variabile
const //costante
```

categorie
``` go
int
string
bool
int8 int16 int32 int64
uint8 uint16 uint32 uint64 uintptr
byte // uint8
rune // int32
float32 float64
complex64 complex128
struct
```
assegnazione
```
x := 5
y := 'test'
z := 0
var a int = 5
var sum int = x + a
```

# conversione
``` go
mean = float32(10)/float32('5.0')
```

# if 
``` go
if x > 6 {
fmt.Println(sum)
} else if x < 0 {
	code
} else {
	code
}
```

# array
``` go
var a [5] int
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
float32 salary = balance[9] // ACCESSO AGLI ARRAY


a := [5]int{1,2,3,4,5} // ARRAY (immutabile di lunghezza)
a[1] = 8 // ARRAY ASSIGNMENT

```

# Slice
``` go
b := [ ]int{1,2,3,4,5} // SLICE (tipo lista di python)
b = append(a,13)

var numbers []int /* a slice of unspecified size */
				  /* numbers == []int{0,0,0,0,0}*/

numbers = make([]int,5,5) /* a slice of length 5 and capacity 
```

Una fetta è un'astrazione su un array. In realtà utilizza gli array come struttura sottostante. La funzione **len()** restituisce gli elementi presenti nella fetta dove la funzione **cap()** restituisce la capacità della fetta (cioè quanti elementi può contenere).

slice **[lower-bound:upper-bound]**
```
numbers := []int{0,1,2,3,4,5,6,7,8}
numbers[1:4]
```

# Mappe (Dizionari)

#### crea dizionario
``` go
var mappa map[string]string
```
   
``` go
mappa = make(map[string]string)
mappa := make(map[string]int)
```

 
#### inserisci valori
``` Go
mappa['a']=1
mappa['b']=2
mappa['c']=3

fmt.Println(dizionario['c'])

delete(dizionario['c'])

```

# For loop
``` go
// FOR LOOP
for i:=0; i<5; i++ {
	fmt.Println(i)
}

```

# While loop
``` go
// WHILE LOOP
i :=0
for i < 5 {
	fmt.Println(i)
	i++
}

for index, value := range array {
	fmt.Println(index, value)
}

for key, value := range dizionario {
	fmt.Println(key, value)
}

```

# Funzioni
``` go
func vals() (int, int) {
    return 3, 7
}

func sum(x int,y int) int {
	return x + y
}

esult = sum(2,3)
a, b := vals()

```

# Error Handling
``` go
type error interface {
   Error() string
}
```

``` go
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("ERROR x<0")
	}
	return math.Sqrt(x), nil
}



func main(){
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err) } 
	else{
		return sqrt(result)
		}
}

```

# Strutture (classi)
``` go 
type Person struct {
	name string
	age int
	height float64
}
func main(){
	person := Person{name:'Jake', age:14, height:1.8}
	person := Person{'Jake',14,1.8}
	
	fmt.Println(person)
	fmt.Println(person.name) //accesso agli attributi
	
}



```

metodi
```go 
import ("strconv")
type Person struct {
	name string
	age int
	height float64
}

// metodo
func (p Person) salute() string{
	strconv.Itoa('ciao ' + p.name)
	//  or return(p.name)
}

```

Point reciver 
riceve valori direttamente al puntatore attraverso una classe
``` go

func (p *Person) aumenta_eta() {
	p.age++
}

func (p *Person) mostra_eta() {
	strconv.Itoa(p.age)
}


```

# Puntatori (indirizzi di memoria)
``` go
i := 1
&i := 2
&i += &i+1 


a :=5
b := &a

*b = 10

```

``` go
i := 1
zeroptr(&i) // 0x42131100
zeroval / zeroptr
```
`zeroval` doesn’t change the `i` in `main`, but `zeroptr` does because it has a reference to the memory address for that variable.

# Funzioni Variadiche (...)
``` go
// VARIADIC FUNCTION
func sum2(nums ...int) {
    for _, num := range nums {
        total += num
    }
    return total
}
sum2(1,2,3)
nums := []int{1,2,3}
sum2(nums...)

```

# Funzioni lambda
```
// FUNZIONI ANONIME (closures / lambda di python)
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

```

# Ricorsione
```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
var fib func(n int) int
```

``` go
fib = func(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2) }
```

# Print
```    
// PRINT OUT
fmt.Println("go" + "lang")
fmt.Println("1+1 =", 1+1)
fmt.Println("7.0/3.0 =", 7.0/3.0)
```


# Booleani
```
// BOOLEANS
true && false   // -> false
treue || false  // -> true
! true          // -> false

```

# Switch
``` go
i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}
```

``` go
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("It's before noon")
default:
	fmt.Println("It's after noon")
}
```

# Interfacce

signatures che rappresentano un set di metodi per strutture 

``` go

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectancle struct {
	x, y float64
}

func (c Circle) area() float64 {
	return math.Pi* c.radius * c.radius
}

func (r Rectancle) area() float64 {
	return math.Pi* r.x * r.y
}


func getArea(s Shape) float64 {
	return s.area()
}


func main() {

	circle = Circle{radius:5}
	rectangle = Rectangle{x:5,y:4}
	
	fmt.Println('area is %f',  getArea(circle) )
	fmt.Println('area is %f',  getArea(rectangle) )
	
}


```

``` go
// esempio con switch
whatAmI := func(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

whatAmI(true)
whatAmI(1)
whatAmI("hey")
```

# Range
``` go
for i, num := range nums {
	if num == 3 {
		fmt.Println("index:", i)
	}

for key, val := range dizionario {}
for key := range dizionario {}
for i, c := range "go" {}

```


``` go
/* print map using keys*/
for country := range countryCapitalMap {
      fmt.Println("Capitalof",country,"is",countryCapitalMap[country])
   }
   ```

``` go
/* print map using key-value*/
for country,capital := range countryCapitalMap {
  fmt.Println("Capital of",country,"is",capital)
}
```

# Built in function
``` go
len(string/array)
cap(array) //capacita di memoria
```

# Web
se vai nel browser al localhost:3000 vedi Hello world per il rendering della pagina
``` go
import ("fmt"
		"net/http"

)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"about")
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/about",index)
	http.HandleFunc("/",index)
	fmt.Println("Connected to server")
	http.ListenAndServe(":3000",nil)
}


```
