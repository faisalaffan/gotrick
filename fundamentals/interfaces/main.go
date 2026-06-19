
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
	var s *MyStruct = nil   // pointer nil bertipe *MyStruct
	var sp Speaker = s      // interface menyimpan (*MyStruct, nil)

	fmt.Println("s == nil:", s == nil)     // true
	fmt.Println("sp == nil:", sp == nil)   // false — karena type info masih ada!

	// Cara aman: cek via type assertion
	_, ok := sp.(*MyStruct)
	fmt.Println("sp.(*MyStruct) ok:", ok)
}

func main() {
	// Interface implicit satisfaction
	fmt.Println("=== Interface Implisit ===")
	var sp Speaker
	sp = Dog{Name: "Rex"}
	fmt.Println(sp.Speak())
	sp = Cat{Name: "Mimi"}
	fmt.Println(sp.Speak())
	fmt.Println()

	// fmt.Stringer
	fmt.Println("=== fmt.Stringer ===")
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // pakai String()
	fmt.Println()

	// Empty interface / any
	fmt.Println("=== any / interface{} ===")
	describe(42)
	describe("hello")
	describe(3.14)
	describe(Dog{Name: "Buddy"})
	fmt.Println()

	// Nil interface trap
	fmt.Println("=== Nil Interface Trap ===")
	nilTrapDemo()
}
