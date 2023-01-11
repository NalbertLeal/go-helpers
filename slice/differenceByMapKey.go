package slice

func DifferenceByMapKey[keyType comparable, valueType comparable](arr1 []map[keyType]valueType, arr2 []map[keyType]valueType, keyComp []keyType) []map[keyType]valueType {
	indexesDontAdd := map[int]bool{}
	for _, key := range keyComp {
		for i, v1 := range arr1 {
			for _, v2 := range arr2 {
				if v1[key] == v2[key] {
					indexesDontAdd[i] = true
					break
				}
			}
		}
	}

	result := []map[keyType]valueType{}
	for i, v := range arr1 {
		_, exists := indexesDontAdd[i]
		if !exists {
			result = append(result, v)
		}
	}

	return result
}
