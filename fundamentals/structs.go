//go:build structs

package main

import "fmt"

// User struct dengan field biasa.
type User struct {
	ID    int
	Name  string
	Email string
}

// Account struct dengan embedded struct (User) dan field tambahan.
type Account struct {
	User                    // embedded — field & method di-promote
	Balance float64
}

// Value receiver — tidak mengubah field asli.
func (u User) Greet() string {
	return "Halo, " + u.Name
}

// Pointer receiver — bisa mengubah field asli.
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

// Method pada embedded struct dipromosikan ke Account.
func (u User) Label() string {
	return u.Name + " <" + u.Email + ">"
}

func main() {
	// Struct literal — named fields
	u1 := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	fmt.Println("=== User ===")
	fmt.Printf("%+v\n", u1)
	fmt.Println(u1.Greet())

	// Pointer receiver — email berubah
	u1.ChangeEmail("alice@new.org")
	fmt.Println("Setelah ganti email:", u1.Email)
	fmt.Println()

	// Embedded struct — field & method User bisa diakses langsung
	acc := Account{
		User: User{
			ID:    10,
			Name:  "Bob",
			Email: "bob@example.com",
		},
		Balance: 250.75,
	}
	fmt.Println("=== Account (embedded User) ===")
	fmt.Printf("ID=%d Name=%s Saldo=%.2f\n", acc.ID, acc.Name, acc.Balance)
	fmt.Println("Label:", acc.Label()) // method User dipromosikan
	fmt.Println()

	// Perbedaan value vs pointer receiver
	fmt.Println("=== Value vs Pointer Receiver ===")
	u2 := User{ID: 2, Name: "Charlie", Email: "c@old.com"}
	u2.Greet()           // value receiver — boleh dipanggil via value
	u2.ChangeEmail("c@new.com") // pointer receiver — Go otomatis ambil &u2
	fmt.Printf("u2.Email = %s\n", u2.Email)
}
