package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) { // since the function have receiver form the syntax is Test[Receiver]_FunctionName 
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid() 
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows doesn't have required field when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non existing field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("a", "123456")
	form = New(postedData)

	form.MinLength("a", 50)
	if form.Valid() {
		t.Error("shows min length of 50 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("b", "abcdefgh")
	form = New(postedData)

	form.MinLength("b", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}

	isError = form.Errors.Get("b")
	if isError != "" {
		t.Error("should have an error, but did not get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("should not have an error, but got one")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@gmail.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}
}