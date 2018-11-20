package main 
import "fmt"
func main() {
	temChave := false
	bateuNaPorta := true
	if temChave == true {
		fmt.Println("Porta aberta com a chave!")
	} else if bateuNaPorta == true {
		fmt.Println("Porta aberta ao bater nela!")
	} else {
		fmt.Println("Vá buscar a chave!")
	}
}
