package myhttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	CountCacheHit = 0
	CountCacheMiss = 0
)

func Curl(url string, referer string, cache bool, js bool) (io.ReadCloser, error) {
	if cache && js {
		return curlCacheJs(url)
	}
	if cache {
		data, hit, err := curlCache(url, referer)
		if err != nil {
			return nil, err
		}
		if !hit {
			log.Printf("[info] sleep 1sec, cache miss. url: %s", url)
			time.Sleep(time.Second)
		}
		return data, err
	}
	if js {		
		return curlJs(url)
	} else {
		resp, err := Get(url)
		return resp.Body, err
	}
}

func curlCache(planeUrl string, planeReferer string) (io.ReadCloser, bool, error) {
	encodedUrl := fmt.Sprintf(
		"%s/?q=%s&referer=%s",
		"http://curlcache",
		base64url.Encode([]byte(planeUrl)),
		base64url.Encode([]byte(planeReferer)),
	)
	resp, err := Get(encodedUrl)
	if err != nil {
		return nil, false, err
	}
	if resp.StatusCode != 200 {
		return nil, false, errors.New(fmt.Sprintf("status code is not 200 but %d", resp.StatusCode))
	}
	hit := resp.Header.Get("Curlcache-Hit") == "true"
	return resp.Body, hit, nil
}

func curlJs(planeUrl string) (io.ReadCloser, error) {
	encodedUrl := "http://curljs/?q=" + base64url.Encode([]byte(planeUrl))
	return getBody(encodedUrl)
}

func curlCacheJs(planeUrl string) (io.ReadCloser, error) {
	encodedUrl := planeUrl
	encodedUrl = "http://curljs/?q=" + base64url.Encode([]byte(encodedUrl))
	encodedUrl = "http://curlcache/?q=" + base64url.Encode([]byte(encodedUrl))
	return getBody(encodedUrl)
}

func getBody(url string) (io.ReadCloser, error) {
	resp, err := Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http response is not 200 but %d", resp.StatusCode))
	}

	return resp.Body, nil
}

func GetQueryUrl(org_url *url.URL) (*url.URL, error) {
	encoded_q := org_url.Query().Get("q")

	if encoded_q == "" {
		return org_url, nil
	}

	decoded_q, err := base64url.Decode(encoded_q)
	if err != nil {
		return nil, err
	}

	decoded_url, err := url.Parse(string(decoded_q))
	if err != nil {
		return nil, err
	}

	return GetQueryUrl(decoded_url)
}

func Get(url string) (*http.Response, error) {
	cli := retryablehttp.NewClient()
	cli.RetryMax = 100
	cli.Logger = nil
	cli.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if b, err := retryablehttp.DefaultRetryPolicy(ctx, resp, err); !b {
			return b, err
		}
		return resp.StatusCode != 200, err
	}

	return cli.Get(url)
}
