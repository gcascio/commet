package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gcascio/commet/internal/git"
	"github.com/gcascio/commet/internal/llm"
)

const defaultLLMUrl = "http://localhost:11434/api/chat"
const defaultModel = "mistral"

var rootCmd = &cobra.Command{
	Use:   "commet",
	Short: "Generate commit messages using an LLM.",
	Long:  "Commit your changes with commet to automatically generate commit messages using LLMs.",
	Run: func(cmd *cobra.Command, args []string) {
		commitAll, _ := cmd.Flags().GetBool("all")

		diff := git.GitDiff(commitAll)

		if len(diff) == 0 {
			fmt.Println("No changes to commit")
			os.Exit(0)
		}

		commitMessage := llm.GenerateCommitMessage(diff)

		git.GitCommit(commitMessage, commitAll)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP("all", "a", false, "Commit all changed files i.e. 'git commit -a'")

	rootCmd.PersistentFlags().String("llm", "", "URL to LLM API i.e. Ollama, defaults to: "+defaultLLMUrl)
	viper.BindPFlag("llm", rootCmd.PersistentFlags().Lookup("llm"))
	viper.SetDefault("llm", defaultLLMUrl)

	rootCmd.PersistentFlags().StringP("model", "m", "", "LLM model to be used, defaults to: "+defaultModel)
	viper.BindPFlag("model", rootCmd.PersistentFlags().Lookup("model"))
	viper.SetDefault("model", defaultModel)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".commet"
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".commet")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
