// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package cmd

import (
    "fmt"
    "os"

    "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "structer",
    Short: "Makes structs from input file(s)",
    Long: `STRUCTER - Struct code generator for Go

Structer is a tool for generating struct types that implement 
the underlying data structure discovered in the input contents 
(the --in flag). Invoke the 'make' command to use structer:

    structer make --in <PATH>

Input required is the path (absolute or relative) to the input 
file(s) or directory (of files), and is specified with the --i 
flag. Supported input formats (flag --format) are:

    - JSON (default) 
    - YAML
    - HCL (Hashicorp Configuration Language)

Default output is a directory called "generated", created in the
current working directory, and can be overridden by setting the 
--out flag to user defined value.'
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.structer.yaml)")
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := homedir.Dir()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        viper.AddConfigPath(home)
        viper.SetConfigName(".structer")
    }

    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
