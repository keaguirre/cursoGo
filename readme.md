Install Go
1. go.dev/doc/install
2. and check with `go version` or `go help`
3. go run [file]

# Data types
- String
- Integer
- Float
- Boolean
- Array
- Slice: Dato flexible como un array pero tambien proporciona mejor control sobre la asignacion de memoria
- Map: Coleccion de pares key-value: "x" -> 30, 1 -> 100, "key" -> "value"

- Go es estaticamente tipado, pero podemos usarlo tambien esperando a que el compilador infiera los tipos
- un ejemplo de esto es usando un 'short' variable declaration, aqui esperamos que el compilador infiera el tipo de dato
```go
    name:= "kevin"
```
## Tipos de integer
- u integer = unsigned integer: lo que significa que solo podran asignarse solo numeros positivos incluyendo 0
- integer = signed integer: positivos y negativos 

| Tipo de datos |                               Memoria                              |
|:-------------:|:------------------------------------------------------------------:|
| uint8         | 8 bits o 1 byte                                                    |
| uint16        | 16 bits o 2 bytes                                                  |
| uint32        | 32 bits o 4 bytes                                                  |
| uint64        | 64 bits o 8 bytes                                                  |
| int8          | 8 bits o 1 byte                                                    |
| int16         | 16 bits o 2 bytes                                                  |
| int32         | 32 bits o 4 bytes                                                  |
| int64         | 64 bits o 8 bytes                                                  |
| int           | 4 bytes para máquinas de 32 bits, 8 bytes para máquinas de 64 bits |

## Tipos de Float
- float32: 32 bits or 4 bytes
- float64: 64 bits or 8 bytes

## Tipos de string
- string: 16 bytes

## Tipos de boolean
- bool: true or false solamente no puede ser 1 o 0 y ocupa 1 byte en memoria

## Tipos de Arrays & Slices

## Tipos de Maps

## (Syntax) Definicion de variables
```go
var [variableName] [dataType] = [value]
var s string = "Hello world"
var i int = 100
var b bool = false
var f float64 = 77.64

//Concatenar
var name string = "kevin"
var user string = "keaguirre"
fmt.Print("Hello ", name, "your username is", user)
//salto de linea es con \n dentro de un string

```
## Print Format specifiers
- %v formats the value in a default format
    - fmt.Printf("Hello %v bienvenido", name)
- %d formats decimal integers
```go
var grades int = 42
fmt.Printf("Marks: %d", grades)
//Result Marks: 42

var name string = "Kevin"
var i integer = 78
fmt.Printf("Hey, %v!, you've scored %d/100 in Physics", name, i)
//output Hey, Kevin! you've scored 78/100 in Physics"

```

| Verb          |                           Description                              |
|:-------------:|:------------------------------------------------------------------:|
| %v            | default format                                                     |
| %T            | type of the value                                                  |
| %d            | integers                                                           |
| %c            | character                                                          |
| %q            | quoted characters/string                                           |
| %s            | plain string                                                       |
| %t            | true or false                                                      |
| %f            | floating numbers                                                   |
| %.2f          | floating numbers up to 2 decimal places                            |

## Assign values
### Declaration + initializacion
```go
var s string
s = "string"
fmt.Println(s)
```

### Shorthand way
```go
var s, t string = "foo", "bar"
fmt.println(s)
fmt.println(t)

//another way
var(
    s string = "foo"
    i int = 5
)
fmt.println(s)
fmt.println(t)

s := "Hello World" //esta variable se puede volver a darle un nuevo valor que sea el mismo que infirio la primera vez
s = "Hello kevin"

```

## Blocks

```go
{
    //outer block
    {
        //inner block
    }
}
```
- inner blocks can access variables declared within outer blocks
- outer blocks cannot access variables declared within inner blocks
### Example
```go

func main(){//outer block
    city:= "London"
    {//inner block
        country := "UK"
        fmt.Println(country)
        fmt.Println(city)
    }
    //fmt.Println(country) este print no puede acceder a country
}
```
## Local vs Global Variables
- Local:    
    - Declared inside a funtion or a block
    - not accessible outside the function or the block
    - can also be declared inside looping and conditional statements

- Global:   
    - Declared outside of a function or a block
    - available throughout the lifetime of a program
    - declared at the top of the program outside all functions or blocks
    - can be accessed from any part of the program

## Zero values
- si estas declarando una variable pero no la estas inicializando con el value la variable recibe un valor por defecto, a este se la llama **Zero value**
-   bool = false
-   int = 0
-   float64 = 0.0
-   string = ""
-   pointers,functions, interfaces, maps = nil

## Usr input
```go
fmt.Scanf("%<format specifier>" (s)", Object_args) //format specifier see table, Object_args = &variable to store the input

func main(){
    var name string
    var is_dev bool
    fmt.Println("Enter your name:  & ru a dev?: (true/false)")
    fmt.Scanf("%s %t", &name, &is_dev)//%s string, %t boolean, specifiers should match the order of the variables
    fmt.Printf("Your name is %s and you are a developer: %t", name, is_dev)
}
```
- scanf returns 2 values, count of the number of items successfully written to the variables and an error if any occurs while reading the input
## Find the type of a variable
```go
fmt.Printf("Type: %T Value: %v\n", name, name)
//or
import (
    "fmt"
    "reflect"
)
func main(){
    name := "kevin"
    fmt.Println(reflect.TypeOf(name))
}
```

