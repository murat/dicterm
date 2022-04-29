package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/murat/go-boilerplate/internal/api"

	"github.com/fatih/color"
)

var (
	key, word string
)

func main() {
	flag.StringVar(&key, "key", "", "dict api key")
	flag.StringVar(&word, "word", "", "word to look up")
	flag.Parse()

	c := api.NewClient(key)
	r, err := c.Get(word)
	if err != nil {
		fmt.Printf("could not get response, err: %v", err)
		os.Exit(1)
	}

	var resp api.Response
	if err := json.Unmarshal(r, &resp); err != nil {
		fmt.Printf("could not unmarshal response, err: %v", err)
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
