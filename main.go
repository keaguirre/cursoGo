// package main

// import "fmt"

// func main() {
// // fmt.Println("Hello world!")
// // name := "kevin" //no usamos ningun tipo de declaracion
// // fmt.Println("name: " + name)
// var name string = "kevin"
// fmt.Printf("Hola %v Bienvenido", name)
// }

// package main

// import "fmt"

// func main() {
// 	var fruit string = "grapes"
// 	if fruit == "apple" {
// 		fmt.Println("The fruit is apple")
// 	} else if fruit == "banana" {
// 		fmt.Println("The fruit is banana")
// 	} else {
// 		fmt.Println("The fruit is not grapes")
// 	}
// }


package main
import "fmt"
func main(){
    for i:=1; i <= 5; i++{
        fmt.Println(i*i)
    }
}