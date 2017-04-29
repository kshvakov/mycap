package web

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	Host string
	Port int

	HeadServerHost string
	HeadServerPort int

	PathTemplates string
	PathStatic    string

	templates *template.Template

	AgentsCollector  AgentsCollector
	QueriesCollector QueriesCollector
}

func (self *Server) StartAgentsCollector() {
	self.AgentsCollector.server = self
	self.AgentsCollector.Collect()
}

func (self *Server) StartQueriesCollector() {
	self.QueriesCollector.server = self
	self.QueriesCollector.Collect()
}

func (self *Server) InitTemplates() {
	if tpl, err := template.ParseGlob(self.PathTemplates + "/*.html"); err == nil {
		self.templates = tpl
	} else {
		log.Fatal(err)
	}
}

func (self *Server) StartWebServer() {
	http.HandleFunc("/", self.HandlerDashboard)
	http.HandleFunc("/all-nodes", self.HandlerAllNodes)
	http.HandleFunc("/all-queries", self.HandlerAllQueries)
	http.HandleFunc("/top-by-avg", self.HandlerTopByAvg)
	http.HandleFunc("/top-by-count", self.HandlerTopByCount)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(self.PathStatic))))
	http.ListenAndServe(fmt.Sprintf("%s:%d", self.Host, self.Port), nil)
}

func (self *Server) HandlerDashboard(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "view/queries/count-per-time", map[string]interface{}{
		"countPerTime": self.QueriesCollector.countPerTime,
	})

	self.templates.ExecuteTemplate(content, "view/queries/all", map[string]interface{}{
		"queries": self.QueriesCollector.queries.Queries.Items,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Dashboard",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) HandlerAllNodes(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)

	self.templates.ExecuteTemplate(content, "view/nodes/all", map[string]interface{}{
		"nodes": self.AgentsCollector.agents,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Все агенты",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) HandlerAllQueries(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "view/queries/all", map[string]interface{}{
		"queries": self.QueriesCollector.queries.Queries.Items,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Все запросы",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) HandlerTopByAvg(w http.ResponseWriter, r *http.Request) {
	self.QueriesCollector.queries.TopAvg.SortDesc()

	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "view/queries/top-by-avg", map[string]interface{}{
		"queries": self.QueriesCollector.queries.TopAvg.Items,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по времени выполнения",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) HandlerTopByCount(w http.ResponseWriter, r *http.Request) {
	self.QueriesCollector.queries.TopCnt.SortDesc()

	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "view/queries/top-by-count", map[string]interface{}{
		"queries": self.QueriesCollector.queries.TopCnt.Items,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по частоте выполнения",
		"content":   template.HTML(content.String()),
	})
}
