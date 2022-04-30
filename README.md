# dicterm

![Lint](https://github.com/murat/dicterm/actions/workflows/lint.yml/badge.svg)
![Tests](https://github.com/murat/dicterm/actions/workflows/test.yml/badge.svg)

dicterm is a cli interface for merriem webster dictionary.

## Requirements

Get your api key from [https://dictionaryapi.com/account/my-keys](https://dictionaryapi.com/account/my-keys)

## Build

```shell
➜ make build
```

## Usage

```shell
➜ ./bin/dicterm -h               
Usage of ./bin/dicterm:
  -key string
        dict api key
  -word string
        word to look up

➜ ./bin/dicterm -key XXXXXX -word dictionary
dictionary(noun)

definitions:
- a reference source in print or electronic form containing words usually alphabetically arranged along with information about their forms, pronunciations, functions, etymologies, meanings, and syntactic and idiomatic uses
- a reference book listing alphabetically terms or names important to a particular subject or activity along with discussion of their meanings and applications
- a reference book listing alphabetically the words of one language and showing their meanings or translations in another language

stems:
- dictionaries, dictionary

====================
```

## Todo

  - [ ] store key in a dotfile
  - [ ] use a beatiful tui
