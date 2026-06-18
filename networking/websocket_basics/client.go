//go:build client

// WebSocket Client (Simulasi)
//
// Di production, client juga pakai gorilla/websocket:
//
//	go get github.com/gorilla/websocket
//
// lalu:
//
//	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
//	c.WriteMessage(websocket.TextMessage, []byte("hello"))
//	_, msg, err := c.ReadMessage()
//
// Simulasi ini melakukan HTTP upgrade handshake manual via raw TCP,
// lalu menggunakan protokol line-based untuk komunikasi.
//
// Cara run (pastikan server sudah jalan):
//
//	go run -tags=client ./networking/websocket_basics/
//
// Ketik pesan di terminal, server akan echo balik.
// Ketik "exit" untuk keluar.
package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// wsConn — definisi sama dengan server.
// Di production: gorilla/websocket.Conn.
type wsConn struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func newWSConn(conn net.Conn) *wsConn {
	return &wsConn{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func (c *wsConn) ReadMessage() ([]byte, error) {
	line, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	return line[:len(line)-1], nil
}

func (c *wsConn) WriteMessage(data []byte) error {
	if _, err := c.writer.Write(data); err != nil {
		return err
	}
	return c.writer.WriteByte('\n')
}

func (c *wsConn) flush() error  { return c.writer.Flush() }
func (c *wsConn) Close() error { return c.conn.Close() }

func main() {
	addr := "localhost:8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	// 1. Dial TCP ke server.
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Printf("Gagal connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 2. Kirim HTTP upgrade request manual.
	// Di production: websocket.DefaultDialer.Dial("ws://...", nil).
	req := fmt.Sprintf("GET /ws HTTP/1.1\r\nHost: %s\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n\r\n", addr)
	if _, err := conn.Write([]byte(req)); err != nil {
		fmt.Printf("Gagal kirim upgrade request: %v\n", err)
		os.Exit(1)
	}

	// 3. Baca response upgrade.
	reader := bufio.NewReader(conn)
	resp, err := http.ReadResponse(reader, nil)
	if err != nil {
		fmt.Printf("Gagal baca response upgrade: %v\n", err)
		os.Exit(1)
	}
	// 101 = Switching Protocols => upgrade berhasil.
	if resp.StatusCode != http.StatusSwitchingProtocols {
		fmt.Printf("Upgrade gagal: %s\n", resp.Status)
		os.Exit(1)
	}

	ws := &wsConn{
		conn:   conn,
		reader: reader,
		writer: bufio.NewWriter(conn),
	}

	fmt.Println("Terhubung ke WebSocket server!")
	fmt.Println("Ketik pesan, server akan echo balik. Ketik 'exit' untuk keluar.")

	// 4. Goroutine untuk baca pesan dari server.
	// Di production: _, msg, err := conn.ReadMessage().
	go func() {
		for {
			msg, err := ws.ReadMessage()
			if err != nil {
				fmt.Printf("\nKoneksi terputus: %v\n", err)
				os.Exit(0)
			}
			fmt.Printf("Server > %s\n", msg)
		}
	}()

	// 5. Baca input dari stdin dan kirim ke server.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		if err := ws.WriteMessage([]byte(text)); err != nil {
			fmt.Printf("Error write: %v\n", err)
			break
		}
		ws.flush()
	}

	fmt.Println("Koneksi ditutup.")
}
