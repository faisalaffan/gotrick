
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
	penggunaBudi := User{
		ID:    1,
		Name:  "Budi",
		Email: "budi@example.com",
	}
	fmt.Println("=== User ===")
	fmt.Printf("%+v\n", penggunaBudi)
	fmt.Println(penggunaBudi.Greet())

	// Pointer receiver — email berubah
	penggunaBudi.ChangeEmail("budi@new.org")
	fmt.Println("Setelah ganti email:", penggunaBudi.Email)
	fmt.Println()

	// Embedded struct — field & method User bisa diakses langsung
	acc := Account{
		User: User{
			ID:    10,
			Name:  "Siti",
			Email: "siti@example.com",
		},
		Balance: 250.75,
	}
	fmt.Println("=== Account (embedded User) ===")
	fmt.Printf("ID=%d Name=%s Saldo=%.2f\n", acc.ID, acc.Name, acc.Balance)
	fmt.Println("Label:", acc.Label()) // method User dipromosikan
	fmt.Println()

	// Perbedaan value vs pointer receiver
	fmt.Println("=== Value vs Pointer Receiver ===")
	penggunaAmin := User{ID: 2, Name: "Amin", Email: "amin@old.com"}
	penggunaAmin.Greet()                      // value receiver — boleh dipanggil via value
	penggunaAmin.ChangeEmail("amin@new.com") // pointer receiver — Go otomatis ambil &penggunaAmin
	fmt.Printf("penggunaAmin.Email = %s\n", penggunaAmin.Email)
}
