package main

import "fmt"

func main() {

	dbOperation()
}

func connectDB() {
	fmt.Println("connect to database")
}

func closeDB() {
	fmt.Println("close db connection")
}

func dbOperation() {
	connectDB()
	defer closeDB()
	fmt.Println("operate database")
}