package doctor

import (
	"fmt"
	"os"

	utils "jus-cli/internal/utils"
)

func CheckGlobalEnvs() {
	fmt.Printf("Checking for necessary basic environment variables... \n\n")

	envVariables := [...]string{"NEXUS_USERNAME", "NEXUS_PASSWORD"}

	for _, value := range envVariables {
		_, present := os.LookupEnv(value)
		if !present {
			// Throw error and show notion docs for each environment variable
		}

		utils.OutputSuccess(value)
	}
}
