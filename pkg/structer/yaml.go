// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
    "encoding/json"
    "fmt"

    "gopkg.in/yaml.v2"
)

// YamlToJsonEncoder
type YamlToJsonDecoder struct {
}

// Decode decodes input from YAML to JSON.
func (e YamlToJsonDecoder) Decode(input string) ([]byte, error) {
    var body interface{}
    if err := yaml.Unmarshal([]byte(input), &body); err != nil {
        fmt.Printf("xxxx")
        panic(err)
    }

    body = remap(body)

    if b, err := json.Marshal(body); err != nil {
        panic(err)
    } else {
        return b, nil
    }
}
