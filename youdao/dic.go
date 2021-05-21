package youdao

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/yulintan/trans/translator"
)

type dic struct {
	cfg translator.Config
}

func NewDic(cfg translator.Config) translator.Dictionary {
	return &dic{
		cfg: cfg,
	}
}

func (d *dic) Translate(word string) (string, error) {
	endpoint := "http://fanyi.youdao.com/openapi.do"
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	parameters := url.Values{}
	parameters.Add("q", word)
	parameters.Add("keyfrom", d.cfg.KeyFrom)
	parameters.Add("key", d.cfg.Key)
	parameters.Add("type", d.cfg.Type)
	parameters.Add("doctype", d.cfg.DocType)
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

	return string(body), nil
}

func (d *dic) PrettyPrint(s string) (string, error) {
	if d.cfg.Brief == false {
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, []byte(s), "", "  "); err != nil {
			return "", nil
		}

		return dst.String(), nil
	}

	var transResult TransResult
	var result string

	err := json.Unmarshal([]byte(s), &transResult)
	if err != nil {
		return "", err
	}

	for _, tran := range transResult.Translation {
		result += tran + "\n"
	}

	for _, ex := range transResult.Basic.Explains {
		result += ex + "\n"
	}

	if len(result) > 0 {
		result = result[:len(result)-1]
	}

	return result, nil
}
