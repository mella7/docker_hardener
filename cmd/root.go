package main

import (
	"docker_hardener/pkg/hardener"
	"docker_hardener/pkg/scanner"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "docker_hardener"}

	// Command to scan an image for vulnerabilities
	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan a Docker image for vulnerabilities",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide an image name to scan.")
				return
			}
			imageName := args[0]
			err := scanner.ScanImage(imageName)
			if err != nil {
				fmt.Println("Error scanning image:", err)
			}
		},
	}

	// Command to harden a Dockerfile
	var hardenCmd = &cobra.Command{
		Use:   "harden",
		Short: "Harden a Dockerfile or Image",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide a Dockerfile path.")
				return
			}
			dockerfilePath := args[0]
			err := hardener.HardenDockerfile(dockerfilePath)
			if err != nil {
				fmt.Println("Error hardening Dockerfile:", err)
				return
			}
			err = hardener.SetNonRootUser(dockerfilePath)
			if err != nil {
				fmt.Println("Error setting non-root user:", err)
				return
			}
			fmt.Println("Dockerfile hardened successfully.")
		},
	}

	// Add subcommands to the root command
	rootCmd.AddCommand(scanCmd, hardenCmd)

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
