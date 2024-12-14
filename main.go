package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/payjp/payjp-go/v1"
)


func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	secretKey := os.Getenv("PAYJP_SECRET_TEST_KEY")
	customerId := os.Getenv("PAYJP_CUSTOMER_ID")
	
	if secretKey == "" {
		log.Fatal("PAYJP_SECRET_KEY is not set in .env file")
	}
	
	// PAY.JPクライアントの初期化
	pay := payjp.New(secretKey, nil)
	fmt.Println("PAY.JP client initialized successfully")
	// 顧客情報の取得
	customer, err := pay.Customer.Retrieve(customerId)
	if err != nil {
		log.Fatalf("Error retrieving customer: %v", err)
	}
	fmt.Printf("Customer: %+v\n", customer)
}
