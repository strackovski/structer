// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
    "encoding/json"

    "github.com/hashicorp/hcl"
)

// HclToJsonDecoder
type HclToJsonDecoder struct {
}

// Decode decodes input from HCL to JSON.
func (e HclToJsonDecoder) Decode(input string) ([]byte, error) {
    var body interface{}
    if err := hcl.Unmarshal([]byte(input), &body); err != nil {
        panic(err)
    }

    body = remap(body)

    if b, err := json.Marshal(body); err != nil {
        panic(err)
    } else {
        return b, nil
    }
}
