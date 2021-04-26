package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockChecker(url string) bool {
	return url != "http://www.4399.com"
}

func TestCheckWebsits(t *testing.T) {
	t.Run("check websit orderly", func(t *testing.T) {
		urls := []string{
			"http://www.baidu.com",
			"http://www.google.com",
			"http://www.4399.com",
		}

		got := CheckWebsits(mockChecker, urls)
		expected := map[string]bool{
			"http://www.baidu.com":  true,
			"http://www.google.com": true,
			"http://www.4399.com":   false,
		}

		if len(got) != len(expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})
}

func slowChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsits(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsits(slowChecker, urls)
	}
}
