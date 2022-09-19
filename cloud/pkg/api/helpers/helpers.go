package helpers

func ToString(x *string) string {
	if x == nil {
		return ""
	}
	return *x
}

func StringPtr(x string) *string {
	if x == "" {
		return nil
	}
	return &x
}
