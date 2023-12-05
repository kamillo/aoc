package utils

import "sort"

type PairStringInt struct {
	Key   string
	Value int
}

type PairList []PairStringInt

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}
	return p[i].Value < p[j].Value
}

func SortMapStringInt(m map[string]int) PairList {

	pairs := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = PairStringInt{k, v}
		i++
	}

	sort.Sort(pairs)

	return pairs
}
