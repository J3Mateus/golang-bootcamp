package main

import (
	"fmt"
)

type Cadastro struct {
	Nome  string
	Cpf   int
	Email string
}

type Pessoa interface {
	ObterNome() string
	ObterCpf() int
	ObterEmail() string
}

func (c Cadastro) ObterNome() string {
	return c.Nome
}

func (c Cadastro) ObterCpf() int {
	return c.Cpf
}

func (c Cadastro) ObterEmail() string {
	return c.Email
}

func main() {
	cadastro := Cadastro{
		Nome:  "João",
		Cpf:   30045678900,
		Email: "jp@email.com",
	}
	var pessoa Pessoa = cadastro

	fmt.Println("Nome:", pessoa.ObterNome())
	fmt.Println("Cpf:", pessoa.ObterCpf())
	fmt.Println("Email:", pessoa.ObterEmail())
}
