package models

import "fmt"

// fieldWithTable return string with table name and field name
//
// Example:
//
//	field := `"Id"`
//	table := "EventTypes"
//	str := fieldWithTable(table, field)
//	print(str)
//
// Output:
//
//	`"EventTypes"."Id"`
func fieldWithTable(table string, field string) string {
	return fmt.Sprintf(`"%s".%s`, table, field)
}
