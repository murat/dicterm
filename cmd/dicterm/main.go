package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/murat/dicterm/internal/api"
	"github.com/murat/dicterm/internal/config"

	"github.com/olekukonko/tablewriter"
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

	var resp []api.Collegiate
	if err := json.Unmarshal(r, &resp); err != nil {
		fmt.Printf("could not unmarshal response, err: %v\n", err)
		os.Exit(1)
	}

	green := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Definition", "Stems", "Etymology"})
	table.SetRowLine(true)

	for _, r := range resp {
		defs := strings.Join(r.Shortdef, ", ")
		stems := strings.Join(r.Meta.Stems, ", ")
		etym := ""
		for _, e := range r.Etymologies {
			etym = strings.TrimSpace(etym + "\n" + strings.Join(e, ", "))
		}
		table.Append([]string{green(word + "(" + r.FunctionalLabel + ")"), defs, stems, etym})
	}
	table.Render()
}
