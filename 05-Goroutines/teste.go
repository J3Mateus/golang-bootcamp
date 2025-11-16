package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("eu sou uma nova goroutine")
}

func main() {
	go hello()                                  //inicia uma nova goroutine, é colocado numa fila de execução
	fmt.Println("eu sou a goroutine principal") //executa primeiro a principal do main
	time.Sleep(time.Second * 1)                 //espera 1 segundo para garantir que a goroutine tenha tempo de executar antes do programa terminar

}
