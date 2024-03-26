package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/tandesko/cards_validator/internal/config"
	"github.com/tandesko/cards_validator/internal/service"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

var (
	app = kingpin.New("cards-validator", "service responsible for validating credit cards")

	runCmd     = app.Command("run", "run command")
	serviceCmd = runCmd.Command("service", "run service")
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		log.Info("started api...")
		service.Run(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
	}
	if err != nil {
		log.WithError(err).Error("failed to run command")
		return false
	}

	return true
}
