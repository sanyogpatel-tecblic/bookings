package handlers

import (
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
	expectedstatuscode int
}{
	// {"home", "/", "GET", http.StatusOK},
	// {"about", "/about", "GET", http.StatusOK},
	// {"gq", "/generals-quarters", "GET", http.StatusOK},
	// {"ms", "/majors-suite", "GET", http.StatusOK},
	// {"sa", "/search-availability", "GET", http.StatusOK},
	// {"contact", "/contact", "GET", http.StatusOK},
	// {"non-existent", "/green/eggs/and/ham", "GET", http.StatusNotFound},
	// {"login", "/user/login", "GET", http.StatusOK},
	// {"logout", "/user/logout", "GET", http.StatusOK},
	// {"dashboard", "/admin/dashboard", "GET", http.StatusOK},
	// {"new res", "/admin/reservations-new", "GET", http.StatusOK},
	// {"all res", "/admin/reservations-all", "GET", http.StatusOK},
	// {"show res", "/admin/reservations/new/1/show", "GET", http.StatusOK},
	// {"show res cal", "/admin/reservations-calendar", "GET", http.StatusOK},
	// {"show res cal with params", "/admin/reservations-calendar?y=2020&m=1", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := GetRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedstatuscode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedstatuscode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedstatuscode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedstatuscode, resp.StatusCode)
			}
		}
	}
}
