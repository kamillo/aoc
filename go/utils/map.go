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
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func SortMapStringInt(m map[string]int, desc bool) PairList {

	pairs := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = PairStringInt{k, v}
		i++
	}

	if desc {
		sort.Sort(sort.Reverse(pairs))
	} else {
		sort.Sort(pairs)
	}

	return pairs
}

type PairIntInt struct {
	Key   int
	Value int
}

type PairIntIntList []PairIntInt

func (p PairIntIntList) Len() int      { return len(p) }
func (p PairIntIntList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairIntIntList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func SortMapIntInt(m map[int]int, desc bool) PairIntIntList {

	pairs := make(PairIntIntList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = PairIntInt{k, v}
		i++
	}

	if desc {
		sort.Sort(sort.Reverse(pairs))
	} else {
		sort.Sort(pairs)
	}

	return pairs
}

func MakeIndexed(text string) map[rune]int {
	ret := make(map[rune]int, len(text))

	for i, r := range text {
		ret[r] = i
	}

	return ret
}
