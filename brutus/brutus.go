package brutus

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/roerohan/bird/logger"
)

// Brute is a struct to define
// a single job to work on
type Brute struct {
	URL  string
	Word string
}

// New returns a new Brute object
func New(url string, word string) *Brute {
	return &Brute{
		URL:  url,
		Word: word,
	}
}

// FormURL forms a URL from a brute object
func (b *Brute) FormURL() string {
	return fmt.Sprintf("%s/%s", b.URL, b.Word)
}

// Try tries to visit a Brute URL and checks the status code
func (b *Brute) Try(success map[string]bool, logs chan logger.Log) {
	url := b.FormURL()

	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error occurred while visiting " + url)
	}

	statusCode := strconv.Itoa(resp.StatusCode)

	if success[statusCode] {
		logs <- logger.Log{
			Message: fmt.Sprintf("%s [Status code %s]", url, statusCode),
			Func:    logger.Info,
		}
	}
}
