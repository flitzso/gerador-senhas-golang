package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Semente do gerador de números aleatórios
    rand.Seed(time.Now().UnixNano())

    // Tamanho da senha
    tamanho := 12

    // Caracteres possíveis na senha
    caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+"

    // Gerar senha aleatória
    senha := make([]byte, tamanho)
    for i := range senha {
        senha[i] = caracteres[rand.Intn(len(caracteres))]
    }

    // Imprimir senha
    fmt.Println(string(senha))
}
