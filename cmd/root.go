// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"
	"os"

	"github.com/Ibotta/sopstool/sopsyaml"
	"github.com/spf13/cobra"
)

//BuildVersion (updated by main)
var BuildVersion string

//BuildCommit (updated by main)
var BuildCommit string

//BuildDate (updated by main)
var BuildDate string

var cfgPath string
var sopsConfig sopsyaml.SopsConfig

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sopstool",
	Short: "Wrapper around sops for multiple files",
	Long: fmt.Sprintf(`sopstool

sops wrapper supporting multiple files and helper commands.`),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgPath, "configpath", ".", "config file path")
}

// initConfig reads in config file
func initConfig() {
	cfgFile, err := sopsyaml.FindConfigFile(cfgPath)
	if err != nil {
		panic(err)
	}
	data, err := sopsyaml.LoadConfigFile(cfgFile)
	if err != nil {
		panic(fmt.Errorf("Error loading config: %s", err))
	}
	encFiles, err := sopsyaml.ExtractConfigEncryptFiles(data)
	if err != nil {
		panic(fmt.Errorf("Error reading config: %s", err))
	}

	sopsConfig = sopsyaml.SopsConfig{
		Path:           cfgFile,
		Tree:           data,
		EncryptedFiles: encFiles,
	}
}
