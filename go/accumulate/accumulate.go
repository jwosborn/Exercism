package accumulate

func Accumulate(list []string, f func(string) string) []string {
	var newList []string
	for _, s := range list {
		newList = append(newList, f(s))
	}
	return newList
}
