# ec: easy csv rendering

## Purpose

GNU-like terminal tool for rendering CSV files

## Usage

## 1. Parameter

```
$ ec -f my_data.csv

┌─────┬─────┬─────┐
│ ABC │ DEF │ GHJ │
└─────┴─────┴─────┘

```

## 2. Stdin

```

$ echo abc,def,ghj | ec

┌─────┬─────┬─────┐
│ ABC │ DEF │ GHJ │
└─────┴─────┴─────┘

```

## Help

```
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
