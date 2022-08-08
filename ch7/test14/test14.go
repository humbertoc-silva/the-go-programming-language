package main

import "fmt"

func main() {
	result := sqlQuote(50)
	fmt.Println(result)
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuoteString(x string) string {
	return x
}
