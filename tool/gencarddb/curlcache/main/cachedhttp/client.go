package cachedhttp

import (
	"io/ioutil"
	"main/exporter"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func Get(
	cacheDir string,
	planeUrl string,
	planeReferer string,
) (data []byte, hit bool, err error) {

	if hit, _ := ReadCache(cacheDir, planeUrl); hit {
		exporter.CacheCounter.With(prometheus.Labels{"isHit": "hit"}).Inc()
	} else {
		exporter.CacheCounter.With(prometheus.Labels{"isHit": "miss"}).Inc()
	}

	if hit, data := ReadCache(cacheDir, planeUrl); hit {
		return data, true, nil
	} else {
		req, err := http.NewRequest("GET", planeUrl, nil)
		if err != nil {
			return nil, false, errors.WithStack(err)
		}
		req.Header.Set("Referer", planeReferer)

		cli := retryablehttp.NewClient()
		cli.Logger = nil
		resp, err := cli.StandardClient().Do(req)
		if err != nil {
			return nil, false, errors.WithStack(err)
		}
		d := resp.Body
		data, err := ioutil.ReadAll(d)
		if err != nil {
			return nil, false, errors.WithStack(err)
		}
		resp.Body.Close()

		WriteCache(cacheDir, planeUrl, data)
		return data, false, nil
	}
}
