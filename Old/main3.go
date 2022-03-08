package main

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso. Saldo:", c.Saldo
	} else {
		return "Saldo insuficiente ou com valor invalido. Saldo:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso. Saldo:", c.Saldo

	} else {
		return "Deposito com valor invalido. Saldo:", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {

	contaSilvia := ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaGustavo := ContaCorrente{Titular: "Gustavo", Saldo: 100}
	fmt.Println(contaSilvia, contaGustavo)

	status := contaGustavo.Transferir(100, &contaSilvia)
	fmt.Println(contaSilvia, contaGustavo)
	fmt.Println(status)

}

// contaSilvia := ContaCorrentente{"Silvia", 100, 63464364, 3000.50}
// fmt.Println(contaSilvia)

// fmt.Println("Saldo inicial:", contaSilvia.saldo)
// fmt.Println(contaSilvia.sacar(300))
// status, valor := contaSilvia.depositar(2000)
// fmt.Println(status, valor)

// contaEnrico := ContaCorrentente{titularConta: "Enrico", numeroAgencia: 1000, numeroConta: 9999999, saldoConta: 250.25}
// contaMarina := ContaCorrentente{"Marina", 3500, 31241412, 1000.3}

// var contaCris *ContaCorrentente
// contaCris = new(ContaCorrentente)
// contaCris.titularConta = "Cris"
// contaCris.saldoConta = 500
