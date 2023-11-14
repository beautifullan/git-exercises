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
	valuesMap := make(map[string]string)
	for _, item := range values {
		valuesMap[item.Key] = item.Val
	}
	for i := range array {
		array[i].Val2 = valuesMap[array[i].Key]
	}
	return array
}

func version1(array []A, values []B) []A {
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
