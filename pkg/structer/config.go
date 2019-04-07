// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

type Config interface {
    Initialisms() map[string]bool
    IntWordMap() []string
    OutputDir() string
}

// Config repesents a set of configuration properties.
type config struct {
    initialisms map[string]bool // A set of common initialisms.
    intWordMap  []string        // Integer (number) to word map,
    outputDir   string          // Path to default output destination directory.
}

// NewConfig constructs a new Config instance with sensisble
// defaults.
func Default() Config {
    c := config{}
    c.initialisms = initialisms
    c.outputDir = "generated"

    return c
}

// Initialisms is a set of common initialisms.
// See "common.go".
func (c config) Initialisms() map[string]bool {
    return c.initialisms
}

// OutputDir is the path to the generator's destination directory.
// The value of OutputDir should be the path to the directory
// where the generator can write to files.
func (c config) OutputDir() string {
    return c.outputDir
}

// IntWordMap is a mapping from integer (0-9) to the English word
// corresponding to the number.
func (c config) IntWordMap() []string {
    return []string{
        "zero",
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }
}
