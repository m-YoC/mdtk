package lib


type VarSlice[T comparable] struct {
	Elem T
}

func Var[T comparable](e T) VarSlice[T] {
	return VarSlice[T]{e}
}

func (s VarSlice[T]) IsContainedIn(arr []T) bool {
	for _, d := range arr {
		if d == s.Elem { return true }
	}
	return false
}



