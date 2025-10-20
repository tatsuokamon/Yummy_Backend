package ytbdtc

type Stringifiable interface {
	String()string
}

func listToString[T Stringifiable](list []T)string{
	result := ""
	for _, t := range list {
		result += t.String()
	}
	return result
}
