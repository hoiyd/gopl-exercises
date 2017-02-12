package bank

type Withdrawl struct {
	amount int
	ok     chan bool
}

var deposits = make(chan int)         // send amount to deposit
var balances = make(chan int)         // receive balance
var withdrawls = make(chan Withdrawl) // withdraw amount
var clear = make(chan struct{})

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ok := make(chan bool)
	withdrawls <- Withdrawl{amount, ok}
	return <-ok
}
func Clear() { clear <- struct{}{} }

func SetBalance(amount int) {
	Clear()
	Deposit(amount)
}

func teller() {
	var balance int //balance is confined to teller goroutine, initialized to 0.

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case <-clear:
			balance = 0
		case w := <-withdrawls:
			if balance >= w.amount {
				balance -= w.amount
				w.ok <- true
			} else {
				w.ok <- false
			}
		}
	}
}

func init() {
	go teller()
}
