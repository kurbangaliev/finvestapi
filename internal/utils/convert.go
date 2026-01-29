package utils

import "fmt"

func ConvertToInt(v interface{}) int {
	switch val := v.(type) {
	case int:
		fmt.Printf("Value %v is a regular int\n", val)
	case int32:
		fmt.Printf("Value %v is an int32, converting to int: %d\n", val, int(val))
	case int64:
		fmt.Printf("Value %v is an int64, converting to int: %d\n", val, int(val))
	case float64:
		fmt.Printf("Value %v is a float64, converting to int: %d\n", val, int(val))
	default:
		fmt.Printf("Unsupported type: %T\n", v)
	}
	return v.(int)
}
