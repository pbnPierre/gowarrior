package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	level   string
	name    string
	rootCmd = &cobra.Command{
		Use:   "gowarrior-cli",
		Short: "Game written in Go for learning Go.",
		Long:  `Build up your warrior through a tower and get the treasure by writing your warrior behavior in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello dear", name, "!")
			fmt.Println("Welcome to the Golang Tower")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&level, "level", "l", "", "--level 1")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "William Wallace", "--name William Wallace")
	viper.BindPFlag("level", rootCmd.PersistentFlags().Lookup("level"))
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}
