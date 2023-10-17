# Golang Starter

## Descrição do Projeto

Este projeto tem como objetivo fornecer um ponto de partida para aprender a linguagem de programação Go (Golang). Golang é conhecida por sua simplicidade, desempenho e eficiência na criação de aplicativos de alto desempenho. Neste repositório, você encontrará exemplos e explicações sobre alguns dos conceitos mais importantes da linguagem.

## Tópicos Cobertos

### 1. Variáveis

Aprenda como declarar e usar variáveis em Go. Explore os tipos de variáveis disponíveis, como inteiros, flutuantes, strings, booleanos e muito mais.

### 2. Funções

Veja como declarar funções em Go, incluindo funções com parâmetros e retorno. Aprenda a estrutura básica de uma função Go e como chamá-la.

### 3. Funções com Retorno de Mais de um Dado

Entenda como retornar múltiplos valores de uma função em Go. Explore exemplos práticos de funções que retornam mais de um dado.

### 4. Estrutura de Controle: For

Saiba como usar loops `for` em Go para iterar por coleções de dados e realizar tarefas repetitivas. Veja exemplos de loops `for` simples e complexos.

### 5. Channels e Goroutines

Introdução às goroutines e channels, recursos poderosos em Go para facilitar a programação concorrente e paralela. Veja como criar goroutines e usar channels para comunicação entre threads.

## Como Usar Este Projeto

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/golang-starter.git
```

```bash
go run main.go
```

## Instalar Go no Ubuntu

```bash
rm -rf /usr/local/go
```

```bash
cd /usr/local/
```

```go
sudo wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
```

```bash
tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```

```go
package main import "fmt" func main() { fmt.Println("👋 Hello World.") }
```

```go
go run helloworld.go
```

## Iniciar e realizar o build de um projeto

```bash
go mod init github.com/JonatasMSantos/golang-starter
```

```go
package main

func main() {
	println("Hello, World!")
}
```

Testar:

```bash
go run main.go
```

Build

```bash
go build
```

```bash
GOOS=windows go build
```

[https://github.com/JonatasMSantos/golang-starter](https://github.com/JonatasMSantos/golang-starter)
