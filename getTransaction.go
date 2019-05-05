package bankstatementreader

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// Credit ...
	Credit = "credit"

	// Debit ...
	Debit = "debit"
)

func getANZTransaction(line []string, firstLine bool) (TransactionRow, error) {
	var amount float64
	var txnType string

	d, err := time.Parse(DateLayout, line[0])
	if err != nil {
		return TransactionRow{}, err
	}

	amount, err = strconv.ParseFloat(line[1], 64)
	if err != nil {
		return TransactionRow{}, err
	}

	if amount > 0 {
		txnType = Credit
	} else {
		txnType = Debit
	}

	atr := TransactionRow{
		Date:   d,
		Desc:   line[2],
		Amount: amount,
		Type:   txnType,
	}

	return atr, nil
}

func getINGTransaction(line []string, firstLine bool) (TransactionRow, error) {
	if firstLine && !validateINGColumns(line) {
		return TransactionRow{}, fmt.Errorf("Invalid Statement file, columns does not match the expected format")
	} else if firstLine {
		return TransactionRow{}, nil
	}

	d, err := time.Parse(DateLayout, line[0])
	if err != nil {
		return TransactionRow{}, err
	}

	amount := 0.00
	txnType := Credit
	if len(line[2]) > 0 {
		amount, err = strconv.ParseFloat(line[2], 64)
		txnType = Credit
		if err != nil {
			return TransactionRow{}, err
		}
	}

	if len(line[3]) > 0 {
		amount, err = strconv.ParseFloat(line[3], 64)
		txnType = Debit
		if err != nil {
			return TransactionRow{}, err
		}
	}

	bl, err := strconv.ParseFloat(line[4], 64)
	if err != nil {
		return TransactionRow{}, err
	}

	tr := TransactionRow{
		Date:    d,
		Desc:    line[1],
		Amount:  amount,
		Type:    txnType,
		Balance: bl,
	}

	return tr, nil
}
