package main

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/lib/pq"
	"github.com/piotrkowalczuk/sklog"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	monitoringEnginePrometheus = "prometheus"
)

func initPrometheus(namespace, subsystem string, constLabels stdprometheus.Labels) func() (*monitoring, error) {
	return func() (*monitoring, error) {
		rpcRequests := prometheus.NewCounter(
			stdprometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "rpc_requests_total",
				Help:        "Total number of RPC requests made.",
				ConstLabels: constLabels,
			},
			monitoringRPCLabels,
		)
		rpcErrors := prometheus.NewCounter(
			stdprometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "rpc_errors_total",
				Help:        "Total number of errors that happen during RPC calles.",
				ConstLabels: constLabels,
			},
			monitoringRPCLabels,
		)

		postgresQueries := prometheus.NewCounter(
			stdprometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "postgres_queries_total",
				Help:        "Total number of SQL queries made.",
				ConstLabels: constLabels,
			},
			monitoringPostgresLabels,
		)
		postgresErrors := prometheus.NewCounter(
			stdprometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "postgres_errors_total",
				Help:        "Total number of errors that happen during SQL queries.",
				ConstLabels: constLabels,
			},
			monitoringPostgresLabels,
		)

		m := &monitoring{}
		m.rpc.requests = rpcRequests
		m.rpc.errors = rpcErrors
		m.postgres.queries = postgresQueries
		m.postgres.errors = postgresErrors

		return m, nil
	}
}

func initPostgres(connectionString string, logger log.Logger) *sql.DB {
	postgres, err := sql.Open("postgres", connectionString)
	if err != nil {
		sklog.Fatal(logger, fmt.Errorf("mnemosyned: postgres connection failure: %s", err.Error()))
	}

	sklog.Info(logger, "postgres connected", "address", connectionString)

	return postgres
}

func initMonitoring(fn func() (*monitoring, error), logger log.Logger) *monitoring {
	m, err := fn()
	if err != nil {
		sklog.Fatal(logger, err)
	}

	return m
}

const (
	loggerAdapterStdOut = "stdout"
	loggerFormatJSON    = "json"
	loggerFormatHumane  = "humane"
	loggerFormatLogFmt  = "logfmt"
)

func initLogger(adapter, format string, level int, context ...interface{}) log.Logger {
	var l log.Logger

	if adapter != loggerAdapterStdOut {
		stdlog.Fatal("service: unsupported logger adapter")
	}

	switch format {
	case loggerFormatHumane:
		l = sklog.NewHumaneLogger(os.Stdout, sklog.DefaultHTTPFormatter)
	case loggerFormatJSON:
		l = log.NewJSONLogger(os.Stdout)
	case loggerFormatLogFmt:
		l = log.NewLogfmtLogger(os.Stdout)
	default:
		stdlog.Fatal("mnemosyned: unsupported logger format")
	}

	l = log.NewContext(l).With(context...)

	sklog.Info(l, "logger has been initialized successfully", "adapter", adapter, "format", format, "level", level)

	return l
}

func initStorage(fn func() (Storage, error), logger log.Logger) Storage {
	s, err := fn()
	if err != nil {
		sklog.Fatal(logger, fmt.Errorf("mnemosyned: storage init failure: %s", err.Error()))
	}

	err = s.Setup()
	if err != nil {
		sklog.Fatal(logger, fmt.Errorf("mnemosyned: storage setup failure: %s", err.Error()))
	}

	return s
}
