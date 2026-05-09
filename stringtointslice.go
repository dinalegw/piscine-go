package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	res := make([]int, 0, len(str))
	for _, r := range str {
		res = append(res, int(r))
	}
	return res
}
