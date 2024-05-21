package main

import "godemo/internal/goconcurrency"

type Node struct {
	FileNodeIP string
	Status     string
}

func main() {
	goconcurrency.PrintFib()
}
