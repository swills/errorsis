package errorsisstruct

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Timestamp struct {
	time.Time
}

type Rate struct {
	Limit int `json:"limit"`

	Remaining int `json:"remaining"`

	Reset Timestamp `json:"reset"`
}

// borrowed from go-github

type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string `json:"message"`
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

func formatRateReset(d time.Duration) string {
	isNegative := d < 0
	if isNegative {
		d *= -1
	}
	secondsTotal := int(0.5 + d.Seconds())
	minutes := secondsTotal / 60
	seconds := secondsTotal - minutes*60

	var timeString string
	if minutes > 0 {
		timeString = fmt.Sprintf("%dm%02ds", minutes, seconds)
	} else {
		timeString = fmt.Sprintf("%ds", seconds)
	}

	if isNegative {
		return fmt.Sprintf("[rate limit was reset %v ago]", timeString)
	}
	return fmt.Sprintf("[rate reset in %v]", timeString)
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, formatRateReset(time.Until(r.Rate.Reset.Time)))
}

func compareHTTPResponse(r1, r2 *http.Response) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 != nil && r2 != nil {
		return r1.StatusCode == r2.StatusCode
	}
	return false
}

func (r *RateLimitError) Is(target error) bool {
	v, ok := target.(*RateLimitError)
	if !ok {
		return false
	}

	return r.Rate == v.Rate &&
		r.Message == v.Message &&
		compareHTTPResponse(r.Response, v.Response)
}

func doerrorsisstuff() {
	err := error(&RateLimitError{Message: "foo"})
	if errors.Is(err, &RateLimitError{}) { // want `incorrect usage of errors`
		fmt.Println("yes, this error is the same as a generic RateLimitError")
	} else {
		fmt.Println("no, this is not the same as a generic RateLimitError")
	}
}

func doerrorsisstuffRight() {
	err := error(&RateLimitError{Message: "foo"})

	var rateLimitErr *RateLimitError
	if errors.As(err, &rateLimitErr) {
		fmt.Println("yes, this is a generic RateLimitError")
	} else {
		fmt.Println("no, this is not a RateLimitError")
	}
}
