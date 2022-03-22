package main

import (
	"foundation/framework"
	"foundation/home"
	"github.com/spf13/cobra"
)

var runCmdCobra = &cobra.Command{
	Use:    "run",
	Short:  "run as a service",
	PreRun: before,
	Run:    execute,
}

func before(cmd *cobra.Command, args []string) {

}

func execute(cmd *cobra.Command, args []string) {
	framework.Run(home.NewActor(1024, 1))
}
