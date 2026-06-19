package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Idiomatic Go ===")
	fmt.Println()
	fmt.Println("go run ./idiomatic/embedding/           # Embedding")
	fmt.Println("go run ./idiomatic/functional_options/  # Functional options")
	fmt.Println("go run ./idiomatic/error_wrapping/      # Error wrapping")
	fmt.Println()
	fmt.Println("Test:")
	fmt.Println("  go test -v ./idiomatic/")
}
