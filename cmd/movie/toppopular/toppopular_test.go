package toppopular

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fschossler/tmdbcli/internal"
	"github.com/stretchr/testify/assert"
)

func TestTopPopular(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/movie/popular", r.URL.Path, "Expected request to /movie/popular")
		assert.Equal(t, "application/json", r.Header.Get("Accept"), "Expected Accept: application/json header")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"page": 1, "results": [{"title": "Movie 1", "vote_average": 8.5, "overview": "Overview of Movie 1"}]}`))
	}))
	defer server.Close()

	// Set the base URL for the internal package to the mock server's URL
	internal.SetBaseURL(server.URL)

	err := TopPopular()
	assert.NoError(t, err)
}
