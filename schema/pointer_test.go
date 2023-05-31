package schema_test

func pointer[T any](in T) *T {
	return &in
}
