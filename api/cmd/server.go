package cmd

import (
	"fmt"

	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/infra/web"
	"github.com/spf13/cobra"
)

func init() {
	var (
		domain    string
		port      string
		jwtConfig config.Jwt
		dbConfig  config.Db
	)
	flagSet := serveCmd.Flags()
	flagSet.StringVarP(&domain, "domain", "d", "localhost", "Domain to listen to")
	flagSet.StringVarP(&port, "port", "p", "8000", "Port to listen to.")

	flagSet.StringVar(&dbConfig.Host, "db.host", "localhost", "Db port to listen to.")
	flagSet.StringVar(&dbConfig.Port, "db.port", "3306", "Db port to listen to.")
	flagSet.StringVar(&dbConfig.Name, "db.name", "zasobar", "Database name.")
	flagSet.StringVar(&dbConfig.Password, "db.password", "user", "Database Password")
	flagSet.StringVar(&dbConfig.User, "db.user", "user", "Database Password")
	flagSet.StringVar(&dbConfig.Type, "db.type", "mysql", "Database Type.")

	flagSet.StringVarP(&jwtConfig.Secret, "jwt.secret", "s", "", "JWT Secret")
	flagSet.StringVarP(&jwtConfig.Issuer, "jwt.issuer", "i", "", "JWT Issuer")
	flagSet.Int64VarP(&jwtConfig.Validity, "jwt.validity", "v", 10*60*24, "JWT Validity in minutes")

	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Serve the application",
	Long:    fmt.Sprintf(`%sTo serve the api, you can use the serve command.`, headerHelp),
	Version: "0.0.1",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print(headerHelp)
	},
	Run: func(cmd *cobra.Command, args []string) {

		config.PrepareDefaults(v)
		config.PrepareJwt(v)
		config.PrepareDefaultServe(v, "localhost", "8232")
		v.BindPFlag("domain", cmd.Flags().Lookup("domain"))
		v.BindPFlag("port", cmd.Flags().Lookup("port"))

		v.BindPFlags(cmd.Flags())
		configuration.LoadConfiguration()

		web.InitServer(configuration)
	},
}
