package agrqueries

type Queries struct {
	Items []Query
}

func (self *Queries) Find(query Query) int {
	for key, val := range self.Items {
		if val.Hash == query.Hash {
			return key
		}
	}
	return -1
}

func (self Queries) Len() int {
	return len(self.Items)
}

func (self Queries) Swap(i, j int) {
	self.Items[i], self.Items[j] = self.Items[j], self.Items[i]
}
