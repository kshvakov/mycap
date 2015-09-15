package gui

import (
	"sort"
)

type TopByCountQueries []TopByCountQuery

func (self *TopByCountQueries) SortAsc() {
	sort.Sort(self)
}

func (self *TopByCountQueries) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *TopByCountQueries) Add(query string, count int) {
	if i := self.FindByQuery(query); i != -1 {
		(*self)[i].Count = count
	} else if len(*self) < 5 {
		(*self) = append((*self), TopByCountQuery{
			Query: query,
			Count: count,
		})
	} else {
		self.SortAsc()
		if (*self)[0].Count < count {
			(*self)[0] = TopByCountQuery{
				Query: query,
				Count: count,
			}
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

func (self TopByCountQueries) FindByQuery(query string) (res int) {
	for res, q := range self {
		if q.Query == query {
			return res
		}
	}
	return -1
}

type TopByCountQuery struct {
	Query string
	Count int
}
