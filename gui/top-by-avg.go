package gui

import (
	"sort"
	"time"
)

type TopByAvgQueries []TopByAvgQuery

func (self *TopByAvgQueries) SortAsc() {
	sort.Sort(self)
}

func (self *TopByAvgQueries) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *TopByAvgQueries) Add(query string, avg time.Duration) {
	if i := self.FindByQuery(query); i != -1 {
		(*self)[i].Avg = avg
	} else if len(*self) < 5 {
		(*self) = append((*self), TopByAvgQuery{
			Query: query,
			Avg:   avg,
		})
	} else {
		self.SortAsc()
		if (*self)[0].Avg < avg {
			(*self)[0] = TopByAvgQuery{
				Query: query,
				Avg:   avg,
			}
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

func (self TopByAvgQueries) FindByQuery(query string) (res int) {
	for res, q := range self {
		if q.Query == query {
			return res
		}
	}
	return -1
}

type TopByAvgQuery struct {
	Query string
	Avg   time.Duration
}
