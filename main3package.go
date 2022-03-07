package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	fmt.Println(contaSilvia, contaGustavo)

	status := contaGustavo.Transferir(100, &contaSilvia)
	fmt.Println(contaSilvia, contaGustavo)
	fmt.Println(status)

	// testando o git
}
