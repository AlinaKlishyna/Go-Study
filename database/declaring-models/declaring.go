package declaringmodels

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           // Стандартное поле для первичного ключа
	Name         string         // Обычное строковое поле
	Email        *string        // Указатель на строку — позволяет хранить NULL в базе
	Age          uint8          // Беззнаковое целое число (8 бит)
	Birthday     *time.Time     // Указатель на дату — может быть NULL
	MemberNumber sql.NullString // Специальный тип для строк, которые могут быть NULL
	ActivatedAt  sql.NullTime   // Специальный тип для дат, которые могут быть NULL
	ignored      string         // Непубличное поле (с маленькой буквы) — GORM его игнорирует

	// gorm.Model  - Встраиваемые поля (ID, CreatedAt, UpdatedAt, DeletedAt)
	CreatedAt  time.Time // Автоматически заполняется GORM при создании записи
	UpdatedAt  time.Time // Автоматически обновляется GORM при изменении записи
	DeleteddAt time.Time // Автоматически заполняется GORM при удалении записи
}

// ->	чтение
// <-	запись, но чтение остаётся включено по умолчанию
// :	уточнение (например, только create/update/false)
// ; — и
// -	исключает из ORM

// ->:false	Явно запрещает чтение
// <-:false	Явно запрещает запись
type UserOther struct {
	Name1 string `gorm:"<-:create"`          // разрешено читать и создавать (create)
	Name2 string `gorm:"<-:update"`          // разрешено читать и обновлять (update)
	Name3 string `gorm:"<-"`                 // разрешено читать и записывать (create и update)
	Name4 string `gorm:"<-:false"`           // разрешено читать, запись запрещена
	Name5 string `gorm:"->"`                 // только для чтения (запись запрещена, если явно не указано)
	Name6 string `gorm:"->;<-:create"`       // разрешено читать и создавать
	Name7 string `gorm:"->:false;<-:create"` // только создание (чтение из базы отключено)
	Name8 string `gorm:"-"`                  // игнорируется при чтении и записи через структуру
	Name9 string `gorm:"-:all"`              // игнорируется при чтении, записи и миграции
	Name0 string `gorm:"-:migration"`        // игнорируется только при миграции
}

// Embedded Struct
type Author struct {
	Name  string `gorm:"embeddedPrefix:info_"` // info_name - Чтобы избежать конфликтов имён
	Email string
}

type Blog1 struct {
	ID      int
	Author  Author `gorm:"embedded"` // явно говорит, что поля нужно встроить
	Upvotes int32  `gorm:"-"`        // исключает поле
}

// equals
type Blog2 struct {
	ID      int64
	Name    string // поле структуры Author
	Email   string // поле структуры Author
	Upvotes int32
}
