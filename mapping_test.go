package exercises

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	array := make([]A, 0)
	array = append(array, A{Key: "1", Val1: "1"})
	array = append(array, A{Key: "2", Val1: "2"})
	array = append(array, A{Key: "3", Val1: "3"})
	array = append(array, A{Key: "4", Val1: "4"})
	array = append(array, A{Key: "5", Val1: "5"})

	values := make([]B, 0)
	values = append(values, B{Key: "1", Val: "11"})
	values = append(values, B{Key: "2", Val: "22"})
	values = append(values, B{Key: "3", Val: "33"})
	values = append(values, B{Key: "4", Val: "44"})
	values = append(values, B{Key: "5", Val: "55"})

	result := AssignValues(array, values)

	for _, value := range result {
		key := value.Key
		expectedValue := key + key
		if value.Val2 != expectedValue {
			t.Fatalf(`Value("") = %s, want "%s"`, value.Val2, expectedValue)
		}
	}
}
