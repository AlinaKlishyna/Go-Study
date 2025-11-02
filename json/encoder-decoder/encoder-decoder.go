package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	src = testStruct{Name: "John Connor", Age: 35} // структура с данными
	dst = testStruct{}                             // структура без данных
	buf = new(bytes.Buffer)                        // буфер для чтения и записи
)

func main() {
	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	enc.Encode(src)
	dec.Decode(&dst)

	fmt.Print(dst) // {John Connor 35}

}
