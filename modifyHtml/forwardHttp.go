package modifyHtml

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func ReForward(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if len(r.Header.Get("content-type")) >= 16 {
			if strings.EqualFold(r.Header.Get("content-type")[:16], "application/json") {

				//获取原url并清理body中加上的原url
				rBBody, _ := ioutil.ReadAll(r.Body)
				rBody := string(rBBody)
				urlIndex := strings.Index(rBody, "&__original_url__=")
				originalRBody := rBody[:urlIndex]
				r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(originalRBody)))
				originalUrl := rBody[urlIndex+18:]

				//进行post转发并获得respond
				client := &http.Client{}
				req, _ := http.NewRequest("POST",
					originalUrl,
					strings.NewReader(originalRBody))

				//设置头部参数
				headers := r.Header
				for k, v := range headers {
					req.Header.Set(k, v[0])
				}

				//设置一些特殊的参数
				//req.Header.Set("Host", originalUrl[getIndex(originalUrl, "/", 2)+1:getIndex(originalUrl, "/", 3)])
				//req.Header.Set("Origin", )
				//req.Header.Del("Referer")

				resp, _ := client.Do(req)
				respBody, _ := ioutil.ReadAll(resp.Body)
				for k, v := range resp.Header {
					w.Header().Set(k, v[0])
				}
				w.Write(respBody)
			}
		}

		if len(r.Header.Get("content-type")) >= 33 {
			if strings.EqualFold(r.Header.Get("content-type")[:33], "application/x-www-form-urlencoded") {
			}
		}
	}
}
