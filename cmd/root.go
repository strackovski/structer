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
    Short: "Convert JSON to struct.",
    Long:  `Structer is a tool for creating struct types from definitions stored as JSON.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    // Defines flags and configuration settings.
    // Supports persistent flags (global to the application) and local flags.
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.structer.yaml)")
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
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
