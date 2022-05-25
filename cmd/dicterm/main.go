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
	if err := run(os.Args); err != nil {
		fmt.Printf("an error occurred, %v\n", err)
		fmt.Println("please run `dicterm -h`")
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing arguments")
	}

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&key, "key", "", "dict api key (will be stored in ~/.dicterm)")
	flags.StringVar(&configPath, "config", "", "path of config file, default ~/.dicterm")
	flags.StringVar(&word, "word", "", "word to look up")
	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("flag parse error, %w", err)
	}

	cfg, err := config.New(configPath)
	if err != nil {
		return fmt.Errorf("could not open config file, %w", err)
	}
	defer cfg.Close()

	if key != "" {
		n, err := cfg.Write([]byte(key))
		if err != nil {
			return fmt.Errorf("could not write key to config file, %w", err)
		}
		if n != len(key) {
			return fmt.Errorf("could not write all bytes to config file, %w", err)
		}
	} else {
		buf := make([]byte, 36) // api key is UUID(36 bytes)
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
		return fmt.Errorf("please run `dicterm -h`")
	}

	if word == "" {
		word = flags.Args()[0]
	}

	c := mwgoapi.NewClient(&http.Client{}, mwgoapi.BaseURL, key)
	r, err := c.Get(word)
	if err != nil {
		return fmt.Errorf("api request failed, %w", err)
	}

	var resp []mwgoapi.Collegiate
	if err := json.Unmarshal(r, &resp); err != nil {
		return fmt.Errorf("marshal response body failed, %w", err)
	}

	if len(resp) == 0 {
		fmt.Printf("word not found: %s\n", word)
	}

	table(resp)

	return nil
}

func table(resp []mwgoapi.Collegiate) {
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
