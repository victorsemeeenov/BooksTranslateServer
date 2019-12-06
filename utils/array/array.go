package array

func Map(array []interface{}, mapFunc func(interface{}) interface{}) []interface{} {
	newArray := make([]interface{}, len(array))
	for i, obj := range array {
		newArray[i] = mapFunc(obj)
	}
	return newArray
}

func Filter(array []interface{}, filterFunc func(interface{}) bool) []interface{} {
	var newArray []interface{}
	for _, obj := range array {
		if filterFunc(obj) {
			newArray = append(newArray, obj)
		}
	}
	return newArray
}