package cachedhttp

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dvsekhvalnov/jose2go/base64url"
)

func cachePath(cacheDir string, planeUrl string) string {
	encodedUrl := base64url.Encode([]byte(planeUrl))
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, 666)
	}
	return cacheDir + "/" + encodedUrl
}

func ReadCache(cacheDir string, planeUrl string) (hit bool, data []byte) {
	if file, err := os.Open(cachePath(cacheDir, planeUrl)); err == nil {
		data, err := ioutil.ReadAll(file)
		file.Close()
		if err != nil {
			log.Printf("キャッシュファイルが読めないので、キャッシュミスとして扱います")
			return false, nil
		}
		return true, data
	}
	return false, nil
}

func WriteCache(cacheDir string, planeUrl string, byteData []byte) error {
	ioutil.WriteFile(cachePath(cacheDir, planeUrl), byteData, os.ModePerm)
	return nil
}
