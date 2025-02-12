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
            