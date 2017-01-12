package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		Withdraw(200)
		fmt.Println("Current balance is: ", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(50)
		Withdraw(50)
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, expect := Balance(), 100; got != expect {
		t.Errorf("Got Balance of %d, expect %d", got, expect)
	}
}

func TestWithdrawSuccessfully(t *testing.T) {
	SetBalance(100)
	b1 := Balance()

	if b1 != 100 {
		t.Errorf("balance = %d, expect %d", b1, 100)
	}

	if ok := Withdraw(50); !ok {
		t.Errorf("ok = false, want true. balance = %d", Balance())
	}
	expect := b1 - 50
	if b2 := Balance(); b2 != expect {
		t.Errorf("balance = %d, expect %d", b2, expect)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	SetBalance(100)

	b1 := Balance()

	if b1 != 100 {
		t.Errorf("balance = %d, expect %d", b1, 100)
	}

	ok := Withdraw(b1 + 1)
	b2 := Balance()
	if ok {
		t.Errorf("ok = true, want false. balance = %d", b2)
	}
	if b2 != b1 {
		t.Errorf("balance = %d, want %d", b2, b1)
	}
}
