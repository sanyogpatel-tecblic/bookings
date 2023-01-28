package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a new form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field is required")
		}

	}
}
func (f *Form) MinLength(field string, length int) bool {
	err := f.Get(field)

	if len(err) < length {
		f.Errors.Add(field, fmt.Sprintf("minimum length reuired is %d", length))
		return false
	}
	return true
}

func (f *Form) MinLength2(length int, r *http.Request, fields ...string) {
	for _, fields := range fields {
		fi := f.Get(fields)

		if len(fi) < length {
			f.Errors.Add(fields, fmt.Sprintf("minimum length criteria dosen't match %d", length))
		}
	}
}
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	if x == "" {
		f.Errors.Add(field, "Field Required!")
		return false
	}
	return true
}
func (f *Form) Isemail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email address please enter the valid email address")

	}
}
