package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateDb() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	dbConectionInfo := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbConectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `words`(" +
		"Id int NOT NULL AUTO_INCREMENT," +
		"Word varchar(255) NOT NULL," +
		"Meaning varchar(255) NOT NULL," +
		"UserId int NOT NULL," +
		"CreatedAt datetime," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users`(" +
		"Id int NOT NULL AUTO_INCREMENT," +
		"Name varchar(255)," +
		"Email varchar(255)," +
		"PassWord varchar(255)," +
		"CreatedAt datetime," +
		"UpdatedAt datetime," +
		"PRIMARY KEY(id))")

	if err != nil {
		log.Println(err)
	}
	defer db.Close()

}
