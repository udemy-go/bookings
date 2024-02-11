package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2024-02-01"},
		{key: "end", value: "2024-02-10"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Thiru"},
		{key: "last_name", value: "T"},
		{key: "email", value: "thiru@gmail.com"},
		{key: "phone", value: "7845098854"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2024-02-01"},
		{key: "end", value: "2024-02-10"},
	}, http.StatusOK},
}

func TestHandler(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url) // ts.URL get the url like localhost:8080 and e.url get the routes that we mentioned 
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			// store the post values in post value format url.Values{}
			values := url.Values{} 
			for _, value := range e.params {
				values.Add(value.key, value.value)
			}

			resp, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}