
package main

import "fmt"

// ============================================================
// Interface declaration
// ============================================================

// Speaker punya method Speak().
type Speaker interface {
	Speak() string
}

// Dog dan Cat implement Speaker secara implisit.
type Dog struct{ Name string }

func (d Dog) Speak() string {
	return d.Name + ": Woof!"
}

type Cat struct{ Name string }

func (c Cat) Speak() string {
	return c.Name + ": Meow!"
}

// Stringer — meng-override fmt output.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d yo)", p.Name, p.Age)
}

// ============================================================
// Any type / empty interface
// ============================================================

func describe(v any) {
	fmt.Printf("(%T) %[1]v\n", v)
}

// ============================================================
// Nil interface trap
// ============================================================

// MyStruct implement Speaker.
type MyStruct struct{}

func (m *MyStruct) Speak() string {
	return "from MyStruct"
}

// Nil interface trap:
// meskipun *MyStruct nil, setelah di-assign ke Speaker,
// interface-nya TIDAK nil karena tetap ada type info.
func nilTrapDemo() {
	var struk *MyStruct = nil   // pointer nil bertipe *MyStruct
	var pembicara Speaker = struk // interface menyimpan (*MyStruct, nil)

	fmt.Println("struk == nil:", struk == nil)             // true
	fmt.Println("pembicara == nil:", pembicara == nil)     // false — karena type info masih ada!

	// Cara aman: cek via type assertion
	_, ok := pembicara.(*MyStruct)
	fmt.Println("pembicara.(*MyStruct) ok:", ok)
}

func main() {
	// Interface implicit satisfaction
	fmt.Println("=== Interface Implisit ===")
	var pembicara Speaker
	pembicara = Dog{Name: "Milo"}
	fmt.Println(pembicara.Speak())
	pembicara = Cat{Name: "Pusi"}
	fmt.Println(pembicara.Speak())
	fmt.Println()

	// fmt.Stringer
	fmt.Println("=== fmt.Stringer ===")
	p := Person{Name: "Sari", Age: 30}
	fmt.Println(p) // pakai String()
	fmt.Println()

	// Empty interface / any
	fmt.Println("=== any / interface{} ===")
	describe(42)
	describe("halo")
	describe(3.14)
	describe(Dog{Name: "Guguk"})
	fmt.Println()

	// Nil interface trap
	fmt.Println("=== Nil Interface Trap ===")
	nilTrapDemo()
}
