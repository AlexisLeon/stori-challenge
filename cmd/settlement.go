package cmd

import (
	"github.com/alexisleon/stori/internal/conf"
	"github.com/alexisleon/stori/internal/dependency"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
)

var settlementCmd = cobra.Command{
	Use: "settlement",
	Run: func(cmd *cobra.Command, args []string) {
		settlement()
	},
}

func settlement() {
	c := conf.LoadConfig(configFile)

	app, err := dependency.CreateAppContext(c)
	if err != nil {
		log.Fatalf("%v", errors.Wrap(err, "failed to init app context: %v"))
	}

	err = app.SettlementHandler.ProcessSettlement()
	if err != nil {
		log.Fatalf("%v", errors.Wrap(err, "processing settlement"))
	}

	log.Println("Done!")
}
