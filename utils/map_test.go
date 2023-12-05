package utils

import "testing"

func TestSortMapStringInt(t *testing.T) {
	m := map[string]int{
		"Australia":  24982688,
		"Qatar":      2781677,
		"Wales":      3139000,
		"Burundi":    11175378,
		"Guinea":     12414318,
		"Niger":      22442948,
		"Brazil":     209469333,
		"Malta":      484630,
		"Peru":       31989256,
		"Yemen":      28498687,
		"Ireland":    4867309,
		"Kenya":      51393010,
		"Montserrat": 5900,
		"Cuba":       11338138,
		"Nicaragua":  6465513,
		"Jordan":     9956011,
		"Gabon":      2119275,
	}

	want := PairList{
		PairStringInt{"Montserrat", 5900},
		PairStringInt{"Malta", 484630},
		PairStringInt{"Gabon", 2119275},
		PairStringInt{"Qatar", 2781677},
		PairStringInt{"Wales", 3139000},
		PairStringInt{"Ireland", 4867309},
		PairStringInt{"Nicaragua", 6465513},
		PairStringInt{"Jordan", 9956011},
		PairStringInt{"Burundi", 11175378},
		PairStringInt{"Cuba", 11338138},
		PairStringInt{"Guinea", 12414318},
		PairStringInt{"Niger", 22442948},
		PairStringInt{"Australia", 24982688},
		PairStringInt{"Yemen", 28498687},
		PairStringInt{"Peru", 31989256},
		PairStringInt{"Kenya", 51393010},
		PairStringInt{"Brazil", 209469333},
	}
	got := SortMapStringInt(m)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got = %v, want %v", got, want)
			break
		}
	}
}
