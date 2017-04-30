package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mycap/libs/agrqueries/countpertime"
	"net/http"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	PathTemplates string `json:"path_templates"`
	PathStatic    string `json:"path_static"`

	templates *template.Template

	AgentsCollector  AgentsCollector  `json:"agents_collector"`
	QueriesCollector QueriesCollector `json:"queries_collector"`

	HeadServerHost string `json:"server_host"`
	HeadServerPort int    `json:"server_port"`
}

func (self *Server) StartAgentsCollector() {
	self.AgentsCollector.server = self
	self.AgentsCollector.Collect()
}

func (self *Server) StartQueriesCollector() {
	self.QueriesCollector.server = self
	self.QueriesCollector.Collect()
}

type PlotItem [2]interface{}
type PlotData []PlotItem

func (self *Server) InitTemplates() {

	funcMap := template.FuncMap{
		"plot": func(counter countpertime.Counter) template.JS {
			result := make(PlotData, len(counter.Items))
			for key, val := range counter.Items {
				result[key] = PlotItem{1000 * (counter.TimeZero + (int64(key) * counter.StepSize)), val}
			}

			if result_js, err := json.Marshal(result); err == nil {
				return template.JS(result_js)
			} else {
				log.Println(err)
				return template.JS("")
			}
		},
	}

	tpl := template.New("asd")
	tpl.Funcs(funcMap)

	if _, err := tpl.ParseGlob(self.PathTemplates + "/*.html"); err == nil {
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

	http.HandleFunc("/counter-queries-per-time", self.CounterQueriesPerTimeAjax)

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
		"content": template.HTML(content.String()),
	})
}

func (self *Server) CounterQueriesPerTimeAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "view/queries/count-per-time", map[string]interface{}{
		"countPerTime": self.QueriesCollector.countPerTime,
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
