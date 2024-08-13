package gupshup

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gabrielcervante/gupshup/interfaces"
	"github.com/gabrielcervante/gupshup/utils"
)

const url = "https://api.gupshup.io/wa/api/v1/msg"

type whatsapp struct {
	Source  string
	AppName string
	Token   string
}

func (w whatsapp) SendMessage(message, destination string) (json.RawMessage, error) {
	req, err := w.prepareRequest(message, destination)
	if err != nil {
		return nil, err
	}

	res, err := w.do(req)
	if err != nil {
		return nil, err
	}

	return w.readBody(res)
}

func (w whatsapp) readBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func (w whatsapp) do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func (w whatsapp) prepareRequest(message, destination string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, w.convertPayload(message, destination))
	if err != nil {
		return req, err
	}

	req.Header = w.addHeaders()

	return req, err
}

func (w whatsapp) convertPayload(message, destination string) *strings.Reader {
	return strings.NewReader(w.generatePayload(message, destination))
}

func (w whatsapp) generatePayload(message, destination string) string {
	return "message=" + utils.EncodeMessage(message) + "&channel=whatsapp&source=" + w.Source + "&destination=" + destination + "&src.name=" + w.AppName + "&encode=true"
}

func (w whatsapp) addHeaders() map[string][]string {
	return map[string][]string{
		"Accept":       {"application/json"},
		"apikey":       {w.Token},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
}

func NewGupshupWhatsapp(source string, appName string, token string) interfaces.Whatsapp {
	return whatsapp{Source: source, AppName: appName, Token: token}
}
