package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/michalfikejs/AdventOfCode23/challenge/day01"
	"github.com/michalfikejs/AdventOfCode23/challenge/day02"
)

var AppName string

var rootCmd = &cobra.Command{
	Use:   AppName,
	Short: "Advent of Code 2023 Solutions",
	Long:  `Golang implementations for the 2022 Advent of Code problems`,
	Args:  cobra.ExactArgs(1),
}

func initialize() {
	day01.AddCommandsTo(rootCmd)
	day02.AddCommandsTo(rootCmd)
}

func Execute() {
	initialize()

	flags := rootCmd.PersistentFlags()
	flags.StringP("input", "i", "", "Input File to read. If not specified, assumes ./challenge/dayN/input.txt for the currently running challenge")

	_ = viper.BindPFlags(flags)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
