package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yulintan/translate/translator"
	"github.com/yulintan/translate/youdao"
)

var dictionary translator.Dictionary

func init() {
	key := os.Getenv("KEY")
	if key == "" {
		key = "60638690"
	}
	cfg := translator.Config{
		KeyFrom: "youdao111",
		Key:     key,
		Type:    "data",
		DocType: "json",
	}

	// hard code to youdo for now
	dictionary = youdao.NewDic(cfg)
}

func main() {
	app := &cli.App{
		Name:  "translate",
		Usage: "translation between English and Chinese",
		Action: func(c *cli.Context) error {
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
