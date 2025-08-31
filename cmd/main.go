package main

import (
	"fmt"
	"net/http"
	"websecscanner/internal/api"
)

func main() {
	http.HandleFunc("/scan", api.ScanHandler)
	fmt.Println("Servidor iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)

	}
}
