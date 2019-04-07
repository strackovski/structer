// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package cmd

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"

    "github.com/spf13/cobra"

    "github.com/nv3-x/structer/pkg/structer"
)

var in string
var out string
var format string

// makeCmd represents the make command
var makeCmd = &cobra.Command{
    Use:   "make",
    Short: "Makes a struct from JSON data.",
    Long:  ``,
    Run: func(cmd *cobra.Command, args []string) {
        if in == "" {
            return
        }

        // Check if path is relative or absolute.
        if strings.HasPrefix(out, "/") != true {
            workDir, err := os.Getwd()
            structer.Check(err)

            out = workDir + "/" + out
        }

        if _, err := os.Stat(out); os.IsNotExist(err) {
            err = os.Mkdir(out, 0755)
        }

        var fileName string
        var contents string

        fi, err := os.Stat(in)
        structer.Check(err)

        // Check if input is a file or a directory.
        switch mode := fi.Mode(); {
        case mode.IsDir():
            files, err := ioutil.ReadDir(in)
            structer.Check(err)

            for _, f := range files {
                fmt.Println(f.Name())
                fileName, contents, err = readFile(in + "/" + f.Name())
                makeFile(string(contents), fileName, "generated")
            }

        case mode.IsRegular():
            fileName, contents, err = readFile(in)
            structer.Check(err)

            makeFile(string(contents), fileName, "generated")
        }
    },
}

// makeFile writes contents to a file.
func makeFile(contents, fileName, packageName string) {
    structName := strings.Title(strings.Split(fileName, ".")[0])
    if outD, e := structer.Generator().Generate(strings.NewReader(string(contents)), structName, packageName); e == nil {
        err := structer.File(outD, fileName, out)
        structer.Check(err)
    } else {
        fmt.Println("An error occurred.", e)
    }
}

// readFile prepares file contents for processing.
func readFile(file string) (string, string, error) {
    contents, err := ioutil.ReadFile(file)
    structer.Check(err)

    fileParts := strings.Split(file, "/")
    fileParts = strings.Split(fileParts[len(fileParts)-1], ".")
    fileName := fileParts[0]

    if format != "json" {
        // convert contents from {format} to json
        // contents = YamlToJson(contents)
        // contents = XmlToJson(contents)
    }

    return fileName, string(contents), nil
}

// init ...
func init() {
    rootCmd.AddCommand(makeCmd)
    makeCmd.PersistentFlags().StringVar(&in, "in", "", "path to input file or directory")
    makeCmd.PersistentFlags().StringVar(&out, "out", "generated", "path to output directory")
    makeCmd.PersistentFlags().StringVar(&format, "format", "json", "the format of the input")
}
