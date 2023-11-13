package exercises

type A struct {
	Key  string
	Val1 string
	Val2 string
}

type B struct {
	Key string
	Val string
}

func AssignValues(array []A, values []B) []A {
	//TODO:: you need to assign B.Val to A.Val2 for the same key value
	for i, _ := range array {
		for _, item := range values {
			if array[i].Key == item.Key {
				array[i].Val2 = item.Val
				break
			}
		}

	}
	return array
}
