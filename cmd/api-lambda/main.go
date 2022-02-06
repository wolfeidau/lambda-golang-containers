package main

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	lmw "github.com/wolfeidau/lambda-go-extras/middleware"
	"github.com/wolfeidau/lambda-go-extras/middleware/raw"
	zlog "github.com/wolfeidau/lambda-go-extras/middleware/zerolog"
	"github.com/wolfeidau/lambda-golang-containers/internal/server"
)

var (
	version = "unknown"

	cli struct {
		Version         kong.VersionFlag
		RawEventLogging bool   `help:"Enable raw event logging." env:"RAW_EVENT_LOGGING"`
		Debug           bool   `help:"Enable debug logging." env:"DEBUG"`
		Stage           string `help:"The development stage." env:"STAGE"`
		Branch          string `help:"The git branch this code originated." env:"BRANCH"`
	}
)

func main() {
	kong.Parse(&cli,
		kong.Vars{"version": version}, // bind a var for version
	)

	// build up a list of fields which will be included in all log messages
	flds := lmw.FieldMap{"version": version}

	ch := lmw.New(
		zlog.New(zlog.Fields(flds)), // assign a logger and bind it in the context
	)

	if cli.RawEventLogging {
		ch.Use(raw.New(raw.Fields(flds))) // if raw event logging is enabled dump everything to the log in and out
	}

	awscfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load SDK config")
	}

	e := echo.New()
	err = server.Setup(awscfg, e)
	if err != nil {
		log.Fatal().Err(err).Msg("server setup failed")
	}

	gw := gateway.NewGateway(e)

	// register our lambda handler with the middleware configured
	lambda.StartHandler(ch.Then(gw))
}
