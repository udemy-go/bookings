package render

import (
	"net/http"
	"testing"

	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
)

// TestAddDefaultData checks whether the added data is in session or not
func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// to check whether the data is in session, first add some value
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 is not found in session ")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = Template(&ww, r, "about.page.gohtml", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = Template(&ww, r, "non-exist.page.gohtml", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that not exist   ")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()                                    // creating the context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session")) // loads the session into the ctx context
	r = r.WithContext(ctx)                                // add the context into our request

	return r, nil

}

func TestNewTemplate(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}