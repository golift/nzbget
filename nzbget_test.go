package nzbget_test

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golift.io/nzbget"
)

func TestNewURLAndAuth(t *testing.T) {
	t.Parallel()

	var (
		gotPath string
		gotAuth string
		gotBody string
	)

	srv := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		gotPath = req.URL.Path
		gotAuth = req.Header.Get("Authorization")
		body, _ := io.ReadAll(req.Body)
		gotBody = string(body)

		if strings.Contains(gotBody, "broken") {
			_, _ = resp.Write([]byte(`{"result":null,"error":"no such method","id":1}`))

			return
		}

		_, _ = resp.Write([]byte(`{"result":{"Version":"21.1"},"error":null,"id":1}`))
	}))
	t.Cleanup(srv.Close)

	client := nzbget.New(&nzbget.Config{URL: srv.URL, User: "nzb", Pass: "get"})

	var version struct {
		Version string
	}

	err := client.GetInto(context.Background(), "version", &version)
	if err != nil {
		t.Fatal(err)
	}

	if version.Version != "21.1" {
		t.Fatalf("version %#v", version)
	}

	if gotPath != "/jsonrpc" {
		t.Fatalf("path %s", gotPath)
	}

	wantAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("nzb:get"))
	if gotAuth != wantAuth {
		t.Fatalf("auth %q", gotAuth)
	}

	if !strings.Contains(gotBody, `"method":"version"`) {
		t.Fatalf("body %s", gotBody)
	}

	err = client.GetInto(context.Background(), "broken", &version)
	if err == nil || !strings.Contains(err.Error(), "no such method") {
		t.Fatalf("err %v", err)
	}
}
