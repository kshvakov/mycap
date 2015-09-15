package gui

import (
	"bytes"
	"crypto/md5"
	"io"

	"html/template"
	"net/http"
	// "sort"
	"time"
)

var (
	tpl        *template.Template
	AllQueries Queries
	TopByAvg   TopByAvgQueries
	TopByCount TopByCountQueries
	err        error
)

type Queries map[string]Query

func (self *Queries) Add(query string, start time.Time, duration time.Duration) {
	q := Query{
		Query: query,
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
		TopByAvg.Add(query, exists.Avg)
		TopByCount.Add(query, exists.Count)
	} else {
		(*self)[q.GetHash()] = q
		TopByAvg.Add(query, q.Avg)
		TopByCount.Add(query, q.Count)
	}
}

type Query struct {
	Query string
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
	q.Hash = string(h.Sum(nil))
}

func (q Query) GetHash() string {
	return q.Hash
}

func init() {
	AllQueries = make(Queries)
	initTemplates()

	http.HandleFunc("/", HandlerMainPage)
	http.HandleFunc("/top-by-avg", HandlerTopByAvg)
	http.HandleFunc("/top-by-count", HandlerTopByCount)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/andrey/GoProjects/mycap/src/mycap/gui/static/"))))
	go http.ListenAndServe("localhost:8080", nil)
}

func initTemplates() {
	tpl, err = template.ParseGlob("/Users/andrey/GoProjects/mycap/src/mycap/gui/templates/*.html")
	if err != nil {
		panic(err)
	}
}

func HandlerMainPage(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/all", map[string]interface{}{
		"queries": AllQueries,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"content": template.HTML(content.String()),
	})
}

func HandlerTopByAvg(w http.ResponseWriter, r *http.Request) {
	TopByAvg.SortDesc()

	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/top-by-avg", map[string]interface{}{
		"queries": TopByAvg,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"content": template.HTML(content.String()),
	})
}

func HandlerTopByCount(w http.ResponseWriter, r *http.Request) {
	TopByCount.SortDesc()

	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/top-by-count", map[string]interface{}{
		"queries": TopByCount,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"content": template.HTML(content.String()),
	})
}
