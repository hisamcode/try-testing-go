package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_form_has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("whatever")

	if has {
		t.Error("form shows has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)
	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func Test_form_required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/wathever", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows post does not have required fields, when it does")
	}
}

func Test_form_check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Valid() returns false, and it should be true when calling check()")
	}
}

func Test_formErrors_get(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")
	if len(s) == 0 {
		t.Error("should have an error returned from get, but do not")
	}

	s = form.Errors.Get("wathever")
	if len(s) != 0 {
		t.Error("should not have an error, but got one")
	}
}
