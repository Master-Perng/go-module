package Alibaba

import (
	log "engine/logsys"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func DingBot(message string , api string)  {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	data := "{\"msgtype\": \"text\",\"text\": {\"content\":\""+message+"\"},\"at\":{\"atMobiles\":[\"17774014757\"]}}"
	req , err := http.NewRequest("POST",api,strings.NewReader(data))
	if err != nil{
		log.Error("Error :",err.Error())
	}
	req.Header.Add("Content-Type","application/json")
	resp , err := client.Do(req)
	if err != nil{
		log.Error("Error :",err.Error())
	}
	context , _:= io.ReadAll(resp.Body)
	fmt.Println(string(context))
}
