package cachedhttp

import (
	"bytes"
	"fmt"
	"main/exporter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFromInternet(t *testing.T) {
	exporter.NewRegistry()

	//tearup http request target server
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	}))

	testUrl := targetServer.URL
	testCacheDir, err := setupCacheDir()
	defer teardownCacheDir()

	// exec test target func
	data, hit, err := Get(testCacheDir, testUrl, "")
	if err != nil {
		t.Fatal(err.Error())
	}
	if hit {
		t.Fatalf("初回のアクセスではキャッシュヒットしないはず")
	}
	// test request body
	if !bytes.Equal(data, []byte("test")) {
		t.Fatalf("%s != %s", data, "test")
	}
}

func TestGetFromCache(t *testing.T) {
	exporter.NewRegistry()

	// setup target server
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	}))

	//setup test
	testUrl := targetServer.URL
	testCacheDir, err := setupCacheDir()
	defer teardownCacheDir()

	// first, get test from internet and write cache
	_, _, err = Get(testCacheDir, testUrl, "")
	if err != nil {
		t.Fatal(err.Error())
	}

	// teardown target server
	targetServer.Close()
	
	// exec test target func
	data, hit, err := Get(testCacheDir, testUrl, "")
	if err != nil {
		t.Fatal(err.Error())
	}
	if !hit {
		t.Fatalf("2回目のアクセスなのでキャッシュヒットしないといけない")
	}
	// test request body
	if !bytes.Equal(data, []byte("test")) {
		t.Fatal(data)
	}
}
