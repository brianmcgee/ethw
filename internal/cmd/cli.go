package cmd

import (
	"time"

	"github.com/charmbracelet/log"
)

var Cli struct {
	Wallet struct {
		Create walletCreateCmd `cmd:"" help:"Create new Ethereum wallets"`
	} `cmd:"" help:"Manage Ethereum wallets"`

	KeyStore struct {
		Create keystoreCreateCmd `cmd:"" help:"Manage Ethereum keystores"`
	} `cmd:"" name:"keystore" help:"Manage Ethereum KeyStores"`

	Seed struct {
		Create seedCreateCmd `cmd:"" help:"Create a new seed"`
	} `cmd:"" help:"Manage cryptographic seeds for Ethereum wallets"`

	Log logOptions `embed:"" prefix:"log-"`

	Version versionFlag `name:"version" short:"v" help:"Display the application version"`
}

type logOptions struct {
	Level string `enum:"debug,info,warn,error,fatal" env:"LOG_LEVEL" default:"info" help:"Configure logging level"`
}

func (l *logOptions) ConfigureLog() {
	log.SetLevel(log.ParseLevel(l.Level))
	log.SetReportTimestamp(true)
	log.SetTimeFormat(time.RFC3339)
}
