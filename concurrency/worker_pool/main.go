// worker_pool.go — Worker Pool Pattern
// N worker mengambil job dari buffered channel, memproses, lalu mengirim hasil.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker memproses job dari channel jobs dan mengirim hasil ke results.
// wg.Done() dipanggil saat worker selesai.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulasi proses yang memakan waktu variatif
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		result := fmt.Sprintf("worker-%d memproses job-%d → hasil: %d", id, job, job*job)
		results <- result
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// Buffered channel agar producer tidak blocking
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	// Start N worker goroutine
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Kirim job ke channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Tidak ada job lagi — worker akan exit dari range loop

	// Tunggu semua worker selesai di goroutine terpisah,
	// lalu close results channel agar consumer bisa exit.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect dan print semua hasil
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Semua worker selesai.")
}
