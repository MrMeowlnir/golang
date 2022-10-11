package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"` // tags/теги для преобразования имени поля в json с маленькой буквы
	Age       int    `json:"age"`
	IsBlocked bool   `json:"is_blocked"`
} // Все поля с большой буквы = экспортируемые. Неэкспортируемые поля не передадутся в json

func main() {
	sv := map[string]interface{}{"filed": "value", "field1": 1, "field2": true}
	bytevar, _ := json.Marshal(sv)
	fmt.Println(string(bytevar))

	sv_struct := User{
		Name:      "Anykey",
		Age:       10,
		IsBlocked: true,
	}
	struct_json, _ := json.Marshal(sv_struct)
	fmt.Println(string(struct_json))
}
