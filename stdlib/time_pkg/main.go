// File: time_pkg.go
// time.Now, time.Date, Format, Parse, Since, Until,
// Ticker, Timer, After, Sleep.


package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. time.Now()
	fmt.Println("=== 1. time.Now() ===")
	now := time.Now()
	fmt.Printf("Waktu sekarang: %v\n", now)
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	fmt.Printf("Tahun: %d, Bulan: %s, Hari: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Jam: %d, Menit: %d, Detik: %d\n", now.Hour(), now.Minute(), now.Second())

	// 2. time.Date() — buat waktu spesifik
	fmt.Println("\n=== 2. time.Date() ===")
	specificDate := time.Date(2025, 12, 25, 10, 30, 0, 0, time.UTC)
	fmt.Printf("Tanggal spesifik: %v\n", specificDate)
	fmt.Printf("Hari: %s, Bulan: %s\n", specificDate.Weekday(), specificDate.Month())

	// 3. Format dengan layout
	fmt.Println("\n=== 3. time.Format() ===")
	// Layout Go: 2006-01-02 15:04:05 (Mon Jan 2 15:04:05 MST 2006 = 1 2 3 4 5 6 7)
	fmt.Printf("Format 1 (RFC3339): %s\n", now.Format(time.RFC3339))
	fmt.Printf("Format 2 (YYYY-MM-DD): %s\n", now.Format("2006-01-02"))
	fmt.Printf("Format 3 (DD/MM/YYYY HH:mm:ss): %s\n", now.Format("02/01/2006 15:04:05"))
	fmt.Printf("Format 4 (Hari, Tanggal): %s\n", now.Format("Monday, 02 January 2006"))
	fmt.Printf("Format 5 (Jam saja): %s\n", now.Format("15:04:05"))

	// 4. Parse — string ke time.Time
	fmt.Println("\n=== 4. time.Parse() ===")
	dateStr := "2025-06-18 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse result: %v\n", parsedTime)
	fmt.Printf("  Tahun: %d, Bulan: %s, Hari: %d\n", parsedTime.Year(), parsedTime.Month(), parsedTime.Day())

	// Parse dengan timezone
	dateStrTZ := "2025-06-18T14:30:00+07:00"
	parsedTZ, err := time.Parse(time.RFC3339, dateStrTZ)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse RFC3339: %v (location: %s)\n", parsedTZ, parsedTZ.Location())

	// 5. time.Since — durasi dari suatu waktu
	fmt.Println("\n=== 5. time.Since() ===")
	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Printf("Durasi Sleep 50ms: %v\n", elapsed)

	// 6. time.Until — durasi sampai suatu waktu
	fmt.Println("\n=== 6. time.Until() ===")
	future := time.Now().Add(2 * time.Hour)
	duration := time.Until(future)
	fmt.Printf("Sisa waktu sampai +2 jam: %v\n", duration.Round(time.Second))

	// 7. time.Ticker — periodic task
	fmt.Println("\n=== 7. time.Ticker ===")
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan struct{})

	go func() {
		ticks := 0
		for {
			select {
			case t := <-ticker.C:
				ticks++
				fmt.Printf("  Ticker tick ke-%d: %s\n", ticks, t.Format("15:04:05.000"))
				if ticks >= 3 {
					ticker.Stop()
					close(done)
					return
				}
			}
		}
	}()
	<-done

	// 8. time.Timer — sekali delay
	fmt.Println("\n=== 8. time.Timer ===")
	timer := time.NewTimer(200 * time.Millisecond)
	<-timer.C
	fmt.Println("  Timer fired setelah 200ms")

	// Timer bisa di-stop sebelum fire
	timer2 := time.NewTimer(1 * time.Hour)
	if timer2.Stop() {
		fmt.Println("  Timer 2 dihentikan sebelum fire (Stop=true)")
	}

	// 9. time.After — return channel untuk select
	fmt.Println("\n=== 9. time.After — select timeout ===")
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  After 100ms: timeout tercapai")
	}

	// Contoh select dengan After sebagai timeout
	fmt.Println("\n  Contoh select dengan timeout:")
	ch := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "hasil"
	}()

	select {
	case msg := <-ch:
		fmt.Printf("  Menerima: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout! Channel tidak merespon dalam 100ms")
	}

	// 10. time.Sleep
	fmt.Println("\n=== 10. time.Sleep ===")
	fmt.Println("  Tidur selama 100ms...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Bangun!")

	// Ringkasan
	fmt.Println("\n--- Ringkasan ---")
	fmt.Println("  time.Now()        : waktu sekarang")
	fmt.Println("  time.Date()       : buat waktu spesifik")
	fmt.Println("  time.Format()     : format time.Time ke string")
	fmt.Println("  time.Parse()      : parse string ke time.Time")
	fmt.Println("  time.Since()      : durasi dari waktu lalu")
	fmt.Println("  time.Until()      : durasi sampai waktu depan")
	fmt.Println("  time.Ticker       : periodic task")
	fmt.Println("  time.Timer        : sekali delay (reset-able)")
	fmt.Println("  time.After()      : channel untuk timeout di select")
	fmt.Println("  time.Sleep()      : delay blok")
	fmt.Println("\nLayout Go (wajib hafal): Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("  01 = bulan, 02 = hari, 03/15 = jam, 04 = menit, 05 = detik")
	fmt.Println("  2006 = tahun, Mon = hari, MST = timezone")
}
