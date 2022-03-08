package main

import (
	// "banco/clientes"
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {

	// clienteBruno := clientes.Titular{"Bruno", "123.123.123-12", "Desenvolvedor Go"}
	// contaBruno := contas.ContaCorrente{clienteBruno, 123, 3231212, 500}
	// fmt.Println(contaBruno)

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	pagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.ObterSaldo())

	contaLuiza := contas.ContaCorrente{}
	contaLuiza.Depositar(500)
	pagarBoleto(&contaLuiza, 400)

	fmt.Println(contaLuiza.ObterSaldo())
}
