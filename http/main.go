package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://randomuser.me/api")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, bs, "", "    "); err != nil {
		return 0, err
	}

	fmt.Println(prettyJSON.String())
	return len(bs), nil
}
