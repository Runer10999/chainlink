//go:build !windows
// +build !windows

package main

import (
	"io/ioutil"
	"testing"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/smartcontractkit/chainlink/core/cmd"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"gopkg.in/guregu/null.v4"
)

func ExampleRun() {
	t := &testing.T{}
	tc := cltest.NewTestGeneralConfig(t)
	tc.Overrides.Dev = null.BoolFrom(false)
	testClient := &cmd.Client{
		Renderer:               cmd.RendererTable{Writer: ioutil.Discard},
		Config:                 tc,
		Logger:                 logger.TestLogger(t),
		AppFactory:             cmd.ChainlinkAppFactory{},
		FallbackAPIInitializer: cltest.NewMockAPIInitializer(t),
		Runner:                 cmd.ChainlinkRunner{},
		HTTP:                   cltest.NewMockAuthenticatedHTTPClient(tc, "session"),
		ChangePasswordPrompter: cltest.MockChangePasswordPrompter{},
	}

	Run(testClient, "core.test", "--help")
	Run(testClient, "core.test", "--version")
	// Output:
	// NAME:
	//    core.test - CLI for Chainlink
	//
	// USAGE:
	//    core.test [global options] command [command options] [arguments...]
	//
	// VERSION:
	//    unset@unset
	//
	// COMMANDS:
	//    admin           Commands for remotely taking admin related actions
	//    attempts, txas  Commands for managing Ethereum Transaction Attempts
	//    blocks          Commands for managing blocks
	//    bridges         Commands for Bridges communicating with External Adapters
	//    config          Commands for the node's configuration
	//    jobs            Commands for managing Jobs
	//    keys            Commands for managing various types of keys used by the Chainlink node
	//    node, local     Commands for admin actions that must be run locally
	//    txs             Commands for handling Ethereum transactions
	//    chains          Commands for handling chain configuration
	//    nodes           Commands for handling node configuration
	//    help, h         Shows a list of commands or help for one command
	//
	// GLOBAL OPTIONS:
	//    --json, -j     json output as opposed to table
	//    --help, -h     show help
	//    --version, -v  print the version
	// core.test version unset@unset
}
