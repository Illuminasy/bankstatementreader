# CSV Reader [![Build Status](https://travis-ci.org/Illuminasy/bankstatementreader.svg?branch=master)](https://travis-ci.org/Illuminasy/bankstatementreader) [![GoDoc](https://godoc.org/github.com/Illuminasy/bankstatementreader?status.svg)](https://godoc.org/github.com/Illuminasy/bankstatementreader) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Illuminasy/bankstatementreader/blob/master/LICENSE.md)

 Simple Golang interface to extract bank transaction data.
 Currently supports ANZ and ING transaction csv files.
 
# Usage

Get the library:

    $ go get -v github.com/Illuminasy/bankstatementreader

```go
package somepackge

import "github.com/Illuminasy/bankstatementreader"

func someFunction() {
    transactions, err := bankstatementreader.ExtractAnzTransactions("some_file_path.csv")
    // or
    transactions, err := bankstatementreader.ExtractIngTransactions("some_file_path.csv")

    for _, transaction := range transactions {
        fmt.Println(transaction.Date)
        fmt.Println(transaction.Desc)
        fmt.Println(transaction.Credit)
        fmt.Println(transaction.Debit)
        fmt.Println(transaction.Balance)
    }
}
```