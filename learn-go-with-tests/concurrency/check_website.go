package concurrency

type WebsitChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsits(wc WebsitChecker, urls []string) map[string]bool {
	d := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		d[result.string] = result.bool
	}
	return d
}
