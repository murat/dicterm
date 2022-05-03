# dicterm

![Lint](https://github.com/murat/dicterm/actions/workflows/lint.yml/badge.svg)
![Tests](https://github.com/murat/dicterm/actions/workflows/test.yml/badge.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/295007f859ca44b5b1a9418bb5685d40)](https://www.codacy.com/gh/murat/dicterm/dashboard?utm_source=github.com&utm_medium=referral&utm_content=murat/dicterm&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/295007f859ca44b5b1a9418bb5685d40)](https://www.codacy.com/gh/murat/dicterm/dashboard?utm_source=github.com&utm_medium=referral&utm_content=murat/dicterm&utm_campaign=Badge_Coverage)

dicterm is a cli interface for merriem webster dictionary.

## Requirements

Get your api key from <https://dictionaryapi.com/account/my-keys>

## Build

```shell
➜ make build
```

## Usage

If it's the first run, you must pass the `-key XXX` argument.
It will be stored in the `~/.dicterm` file.
You can specify a custom path via `-config /path/of/file`.
But if you set a custom path, you will need to pass it by the `-config /path/of/file` argument all time.

If you did not set a custom config path, you will be able to run the
simpler command like `dicterm word` after the first command.

```shell
➜ ./bin/dicterm -h
Usage of ./bin/dicterm:
  -config string
        path of config file, default ~/.dicterm
  -key string
        dict api key (will be stored in ~/.dicterm)
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

➜ ./bin/dicterm word
word(noun)

definitions:
- a speech sound or series of speech sounds that symbolizes and communicates a meaning usually without being divisible into smaller units capable of independent use
- the entire set of linguistic forms produced by combining a single base with various inflectional elements without change in the part of speech elements
- a written or printed character or combination of characters representing a spoken word —sometimes used with the first letter of a real or pretended taboo word prefixed as an often humorous euphemism      

stems:
- a good word, good word, in a word, in so many words, of few words, of her word, of his word, of its words, of my word, of one's word, of our word, of their word, of your word, the good word, upon my word, word, words

====================
```

## Todo

-   [x] store key in a dotfile
-   [ ] use a beatiful tui
