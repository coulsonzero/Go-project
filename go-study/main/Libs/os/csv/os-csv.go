package main

import (
	"encoding/csv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

const (
	dbname = "go_study"
)

type Bill struct {
	gorm.Model
	TradeAt        string
	TradeType      string
	Trader         string
	Product        string
	ReceiptPayment string
	Amount         string
	PayMethod      string
	CurStatus      string
	TransactionID  string
	MerchantID     string
	Remarks        string
}

func main() {
	db := SetUpDB()
	defer CloseDB(db)
	ReadCsv("billMonthly.csv")
}

func SetUpDB() *gorm.DB {
	dsn := fmt.Sprintf("root:root@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True", dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connect mysql success!")
	db.Exec("DROP TABLE bills")
	db.AutoMigrate(&Bill{})
	DB = db
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

func WriteCsv(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	// file.WriteString("\xEF\xBB\xBF")	// 写入utf-8 BOM
	w := csv.NewWriter(file)
	_ = w.WriteAll([][]string{
		{"id", "name", "score"},
		{"1", "Barry", "97"},
		{"2", "Shirdon", "89"},
		{"3", "Jack", "92"},
		{"4", "Tom", "78"},
	})
	w.Flush()
}

func ReadCsv(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	// for {
	// 	record, _ := reader.Read()
	// }

	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Failed to read csv file, err: %s", err)
	}
	for _, record := range records {
		// fmt.Printf("%#v\n", record)
		DB.Create(&Bill{
			TradeAt:        record[0],
			TradeType:      record[1],
			Trader:         record[2],
			Product:        record[3],
			ReceiptPayment: record[4],
			Amount:         record[5],
			PayMethod:      record[6],
			CurStatus:      record[7],
			TransactionID:  record[8],
			MerchantID:     record[9],
			Remarks:        record[10],
		})
	}
	log.Printf("Read csv file successfully, the data is updated")
}
