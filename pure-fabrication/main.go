package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID     string
	Amount float64
}

type LogService struct{}

func (ls *LogService) Log(message string) {
	fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339), message)
}

type BankAccount struct {
	ID           string
	transactions []Transaction
	logService   *LogService
}

func NewBankAccount(id string, logService *LogService) *BankAccount {
	return &BankAccount{ID: id, logService: logService}
}

func (b *BankAccount) DisplayTransactions() {
	b.logService.Log("Printing all transactions")
	for _, transaction := range b.transactions {
		fmt.Printf("ID:%s, Amount: $%.2f\n", transaction.ID, transaction.Amount)
	}
}

func (b *BankAccount) AddTransaction(transaction Transaction) {
	b.transactions = append(b.transactions, transaction)
	b.logService.Log(fmt.Sprintf("Added Transaction ID: %s,Amount: $%2.f", transaction.ID, transaction.Amount))
}

func main() {
	logService := &LogService{}
	account := NewBankAccount("123", logService)
	t1 := Transaction{ID: "1", Amount: 199.0}
	t2 := Transaction{ID: "2", Amount: 299.0}
	account.AddTransaction(t1)
	account.AddTransaction(t2)

}
