package main

import (
	"net/url"
	"strings"
)

type FormErrors map[string][]string

func (fe FormErrors) Get(field string) string {
	errorSlice := fe[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

func (fe FormErrors) Add(field, message string) {
	fe[field] = append(fe[field], message)
}

type Form struct {
	Data   url.Values
	Errors FormErrors
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	return x != ""
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.Errors.Add(key, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
