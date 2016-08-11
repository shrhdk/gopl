// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // send amount to withdraw
var results = make(chan bool)  // receive result of withdraw

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-results
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if balance < amount {
				results <- false
			} else {
				balance -= amount
				results <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
