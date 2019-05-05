# CSV Reader [![Build Status](https://travis-ci.org/illuminasy/bankstatementreader.svg?branch=master)](https://travis-ci.org/illuminasy/bankstatementreader) [![GoDoc](https://godoc.org/github.com/illuminasy/bankstatementreader?status.svg)](https://godoc.org/github.com/illuminasy/bankstatementreader) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/illuminasy/bankstatementreader/blob/master/LICENSE.md)

 Simple Golang interface to extract bank transaction data.
 Currently supports ANZ and ING transaction csv files.
 
# Usage

Get the library:

    $ go get -v github.com/illuminasy/bankstatementreader

```go
package somepackge

import "github.com/illuminasy/bankstatementreader"

func someFunction() {
    transactions, err := bankstatementreader.ExtractAnzTransactions("some_file_path.csv")
    // or
    transactions, err := bankstatementreader.ExtractIngTransactions("some_file_path.csv")

    for _, transaction := range transactions {
        fmt.Println(transaction.Date)
        fmt.Println(transaction.Desc)
        fmt.Println(transaction.Amount)
        fmt.Println(transaction.Type)
        fmt.Println(transaction.Balance)
    }
}
```