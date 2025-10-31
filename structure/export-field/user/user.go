package user // Пакет user

// User - экспортируемая структура (с большой буквы)

type User struct {
	// i campi
	FirstName string
	LastName  string
	Birthday  int
	password  string // закрытое поле (с маленькой буквы)
}

// GetPassword - экспортируемый метод для получения закрытого поля password
func (p *User) GetPassword() string {
	return p.password
}

// SetPassword - экспортируемый метод для установки значения закрытого поля password
func (p *User) SetPassword(newPassword string) {
	p.password = newPassword
}
