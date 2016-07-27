package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jessevdk/go-flags"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"

	"cred-alert/ingestor"
	"cred-alert/metrics"
	"cred-alert/queue"
)

type AWSOpts struct {
	AwsAccessKey       string `long:"aws-access-key" description:"access key for aws SQS service" env:"AWS_ACCESS_KEY" value-name:"ACCESS_KEY"`
	AwsSecretAccessKey string `long:"aws-secret-key" description:"secret access key for aws SQS service" env:"AWS_SECRET_ACCESS_KEY" value-name:"SECRET_KEY"`
	AwsRegion          string `long:"aws-region" description:"aws region for SQS service" env:"AWS_REGION" value-name:"REGION"`
	SqsQueueName       string `long:"sqs-queue-name" description:"queue name to use with SQS" env:"SQS_QUEUE_NAME" value-name:"QUEUE_NAME"`
}

type Opts struct {
	Port      uint16   `short:"p" long:"port" description:"the port to listen on" default:"8080" env:"PORT" value-name:"PORT"`
	Whitelist []string `short:"i" long:"ignore-repos" description:"comma separated list of repo names to ignore. The names may be regex patterns." env:"IGNORED_REPOS" env-delim:"," value-name:"REPO_LIST"`

	GitHub struct {
		WebhookToken string `short:"w" long:"webhook-token" description:"github webhook secret token" env:"GITHUB_WEBHOOK_SECRET_KEY" value-name:"TOKEN" required:"true"`
	} `group:"GitHub Options"`

	Datadog struct {
		APIKey      string `long:"datadog-api-key" description:"key to emit to datadog" env:"DATA_DOG_API_KEY" value-name:"KEY"`
		Environment string `long:"datadog-environment" description:"environment tag for datadog" env:"DATA_DOG_ENVIRONMENT_TAG" value-name:"NAME" default:"development"`
	} `group:"Datadog Options"`

	AWS AWSOpts `group:"AWS Options"`
}

func main() {
	var opts Opts

	logger := lager.NewLogger("cred-alert-ingestor")
	logger.Info("starting")

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		logger.Error("failed", err)
		os.Exit(1)
	}

	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))

	taskQueue, err := createQueue(opts, logger)
	if err != nil {
		logger.Error("failed", err)
		os.Exit(1)
	}

	emitter := metrics.BuildEmitter(opts.Datadog.APIKey, opts.Datadog.Environment)
	repoWhitelist := ingestor.BuildWhitelist(opts.Whitelist...)
	generator := queue.NewGenerator()

	in := ingestor.NewIngestor(taskQueue, emitter, repoWhitelist, generator)

	router := http.NewServeMux()
	router.Handle("/webhook", ingestor.Handler(logger, in, opts.GitHub.WebhookToken))

	members := []grouper.Member{
		{"api", http_server.New(
			fmt.Sprintf(":%d", opts.Port),
			router,
		)},
		{"debug", http_server.New(
			"127.0.0.1:6060",
			debugHandler(),
		)},
	}

	runner := sigmon.New(grouper.NewParallel(os.Interrupt, members))

	serverLogger := logger.Session("server", lager.Data{
		"port": opts.Port,
	})
	serverLogger.Info("starting")

	err = <-ifrit.Invoke(runner).Wait()
	if err != nil {
		serverLogger.Error("failed", err)
	}
}

func createQueue(opts Opts, logger lager.Logger) (queue.Queue, error) {
	logger = logger.Session("create-queue")
	logger.Info("starting")

	if sqsValuesExist(opts) {
		logger.Info("done")
		return createSqsQueue(logger, opts.AWS)
	}

	logger.Info("done")
	return queue.NewNullQueue(logger), nil
}

func sqsValuesExist(opts Opts) bool {
	if opts.AWS.AwsAccessKey != "" &&
		opts.AWS.AwsSecretAccessKey != "" &&
		opts.AWS.AwsRegion != "" &&
		opts.AWS.SqsQueueName != "" {

		return true
	}

	return false
}

func createSqsQueue(logger lager.Logger, awsOpts AWSOpts) (queue.Queue, error) {
	logger = logger.Session("create-sqs-queue")

	creds := credentials.NewStaticCredentials(awsOpts.AwsAccessKey, awsOpts.AwsSecretAccessKey, "")
	config := aws.NewConfig().WithCredentials(creds).WithRegion(awsOpts.AwsRegion)
	service := sqs.New(session.New(config))

	queue, err := queue.BuildSQSQueue(service, awsOpts.SqsQueueName)
	if err != nil {
		logger.Error("failed", err)
		return nil, err
	}

	logger.Info("done")
	return queue, nil
}

func debugHandler() http.Handler {
	debugRouter := http.NewServeMux()
	debugRouter.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	debugRouter.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	debugRouter.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	debugRouter.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	debugRouter.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	return debugRouter
}
