package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	t.Run("retrun an error if a server doesn't respond with 10s", func(t *testing.T) {

		serverA := makeDelayServer(time.Second * 1)
		serverB := makeDelayServer(time.Second * 2)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Race(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("expected an error, but don't get one")
		}

	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
