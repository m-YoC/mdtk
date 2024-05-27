package lib


type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
    *s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(*s) == 0 {
        var zero T
        return zero, false
    }
    
    top := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]

    return top, true
}

func (s Stack[T]) Size() int {
    return len(s)
}

func (s Stack[T]) Top() (T, bool) {
    if len(s) == 0 {
        var zero T
        return zero, false
    }

    return (s)[len(s)-1], true
}
