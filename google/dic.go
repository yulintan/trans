package google

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
	endpoint := "https://translate.googleapis.com/translate_a/single"
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	parameters := url.Values{}
	parameters.Add("client", "gtx")
	parameters.Add("sl", "auto")
	parameters.Add("tl", d.cfg.TargetLanguage)
	parameters.Add("dt", "t")
	parameters.Add("q", word)
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

	var wrapper []interface{}
	var transResult [][]interface{}
	var result string

	err := json.Unmarshal([]byte(s), &wrapper)
	if err != nil {
		return "", err
	}
	temp, err := json.Marshal(wrapper[0])
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(temp, &transResult)
	if err != nil {
		return "", err
	}

	// for _, r := range transResult[0] {
	// 	switch v := r.(type) {
	// 	case string:
	// 		result += v + "\n"
	// 	}
	// }

	// if len(result) > 0 {
	// 	result = result[:len(result)-1]
	// }

	result, ok := transResult[0][0].(string)
	if !ok {
		return "", err
	}

	return result, nil
}
