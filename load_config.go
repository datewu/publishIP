package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	mode      = flag.String("mode", "dev", "environment default to develop environment")
	thirdPart = flag.String("echo", "https://ipecho.net/plain", "the thirdparty service for ipECHO")
	server    = flag.String("server", "https://wutuofu.com/api/v1/channel/ip", "the server side channel")
	name      = flag.String("name", "cd-router-publicIP", "the client name, you should set this identity")
)

func init() {
	flag.Parse()
	switch *mode {
	case "dev":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = ""
	default:
		log.Panic().
			Str("mode", *mode).
			Msg("unsupport mode")
	}
	log.Debug().
		Str("mode config", *mode).
		Str("gitCmit", GitCommit).
		Str("version", SemVer).
		Msg("runing env")

}
