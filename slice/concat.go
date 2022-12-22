package slice

func Concat[data any](a []data, b []data) []data {
    result := make([]data, len(a)+len(b))
    
    for i, v := range a {
        result[i] = v
    }

    for i, v := range b {
        result[len(a)+i] = v
    }

    return result
}
