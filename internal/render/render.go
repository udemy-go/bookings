package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/config"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderedTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	// We say if UseCache is false call CreateTemplateCache()(means we're in production so create template cache for every request), 
	if app.UseCache {
		// get the template cache from the app config 
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	

	// get requested from template cache
	t, ok := tc[tmpl]    // checking the given tmpl is there in the template cache
	if !ok {
		// log.Fatal("Could not get template from template cache")
		return errors.New("could not get template from template cache")
	}

	buf := new(bytes.Buffer) // this is an optional line, it is for finding error 

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)
	

	// render the template 
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil 	
} 

// CreateTemplateCache creates a template cache as a map 
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // first you the parse the template file and then the layout file 
	

	// get all the files names *.page.gohtml from ./templates 
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all the pages ending with *.page.gohtml 
	for _, page := range pages {
		// Base return the last element of the path, that is the file name here
		name := filepath.Base(page)
		// ts (template set) stores the template called name, parsed by the file page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts 
	}

	return myCache, nil
}

