package cmd

import (
	trivylog "github.com/aquasecurity/trivy/pkg/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "Outil CLI pour interroger la base Trivy",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(idCmd)
	rootCmd.AddCommand(searchCmd)
	trivylog.InitLogger(false, false)
}
