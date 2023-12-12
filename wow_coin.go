package wow_coin

import "fmt"

type Wallet struct {
	goldCoins   int
	silverCoins int
	copperCoins int
}

func (wallet Wallet) GetValues() string {
	return fmt.Sprintf(
		"%dG%dS%dC",
		wallet.goldCoins,
		wallet.silverCoins,
		wallet.copperCoins)
}

func (wallet *Wallet) AddCopperCoins(amount int) {
	wallet.copperCoins += amount

	if wallet.copperCoins > 99 {
		extraSilver := wallet.copperCoins / 100
		wallet.AddSilverCoins(extraSilver)
		wallet.copperCoins -= extraSilver * 100
	}
}

func (wallet *Wallet) AddSilverCoins(amount int) {
	wallet.silverCoins += amount

	if wallet.silverCoins > 99 {
		extraGold := wallet.silverCoins / 100
		wallet.AddGoldCoins(extraGold)
		wallet.silverCoins -= extraGold * 100
	}
}

func (wallet *Wallet) AddGoldCoins(amount int) {
	wallet.goldCoins += amount
}
