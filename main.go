package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

const (
	jsonyamlUsage = "json2yaml converts JSON to YAML.\nusage: json2yaml < file.json > file.yaml"
	yamljsonUsage = "yaml2json converts YAML to JSON.\nusage: yaml2json < file.yaml > file.json"
)

// converter is the function that will execute the conversion.  It is set in
// init, depending on the executable name and defaults to jsonyaml.
var (
	converter = jsonyaml
	helpstr   = jsonyamlUsage
)

func init() {
	if isYamlMode() {
		converter = yamljson
		helpstr = yamljsonUsage
	}
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), helpstr)
		flag.PrintDefaults()
	}
}

func main() {
	// the only flag supported is -h
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	in, out := os.Stdin, os.Stdout
	defer out.Close()

	return converter(in, out)
}

func jsonyaml(in io.Reader, out io.Writer) error {
	dec, enc := json.NewDecoder(in), yaml.NewEncoder(out)
	return convert(dec, enc)
}

func yamljson(in io.Reader, out io.Writer) error {
	dec, enc := yaml.NewDecoder(in), json.NewEncoder(out)
	return convert(dec, enc)
}

type encoder interface {
	Encode(v any) error
}

type decoder interface {
	Decode(v any) error
}

func convert(dec decoder, enc encoder) error {
	var buf any
	for {
		if err := dec.Decode(&buf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if err := enc.Encode(buf); err != nil {
			return err
		}
	}
}

func isYamlMode() bool {
	exe := filepath.Base(os.Args[0])
	return exe == "yaml2json" || exe == "y2j" || exe == "y2j.exe" || exe == "yaml2json.exe"
}
