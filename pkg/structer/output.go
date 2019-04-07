// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import (
    "io/ioutil"
)

// Builder is a function type that builds a resource named outputName at address
// outputPath, and populates it with input.
type Builder func(input []byte, outputName, outputPath string) error

// File is a concrete Builder implementation for processing file resources.
func File(input []byte, outputName, outputPath string) error {
    err := ioutil.WriteFile(outputPath+"/"+outputName+".go", input, 0755)
    Check(err)

    return nil
}
