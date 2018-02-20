package golang_playground

import (
	"crypto/tls"
	"github.com/certifi/gocertifi"
	"net/http"
	"testing"
	"time"
)

func Test_HTTPSGet(t *testing.T) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	_, err := client.Get("https://github.com/GoogleChrome/puppeteer.git/info/refs?service=git-upload-pack")
	if err != nil {
		t.Error(err)
	}
}

func Test_HTTPGet(t *testing.T) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	_, err := client.Get("http://github.com/GoogleChrome/puppeteer.git/info/refs?service=git-upload-pack")
	if err != nil {
		t.Error(err)
	}
}

func Test_HTTPSGetWithCustomRootCertificates(t *testing.T) {
	certPool, err := gocertifi.CACerts()
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: certPool},
		},
		Timeout: 60 * time.Second,
	}
	_, err = client.Get("https://github.com/GoogleChrome/puppeteer.git/info/refs?service=git-upload-pack")
	if err != nil {
		t.Error(err)
	}
}

func Test_HTTPGetWithCustomRootCertificates(t *testing.T) {
	certPool, err := gocertifi.CACerts()
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: certPool},
		},
		Timeout: 60 * time.Second,
	}
	_, err = client.Get("http://github.com/GoogleChrome/puppeteer.git/info/refs?service=git-upload-pack")
	if err != nil {
		t.Error(err)
	}
}
