// Package bankstatementreader ...
// Maintainer Illuminasy
// Package to retrieve data from bank statements.
package bankstatementreader

import (
	"fmt"
	"time"

	"github.com/illuminasy/csvreader"
)

const (
	// DateLayout ...
	DateLayout = "02/01/2006"

	// ANZ ...
	ANZ = "anz"

	// ING ...
	ING = "ing"
)

// TransactionRow transaction record
type TransactionRow struct {
	Date    time.Time
	Desc    string
	Credit  float64
	Debit   float64
	Balance float64
}

// Transactions array of transactions
type Transactions []TransactionRow

// ExtractAnzTransactions extracts array of lines from ANZ transaction file
func ExtractAnzTransactions(fileName string) ([]TransactionRow, error) {
	return extractTransactions(ANZ, fileName)
}

// ExtractIngTransactions extracts array of lines from ING transaction file
func ExtractIngTransactions(fileName string) ([]TransactionRow, error) {
	return extractTransactions(ING, fileName)
}

func extractTransactions(bank string, fileName string) (Transactions, error) {
	lines, err := csvreader.ExtractContents(fileName)
	if err != nil {
		return Transactions{}, err
	}

	var t Transactions

	// Loop through lines & turn into object
	for i, line := range lines {
		firstLine := false
		if i == 0 {
			firstLine = true
		}

		var tr TransactionRow
		switch bank {
		case ANZ:
			tr, err = getANZTransaction(line, firstLine)
		case ING:
			tr, err = getINGTransaction(line, firstLine)
		default:
			return Transactions{}, fmt.Errorf("Unkown Bank Statement")
		}

		if err != nil {
			return Transactions{}, err
		}

		if tr != (TransactionRow{}) {
			t = append(t, tr)
		}
	}

	return t, nil
}
