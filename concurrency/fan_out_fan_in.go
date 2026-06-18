// fan_out_fan_in.go — Fan-Out / Fan-In Pattern
// Fan-out: satu producer mengirim data ke banyak worker.
// Fan-in: banyak worker mengirim hasil ke satu merged channel.
//go:build fan_out_fan_in

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// producer mengirim data ke channel output.
// Setelah selesai, channel ditutup.
func producer(out chan<- int, count int) {
	for i := 1; i <= count; i++ {
		out <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(out)
}

// worker membaca dari input, memproses, dan mengirim ke output.
// Setelah input habis (channel di-close producer), worker close channel output-nya sendiri.
func worker(id int, in <-chan int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out) // Tutup channel output agar fanIn tidak deadlock
	for val := range in {
		// Simulasi proses berat
		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
		out <- fmt.Sprintf("worker-%d: input=%d → output=%d", id, val, val*val)
	}
}

// fanIn menggabungkan beberapa channel input ke satu channel output.
// WaitGroup menunggu semua goroutine selesai, lalu close merged.
func fanIn(inputs ...<-chan string) <-chan string {
	merged := make(chan string, 100)
	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for msg := range c {
				merged <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	const numJobs = 8
	const numWorkers = 3

	jobs := make(chan int, numJobs)

	// Fan-out: spread data ke N worker
	workerChannels := make([]chan string, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		workerChannels[i] = make(chan string, numJobs)
		wg.Add(1)
		go worker(i+1, jobs, workerChannels[i], &wg)
	}

	// Producer (akan close jobs setelah selesai)
	go producer(jobs, numJobs)

	// Konversi []chan string ke []<-chan string untuk fanIn
	inputs := make([]<-chan string, numWorkers)
	for i, ch := range workerChannels {
		inputs[i] = ch
	}

	// Fan-in: gabungkan semua hasil worker
	merged := fanIn(inputs...)

	// Consumer: baca hasil dari merged channel
	for res := range merged {
		fmt.Println(res)
	}

	fmt.Println("Fan-out/fan-in selesai.")
}
