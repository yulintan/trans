package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yulintan/translate/google"
	"github.com/yulintan/translate/translator"
	"github.com/yulintan/translate/youdao"
)

var cfg translator.Config
var dictionary translator.Dictionary

func config(tl, provider, brief string) {
	key := os.Getenv("KEY")
	if key == "" {
		key = "60638690"
	}

	if tl == "" {
		tl = "en"
	}
	bf := true
	if brief == "false" {
		bf = false
	}

	cfg = translator.Config{
		KeyFrom:        "youdao111",
		Key:            key,
		Type:           "data",
		DocType:        "json",
		TargetLanguage: tl,
		Brief:          bf,
	}

	switch provider {
	case "":
	case "google":
		dictionary = google.NewDic(cfg)
	case "youdao":
		dictionary = youdao.NewDic(cfg)
	default:
		log.Fatal("current supported translate providers are google and youdao")
	}
}

func main() {
	var tl, provider, brief string

	app := &cli.App{
		Name:  "translate",
		Usage: "translation between English and Chinese",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "target",
				Aliases:     []string{"t"},
				Value:       "en",
				Usage:       "target language which will be translated to, list of supported language codes at https://cloud.google.com/translate/docs/languages",
				Destination: &tl,
			},
			&cli.StringFlag{
				Name:        "provider",
				Aliases:     []string{"p"},
				Value:       "google",
				Usage:       "translator providers, e.g. google, youdao",
				Destination: &provider,
			},
			&cli.StringFlag{
				Name:        "brief",
				Aliases:     []string{"b"},
				Value:       "true",
				Usage:       "return brief defination of the given word(s)",
				Destination: &brief,
			},
		},
		Action: func(c *cli.Context) error {
			config(tl, provider, brief)
			word := c.Args().Get(0)
			if word == "" {
				log.Fatal("Please enter the word you want to translate")
			}

			trans := translator.NewTranslator(dictionary)
			result, err := trans.Translate(word)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
