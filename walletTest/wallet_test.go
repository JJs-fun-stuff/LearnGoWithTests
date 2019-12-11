package main

import (
	"testing"
	// "fmt"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	// assertError := func(t *testing.T, err error){
	// 	t.Helper()
	// 	if err == nil{
	// 		t.Error("wanted an error but didn't get one")
	// 	}
	// }

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		// got.Error() return error message as a string
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error){
		t.Helper()
		if got != nil{
			t.Fatal("got an error but didn't want one")
			
		}
	}
	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		//for checking the address
		// fmt.Printf("address of balace in test is %v \n", &wallet.balance)

		//2nd refactorization no need these codes below
		//becuz we use assertBalance
		// got := wallet.Balance()
		// want := Bitcoin(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
