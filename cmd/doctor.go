package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	doctor "jus-cli/internal/doctor"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check for necessary Jusbrasil configuration",
	Long:  `Check for necessary Jusbrasil configuration, which are supposed to be done on onboarding`,
	Run: func(cmd *cobra.Command, args []string) {
		doctor.CheckGlobalEnvs()

		gcloud := doctor.CliType{
			Name:       "gcloud",
			Command:    []string{"gcloud", "config", "list"},
			NotionLink: "https://www.notion.so/jusbrasil/Google-Cloud-e55cbd1e9e9c4a928fd974b1eb5d1e97",
		}
		doctor.CheckConfig(gcloud)

		kubectl := doctor.CliType{
			Name:       "kubectl",
			Command:    []string{"kubectl", "config", "view"},
			NotionLink: "https://www.notion.so/jusbrasil/Kubernetes-63bf1083f9b44888ade0a3ae246fd036",
		}
		doctor.CheckConfig(kubectl)

		fmt.Println("\nAll configuration is set properly!")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doctorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doctorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
