package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres" // go get -u gorm.io/driver/postgres - драйвер для Postgres
	"gorm.io/gorm"
)

type Player struct {
	PlayerID uint   `gorm:"primaryKey"`
	Username string `gorm:"size:50"`
	Name     string `gorm:"size:50"`
	Age      int    `gprm:"size:3"`
	Email    string `gorm:"unique"`
}

func main() {
	// CONNECTION TO DATABASE
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/postgres")) // Подключаемся к базе данных Postgres
	if err != nil {
		panic("failed to connect database")
	}

	player := Player{
		Username: "vovk_2014",
		Email:    "vovk12@gmail.com",
	}

	card := CreditCard{
		Bank: "BPER",
	}

	// CREATE TABLE
	db.AutoMigrate(&Player{})     // create table
	db.AutoMigrate(&CreditCard{}) // create table

	// CREATE RECORD
	//db.Create(&player)
	//db.Create(&card)
	fmt.Println(player, card)

	// SEARCH
	// Получение первой записи при сортировке по первичному ключу "id"
	FirstPlayer := Player{}
	// // db.First(&user, 10)   == db.First(&user, "10")
	resultFirst := db.First(&FirstPlayer) // SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println("Ошибок нет:", resultFirst.Error == nil)
	fmt.Println("Найдено строк:", resultFirst.RowsAffected)
	fmt.Printf("Результат FirstPlayer: %v\n\n", FirstPlayer)

	// // Получение одной записи без указания порядка сортировки
	TakedPlayerWithoutSort := Player{}
	resultTakePlayer := db.Take(&TakedPlayerWithoutSort) // // SELECT * FROM users LIMIT 1;
	fmt.Println("Ошибок нет:", resultTakePlayer.Error == nil)
	fmt.Println("Найдено строк:", resultTakePlayer.RowsAffected)
	fmt.Printf("Результат TakedPlayerWithoutSort: %v\n\n", TakedPlayerWithoutSort)

	// Получение последней записи, отсортированные по первичному ключу "id" по убыванию
	LastPLayer := Player{}
	resultLastPlayer := db.Last(&LastPLayer) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("Ошибок нет:", resultLastPlayer.Error == nil)
	fmt.Println("Найдено строк:", resultLastPlayer.RowsAffected)
	fmt.Printf("Результат LastPLayer: %v\n\n", LastPLayer)

	// проверка ошибки на ErrRecordNotFound
	if errors.Is(resultLastPlayer.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Record not found")
	} else {
		fmt.Println("Record found!")
	}

	// Если вы хотите избежать ошибки ErrRecordNotFound, используйте метод Find, например db.Limit(1).Find(&user)
	// метод Find принимает как структурные, так и срезанные данные
	PlayerFound := Player{}
	db.Limit(2).Find(&PlayerFound)
	fmt.Printf("Результат LastPLayer: %v\n\n", PlayerFound)

	Players := []Player{}
	db.Find(&Players, []int{3, 1, 2})                             // SELECT * FROM users WHERE id IN (1,2,3);
	fmt.Printf("Результат Players(id: 1, 2, 3): %v\n\n", Players) // sorted

	CardFind := CreditCard{}
	// Функция uuid.MustParse(...) превращает строку "b3704d32-806c-4138-9d46-47a6cd4cd892" в объект типа uuid.UUID
	db.Find(&CardFind, "uuid=?", uuid.MustParse("b3704d32-806c-4138-9d46-47a6cd4cd892"))
	fmt.Printf("Результат CardFind(uuid: b3704d32-806c-4138-9d46-47a6cd4cd892): %v\n\n", CardFind) // sorted

	// CONDITION
	// Get FIRST matched record
	CardFirstFineco := CreditCard{}
	db.Where("bank=?", "FINECO").First(&CardFirstFineco)                          // SELECT * FROM credit_cards WHERE bank = 'FINECO' ORDER BY id LIMIT 1;
	fmt.Printf("Результат CardFirstFineco(bank=FINECO): %v\n\n", CardFirstFineco) // {71db14a5-74e3-4812-9b96-28282e0df7ed FINECO}

	// Get ALL matched records
	CardsUnicredit := []CreditCard{}
	db.Where("bank=?", "UNICREDIT").Find(&CardsUnicredit)                          // SELECT * FROM credit_cards WHERE bank <> 'UNICREDIT';
	fmt.Printf("Результат CardsUnicredit(bank=UNICREDIT): %v\n\n", CardsUnicredit) // [{b3704d32-806c-4138-9d46-47a6cd4cd892 UNICREDIT} {603fcc0f-b...

	// IN - Дай мне все карточки, у которых банк — любой из этих: FINECO или BPM (перечисление возможных значений)
	CardsItaliani := []CreditCard{}
	db.Where("bank IN ?", []string{"FINECO", "BPM"}).Find(&CardsItaliani)
	fmt.Printf("Результат cardsItaliani(bank IN \"FINECO\", \"BPM\"): %v\n\n", CardsItaliani)

	// LIKE ИМЕЕТ
	// LIKE 'BAN%' - всё, что начинается с BAN	BANCA, BANCO, BANANA
	// LIKE '%CREDIT' - всё, что заканчивается на CREDIT	UNICREDIT, SUPER CREDIT
	// LIKE '%PER%'	- всё, где внутри есть PER	BPER, SUPERBANK
	// LIKE 'B_P%'	- где 2-я буква любая	BPM, BPER
	CardBP := []CreditCard{}
	db.Where("bank LIKE ?", "BP%").Find(&CardBP)
	fmt.Printf("Результат cardBP(ank LIKE BP%%): %v\n\n", CardBP)

	// AND
	playerSpecific := []Player{}
	db.Where("email LIKE ? AND age>=18", "%@gmail.com").Find(&playerSpecific)
	fmt.Printf("Результат playerSpecific(email LIKE \"%%@gmail.com\" AND age>=18): %v\n\n", playerSpecific)

	// Query conditions can be inlined into methods like First and Find in a similar way to Where
	playerMinor := Player{}
	db.First(&playerMinor, "age < ?", 18)
	fmt.Printf("Результат playerMinor(age<18): %v\n\n", playerMinor)

	playersAdult := []Player{}
	db.Find(&playersAdult, "age >= ?", 18)
	fmt.Printf("Результат playersAdult(age>=18): %v\n\n", playersAdult)

	playersAdultNotAnya := []Player{}
	db.Find(&playersAdultNotAnya, "name <> ? AND age >= ?", "Anya", 18)
	fmt.Printf("Результат playersAdultNotAnya(name <> Anya AND age>=18): %v\n\n", playersAdultNotAnya)

	// NOT
	PlayersNotUsername := []Player{}
	db.Not("username = ?", "ranetki").Find(&PlayersNotUsername)
	fmt.Printf("Результат PlayersNotUsername(NOT username = ranetki): %v\n\n", PlayersNotUsername)

	// OR
	PlayerNotRuOrAgeAdult := []Player{}
	db.Where("age > ?", 18).Or("email LIKE ?", "%ru").Find(&PlayerNotRuOrAgeAdult)
	fmt.Printf("Результат PlayerNotRuOrAgeAdult(age > 18 OR LIKE email <> %%ru): %v\n\n", PlayerNotRuOrAgeAdult)

	// SELECT
	players := []Player{}
	db.Select("name", "age").Find(&players)
	fmt.Printf("Результат players(SELECT name * age): %v\n\n", players)

	// Order
	playersNameAsc := []Player{}
	db.Order("name asc, age desc").Find(&playersNameAsc)
	fmt.Printf("Результат playersNameAsc(name asc, age desc): %v\n\n", playersNameAsc)

	// Limit (сколько всего записей) и Offset (сколько пропустить(не показать))
	LimitOffset := Player{}
	db.Limit(3).Offset(1).Find(&LimitOffset)
	fmt.Printf("Результат LimitOffset(Limit(3).Offset(1)): %v\n\n", LimitOffset)
}

type CreditCard struct {
	UUID uuid.UUID `gorm:"primaryKey"` // первичным ключом является строка
	Bank string
}

// Автоматически при создании (через GORM hook)
func (c *CreditCard) BeforeCreate(tx *gorm.DB) (err error) { // при каждом db.Create() UUID сгенерируется автоматически
	c.UUID = uuid.New()
	return
}
