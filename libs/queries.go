package libs

type Queries []Query

func (self *Queries) Find(query Query) int {
	for key, val := range *self {
		if val.ID == query.ID {
			return key
		}
	}
	return -1
}

func (self *Queries) AddIfNotExists(query Query) {
	if index := self.Find(query); index == -1 {
		(*self) = append(*self, query)
	} else {
		(*self)[index] = query
	}
}
