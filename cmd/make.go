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
    "time"

    tm "github.com/buger/goterm"
    "github.com/fatih/color"
    "github.com/spf13/cobra"

    "github.com/nv3-x/structer/pkg/structer"
)

var in string
var out string
var format string
var initTasks []string

// makeCmd represents the make command
var makeCmd = &cobra.Command{
    Use:   "make",
    Short: "Makes structs from input file(s)",
    Long: `Generates struct implementations and writes them to files.

Make takes a file or a directory of files as input, reads their contents, 
and generates source code that implements Go structs with data structure 
corresponding to the structure discovered in the contents of the input 
file(s).`,
    Run: func(cmd *cobra.Command, args []string) {
        if _, err := os.Stat(in); os.IsNotExist(err) {
            tm.Clear()
            tm.MoveCursor(1, 1)
            text := fmt.Sprintf("%-48v", "  ➭  No such file or directory.")
            alert(0, 1, tm.Background(text, 1))
            tm.Flush()
            return
        }

        if format != "yaml" && format != "json" && format != "hcl" {
            tm.Clear()
            tm.MoveCursor(1, 1)
            alert(0, 1, fmt.Sprintf("  ➭  Format '%s' is not supported.", format))
            tm.Flush()
            return
        }

        tm.Clear()
        tm.MoveCursor(1, 1)
        alert(2, 0, "⚡  Structing")

        _, _ = tm.Println(tm.Bold(tm.Color("▶  "+in, tm.GREEN)))
        tm.Flush()

        for _, x := range initTasks {
            _, _ = tm.Println(tm.Color(x, tm.CYAN))
            tm.Flush()
            time.Sleep(time.Millisecond * 100)
        }

        tm.MoveCursor(28, 3)
        _, _ = tm.Println(tm.Color("✓", tm.GREEN))
        tm.MoveCursor(1, 3)
        tm.Flush()

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
        var contents []byte

        fi, err := os.Stat(in)
        structer.Check(err)

        // Check if input is a file or a directory.
        switch mode := fi.Mode(); {
        case mode.IsDir():
            files, err := ioutil.ReadDir(in)
            structer.Check(err)
            var tasks2 []string
            tasks2 = append(tasks2, fmt.Sprintf("➜  Format specified is '%s'", format))

            if len(files) < 0 {
                errString := fmt.Sprintf("\n➜  ERROR: No %s files found in the specified path, aborting   ✗\n", format)
                errString = tm.Bold(tm.Color(errString, tm.WHITE))
                tm.Flush()
                _, _ = tm.Println(tm.Background("ddd", tm.RED))

                return

            } else {
                tasks2 = append(tasks2, tm.Color(fmt.Sprintf("➜  Found %d %s files.", len(files), format), tm.GREEN))
            }

            tasks2 = append(tasks2, tm.Color("➜  Veryfing...", tm.WHITE))
            tasks2 = append(tasks2, tm.Color(fmt.Sprintf("➜  %d/%d files OK. ", len(files), len(files)), tm.GREEN))

            for _, x := range tasks2 {
                _, _ = tm.Println(x)
                tm.Flush()
            }

            for i, f := range files {
                fileName, contents, err = readFile(in + "/" + f.Name())
                tm.MoveCursor(1, 7)
                _, _ = tm.Printf("➜  Processing file %d/%d\n", i+1, len(files))
                _, _ = tm.Printf(tm.Color("➜  %s.%s", tm.YELLOW), fileName, format)
                tm.Flush()
                makeFile(contents, fileName, "generated")
                time.Sleep(time.Millisecond * 100)
            }

            tm.MoveCursor(1, 7)
            _, _ = tm.Println(tm.Color("✓", tm.GREEN))
            tm.MoveCursor(1, 8)
            alert(1, 1, "  ✓✓  Structing complete.")
            tm.Flush()

        case mode.IsRegular():
            color.HiWhite("   ➜  File")
            fileName, contents, err = readFile(in)
            structer.Check(err)
            makeFile(contents, fileName, "generated")

        default:
            alert(0, 1, "No such file or directory.")
        }
    },
}

// makeFile writes contents to a file.
func makeFile(contents []byte, fileName, packageName string) []byte {
    structName := strings.Title(strings.Split(fileName, ".")[0])
    if outD, e := structer.Generator().Generate(contents, structName, packageName); e == nil {
        err := structer.File(outD, fileName, out)
        structer.Check(err)
        return outD
    } else {
        fmt.Println("An error occurred.", e)
    }

    return nil
}

// readFile prepares file contents for processing.
func readFile(file string) (string, []byte, error) {
    contents, err := ioutil.ReadFile(file)
    structer.Check(err)

    fileParts := strings.Split(file, "/")
    fileParts = strings.Split(fileParts[len(fileParts)-1], ".")
    fileName := fileParts[0]

    if format != "json" {
        // Get decoder factory if format is valid.
        factory := structer.NewDecoderFactory()
        decoder, err := factory.Create(format)
        structer.Check(err)

        // Decode {format} to json
        contents, err = decoder.Decode(string(contents))
        structer.Check(err)
    }

    return fileName, contents, nil
}

// init ...
func init() {
    rootCmd.AddCommand(makeCmd)
    makeCmd.PersistentFlags().StringVar(&in, "in", "", "path to input file or directory")
    makeCmd.PersistentFlags().StringVar(&out, "out", "generated", "path to output directory")
    makeCmd.PersistentFlags().StringVar(&format, "format", "json", "the format of the input json|yaml|hcl")
}

func alert(level, padding int, text string) {

    var fg, bg int

    if level == 0 {
        fg = tm.WHITE
        bg = tm.RED
    } else if level == 1 {
        fg = tm.WHITE
        bg = tm.GREEN
    } else if level == 2 {
        fg = tm.WHITE
        bg = tm.BLUE
    } else {
        os.Exit(1)
    }

    s := "%-48v"

    if padding > 0 {
        _, _ = tm.Println(tm.Background(tm.Bold(tm.Color(fmt.Sprintf(s, " "), fg)), bg))
    }

    if level == 0 {
        _, _ = tm.Println(tm.Background(tm.Bold(tm.Color(fmt.Sprintf(s, "  ✗  ERROR"), fg)), bg))
    }

    _, _ = tm.Println(tm.Background(tm.Bold(tm.Color(fmt.Sprintf(s, text), fg)), bg))

    if padding > 0 {
        _, _ = tm.Println(tm.Background(tm.Bold(tm.Color(fmt.Sprintf(s, " "), fg)), bg))
    }

}
