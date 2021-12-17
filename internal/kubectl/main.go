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
	fmt.Println("\nPlease set up kubectl configuration properly")
	fmt.Println("\thttps://www.notion.so/jusbrasil/Kubernetes-63bf1083f9b44888ade0a3ae246fd036\n")
	os.Exit(1)
}

func parseKubectlOutput(output string)  {
	hasProject := strings.Contains(output, "jusbrasil-155317")

	if !hasProject {
		gcloudErrorThrow(nil)
	}

	OutputSuccess("kubectl has proper access")
}

func CheckKubectlConfig()  {
	fmt.Printf("\nChecking for kubectl configuration... \n\n")

	path, err := exec.LookPath("kubectl")
	if err != nil {
		gcloudErrorThrow(err)
	}

	response := fmt.Sprintf("kubectl installed in %s", path)
	OutputSuccess(response)

	cmd := exec.Command("kubectl", " config", "view")

	output, err := cmd.Output()
	if err != nil {
		gcloudErrorThrow(err)
	}

	parseGCloudOutput(string(output))
}
