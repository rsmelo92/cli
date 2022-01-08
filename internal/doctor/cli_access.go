package doctor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"

	utils "jus-cli/internal/utils"
)

type CliType struct {
	Name       string
	Command    []string
	NotionLink string
}

func throwError(err error, link string) {
	fmt.Println("Error:", err)
	fmt.Println("\nPlease set up configuration properly")
	fmt.Printf("\t %s \n", link)
	os.Exit(1)
}

func validateOutput(output string, cliType CliType) {
	var isValid bool
	if cliType.Name == "gcloud" {
		isValid = strings.Contains(output, "@jusbrasil.com.br") && strings.Contains(output, "jusbrasil-155317")
	}
	if cliType.Name == "kubectl" {
		isValid = strings.Contains(output, "jusbrasil-155317")
	}

	if !isValid {
		throwError(nil, cliType.NotionLink)
	}

	utils.OutputSuccess(fmt.Sprintf("%s configuration is set properly", cliType.Name))
}

func CheckConfig(cliType CliType) {
	fmt.Printf("\nChecking for %s configuration... \n\n", cliType.Name)

	s := spinner.New(spinner.CharSets[6], 100*time.Millisecond)
	s.Start()

	path, err := exec.LookPath(cliType.Name)
	if err != nil {
		throwError(err, cliType.NotionLink)
	}

	cmd := exec.Command(cliType.Command[0], cliType.Command[1], cliType.Command[2])

	output, err := cmd.Output()
	if err != nil {
		throwError(err, cliType.NotionLink)
	}

	s.Stop()

	utils.OutputSuccess(fmt.Sprintf("%s installed in %s", cliType.Name, path))
	validateOutput(string(output), cliType)
}
