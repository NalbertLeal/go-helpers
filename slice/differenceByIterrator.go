package slice

func DifferenceByIterrator[dataIn any, dataComp comparable](arr1 []dataIn, arr2 []dataIn, iteratee func(dataIn) dataComp) []dataIn {
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