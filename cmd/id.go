package cmd

import (
	"fmt"
	"os"
	"trivy-db/internal"

	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/spf13/cobra"

	"github.com/aquasecurity/table"
)

var idCmd = &cobra.Command{
	Use:   "id <CVE-ID>",
	Short: "Recherche une vulnérabilité par identifiant CVE",
	Args:  cobra.ExactArgs(1),
	Example: `  # Recherche une vulnérabilité par son identifiant CVE
  $ trivy db id CVE-1999-0163`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		dbPath := internal.GetTrivyDBPath()

		vuln, err := internal.GetCVE(dbPath, id)
		if err != nil {
			log.Info("CVE non trouvée", id)
			return
		}

		log.Info("CVE trouvée")

		t := table.New(os.Stdout)
		t.SetHeaders("Champ", "Valeur")

		orderedKeys := []string{
			"CVE-ID",
			"Title",
			"Description",
			"PublishedDate",
			"LastModifiedDate",
			"Severity",
			"VendorSeverity",
			"CVSS",
			"References",
		}

		for _, key := range orderedKeys {
			if val, ok := vuln[key]; ok {
				t.AddRow(key, fmt.Sprintf("%v", val))
			}
		}

		t.Render()
	},
}
