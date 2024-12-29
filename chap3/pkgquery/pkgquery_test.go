package pkgquery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestPackageServer() *httptest.Server {
	pkgData := `[
		{"name": "pkg1", "version": "1.0.0"},
		{"name": "pkg2", "version": "2.0.0"}]`

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, pkgData)
		}),
	)

	return ts
}

func TestFetchPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()

	pkgs, err := fetchPackageData(ts.URL)
	if err != nil {
		t.Fatalf("fetchPackageData(%q) returned error: %v", ts.URL, err)
	}

	if len(pkgs) != 2 {
		t.Fatalf("fetchPackageData(%q) = %d; want 2", ts.URL, len(pkgs))
	}

	if pkgs[0].Name != "pkg1" {
		t.Errorf("fetchPackageData(%q) = %q; want %q", ts.URL, pkgs[0].Name, "pkg1")
	}

	if pkgs[1].Name != "pkg2" {
		t.Errorf("fetchPackageData(%q) = %q; want %q", ts.URL, pkgs[1].Name, "pkg2")
	}
}
