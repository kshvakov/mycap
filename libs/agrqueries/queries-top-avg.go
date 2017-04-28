package agrqueries

import "sort"

type QueriesTopByAvg []Query

const MAX_BY_AVG = 30

func (self *QueriesTopByAvg) SortAsc() {
	sort.Sort(self)
}

func (self *QueriesTopByAvg) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *QueriesTopByAvg) Add(q Query) {
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

func (self QueriesTopByAvg) Len() int {
	return len(self)
}

func (self QueriesTopByAvg) Less(i, j int) bool {
	return self[i].Avg < self[j].Avg
}

func (self QueriesTopByAvg) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self QueriesTopByAvg) FindByQuery(fnd Query) (res int) {
	for res, q := range self {
		if q.Hash == fnd.Hash {
			return res
		}
	}
	return -1
}
