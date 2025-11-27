package cmd

import (
	"strconv"

	"pbnPierre/gowarrior/app/game"

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
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	i_level, err := strconv.Atoi(level)
	if err != nil {
		panic(err)
	}
	player := game.NewPlayer(name)
	game := game.NewGame(player, i_level)
	game.Run()
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&level, "level", "l", "1", "--level 1")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "--name William Wallace")
	viper.BindPFlag("level", rootCmd.PersistentFlags().Lookup("level"))
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}
