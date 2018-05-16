package account

import "sync"

type Account struct {
	balance int
	mu      sync.Mutex
	isOpen  bool
}

func Open(initialDeposit int) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := Account{balance: initialDeposit, isOpen: true}
	return &a
}

func (a *Account) Close() (int, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen {
		return 0, false
	}
	a.isOpen = false
	return a.balance, true
}

func (a *Account) Balance() (int, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int) (int, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}