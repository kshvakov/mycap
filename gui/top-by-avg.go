package gui

import (
	"sort"
)

const MAX_BY_AVG = 30

type TopByAvgQueries []Query

func (self *TopByAvgQueries) SortAsc() {
	sort.Sort(self)
}

func (self *TopByAvgQueries) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *TopByAvgQueries) Add(q Query) {
	if i := self.FindByQuery(q); i != -1 {
		(*self)[i].Avg = q.Avg
	} else if len(*self) < MAX_BY_AVG {
		(*self) = append((*self), q)
	} else {
		self.SortAsc()
		if (*self)[0].Avg < q.Avg {
			(*self)[0] = q
		}
	}
}

func (self TopByAvgQueries) Len() int {
	return len(self)
}

func (self TopByAvgQueries) Less(i, j int) bool {
	return self[i].Avg < self[j].Avg
}

func (self TopByAvgQueries) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self TopByAvgQueries) FindByQuery(fnd Query) (res int) {
	for res, q := range self {
		if q.Hash == fnd.Hash {
			return res
		}
	}
	return -1
}
