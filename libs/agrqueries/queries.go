package agrqueries

type Queries struct {
	Queries map[string]Query
	TopAvg  QueriesTopByAvg
	TopCnt  QueriesTopByCount
}

func (self *Queries) Add(query Query) {
	if self.Queries == nil {
		self.Queries = make(map[string]Query)
	}

	if len(query.Query) < 1 {
		return
	}

	if exists, ok := self.Queries[query.GetHash()]; ok {
		exists.Avg = (exists.Avg + query.Avg) / 2
		exists.Count += query.Count

		if exists.Min > query.Min {
			exists.Min = query.Min
		}

		if exists.Max < query.Max {
			exists.Max = query.Max
		}

		self.Queries[query.GetHash()] = exists
		self.TopAvg.Add(exists)
		self.TopCnt.Add(exists)
	} else {
		self.Queries[query.GetHash()] = query
		self.TopAvg.Add(query)
		self.TopCnt.Add(query)
	}
}
