package main

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Calculate struct {
	value int
	mu    sync.Mutex
}

func (c *Calculate) Increase() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Calculate) Decrease() {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
}

func (c *Calculate) Multiply() {
	c.mu.Lock()
	c.value = c.value * 2
	c.mu.Unlock()
}

func (c *Calculate) Divide() {
	c.mu.Lock()
	c.value = c.value / 2
	c.mu.Unlock()
}

func (c *Calculate) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &Calculate{}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})
	r.Post("/increase", func(w http.ResponseWriter, _ *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))
		counter.Increase()
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})
	r.Post("/decrease", func(w http.ResponseWriter, _ *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))
		counter.Decrease()
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)

	})
	r.Post("/multiply", func(w http.ResponseWriter, _ *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))
		counter.Multiply()
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)

	})
	r.Post("/divide", func(w http.ResponseWriter, _ *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))
		counter.Divide()
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})

	http.ListenAndServe("localhost:3000", r)
}
