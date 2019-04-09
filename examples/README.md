<h3 align="center">
	<img
		width="40"
		alt="NV3"
		src="https://s3.eu-central-1.amazonaws.com/clogger/media/nv3-bw.png">
</h3>

<h1 align="center">
	GO STRUCTER EXAMPLES
</h1>

<p align="center">
	<strong>
		<a href="https://github.com/strackovski/structer">Home</a>
		•
		<a href="https://github.com/strackovski/structer/wiki">Wiki</a>
		•
		Examples
	</strong>
</p>

# Introduction
The structure of the data provided as JSON should be an object:

```json
{
  "id": 1,
  "name": "Name",
  "age": 30
}
```

The key names become struct's property names. The sample values are used to determine the property type.

Because an input file should represent a single Go struct type, an array of objects will result in errors. For example, this will **not** work:

```json
[
   {
   "id": 1,
     "name": "Name",
     "age": 30
   }
]
```

The examples below cover common use cases:


> #### 1. [Basic example](#1-basic-example)
> #### 2. [Batch operation example](#2-batch-operation-example)
> #### 3. [Specifying file path](#3-specifying-file-path)
> #### 4. [Specifying file format explicitly](#4-specifying-file-format-explicitly)
> #### 5. [Handling errors](#5-handling-errors)

## 1. Basic example

With a [~/user.json](input/json/user.json) file containing a json definition for the `user` type:

```json
{
  "id": 1,
  "email": "test@email.net",
  "name": "Name",
  "role": "ROLE_USER"
}
```

Executing the following command:

```bash
structer make --in ~/user.json
```

Results in `generated/model/user.go`:

```go
package main

type MyJsonName struct {
  Age  int    `json:"age"`
  Name string `json:"name"`
}
```

## 2. Batch operation example

The `--in` flag can be used to process directories as well. Like with a single file, the directory path must be absolute as well. If a relative path is given, the resolver will assume current directory as the root.

```bash
structer make --in ~/models
```

## 3. Specifying output path

By default, structer will write generated files to the `generated` directory under project root. Use the `--out` flag to direct output to a custom path. Requires absolute path, or project root is assumed as root when resolving.

```bash
structer make --in ~/models --out ~/models/generated

ls -al ~/models/generated
total 3
drwxr-xr-x   4 ubuntu  ubuntu 128 Apr  7 10:36 .
drwxr-xr-x  22 ubuntu  ubuntu 704 Apr  7 16:47 ..
-rw-r--r--   1 ubuntu  ubuntu  94 Apr  7 10:06 user.go
-rw-r--r--   1 ubuntu  ubuntu  94 Apr  7 10:06 role.go
-rw-r--r--   1 ubuntu  ubuntu  94 Apr  7 10:06 team.go
```


## 4. Specifying file format explicitly

Any format which can be decoded to json can be provided as the `--format` parameter, if a suitable decoder exists:

```bash
structer make --in ~/models --format yaml
```

A structer decoder is a simple interface:

```go
type ToJsonDecoder interface {
    Decode(input string) ([]byte, error)
}
```

YAML is the only built-in decoder by default. To implement additional decoders, take a look at the available one 
below for reference.

**A concrete implementation for YAML** (file `pkg/structer/yaml.go`):

```go
package structer

import (
    "encoding/json"

    "gopkg.in/yaml.v2"
)

type YamlToJsonDecoder struct {
}

func (e YamlToJsonDecoder ) Decode(input string) ([]byte, error) {
    var body interface{}
    if err := yaml.Unmarshal([]byte(input), &body); err != nil {
        panic(err)
    }

    body = e.convert(body)

    if b, err := json.Marshal(body); err != nil {
        panic(err)
    } else {
        return b, nil
    }
}

func (e YamlToJsonDecoder) convert(i interface{}) interface{} {
    // decoder internal
}
```

## 5. Errors

On failure structer will return an error message if possible, or panic otherwise.

<h1 align="right">
	<img
		width="20"
		alt="Logo"
		src="https://s3.eu-central-1.amazonaws.com/clogger/media/nv3-bw.png">
</h1>
