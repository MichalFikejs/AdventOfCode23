package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/michalfikejs/AdventOfCode23/challenge/day01"
	"github.com/michalfikejs/AdventOfCode23/challenge/day02"
	"github.com/michalfikejs/AdventOfCode23/challenge/day03"
	"github.com/michalfikejs/AdventOfCode23/challenge/day04"
	"github.com/michalfikejs/AdventOfCode23/challenge/day05"
	"github.com/michalfikejs/AdventOfCode23/challenge/day06"
	"github.com/michalfikejs/AdventOfCode23/challenge/day07"
	"github.com/michalfikejs/AdventOfCode23/challenge/day08"
	"github.com/michalfikejs/AdventOfCode23/challenge/day09"
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
	day03.AddCommandsTo(rootCmd)
	day04.AddCommandsTo(rootCmd)
	day05.AddCommandsTo(rootCmd)
	day06.AddCommandsTo(rootCmd)
	day07.AddCommandsTo(rootCmd)
	day08.AddCommandsTo(rootCmd)
	day09.AddCommandsTo(rootCmd)
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
