// fan_out_fan_in.go — Fan-Out / Fan-In Pattern
// Fan-out: satu sumber gambar dikirim ke banyak prosesor.
// Fan-in: banyak prosesor mengirim hasil ke satu channel merged.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// pembacaGambar mengirim resolusi gambar ke channel output.
// Setelah selesai, channel ditutup.
func pembacaGambar(out chan<- int, count int) {
	for i := 1; i <= count; i++ {
		out <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(out)
}

// prosesor membaca resolusi dari input, menghitung piksel, dan mengirim ke output.
// Setelah input habis (channel di-close pembacaGambar), prosesor close channel output-nya sendiri.
func prosesor(id int, in <-chan int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out) // Tutup channel output agar fanIn tidak deadlock
	for resolusi := range in {
		// Simulasi proses render gambar
		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
		out <- fmt.Sprintf("prosesor-%d: gambar %dpx → total piksel: %d", id, resolusi, resolusi*resolusi)
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
	const jumlahGambar = 8
	const jumlahProsesor = 3

	gambar := make(chan int, jumlahGambar)

	// Fan-out: spread data ke N prosesor
	channelProsesor := make([]chan string, jumlahProsesor)
	var wg sync.WaitGroup

	for i := 0; i < jumlahProsesor; i++ {
		channelProsesor[i] = make(chan string, jumlahGambar)
		wg.Add(1)
		go prosesor(i+1, gambar, channelProsesor[i], &wg)
	}

	// PembacaGambar (akan close gambar setelah selesai)
	go pembacaGambar(gambar, jumlahGambar)

	// Konversi []chan string ke []<-chan string untuk fanIn
	inputs := make([]<-chan string, jumlahProsesor)
	for i, ch := range channelProsesor {
		inputs[i] = ch
	}

	// Fan-in: gabungkan semua hasil prosesor
	merged := fanIn(inputs...)

	// Consumer: baca hasil dari merged channel
	for res := range merged {
		fmt.Println(res)
	}

	fmt.Println("Semua gambar selesai diproses.")
}
