package protomockstring

import (
	"fmt"
)

// Q quotas the passed string
func Q(obj any) string {
	return "\"" + S(obj) + "\""
}

// S transforms an object into a string
func S(obj any) string {
	return fmt.Sprint(obj)
}
