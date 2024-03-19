package main

import "testing"

// func TestWallet(t *testing.T) {

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})
// }

func TestWallet(t *testing.T) {

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}