## Convertir tipos de datos (Type casting) //convertir un tipo de dato a otro no garantiza que el valor sea el mismo
```go
//integer to float
 var i int = 90
 var f float64 = float64(i)
 fmt.Printf("%.2f\n", f)

//float to integer
var f float64 = 45.89
var i int = int(f)
fmt.Printf("%v\n", i)
```
### strcpmv package
- strcpmv package provides functions to convert strings to other data types
```go
//Itoa() converts integer to string
import (
    "fmt"
    "strconv"
)
func main(){
    var i int = 42
    var s string = strconv.Itoa(i)//convertir integer a string
    fmt.Printf("%q", s)
}
//Atoi() converts string to integer
var s string = "200"
i, err := strconv.Atoi(s)
fmt.Printf("%v, %T \n", i, i)
fmt.Printf("%v, %T", err, err)
```
## Constants
```go
const [name] [data_type] = [value]

//untyped constant
//Constants are untyped unless they are explicitly typed
//allow more flexibility
```go
const age = 12
const h_name, h_age = "kevin", 27
```
## typed constant
- Constants are typed when you explicitly specify the type in the declaration
- flexibility that comes with untyped constants is lost
```go
const string name
name = "kevin" //error, el por defecto es untyped, por lo que arrojara error de declaracion, ya que pide que se declare el valor al momento de declarar la constante
const name := "kevin" //esto es incorrecto, ya que no se puede usar := para declarar constantes
```

## Operators
- Comparison operators
    - ==, !=, <, >, <=, >=
- Arithmetic operators
    - +, -, *, /, %, ++, --
- Assignment operators
    - =, +=, -=, *=, /=, %=
- Bitwise operators
    - &, |, ^, <<, >>
- Logical operators
    - &&, ||, !

## Bitwise ops
- Trabajan a nivel de bit
- & takes two number as operands and does **and** on every bit of two numbers
```binary
12 = 00001100
25 = 00011001
&   ---------
     00001000 = 8 (In decimal)
```
```go
package main

import "fmt"
func main(){
  var x, y int = 12, 25
  z := x & y
  fmt.Println(z)
}
//Result = 8
```

- | (OR) takes two number as operands and does **or** on every bit of two numbers
```binary
12 = 00001100
25 = 00011001
|   ---------
     00011101 = 29 (In decimal)
```


## If-else statements
```go
package main
import "fmt"
func main(){
    var fruit string = "grapes"
    if fruit == "apple"{
        fmt.Println("The fruit is apple")
    }else if fruit == "banana"{
        fmt.Println("The fruit is banana")
    }else{
        fmt.Println("The fruit is not grapes")
    }
}
``` 
## Switch statement
```go
package main
import "fmt"
func main(){
    var i int = 800
    switch i{
        case 10:
            fmt.Println("i is 10")
        case 100, 200:
            fmt.Println("i is either 100 or 200")
        default:
            fmt.Println("i is neither 0, 100 or 200")
    }
}
```
- **fallthrough** keyword is in switch-case to force the execution flow to fall through the successive case block, even if the case block doesn't match the expression and until a case block with fallthrough is found or the switch block ends 

## Switch with conditions
```go
package main
import "fmt"
func main(){
    var a, b int = 10, 20
    switch {
        case a+b == 30:
            fmt.Println("Equal to 30")
        case a+b <= 30:
            fmt.Println("less than or equal to 30")
        default:
            fmt.Println("Greater than 30")    
    }
}
```

## Loops
```go
for [initialization]; [condition]; [increment/decrement]{
    //the for loop initialization, condition, and increment/decrement are optional and can be omitted
}
//Example:
package main
import "fmt"
func main(){
    for i:=1; i <= 5; i++{
        fmt.Println(i*i)
    }
}
```

## Arrays
- Arrays en Go son una coleccion de elementos del mismo tipo y se almacenan en memoria contigua secuencialmente dependiendo del tipo de dato
- El tamaño de un array es fijo y no puede ser cambiado una vez que se declara
- El tamaño de un array es parte de su tipo, por lo que [5]int y [10]int son dos tipos diferentes
- Los arrays en Go son de tipo 0-based, lo que significa que el primer elemento del array tiene un indice de 0
- Para pasar un array a una funcion por referencia, se debe pasar un puntero al array
- Los arrays en Go son valores, no punteros, por lo que cuando se pasa un array a una funcion, se pasa una copia del array, no una referencia al array original
- Para pasar un array a una funcion por referencia, se debe pasar un puntero al array
- para ocupar un array se debe inicializar con un valor, si no se inicializa el valor por defecto es zero-value
- si se le pasan 3 elementos a un array de 5, los otros dos elementos seran zero-value

```go
var [arrayName] [size] [dataType]
//ex:
var grades [5] int
var fruits [3] string
```
### Example:
```go
package main
import "fmt"
func main(){

//array initialization
    var fruits [2] string = [2]string{"apple", "banana"}

// si se le pasa 5 elementos, el tamaño del array sera 5 en la short declaration
    marks := [3]int{10,20,30}
    names := [...] string{"kevin", "jose", "maria"}
// [...] esto es un ellipsis, que le dice al compilador que infiera el tamaño del array basado en el numero de elementos que se le pasan
//length of an array
    fmt.Println(len(fruits))
    fmt.Println(fruits[0])
    //podemos asignar un nuevo valor a una posicion del array con: fruits[1] = "grapes"
    
}
```
### Looping through an array
```go
package main
import "fmt"
func main(){
    for i := 0; i < len(marks); i++{
        fmt.Println(marks[i])
    }
    for index, element := range marks{
        fmt.Println(index,"res=>", element)
    }
    //range establece el scope de iteracion hasta la longitud del array y devuelve el indice y el valor del elemento en cada iteracion

}
```

### Multi-dimensional arrays
```go
package main
import "fmt"
func main(){
    arr :=[3][2] int {{2,4}, {4,16}, {8,64}}
    fmt.Println(arr[2][1]) //res => 64

}
```