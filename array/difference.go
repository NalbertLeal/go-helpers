package array

func Difference[data comparable](a []data, b []data) []data {
    result := []data{}
    cache := map[data]bool{}

    for _, va := range a {
        _, exists := cache[va]
        if exists {
            result = append(result, va)
            continue
        }

        addToResult := true
        for j := 0; j < len(b); j++ {
            if va == b[j] {
                addToResult = false
                break
            }
        }
        if addToResult {
            cache[va] = true
            result = append(result, va)
        }
    }

    return result
}
