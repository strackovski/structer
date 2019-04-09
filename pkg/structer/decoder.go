// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
    "errors"
    "fmt"
)

// ToJsonDecoder is the interface that describes the decoder types
// for decoding input data in formats different from default (JSON),
// in order to feed the JSON-based Generator. Because the generator
// accepts JSON as the input, all decoders decode to JSON format only.
type ToJsonDecoder interface {
    Decode(input string) ([]byte, error)
}

// remapper builds the key:value map by unmarshaling the input into a value
// of type interface{}, then looping the result recursively, converting
// each encountered map[interface{}]interface{} to a map[string]interface{}
// value in the process.
type Remapper func(i interface{}) interface{}

// DecoderFactory is a factory of ToJsonDecoder instances.
type DecoderFactory struct {
    supportedTypes []string
}

// NewDecoderFactory is a constructor method for DecoderFactory type.
func NewDecoderFactory() DecoderFactory {
    types := make([]string, 1)
    types = append(types, "yaml")
    types = append(types, "hcl")
    return DecoderFactory{supportedTypes: types}
}

// Create checks if the given format is supported and delegates finding
// the correct decoder to the internal find method.
func (f DecoderFactory) Create(format string) (ToJsonDecoder, error) {
    set := make(map[string]bool)
    for _, v := range f.supportedTypes {
        set[v] = true
    }

    if set[format] {
        // Format is valid.
        return f.find(format)
    }

    return nil, errors.New(fmt.Sprintf("the given type %s is not supported", format))
}

// find is a method that creates and returns the decoder instance if
// a decoder supporting the given format exists.
func (f DecoderFactory) find(format string) (ToJsonDecoder, error) {
    var decoder ToJsonDecoder
    switch format {
    case "yaml":
        decoder = YamlToJsonDecoder{}
        break
    case "hcl":
        decoder = HclToJsonDecoder{}
        break
    default:
        // A format is registered as supported, but no suitable decoder exists, this
        // should not happen (todo panic?).
        return nil, errors.New(fmt.Sprintf("no decoder for format type %s found", format))
    }

    return decoder, nil
}

// remap is the default concrete implementation of the remapper type.
func remap(i interface{}) interface{} {
    switch x := i.(type) {
    case map[interface{}]interface{}:
        m2 := map[string]interface{}{}
        for k, v := range x {
            m2[k.(string)] = remap(v)
        }
        return m2
    case []interface{}:
        for i, v := range x {
            x[i] = remap(v)
        }
    }
    return i
}
