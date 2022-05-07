package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	defer cfg.Close()

	if key != "" {
		n, err := cfg.Write([]byte(key))
		if err != nil {
			fmt.Printf("could not write key to config file, %v\n", err)
			os.Exit(1)
		}
		if n != len(key) {
			fmt.Printf("could not write all bytes to config file, %v\n", err)
			os.Exit(1)
		}
	} else {
		buf := make([]byte, 36) // api key is UUID(32 bytes)
		for {
			n, err := cfg.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("could not read key from config file, %v\n", err)
				break
			}
			if n > 0 {
				key = string(buf[:n])
			}
		}
	}

	if key == "" {
		fmt.Println("please run `dicterm -h`")
		os.Exit(1)
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
	if err := json.Unmarshal(r, &resp); err != nil {
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
	header := []string{"", "Definition", "Stems", "Etymology"}
	data := [][]string{}

	for _, r := range resp {
		defs := strings.Join(r.Shortdef, ",")
		fl := ""
		if r.FunctionalLabel != "" {
			fl = "(" + r.FunctionalLabel + ")"
		}
		stems := strings.Join(r.Meta.Stems, ", ")

		etym := ""
		for _, e := range r.Etymologies {
			etym = strings.TrimSpace(etym + "\n" + strings.Join(e, ", "))
		}

		data = append(data, []string{green(word + fl), defs, stems, etym})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(data)
	table.Render()
}
