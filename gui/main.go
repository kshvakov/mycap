package gui

import (
	"bytes"
	"html/template"
	"net/http"
)

var (
	tpl        *template.Template
	AllQueries Queries
	TopByAvg   TopByAvgQueries
	TopByCount TopByCountQueries
	err        error
	guiPath    string
)

func InitGui(gui string) {
	guiPath = gui
	AllQueries = make(Queries)
	initTemplates()

	http.HandleFunc("/", HandlerDashboard)
	http.HandleFunc("/all-queries", HandlerAllQueries)
	http.HandleFunc("/top-by-avg", HandlerTopByAvg)
	http.HandleFunc("/top-by-count", HandlerTopByCount)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(guiPath+"/static/"))))
	go http.ListenAndServe("localhost:8080", nil)
}

func initTemplates() {
	tpl, err = template.ParseGlob(guiPath + "/templates/*.html")
	if err != nil {
		panic(err)
	}
}

func HandlerDashboard(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/all", map[string]interface{}{
		"queries": AllQueries,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Dashboard",
		"content":   template.HTML(content.String()),
	})
}

func HandlerAllQueries(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/all", map[string]interface{}{
		"queries": AllQueries,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Все запросы",
		"content":   template.HTML(content.String()),
	})
}

func HandlerTopByAvg(w http.ResponseWriter, r *http.Request) {
	TopByAvg.SortDesc()

	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/top-by-avg", map[string]interface{}{
		"queries": TopByAvg,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по времени выполнения",
		"content":   template.HTML(content.String()),
	})
}

func HandlerTopByCount(w http.ResponseWriter, r *http.Request) {
	TopByCount.SortDesc()

	content := new(bytes.Buffer)
	tpl.ExecuteTemplate(content, "view/queries/top-by-count", map[string]interface{}{
		"queries": TopByCount,
	})

	tpl.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по частоте выполнения",
		"content":   template.HTML(content.String()),
	})
}
