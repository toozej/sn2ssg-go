package sn2ssg

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func exportFromSimplenote() {
	sncliBinaryPath, err := exec.LookPath("sncli")
	if err != nil {
		sncliBinaryPath = "/home/james/.local/bin/sncli"
	}

	command := []string{sncliBinaryPath, "--config=/dev/null", "-r", "dump", tagToDownload}
	inputFilename := inputDir + "/sn_dump.md"

	err = os.WriteFile(inputFilename, []byte{}, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(command[0], command[1:]...)
	outFile, err := os.Create(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	cmd.Stdout = outFile

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Dumping of notes via sncli was successful.")

}
