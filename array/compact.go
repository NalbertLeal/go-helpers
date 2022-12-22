package array

func Compact[data comparable](arr []data) []data {
    var zero data
    result := []data{}

    for i := 0; i < len(arr); i++ {
        if arr[i] != zero {
            result = append(result, arr[i])
        }
    }

    return result
}
