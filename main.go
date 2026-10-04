// Copyright (c) arkade author(s) 2022. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"os"

	"github.com/lmasiero/arkade/cmd"
	"github.com/lmasiero/arkade/cmd/chart"
	"github.com/lmasiero/arkade/cmd/docker"
	"github.com/lmasiero/arkade/cmd/fstail"
	"github.com/lmasiero/arkade/cmd/gha"
	"github.com/lmasiero/arkade/cmd/oci"
	"github.com/lmasiero/arkade/cmd/system"
	"github.com/spf13/cobra"
)

func main() {
	printarkadeASCIIArt := cmd.PrintArkadeASCIIArt

	var rootCmd = &cobra.Command{
		Use: "arkade",
		Run: func(cmd *cobra.Command, args []string) {
			printarkadeASCIIArt()
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.MakeInstall())
	rootCmd.AddCommand(cmd.MakeVersion())
	rootCmd.AddCommand(cmd.MakeInfo())
	rootCmd.AddCommand(cmd.MakeUpdate())
	rootCmd.AddCommand(cmd.MakeGet())
	rootCmd.AddCommand(cmd.MakeUninstall())
	rootCmd.AddCommand(cmd.MakeShellCompletion())

	rootCmd.AddCommand(cmd.MakeRelease())

	rootCmd.AddCommand(chart.MakeChart())
	rootCmd.AddCommand(docker.MakeDocker())
	rootCmd.AddCommand(fstail.MakeFstail())
	rootCmd.AddCommand(gha.MakeGHA())
	rootCmd.AddCommand(system.MakeSystem())
	rootCmd.AddCommand(oci.MakeOci())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
