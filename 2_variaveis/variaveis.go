//

package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // Declaramos o tipo da variavel
	variavel2 := "Variável 2"           // Deixamos o tipo implícito
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

}
