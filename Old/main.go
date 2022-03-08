package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *contaCorrente) sacar(valorSaque float64) string {

	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente ou com valor invalido."
	}
}

func main() {

	contaSilvia := contaCorrente{}
	contaSilvia.titular = "Silvia"
	contaSilvia.saldo = 500

	fmt.Println(contaSilvia.saldo)

	fmt.Println(contaSilvia.sacar(200))

	fmt.Println(contaSilvia.saldo)

	// contaSilvia := contaCorrente{"Silvia", 100, 63464364, 3000.50}
	// fmt.Println(contaSilvia)

	// contaEnrico := contaCorrente{titularConta: "Enrico", numeroAgencia: 1000, numeroConta: 9999999, saldoConta: 250.25}
	// contaMarina := contaCorrente{"Marina", 3500, 31241412, 1000.3}

	// var contaCris *contaCorrente
	// contaCris = new(contaCorrente)
	// contaCris.titularConta = "Cris"
	// contaCris.saldoConta = 500

}
