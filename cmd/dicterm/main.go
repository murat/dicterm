package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/murat/go-boilerplate/internal/config"
	"net/http"
	"os"
	"strings"

	"github.com/murat/go-boilerplate/internal/api"

	"github.com/fatih/color"
)

var (
	key, configPath, word string
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please run `dicterm -h`")
		os.Exit(1)
	}

	flag.StringVar(&key, "key", "", "dict api key (will be stored in ~/.dicterm)")
	flag.StringVar(&configPath, "config", "", "path of config file, default ~/.dicterm")
	flag.StringVar(&word, "word", "", "word to look up")
	flag.Parse()

	cfg, err := config.New(configPath)
	if err != nil {
		fmt.Printf("could not access the config file, %v\n", err)
		os.Exit(1)
	}
	defer cfg.File.Close()

	if key != "" {
		err := cfg.Write(key)
		if err != nil {
			fmt.Printf("could not write key to config file, %v\n", err)
			os.Exit(1)
		}
	} else {
		k, err := cfg.Read()
		if err != nil {
			switch {
			case errors.Is(err, config.ErrEmptyFile):
				fmt.Println("please run `dicterm -h`")
			default:
				fmt.Printf("could not read key, %v\n", err)
			}
			os.Exit(1)
		}
		key = *k
	}

	if word == "" {
		word = os.Args[1]
	}

	c := api.NewClient(&http.Client{}, api.BaseURL, key)
	r, err := c.Get(word)
	if err != nil {
		fmt.Printf("could not get response, err: %v\n", err)
		os.Exit(1)
	}

	var resp api.Response
	if err := json.Unmarshal(r, &resp); err != nil {
		fmt.Printf("could not unmarshal response, err: %v\n", err)
		os.Exit(1)
	}

	highlight := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	bold := color.New(color.Bold, color.Faint).SprintFunc()

	for _, r := range resp {
		defs := strings.Join(r.Shortdef, "\n- ")
		stems := strings.Join(r.Meta.Stems, ", ")
		fmt.Printf("%s\n\n%s:\n- %s\n\n%s:\n- %s\n\n",
			highlight(word+"("+r.FL+")"),
			bold("definitions"),
			defs,
			bold("stems"),
			stems)
		fmt.Printf("%s\n\n", bold(strings.Repeat("=", 20)))
	}
}
