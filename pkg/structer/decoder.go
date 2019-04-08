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

// ToJsonDecoder is the interface that describes the simple decoder types
// that take care of decoding input data when it's not in json format,
// in order to feed the json-based Generator.
//
// Because the generator accepts json as the input, all decoders decode to
// json format only.
type ToJsonDecoder interface {
    Decode(input string) ([]byte, error)
}

// DecoderFactory is a factory of ToJsonDecoder instances.
type DecoderFactory struct {
    supportedTypes []string
}

// NewDecoderFactory is a ...
func NewDecoderFactory() DecoderFactory {
    types := make([]string, 1)
    types = append(types, "yaml")
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
    default:
        // A format is registered as supported, but no suitable decoder exists, this
        // should not happen (todo panic?).
        return nil, errors.New(fmt.Sprintf("no decoder for format type %s found", format))
    }

    return decoder, nil
}
