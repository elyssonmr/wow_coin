package wow_coin

import (
	"testing"
)

func TestGetValues(t *testing.T) {
	expected := "1G2S3C"
	wallet := Wallet{1, 2, 3}

	result := wallet.GetValues()

	if result != expected {
		t.Fatalf("Expected: %s Result: %s", expected, result)
	}
}

func TestGetBigValues(t *testing.T) {
	expected := "100G35S99C"
	wallet := Wallet{100, 35, 99}

	result := wallet.GetValues()

	if result != expected {
		t.Fatalf("Expected: %s Result: %s", expected, result)
	}
}

func TestAddCopper(t *testing.T) {
	expected := "0G0S45C"
	wallet := Wallet{0, 0, 0}

	wallet.AddCopperCoins(45)

	result := wallet.GetValues()

	if result != expected {
		t.Fatalf("Expected: %s Result: %s", expected, result)
	}
}

func TestAddCopperConvertingToSilver(t *testing.T) {
	expected := "0G1S5C"
	amount := 45
	wallet := Wallet{0, 0, 60}

	wallet.AddCopperCoins(amount)

	if wallet.GetValues() != expected {
		t.Fatalf("Expected: %s Result: %s", expected, wallet.GetValues())
	}
}

func TestAddSilver(t *testing.T) {
	expected := "0G1S0C"
	amount := 1
	wallet := Wallet{0, 0, 0}

	wallet.AddSilverCoins(amount)

	result := wallet.GetValues()

	if result != expected {
		t.Fatalf("Expected: %s Result: %s", expected, result)
	}
}

func TestAddSilverConvertingToGold(t *testing.T) {
	expected := "1G5S0C"
	amount := 45
	wallet := Wallet{0, 60, 0}

	wallet.AddSilverCoins(amount)

	if wallet.GetValues() != expected {
		t.Fatalf("Expected: %s Result: %s", expected, wallet.GetValues())
	}
}

func TestAddGoldCoins(t *testing.T) {
	expected := "1G0S0C"
	amount := 1
	wallet := Wallet{0, 0, 0}

	wallet.AddGoldCoins(amount)

	if wallet.GetValues() != expected {
		t.Fatalf("Expected: %s Result: %s", expected, wallet.GetValues())
	}
}

func TestAddBigGoldCoins(t *testing.T) {
	expected := "150G0S0C"
	amount := 70
	wallet := Wallet{80, 0, 0}

	wallet.AddGoldCoins(amount)

	if wallet.GetValues() != expected {
		t.Fatalf("Expected: %s Result: %s", expected, wallet.GetValues())
	}
}

func TestAddMultipleConvertions(t *testing.T) {
	expected := "1G0S49C"
	amount := 50
	wallet := Wallet{0, 99, 99}

	wallet.AddCopperCoins(amount)

	if wallet.GetValues() != expected {
		t.Fatalf("Expected: %s Result: %s", expected, wallet.GetValues())
	}
}
