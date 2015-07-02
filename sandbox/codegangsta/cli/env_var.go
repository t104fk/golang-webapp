package cli
import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func Start() {
	app := cli.NewApp()
	app.Name = "sandbox"
	app.Usage = "commend out other code in main.go, and set ENVs"
	app.Commands = []cli.Command{
		StartCmd,
	}
	app.Run(os.Args)
}

var StartCmd = cli.Command{
		Name: "start",
		Usage: "ENVs start",
		Description: "Start server",
		Action: action,
		Flags: Flags,
	}

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Usage:  "Location of configuration file, e.g: ./app.toml",
		EnvVar: "HAMPTON_CONFIG",
	},
	cli.StringFlag{ Name: "host", EnvVar: "HOST", },
	cli.StringFlag{ Name: "port", EnvVar: "PORT", },

	cli.StringFlag{ Name: "database.host", EnvVar: "DB_HOST", },
	cli.StringFlag{ Name: "database.port", EnvVar: "DB_PORT", },
	cli.StringFlag{ Name: "database.name", EnvVar: "DB_NAME", },
	cli.StringFlag{ Name: "database.user", EnvVar: "DB_USER", },
	cli.StringFlag{ Name: "database.pass", EnvVar: "DB_PASS", },
	cli.StringFlag{ Name: "database.sslmode", EnvVar: "DB_SSL_MODE", },
	cli.StringFlag{ Name: "database.contimeout", EnvVar: "DB_CON_TIMEOUT", },

	cli.StringFlag{ Name: "ldap.url", EnvVar: "LDAP_URL", },
	cli.StringFlag{ Name: "ldap.port", EnvVar: "LDAP_PORT", },
	cli.StringFlag{ Name: "ldap.dn", EnvVar: "LDAP_DN", },

	cli.StringFlag{ Name: "authKey.publicKey", EnvVar: "AUTH_KEY_PUB", },
	cli.StringFlag{ Name: "authKey.privateKey", EnvVar: "AUTH_KEY_PRI", },
}

func action(c *cli.Context) {
	log.Print("cli action")
	log.Print(c.FlagNames())
	for _, v := range c.FlagNames() {
		if value := c.String(v); value != "" {
			log.Print(value)
		}
	}
}
