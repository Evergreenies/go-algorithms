/**
* Implement a URL shortener with the following methods:
*
* shorten(url), which shortens the url into a six-character alphanumeric string, such as zLg6wl.
* restore(short), which expands the shortened string into the original url. If no such shortened string exists, return null.
*
* Hint: What if we enter the same URL twice?
* */
package dailycodingproblems

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

type urlShortnerStrct struct {
	urlToShort map[string]string
	shortToURL map[string]string

	baseURL string
	length  int
}

func newURLShortner() *urlShortnerStrct {
	return &urlShortnerStrct{
		urlToShort: make(map[string]string),
		shortToURL: make(map[string]string),

		baseURL: "http://short.ly/",
		length:  6,
	}
}

func (u *urlShortnerStrct) generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, u.length)

	for index := range b {
		b[index] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func (u *urlShortnerStrct) shorten(url string) string {
	if short, exists := u.urlToShort[url]; exists {
		return u.baseURL + short
	}

	var short string
	for {
		short = u.generateShortURL()
		if _, ok := u.shortToURL[url]; !ok {
			break
		}
	}

	u.urlToShort[url] = short
	u.shortToURL[short] = url

	return u.baseURL + short
}

func (u *urlShortnerStrct) restore(short string) string {
	key := strings.TrimPrefix(short, u.baseURL)
	if original, ok := u.shortToURL[key]; ok {
		return original
	}

	return ""
}

func TestURLShortener(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	// assert := assert.New(t)

	shortener := newURLShortner()

	// Shorten a URL
	shortURL := shortener.shorten("https://www.example.com")
	fmt.Println("Shortened URL:", shortURL)

	// Restore the original URL
	originalURL := shortener.restore(shortURL)
	fmt.Println("Original URL:", originalURL)

	// Shorten the same URL again
	shortURLDuplicate := shortener.shorten("https://www.example.com")
	fmt.Println("Shortened URL (duplicate):", shortURLDuplicate)

	// Restore an invalid shortened URL
	invalidRestore := shortener.restore("http://short.ly/invalid")
	if invalidRestore == "" {
		fmt.Println("Invalid Restore: None")
	} else {
		fmt.Println("Invalid Restore:", invalidRestore)
	}
}
