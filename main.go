package main

import (
	"log"

	endure "github.com/spiral/endure/pkg/container"
	"github.com/spiral/roadrunner/v2/plugins/gzip"
	"github.com/spiral/roadrunner/v2/plugins/headers"
	"github.com/spiral/roadrunner/v2/plugins/static"
	"github.com/spiral/roadrunner/v2/plugins/status"

	// plugins
	"github.com/spiral/roadrunner-binary/v2/cli"
	httpPlugin "github.com/spiral/roadrunner/v2/plugins/http"
	"github.com/spiral/roadrunner/v2/plugins/informer"
	"github.com/spiral/roadrunner/v2/plugins/logger"
	"github.com/spiral/roadrunner/v2/plugins/metrics"
	"github.com/spiral/roadrunner/v2/plugins/reload"
	"github.com/spiral/roadrunner/v2/plugins/resetter"
	"github.com/spiral/roadrunner/v2/plugins/rpc"
	"github.com/spiral/roadrunner/v2/plugins/server"
	"github.com/temporalio/roadrunner-temporal/activity"
	temporalClient "github.com/temporalio/roadrunner-temporal/client"
	"github.com/temporalio/roadrunner-temporal/workflow"
)

func main() {
	var err error
	cli.Container, err = endure.NewContainer(nil, endure.SetLogLevel(endure.ErrorLevel), endure.RetryOnFail(false))
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Container.RegisterAll(
		// logger plugin
		&logger.ZapLogger{},
		// metrics plugin
		&metrics.Plugin{},
		// http server plugin
		&httpPlugin.Plugin{},
		// reload plugin
		&reload.Plugin{},
		// informer plugin (./rr workers, ./rr workers -i)
		&informer.Plugin{},
		// resetter plugin (./rr reset)
		&resetter.Plugin{},
		// rpc plugin (workers, reset)
		&rpc.Plugin{},
		// server plugin (NewWorker, NewWorkerPool)
		&server.Plugin{},

		// static
		&static.Plugin{},
		// headers
		&headers.Plugin{},
		// checker
		&status.Plugin{},
		// gzip
		&gzip.Plugin{},

		// temporal plugins
		&activity.Plugin{},
		&workflow.Plugin{},
		&temporalClient.Plugin{},
	)
	if err != nil {
		log.Fatal(err)
	}

	cli.Execute()
}
