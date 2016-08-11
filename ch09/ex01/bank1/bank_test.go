package bank_test

import (
	"ch09/ex01/bank1"
	"testing"
)

func reset() {
	bank.Withdraw(bank.Balance())
}

func TestWithDraw(t *testing.T) {
	reset()

	successCount := 1000
	failureCount := 100

	bank.Deposit(successCount)

	results := make(chan bool, successCount+failureCount)

	for i := 0; i < successCount+failureCount; i++ {
		go func() {
			results <- bank.Withdraw(1)
		}()
	}

	counts := make(map[bool]int)
	for i := 0; i < successCount+failureCount; i++ {
		counts[<-results]++
	}

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if got, want := counts[true], successCount; got != want {
		t.Errorf("Success Count = %d, want %d", got, want)
	}

	if got, want := counts[false], failureCount; got != want {
		t.Errorf("Failure Count = %d, want %d", got, want)
	}
}

func TestWithDraw2(t *testing.T) {
	for range [100]struct{}{} {
		c := 100

		reset()
		bank.Deposit(c)

		results := make(chan bool, c)

		for i := 0; i < c; i++ {
			go func() {
				results <- bank.Withdraw(1)
			}()
		}

		if bank.Withdraw(c + 1) {
			t.Errorf("This withdraw request must be failed.")
		}

		// wait for all goroutine
		for i := 0; i < c; i++ {
			<-results
		}
	}
}
