package interfaces

import "fmt"

// Base theory
type AnyStruct struct {
	N1, N2 int
}

type AnyInterface interface {
	Sum() int
}

func (s *AnyStruct) Sum() int {
	return s.N1 + s.N2
}

type OtherStruct struct {
	A, B int
}

func (o OtherStruct) Sum() int {
	return o.A + o.B
}

// --- end Base theory

// future theory
// small tip
var _ User = &user{}

// end tip

type user struct {
	FIO, Address, Phone string
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func (u *user) ChangeFIO(newFIO string) {
	fmt.Println("Old FIO:", u.FIO)
	u.FIO = newFIO
	fmt.Println("New FIO:", u.FIO)
}

func (u *user) ChangeAddress(newAddress string) {
	fmt.Println("Old Address:", u.Address)
	u.Address = newAddress
	fmt.Println("New Address:", u.Address)
}

func NewUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u // !return &pointer to relize interface User
}

func main() {
	// base theory
	fmt.Println("Base theory")
	var a AnyInterface
	sh := AnyStruct{1, 2}
	os := OtherStruct{3, 4}

	a = &sh
	fmt.Println(a.Sum())
	a = os
	fmt.Println(a.Sum())

	// further theory
	fmt.Println()
	fmt.Println("Further Theory")
	u := NewUser("fio", "address", "phone") // "Конструктор класса" (почти)
	u.ChangeFIO("newFIO")                   // ! нет прямого доступа к полям u !
	fmt.Println()
	// empty interfaces
	fmt.Println("Empty Interfaces")
	var b interface{} // может хранить значения любого типа, как переменная в питоне
	b = "String Value"
	fmt.Println(b)
	fmt.Printf("(%v, %T)\n\n", b, b)
	b = 0
	fmt.Println(b)
	fmt.Printf("(%v, %T)\n\n", b, b)

	// Type Assertion
	fmt.Println("Type assertion")
	var c interface{} = true
	s, ok := c.(string)
	fmt.Printf("c: (%v, %T)\n", c, c)
	fmt.Printf("s: (%v, %T)\n\n", s, s)
	fmt.Println("s:", s, ok)
	fmt.Println()
	f, ok := c.(float32)
	fmt.Println("f:", f, ok)
	fmt.Println()

	//Type switching
	switch c.(type) {
	case int:
		fmt.Println("c is int")
	case float32:
		fmt.Println("c is float32")
	case string:
		fmt.Println("c is string")
	default:
		fmt.Printf("Unknown type %T\n", c)
	}
}
