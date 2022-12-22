package array

func DiffenrenceByIterrator[dataIn any, dataComp comparable](arr1 []dataIn, arr2 []dataIn, iteratee func(dataIn) dataComp) []dataIn {
	result := []dataIn{}

	for _, v := range arr1 {
		vComp := iteratee(v)
		canAdd := true
		for _, r := range arr2 {
			rComp := iteratee(r)
			if rComp == vComp {
				canAdd = false
				break
			}
		}
		if canAdd {
			result = append(result, v)
		}
	}

	return result
}

func DiffenrenceByMapKey[keyType comparable, valueType comparable](arr1 []map[keyType]valueType, arr2 []map[keyType]valueType, keyComp []keyType) []map[keyType]valueType {
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
