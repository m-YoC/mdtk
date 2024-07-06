package lib

type ArraySlice[T comparable] struct {
	A []T
}

func Slice[T comparable](a []T) ArraySlice[T] {
	return ArraySlice[T]{a}
}

func (s ArraySlice[T]) Have(e T) bool {
	for _, d := range s.A {
		if d == e { return true }
	}
	return false
}

// Get first data
func (s ArraySlice[T]) HaveFunc(f func(T)bool) (T, bool) {
	for _, d := range s.A {
		if f(d) { return d, true }
	}
	var zero T
	return zero, false
}


// ------------------------------------

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



