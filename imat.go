package main

type Imat int

const (
	Wd Imat = iota
	Stn
	Fe
	C
	Cu
	Sn
	Ni
)

type ISet [7]int

func (i *ISet) Pay(cost_set ISet) bool {
	for idx := 0; idx < 7; idx++ {
		i[idx] -= cost_set[idx]
		if i[idx] < 0 {
			i[idx] += cost_set[idx]
			return false
		}
	}

	return true
}

func (s *ISet) Add(add_set ISet) {
	for i := 0; i < 7; i++ {
		s[i] += add_set[i]
	}
}
