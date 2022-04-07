package db

import (
	"fmt"
	"gopgsql/models"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := getEnv("PG_USER", "postgres")
	password := getEnv("PG_PASSWORD", "postgres")
	host := getEnv("PG_HOST", "project_gopgsql_postgresql")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "postgres")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	// defer CloseDB()
	log.Println("Database connected")

	if !db.HasTable(&models.Task{}) {
		err := db.CreateTable(&models.Task{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Task{})

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

// サンプルデータを入力する
func InsertRecord() error {
	var err error

	u1 := uuid.Must(uuid.NewV4(), err)
	if err != nil {
		fmt.Printf("u1 %s\n", err)
		return err
	}
	// u2 := uuid.Must(uuid.NewV4(), err)
	// if err != nil {
	// 	fmt.Printf("u2 %s\n", err)
	// 	return err
	// }

	// var tasks = []models.Task{
	// 	{ID: u1, Title: "Tanaka", CreatedAt: time.Now(), UpdatedAt: time.Now(), Completed: true},
	// 	{ID: u2, Title: "Sakamoto", CreatedAt: time.Now(), UpdatedAt: time.Now(), Completed: false},
	// }
	task := models.Task{ID: u1, Title: "Tanaka", CreatedAt: time.Now(), UpdatedAt: time.Now(), Completed: true}

	db = GetDB()
	// defer CloseDB()

	db.Create(&task)

	// for _, task := range tasks {
	// 	fmt.Println("[task]", task.ID)
	// }

	return nil
}
