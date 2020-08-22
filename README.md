# ec: easy csv rendering

## Purpose

GNU-like terminal tool for rendering CSV files

## Usage

## 1. Parameter

```bash
$ ec -f my_data.csv

┌─────┬─────┬─────┐
│ ABC │ DEF │ GHJ │
└─────┴─────┴─────┘

```

## 2. Stdin

```bash

$ echo abc,def,ghj | ec

┌─────┬─────┬─────┐
│ ABC │ DEF │ GHJ │
└─────┴─────┴─────┘

```

It is designed to be used in a series of GNU commands:

```bash
echo "great,row,full,of,goodies\nboring,row,that,i,dislike" | grep great | ec

┌───────┬─────┬──────┬────┬─────────┐
│ GREAT │ ROW │ FULL │ OF │ GOODIES │
└───────┴─────┴──────┴────┴─────────┘
```



## Help

```bash
$ ec --help
Usage of ec:
  -f string
        filename
```

# Roadmap

- Filter rows by search, retaining headers
- tsv support
- autodetect delimiter
- styling parameter
