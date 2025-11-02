package main

import (
	"encoding/json"
	"fmt"
)

// неэкспортируемые поля (имена которых начинаются с маленькой буквы)
// не участвуют в кодировании / декодировании
// name string // не экспортируется
// Age  int    // экспортируется

// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`

type PC struct {
	// при кодировании / декодировании будет использовано имя model, а не Model
	Model string `json:"model"`
	Color string `json:"color"`

	// Если это поле пустое (пустое значение: false, nil, пустой слайс и пр.),
	// просто не включай его в JSON
	DataCreated int `json:",omitempty"`

	// при кодировании / декодировании поле всегда игнорируется
	Monitor bool `json:"-"`
}

func main() {
	// пример использования аннотаций
	pc := PC{
		Model:       "MacBook Pro",
		Color:       "Space Gray",
		DataCreated: 0,
		Monitor:     true,
	}

	data, err := json.Marshal(pc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data) // {"model":"MacBook Pro","color":"Space Gray"}
}
