package main

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

type (
	reader struct {
		regex string
	}
)

func NewReader() *reader {
	return &reader{
		regex: `\t(modified|deleted|"new file"):\s*(.*)`,
	}
}

func (r reader) Status() ([]Status, error) {
	list := []Status{}

	// TODO add "new file" also
	regularExpression := regexp.MustCompile(r.regex)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res := regularExpression.FindStringSubmatch(scanner.Text())

		if len(res) == 0 {
			continue
		}

		list = append(list, Status{
			Operation: res[1],
			File:      res[2],
		})
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return list, nil
}

func (r reader) Configurations() ([]Configuration, error) {
	f, err := os.Open("./check.config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Configurations
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg.Configuration, nil
}
