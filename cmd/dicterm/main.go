package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/murat/dicterm/internal/config"

	"github.com/fatih/color"
	"github.com/murat/mwgoapi"
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
			case errors.Is(err, mwgoapi.ErrEmptyConfig):
				fmt.Println("please run `dicterm -h`")
			default:
				fmt.Printf("could not read key, %v\n", err)
			}
			os.Exit(1)
		}
		key = *k
	}

	if word == "" {
		word = flag.Args()[0]
	}

	c := mwgoapi.NewClient(&http.Client{}, mwgoapi.BaseURL, key)
	r, err := c.Get(word)
	if err != nil {
		fmt.Printf("could not get response, err: %v\n", err)
		os.Exit(1)
	}

	var resp []mwgoapi.Collegiate
	if err := c.UnmarshalResponse(r, &resp); err != nil {
		fmt.Printf("could not unmarshal response, err: %v\n", err)
		os.Exit(1)
	}

	if len(resp) == 0 {
		fmt.Printf("could not find word: %s\n", word)
		os.Exit(1)
	}

	Print(resp)
}

func Print(resp []mwgoapi.Collegiate) {
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
