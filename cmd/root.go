package cmd

import (
	"fmt"
	"strings"

	"github.com/elliot-token/api/config"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     config.Config
)

const (
	configFlag = "config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v", cfg)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, configFlag, "", "")
	if err := rootCmd.MarkPersistentFlagRequired(configFlag); err != nil {
		panic(fmt.Sprintf("failed to init flag: %s", err.Error()))
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config file '%s': %s", viper.ConfigFileUsed(), err.Error()))
	}

	// read in environment variables that match
	viper.SetEnvPrefix("ELLIOT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("failed to unmarshal config: %s", err.Error()))
	}
}
