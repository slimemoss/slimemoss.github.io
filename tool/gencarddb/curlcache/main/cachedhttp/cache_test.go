package cachedhttp

import (
	"bytes"
	"os"
	"testing"
)

func setupCacheDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil { return "", err }
	testCacheDir := pwd + "/run_test"
	os.RemoveAll(testCacheDir)
	os.Mkdir(testCacheDir, os.ModePerm)
	return testCacheDir, err
}

func teardownCacheDir() {
	pwd, _ := os.Getwd()
	testCacheDir := pwd + "/run_test"
	os.RemoveAll(testCacheDir)
}

func TestAll(t *testing.T) {
	// setup test
	testCacheDir, err := setupCacheDir()
	defer teardownCacheDir()
	if err != nil { t.Fatal(err) }
	testUrl := "http://a.b.c/?a=b&c=d"
	testData := "testdata"

	// cache is empty
	hit, _ := ReadCache(testCacheDir, testUrl)
	if hit {
		t.Fatal("空のキャッシュにヒットしてる")
	}

	// write cache
	WriteCache(testCacheDir, testUrl, []byte(testData))

	// read cache
	hit, data := ReadCache(testCacheDir, testUrl)
	if !hit {
		t.Fatal("")
	}
	if !bytes.Equal(data, []byte(testData)) {
		t.Fatal("")
	}
}
