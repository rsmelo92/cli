package gcloud

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "internal/utils"
)

func gcloudErrorThrow(err error)  {
	fmt.Println("Error:", err)
	fmt.Println("\nPlease set up gcloud configuration properly")
	fmt.Println("\thttps://www.notion.so/jusbrasil/Google-Cloud-e55cbd1e9e9c4a928fd974b1eb5d1e97\n")
	os.Exit(1)
}

func parseGCloudOutput(output string)  {
	hasAccount := strings.Contains(output, "@jusbrasil.com.br")
	hasProject := strings.Contains(output, "jusbrasil-155317")

	if !hasAccount || !hasProject {
		gcloudErrorThrow(nil)
	}

	OutputSuccess("gcloud account and project are set properly")
}

func CheckGCloudConfig()  {
	fmt.Printf("\nChecking for gcloud configuration... \n\n")

	path, err := exec.LookPath("gcloud")
	if err != nil {
		gcloudErrorThrow(err)
	}

	response := fmt.Sprintf("gcloud installed in %s", path)
	OutputSuccess(response)

	cmd := exec.Command("gcloud", "config", "list")

	output, err := cmd.Output()
	if err != nil {
		gcloudErrorThrow(err)
	}

	parseGCloudOutput(string(output))
}
