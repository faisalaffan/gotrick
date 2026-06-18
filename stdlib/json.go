// File: json.go
// Marshal, MarshalIndent, Unmarshal, struct tags, omitempty,
// dan custom MarshalJSON/UnmarshalJSON.

//go:build json

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// --- Struct dengan JSON tags ---

// Employee struct dengan json tags.
type Employee struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email,omitempty"`             // omit jika empty string
	Salary    *float64 `json:"salary,omitempty"`            // omit jika nil pointer
	Skills    []string `json:"skills,omitempty"`            // omit jika nil/empty slice
	IsActive  bool     `json:"is_active"`
	CreatedAt CustomTime `json:"created_at"`                // custom marshal/unmarshal
}

// CustomTime wrapper time.Time dengan format JSON custom.
type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05"

// MarshalJSON implementasi json.Marshaler untuk CustomTime.
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + ct.Format(customTimeFormat) + `"`), nil
}

// UnmarshalJSON implementasi json.Unmarshaler untuk CustomTime.
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	parsed, err := time.Parse(customTimeFormat, s)
	if err != nil {
		return fmt.Errorf("parse custom time: %w", err)
	}
	ct.Time = parsed
	return nil
}

func main() {
	// 1. Marshal — serialize struct ke JSON
	fmt.Println("=== 1. json.Marshal ===")
	salary := 75000.0
	emp := Employee{
		ID:        1,
		Name:      "Budi Santoso",
		Email:     "budi@example.com",
		Salary:    &salary,
		Skills:    []string{"Go", "Docker", "K8s"},
		IsActive:  true,
		CreatedAt: CustomTime{Time: time.Now()},
	}
	data, err := json.Marshal(emp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// 2. MarshalIndent — pretty print
	fmt.Println("\n=== 2. json.MarshalIndent ===")
	dataIndent, _ := json.MarshalIndent(emp, "", "  ")
	fmt.Println(string(dataIndent))

	// 3. Omitempty demo — field kosong akan dihilangkan
	fmt.Println("\n=== 3. Omitempty — field kosong ===")
	emp2 := Employee{
		ID:       2,
		Name:     "Siti Nurhaliza",
		IsActive: false, // bool false tetap muncul (bukan omitempty)
		CreatedAt: CustomTime{Time: time.Now()},
	}
	data2, _ := json.MarshalIndent(emp2, "", "  ")
	fmt.Println(string(data2))

	// 4. Unmarshal — parse JSON ke struct
	fmt.Println("\n=== 4. json.Unmarshal ===")
	jsonStr := `{
		"id": 3,
		"name": "Agus Wijaya",
		"email": "agus@example.com",
		"salary": 85000.5,
		"skills": ["Python", "Go"],
		"is_active": true,
		"created_at": "2025-06-15 10:30:00"
	}`
	var emp3 Employee
	if err := json.Unmarshal([]byte(jsonStr), &emp3); err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d, Nama: %s, Email: %s, Salary: %.0f, Skills: %v, Aktif: %t, Dibuat: %s\n",
		emp3.ID, emp3.Name, emp3.Email, *emp3.Salary, emp3.Skills, emp3.IsActive, emp3.CreatedAt.Format(customTimeFormat))

	// 5. Omitempty dengan pointer — nil tidak muncul
	fmt.Println("\n=== 5. Omitempty dengan pointer nil ===")
	emp4 := Employee{
		ID:        4,
		Name:      "Dewi Lestari",
		IsActive:  true,
		CreatedAt: CustomTime{Time: time.Now()},
	}
	data4, _ := json.MarshalIndent(emp4, "", "  ")
	fmt.Println(string(data4))
	fmt.Println("(Catatan: Salary tidak muncul karena nil)")

	// 6. Unmarshal JSON array
	fmt.Println("\n=== 6. Unmarshal JSON array ===")
	jsonArray := `[
		{"id":10, "name":"User A", "is_active":true, "created_at":"2025-01-01 00:00:00"},
		{"id":11, "name":"User B", "is_active":false, "created_at":"2025-06-01 12:00:00"}
	]`
	var users []Employee
	if err := json.Unmarshal([]byte(jsonArray), &users); err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("  - %d: %s (aktif: %t)\n", u.ID, u.Name, u.IsActive)
	}
}
