# JSON to YAML and back converter

This is a simple tool to convert JSON to YAML and back.

## Usage

```shell
json2yaml < input.json > output.yaml
```
To convert YAML to JSON, create a symlink to the executable named `yaml2json`
and use it the same way:

```shell
ln -s json2yaml yaml2json
yaml2json < input.yaml > output.json
```

## Installation

```sh
go install github.com/rusq/json2yaml@latest
```
