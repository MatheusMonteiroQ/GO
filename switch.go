package main 
import "fmt"
func main() {
	valorProduto := 5
	switch valorProduto {
		case 10:
			fmt.Println("Caros")
		case 5:	
			fmt.Println("Populares")
		default:
			fmt.Println("Sem classificação")
	}
	//OUTPUT: Populares
}