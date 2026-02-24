/*panic:
- Interrompe o fluxo normal do programa
- Encerra a execução se não for tratado
- Usado para erros críticos e irrecuperáveis
*/

/*recover:
- Recupera um panic
- Só funciona dentro de uma função defer
- Evita que o programa quebre
*/

package main

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println((x))
	}()
	panic("Pânico")
}
