package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/martinjirku/zasobar/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	defaultCfgFile = "spajza"
	v              = viper.NewWithOptions(viper.KeyDelimiter("."))
	configuration  config.Configuration
	headerHelp     = `
   ┌───────────────────────────────────────┐
   │                                       │
   │      _\_/              _              │
   │    / ___| _ __   __ _ (_)______ _     │
   │    \___ \| '_ \ / _' || |_  / _' |    │
   │     ___) | |_) | (_| || |/ / (_| |    │
   │    |____/| .__/ \__,_|/ /___\__,_|    │
   │          |_|        |__/              │
   │                                       │
   └───────────────────────────────────────┘

`
	rootCmd = &cobra.Command{
		Use:   "spajza",
		Short: "Spajza is application to managa home resources",
		Long: fmt.Sprintf(`%sTo properly manage resources in you home storage you should use this tool.

`, headerHelp),
		Run: func(cmd *cobra.Command, args []string) {
			configuration.LoadConfiguration()
			fmt.Print("\n>>>>\n\n")
			v.Debug()

			fmt.Print(configuration)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	log.Default().Println("Initializing the root")

	if cfgFile != "" {
		log.Default().Printf("setting configuration file: %s\n", cfgFile)
		v.SetConfigFile(cfgFile)
	} else {
		log.Default().Printf("setting default configuraiton file: %s\n", defaultCfgFile)
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME")
		v.SetConfigName(defaultCfgFile)
	}

	v.SetEnvPrefix("SPAJZA")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	// config.PrepareDefaults(v)

	v.ReadInConfig()
	configuration = config.NewConfiguration(v)
}
