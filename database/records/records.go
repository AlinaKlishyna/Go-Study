package main

import (
	"fmt"

	"errors"

	"github.com/google/uuid"
	"gorm.io/driver/postgres" // go get -u gorm.io/driver/postgres - драйвер для Postgres
	"gorm.io/gorm"
)

type User struct {
	ID       int64     `gorm:"primaryKey"`
	UUID     uuid.UUID `gorm:"type:uuid"` // `gorm:"type:uuid"` - Создавай в таблице это поле с типом UUID (а не TEXT, VARCHAR или BYTEA)
	Role     string
	Name     string
	LastName string
	Age      int
}

// Hook - ищет методы с определёнными именами: BeforeCreate, AfterCreate, BeforeSave, AfterFind, AfterSave
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()

	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
}

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/postgres"), &gorm.Config{ // Подключаемся к базе данных Postgres
		CreateBatchSize: 100, // “Если ты вставляешь много записей за раз — дели их на пакеты размером не больше N штук.”
	})
	if err != nil {
		panic("failed to connect database")
	}

	// ✅ Автоматически создаёт таблицу, если её нет
	db.AutoMigrate(&User{})

	// selectOmitFields(db)
	// batchInsert(db)
	createFromMap(db)
}

func selectOmitFields(db *gorm.DB) {
	user := User{
		Name:     "Larisa",
		LastName: "Hudakova",
		Age:      35,
	}

	result := db.Create(&user) //  - Создай новую запись в таблице
	fmt.Println(user.ID)
	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count
	fmt.Println(result.RowsAffected)

	// Select — “бери только это”,
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// // INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("Larisa", 35, "2020-07-04 11:05:21.775")

	// Omit — “всё, кроме этого”.
	// db.Omit("Age").Create(&user)
}

func batchInsert(db *gorm.DB) {
	users := []User{
		{Name: "Olga", LastName: "Pocrovka", Age: 23},
		{Name: "Nonna", LastName: "Vovk", Age: 55},
		{Name: "Pavel", LastName: "Rozetka", Age: 32},
	}

	db.CreateInBatches(users, 100) // // batch size 100
}

func createFromMap(db *gorm.DB) {
	result := db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "Kukushka",
			"Age": 23},
		{"Name": "Vasilisa",
			"Age": 34},
	})

	if result.Error != nil {
		fmt.Println("Ошибка вставки:", result.Error)
	} else {
		fmt.Println("Успешно вставлено:", result.RowsAffected)
	}
}
