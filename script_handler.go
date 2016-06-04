package traffic_api

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"

	traffic "github.com/DStatIt/traffic"
)

var (
	tmpl *template.Template
)

type script struct {
	Redirect bool
	URL      string
}

func init() {
	var err error
	tmpl, err = compileTemplates("templates/script.js")
	if err != nil {
		panic(err)
	}

}

func GetScript(w http.ResponseWriter, r *http.Request) (int, error) {

	x := traffic.RequestInfo{
		Request: r,
	}

	if err := x.BuildUser(); err != nil {
		return 500, err
	}

	if err := x.SaveUser(); err != nil {
		return 500, err
	}

	log.Println("New User: ", x.IP)

	// http.SetCookie(w, x.Cookie)

	w.Header().Set("Content-Type", "application/javascript")

	if err := tmpl.Execute(w, script{
		Redirect: false,
		URL:      "http://google.com",
	}); err != nil {
		return 500, err

	}

	return 200, nil
}

func compileTemplates(filenames ...string) (*template.Template, error) {
	m := minify.New()
	m.AddFunc("application/javascript", js.Minify)

	var tmpl *template.Template
	for _, filename := range filenames {
		name := filepath.Base(filename)
		if tmpl == nil {
			tmpl = template.New(name)
		} else {
			tmpl = tmpl.New(name)
		}

		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		mb, err := m.Bytes("application/javascript", b)
		if err != nil {
			return nil, err
		}
		tmpl.Parse(string(mb))
	}
	return tmpl, nil
}
