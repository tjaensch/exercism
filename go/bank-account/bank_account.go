package account

import "sync"

type Account struct {
	sync.Mutex
	open    bool
	balance int64
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{open: true, balance: amount}
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	balance, ok = a.balance, a.open
	a.Unlock()
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return a.balance, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	if a.balance < 0 {
		return 0, true
	}
	payout = a.balance
	return payout, true
}
