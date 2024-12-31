package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID               int
	Name             string
	Balance          float64
	TransactionHistory []string
}

var accounts []Account

func findAccountByID(id int) (*Account, error) {
	for i, acc := range accounts {
		if acc.ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}

	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

func withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

func viewTransactionHistory(accountID int) error {
	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}

	fmt.Printf("Transaction history for account %d (%s):\n", account.ID, account.Name)
	for _, transaction := range account.TransactionHistory {
		fmt.Println(transaction)
	}
	return nil
}

func main() {
	const (
		MenuDeposit    = 1
		MenuWithdraw   = 2
		MenuViewBalance = 3
		MenuViewHistory = 4
		MenuExit       = 5
	)

	// Initialize sample accounts
	accounts = append(accounts, Account{ID: 1, Name: "Alice", Balance: 1000})
	accounts = append(accounts, Account{ID: 2, Name: "Bob", Balance: 500})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case MenuDeposit:
			fmt.Print("Enter account ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter deposit amount: ")
			var amount float64
			fmt.Scan(&amount)
			if err := deposit(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case MenuWithdraw:
			fmt.Print("Enter account ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter withdrawal amount: ")
			var amount float64
			fmt.Scan(&amount)
			if err := withdraw(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case MenuViewBalance:
			fmt.Print("Enter account ID: ")
			var id int
			fmt.Scan(&id)
			account, err := findAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account Balance for %s: %.2f\n", account.Name, account.Balance)
			}

		case MenuViewHistory:
			fmt.Print("Enter account ID: ")
			var id int
			fmt.Scan(&id)
			if err := viewTransactionHistory(id); err != nil {
				fmt.Println("Error:", err)
			}

		case MenuExit:
			fmt.Println("Exiting the system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
