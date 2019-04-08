// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
"encoding/json"

"gopkg.in/yaml.v2"
)

// YamlToJsonEncoder
type YamlToJsonDecoder struct {
}

// Decode is a ...
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

// convert lets yaml unmarshal the input into a value of type interface{}, then
// loops the result recursively, converting each encountered map[interface{}]interface{}
// to a map[string]interface{} value in the process.
func (e YamlToJsonDecoder) convert(i interface{}) interface{} {
    switch x := i.(type) {
    case map[interface{}]interface{}:
        m2 := map[string]interface{}{}
        for k, v := range x {
            m2[k.(string)] = e.convert(v)
        }
        return m2
    case []interface{}:
        for i, v := range x {
            x[i] = e.convert(v)
        }
    }
    return i
}
