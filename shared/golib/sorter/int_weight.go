package sorter

type IntWeight struct {
	Id     int
	Weight int
}

//sort descentigd based on weight
type IntWeightSlice []IntWeight

func (s IntWeightSlice) Len() int {
	return len(s)
}

func (s IntWeightSlice) Less(i, j int) bool {
	return s[i].Weight > s[j].Weight
}

func (s IntWeightSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
