package bankstatementreader

import "reflect"

var ingColumns = []string{"Date", "Description", "Credit", "Debit", "Balance"}

func validateINGColumns(header []string) bool {
	return reflect.DeepEqual(ingColumns, header)
}
