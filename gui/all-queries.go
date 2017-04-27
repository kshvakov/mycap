package gui

import (
	"crypto/md5"
	"io"
	"time"
)

type Queries map[string]Query

func (self *Queries) Add(from string, to string, query string, start time.Time, duration time.Duration) {
	if len(query) < 1 {
		// fucked empty queryies :|
		return
	}

	q := Query{
		Query: query,
		From:  from,
		To:    to,
		Avg:   duration,
		Min:   duration,
		Max:   duration,
		Count: 1,
	}
	q.CalcHash()

	if exists, ok := (*self)[q.GetHash()]; ok {
		exists.Avg = (exists.Avg + duration) / 2
		exists.Count += 1
		if exists.Min > duration {
			exists.Min = duration
		}
		if exists.Max < duration {
			exists.Max = duration
		}
		(*self)[q.GetHash()] = exists
		TopByAvg.Add(exists)
		TopByCount.Add(exists)
	} else {
		(*self)[q.GetHash()] = q
		TopByAvg.Add(q)
		TopByCount.Add(q)
	}
}

type Query struct {
	Query string
	From  string
	To    string
	Hash  string
	Avg   time.Duration
	Min   time.Duration
	Max   time.Duration
	Count int
}

func (q Query) GetQuery() string {
	return q.Query
}

func (q *Query) CalcHash() {
	h := md5.New()
	io.WriteString(h, q.Query)
	io.WriteString(h, q.From)
	io.WriteString(h, q.To)
	q.Hash = string(h.Sum(nil))
}

func (q Query) GetHash() string {
	return q.Hash
}
