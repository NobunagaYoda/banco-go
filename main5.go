package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteBruno := clientes.Titular{"Bruno", "123.123.123-12", "Desenvolvedor Go"}
	contaBruno := contas.ContaCorrente{clienteBruno, 123, 3231212, 500}

	fmt.Println(contaBruno)

}
