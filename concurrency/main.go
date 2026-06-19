package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Concurrency Patterns ===")
	fmt.Println()
	fmt.Println("go run ./concurrency/worker_pool/     # Worker pool")
	fmt.Println("go run ./concurrency/fan_out_fan_in/  # Fan-out / fan-in")
	fmt.Println("go run ./concurrency/pipeline/        # Pipeline")
	fmt.Println("go run ./concurrency/select_pattern/  # Select patterns")
	fmt.Println("go run ./concurrency/atomic_vs_mutex/ # Atomic vs mutex")
	fmt.Println("go run ./concurrency/race_demo/       # Race condition")
}
