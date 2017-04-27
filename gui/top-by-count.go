package gui

import (
	"sort"
)

const MAX_BY_CNT = 30

type TopByCountQueries []Query

func (self *TopByCountQueries) SortAsc() {
	sort.Sort(self)
}

func (self *TopByCountQueries) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *TopByCountQueries) Add(q Query) {
	if i := self.FindByQuery(q); i != -1 {
		(*self)[i].Count = q.Count
	} else if len(*self) < MAX_BY_CNT {
		(*self) = append((*self), q)
	} else {
		self.SortAsc()
		if (*self)[0].Count < q.Count {
			(*self)[0] = q
		}
	}
}

func (self TopByCountQueries) Len() int {
	return len(self)
}

func (self TopByCountQueries) Less(i, j int) bool {
	return self[i].Count < self[j].Count
}

func (self TopByCountQueries) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self TopByCountQueries) FindByQuery(fnd Query) (res int) {
	for res, q := range self {
		if q.Hash == fnd.Hash {
			return res
		}
	}
	return -1
}
