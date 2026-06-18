//go:build server

// WebSocket Echo Server (Simulasi)
//
// Di production, WebSocket butuh library eksternal:
//
//	go get github.com/gorilla/websocket
//
// lalu gunakan gorilla/websocket.Upgrade() untuk upgrade HTTP ke WebSocket,
// dan conn.ReadMessage() / conn.WriteMessage() untuk komunikasi.
//
// Simulasi ini menggunakan:
//   - http.Hijacker untuk mengambil alih koneksi HTTP setelah handshake
//   - Protocol line-based (\n delimiter) sebagai ganti WebSocket frame
//   - Pattern: upgrade handshake → read/write loop → close
//
// Cara run:
//
//	go run -tags=server ./networking/websocket_basics/
//
// Lalu jalankan client di terminal lain:
//
//	go run -tags=client ./networking/websocket_basics/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

// wsConn adalah simulasi WebSocket connection.
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

// ReadMessage membaca satu pesan.
// Di WebSocket asli: read frame, unmask payload, handle fragmentasi.
func (c *wsConn) ReadMessage() ([]byte, error) {
	line, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	return line[:len(line)-1], nil // buang newline
}

// WriteMessage mengirim satu pesan.
// Di WebSocket asli: frame + mask (client), frame saja (server).
func (c *wsConn) WriteMessage(data []byte) error {
	if _, err := c.writer.Write(data); err != nil {
		return err
	}
	return c.writer.WriteByte('\n')
}

func (c *wsConn) flush() error  { return c.writer.Flush() }
func (c *wsConn) Close() error { return c.conn.Close() }

// handleConnection adalah echo handler: baca pesan → echo balik.
func handleConnection(ws *wsConn) {
	defer ws.Close()
	log.Println("Client terhubung")

	for {
		msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Client disconnect: %v", err)
			return
		}
		log.Printf("Menerima: %s", msg)

		// Echo pesan.
		echo := append([]byte("Echo: "), msg...)
		if err := ws.WriteMessage(echo); err != nil {
			log.Printf("Error write: %v", err)
			return
		}
		ws.flush()
	}
}

// upgradeHandler menangani HTTP → WebSocket upgrade.
// Di WebSocket asli: conn, err := upgrader.Upgrade(w, r, nil).
func upgradeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") != "websocket" {
		http.Error(w, "bukan request WebSocket", http.StatusBadRequest)
		return
	}

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "hijacking tidak didukung", http.StatusInternalServerError)
		return
	}

	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim response upgrade secara manual.
	// Di production: gorilla/websocket melakukan ini otomatis di Upgrader.Upgrade().
	response := "HTTP/1.1 101 Switching Protocols\r\n" +
		"Upgrade: websocket\r\n" +
		"Connection: Upgrade\r\n" +
		"\r\n"
	if _, err := bufrw.WriteString(response); err != nil {
		log.Printf("Error write upgrade response: %v", err)
		conn.Close()
		return
	}
	bufrw.Flush()

	ws := &wsConn{
		conn:   conn,
		reader: bufrw.Reader,
		writer: bufrw.Writer,
	}

	handleConnection(ws)
}

func main() {
	http.HandleFunc("/ws", upgradeHandler)

	addr := ":8080"
	fmt.Printf("WebSocket echo server listening di %s/ws\n", addr)
	fmt.Println("Jalankan client:")
	fmt.Println("  go run -tags=client ./networking/websocket_basics/")
	log.Fatal(http.ListenAndServe(addr, nil))
}
