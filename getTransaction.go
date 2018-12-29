package bankstatementreader

import (
	"fmt"
	"strconv"
	"time"
)

func getANZTransaction(line []string, firstLine bool) (TransactionRow, error) {
	var credit float64
	var debit float64

	d, err := time.Parse(DateLayout, line[0])
	if err != nil {
		return TransactionRow{}, err
	}

	a, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return TransactionRow{}, err
	}

	if a > 0 {
		credit = a
	} else {
		debit = a
	}

	atr := TransactionRow{
		Date:   d,
		Desc:   line[2],
		Credit: credit,
		Debit:  debit,
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

	cr := 0.00
	if len(line[2]) > 0 {
		cr, err = strconv.ParseFloat(line[2], 64)
		if err != nil {
			return TransactionRow{}, err
		}
	}

	db := 0.00
	if len(line[3]) > 0 {
		db, err = strconv.ParseFloat(line[3], 64)
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
		Credit:  cr,
		Debit:   db,
		Balance: bl,
	}

	return tr, nil
}
