package runner

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"time"
)

func TestRunner(t *testing.T) {

	t.Run("returns the fastest server", func(t *testing.T) {
		slowService := createService(1 * time.Millisecond)
		fastService := createService(0 * time.Millisecond)
		
		// defer chama a função no final da execução.
		defer slowService.Close()
		defer fastService.Close()

		expected := fastService.URL
		result, err := ConfigRunner(slowService.URL, fastService.URL, 2 * time.Millisecond)
		
		if err != nil {
			t.Fatalf("Don't expected a error: %v", err)
		}
		if expected != result {
			t.Errorf("expected '%s', result '%s'", expected, result)
		}
	})

	t.Run("returns if no server responds in 10s", func(t *testing.T) {
		slowService := createService(3 * time.Millisecond)
		fastService := createService(2 * time.Millisecond)
		
		// defer chama a função no final da execução.
		defer slowService.Close()
		defer fastService.Close()

		_, err := ConfigRunner(slowService.URL, fastService.URL, 1 * time.Millisecond)
		
		if err == nil {
			t.Error("Expected in return in less than 10s")
		}
	})
}

func createService(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
