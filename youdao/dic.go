package youdao

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/yulintan/translate/translator"
)

type dic struct {
	cfg translator.Config
}

func NewDic(cfg translator.Config) translator.Dictionary {
	return &dic{
		cfg: cfg,
	}
}

func (s *dic) Translate(word string) (string, error) {
	endpoint := "http://fanyi.youdao.com/openapi.do"
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	parameters := url.Values{}
	parameters.Add("q", word)
	parameters.Add("keyfrom", s.cfg.KeyFrom)
	parameters.Add("key", s.cfg.Key)
	parameters.Add("type", s.cfg.Type)
	parameters.Add("doctype", s.cfg.DocType)
	parameters.Add("version", "1.1")
	u.RawQuery = parameters.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, body, "", "  "); err != nil {
		return "", nil
	}

	return dst.String(), nil
}
