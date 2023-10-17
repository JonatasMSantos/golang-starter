// TUDO que é criado com letra maiúcula é publico, e minúscula é privado
// Paradigma: Go Way (Não tem paradigma)
// Toda função que tem go na frente gera uma nova Thread
// Channels: Criação de multithreads assincronas

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func funcaoRetornoSimples(x int, y int) int {
	return x + y
}

/*
No Go é permitido o retorno de mais de um valor por tipo (int, bool)
*/
func funcaoRetornoDuplo(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}

/*
Criação do equivalente a classe no Go.
sendo “json:"name"` = TAG. O apelido na construção do objeto JSON.
*/
type Course struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func variaveisEFuncoes() {
	var varTipada string = "Hello, World!"
	varImplicita := "Hello, World!" //string

	println("Resultado 1: ", varTipada)
	println("Resultado 1: ", varImplicita)
	println("Resultado 2: ", funcaoRetornoSimples(1, 2))

	resultado, status := funcaoRetornoDuplo(10, 20)
	println("Resultado 3: ", resultado, status)
}

func criarCurso() Course {
	course := Course{
		Name:        "Golang",
		Description: "Golang Course",
		Price:       100,
	}

	return course
}

func getFullInfo(c Course) string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func iniciarServidorWeb() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

/*Estrutura for */
func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// go routine
func executarComThread() {
	go counter()
	go counter()
	go counter()
}

// Channels: Enviado dados de uma thread para outra
func channels() {
	channel := make(chan string)

	go func() {
		channel <- "Hello To another func"
	}()

	fmt.Println(<-channel)
}

// Worker: processar as requisições por balanceamento de carga
func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func execWorker() {
	channel := make(chan int)
	go worker(1, channel) //T2
	go worker(2, channel) //T3

	for i := 0; i < 10; i++ {
		channel <- i
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	course := criarCurso()
	course.Price = 150
	fmt.Println(getFullInfo(course))

	json.NewEncoder(w).Encode(course)
	// w.Write([]byte("Hello World"))
}

func main() {
	executarComThread()
	variaveisEFuncoes()
	iniciarServidorWeb()
	execWorker()
	channels()
}
