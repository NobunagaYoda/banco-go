package main

import "fmt"

type Titular struct {
	Nome      string
	CPF       string
	Profissao string
}

type ContaCorrente struct {
	Titular       Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func main() {
	// contaBruno := ContaCorrente{Titular: Titular{
	// 	Nome:      "Bruno",
	// 	CPF:       "487.099.668-63",
	// 	Profissao: "Engenheiro"},
	// 	NumeroAgencia: 123, NumeroConta: 312312, Saldo: 500}

	// fmt.Println(contaBruno)

	clienteBruno := Titular{"Bruno", "123.123.123-12", "Desenvolvedor Go"}
	contaBruno := ContaCorrente{clienteBruno, 123, 3231212, 500}

	fmt.Println(contaBruno)

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

func (c *ContaCorrente) obterSaldo() float64 {
	return c.Saldo

}
