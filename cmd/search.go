package cmd

import (
    "fmt"
    "os"
    "trivy-db/internal"
    "github.com/aquasecurity/trivy/pkg/log"
    "go.uber.org/zap"
    "github.com/spf13/cobra"
    "github.com/aquasecurity/table"
)

var searchCmd = &cobra.Command{
    Use:   "search <mot-clé>",
    Short: "Recherche des vulnérabilités par mot-clé",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        keyword := args[0]
        dbPath := internal.GetTrivyDBPath()

        results, err := internal.SearchByKeyword(dbPath, keyword)
        if err != nil {
            log.Error("Erreur lors de la recherche",
                zap.String("keyword", keyword),
                zap.Error(err),
            )
            return
        }

        log.Info("Résultats trouvés",
            zap.String("keyword", keyword),
            zap.Int("count", len(results)),
        )

        fmt.Printf("🔍 %d résultat(s) trouvé(s) pour '%s'\n", len(results), keyword)

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

        for _, vuln := range results {
            fmt.Println(" ")
            t := table.New(os.Stdout)
            t.SetHeaders("Champ", "Valeur")
            for _, key := range orderedKeys {
                if val, ok := vuln[key]; ok {
                    t.AddRow(key, fmt.Sprintf("%v", val))
                }
            }

            t.Render()
        }
    },
}