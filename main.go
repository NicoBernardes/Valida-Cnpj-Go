package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o CNPJ: ")
	input, _ := reader.ReadString('\n')
	cnpj := strings.TrimSpace(input)

	if IsValidaCnpj(cnpj) {
		fmt.Println("\u2705 CNPJ Válido.")
	} else {
		fmt.Println("\u274C CNPJ Inválido.")
	}
}
