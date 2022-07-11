package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dvsekhvalnov/jose2go/base64url"
)

func setupCacheDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil { return "", err }
	testCacheDir := pwd + "/run_test"
	os.Mkdir(testCacheDir, os.ModePerm)
	return testCacheDir, err
}

func teardownCacheDir() {
	pwd, _ := os.Getwd()
	testCacheDir := pwd + "/run_test"
	os.RemoveAll(testCacheDir)
}

func TestRoot(t *testing.T) {
	testCacheDir, err := setupCacheDir()
	if err != nil { t.Fatal(err) }
	defer teardownCacheDir()

	//tearup curlcache
	ts := httptest.NewServer(setupGin(testCacheDir))
	defer ts.Close()

	//tearup http request target server
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "yeah")
	}))


	//throw request and validate
	resp, err := http.Get(fmt.Sprintf("%s/?q=%s", ts.URL, base64url.Encode([]byte(targetServer.URL))))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	if resp.Header.Get("Curlcache-Hit") != "false" {
		t.Fatalf("キャッシュミスを表すヘッダーの値が正しくない\n expect:%s actula:%s", "false", resp.Header.Get("Curlcache-Hit"))
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if string(body) != "yeah" {
			t.Fatalf("Expected data is %s, but body is %s", "yeah", body)
		}
	} else {
		t.Fatalf("Can not read body")
	}


	// teardown http request target server for test using cache
	targetServer.Close()

	//throw request and validate
	resp, err = http.Get(fmt.Sprintf("%s/?q=%s", ts.URL, base64url.Encode([]byte(targetServer.URL))))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	if resp.Header.Get("Curlcache-Hit") != "true" {
		t.Fatalf("キャッシュヒットを表すヘッダーの値が正しくない\n expect:%s actula:%s", "true", resp.Header.Get("Curlcache-Hit"))
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if string(body) != "yeah" {
			t.Fatalf("Expected data is %s, but body is %s", "yeah", body)
		}
	} else {
		t.Fatalf("Can not read body")
	}
}

func TestReferer(t *testing.T) {
	testCacheDir, err := setupCacheDir()
	if err != nil { t.Fatal(err) }
	defer teardownCacheDir()

	referer := "http://sample.dammy/r"

	sample_dammy := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, r.Referer())
	}))

	ts := httptest.NewServer(setupGin(testCacheDir))
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf(
		"%s/?q=%s&referer=%s",
		ts.URL,
		base64url.Encode([]byte(sample_dammy.URL)),
		base64url.Encode([]byte(referer)),
	))
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != referer {
		t.Fatalf("refereをうまく受け渡せていません\n expect: %s\n actual %s ", referer, string(data))
	}
}

func TestMetrics(t *testing.T) {
	//tearup curlcache
	testCacheDir, err := setupCacheDir()
	if err != nil { t.Fatal(err) }
	defer teardownCacheDir()

	ts := httptest.NewServer(setupGin(testCacheDir + "/run_test"))
	defer ts.Close()

	_, err = http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resp, err := http.Get(fmt.Sprintf("%s/metrics", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.Contains(string(data), "request_counter") {
		t.Fatalf("request_counterメトリクスを得られませんでした\n %s", data)
	}
}
