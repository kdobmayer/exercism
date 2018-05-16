package flatten

func Flatten(a interface{}) []interface{} {
	flattened := []interface{}{}
	for _, v := range a.([]interface{}) {
		if _, ok := v.([]interface{}); ok {
			flattened = append(flattened, Flatten(v)...)
		}
		if _, ok := v.(int); ok {
			flattened = append(flattened, v)
		}
	}
	return flattened
}