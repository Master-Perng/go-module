package tyc

import (
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func TycGZH(page int, name string, token string) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	TycUrl := fmt.Sprintf(apiSub, 20, url.PathEscape(name), page)
	req, err := http.NewRequest("GET", TycUrl, strings.NewReader(""))
	req.Header.Add("Authorization", token)

	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	return string(result), err

}
