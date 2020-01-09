# go-converter
package main

import "fmt"

func main() {
    fmt.Print("Конвертируй футы в метры: ")
    var input float64
    fmt.Scanf("%f", &input)
 	 const fut = 0.3048
    output := input*fut

    fmt.Println(output)    
}