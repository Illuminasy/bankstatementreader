package bankstatementreader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractANZTransactions(tt *testing.T) {
	d1, _ := time.Parse(DateLayout, "31/10/2018")
	d2, _ := time.Parse(DateLayout, "30/10/2018")
	d3, _ := time.Parse(DateLayout, "22/10/2018")
	d4, _ := time.Parse(DateLayout, "19/10/2018")
	d5, _ := time.Parse(DateLayout, "18/10/2018")
	type args struct {
		filePath string
	}

	type want struct {
		lines []TransactionRow
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Successfully read csv",
			args: args{
				filePath: "test_anz_file.csv",
			},
			want: want{
				lines: []TransactionRow{
					TransactionRow{
						Date:    d1,
						Desc:    "REVERSAL OF ACCOUNT SERVICING FEE MINIMUM $2000 IN DEPOSITS RECEIVED",
						Amount:  5,
						Type:    "credit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d1,
						Desc:    "ACCOUNT SERVICING FEE",
						Amount:  -5,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d1,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 SHOP1 BRISBANE",
						Amount:  -14.6,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d2,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 NEW FARM CINEMAS NEWFARM",
						Amount:  -18.7,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d3,
						Desc:    "ANZ INTERNET BANKING FUNDS TFER TRANSFER 111111 TO RENTAL MANGER",
						Amount:  -900,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d3,
						Desc:    "ANZ M-BANKING PAYMENT TRANSFER 111111 TO MY SAVING ACCOUNT",
						Amount:  -700,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d3,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 SHOP1 BRISBANE",
						Amount:  -51.88,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d4,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 SOME INTERNET PTY LTD BRISBANE",
						Amount:  -59.99,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d4,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 SHOP2 BRISBANE",
						Amount:  -63.8,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d4,
						Desc:    "VISA DEBIT PURCHASE CARD 1111 NETFLIX COM MELBOURNE",
						Amount:  -13.99,
						Type:    "debit",
						Balance: 0,
					},
					TransactionRow{
						Date:    d5,
						Desc:    "TRANSFER FROM COMPANY1 WAGES 11111111",
						Amount:  2707.85,
						Type:    "credit",
						Balance: 0,
					},
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			lines, err := ExtractAnzTransactions(test.args.filePath)
			assert.Equal(t, test.want.lines, lines)
			assert.Equal(t, test.want.err, err)
		})
	}
}

func TestExtractINGTransactions(tt *testing.T) {
	d1, _ := time.Parse(DateLayout, "21/12/2018")
	d2, _ := time.Parse(DateLayout, "07/12/2018")
	d3, _ := time.Parse(DateLayout, "02/12/2018")
	d4, _ := time.Parse(DateLayout, "28/11/2018")
	d5, _ := time.Parse(DateLayout, "16/11/2018")
	d6, _ := time.Parse(DateLayout, "18/07/2016")
	d7, _ := time.Parse(DateLayout, "13/07/2016")
	type args struct {
		filePath string
	}

	type want struct {
		lines []TransactionRow
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Successfully read csv",
			args: args{
				filePath: "test_ing_file.csv",
			},
			want: want{
				lines: []TransactionRow{
					TransactionRow{
						Date:    d1,
						Desc:    "SHOP1 - Direct Debit - Receipt 111111 11111111",
						Amount:  -26.5,
						Type:    "debit",
						Balance: 829.59,
					},
					TransactionRow{
						Date:    d2,
						Desc:    "SHOP1 - Direct Debit - Receipt 111111 11111111",
						Amount:  -26.5,
						Type:    "debit",
						Balance: 876.09,
					},
					TransactionRow{
						Date:    d3,
						Desc:    "SHOP2 - Visa Purchase - Receipt 111111In BRISBANE Date 29 Nov 2018 Card 111111xxxxxx1111",
						Amount:  -20,
						Type:    "debit",
						Balance: 929.09,
					},
					TransactionRow{
						Date:    d4,
						Desc:    "Shop3 - Direct Debit - Receipt 111111 11111111",
						Amount:  -69.25,
						Type:    "debit",
						Balance: 969.09,
					},
					TransactionRow{
						Date:    d5,
						Desc:    "PAYPAL *EBAY AU GST - Visa Purchase - Receipt 111111In 1111111111 Date 14 Nov 2018 Card 111111xxxxxx1111",
						Amount:  -5.6,
						Type:    "debit",
						Balance: 1038.34,
					},
					TransactionRow{
						Date:    d5,
						Desc:    "PAYPAL *SHOP4 - Visa Purchase - Receipt 111111In 1111111111 Date 14 Nov 2018 Card 111111xxxxxx1111",
						Amount:  -56.04,
						Type:    "debit",
						Balance: 1043.94,
					},
					TransactionRow{
						Date:    d6,
						Desc:    "SOMEONE             To my ING Personal  - Receipt 111111",
						Amount:  1000,
						Type:    "credit",
						Balance: 1100,
					},
					TransactionRow{
						Date:    d7,
						Desc:    "MR SOMEONE to me  - Receipt 111111",
						Amount:  100,
						Type:    "credit",
						Balance: 100,
					},
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			lines, err := ExtractIngTransactions(test.args.filePath)
			assert.Equal(t, test.want.lines, lines)
			assert.Equal(t, test.want.err, err)
		})
	}
}

func BenchmarkExtractAnzTransactions(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ExtractAnzTransactions("test_anz_file.csv")
	}
}

func BenchmarkExtractINGTransactions(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ExtractIngTransactions("test_ing_file.csv")
	}
}
