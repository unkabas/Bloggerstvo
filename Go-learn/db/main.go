package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	ID   uint   `gorm:"primaryKey"`
	User string `gorm:"column:my_user"`
	Mess string
}

func main() {
	dsn := "host=localhost user=unkabas password=localhost2002 dbname=test port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Message{})

	// Пример добавления данных
	db.Create(&Message{User: "moverq", Mess: "Hello!"})
	db.Create(&Message{User: "gor", Mess: "Ass!"})

	// Поиск по user
	var chat []Message
	res := db.Find(&chat, "my_user = ?", "moverq")
	count := res.RowsAffected
	if res.Error != nil {
		fmt.Println("Ошибка:", res.Error)
		return
	}
	fmt.Println("Количество записей с user 'moverq':", count)
}
