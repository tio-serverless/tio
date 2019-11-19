package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func build(name string) error {

	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("bin/%s", b.Name))
	//cmd := exec.Command("go","build","main.go")
	cmd.Dir = fmt.Sprintf("%s/%s", b.Root, zipName)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	logrus.Infof("Work Dir: %s", cmd.Dir)
	logrus.Infof("Command: %s %v", cmd.Path, cmd.Args)
	logrus.Info("===========Build Log===========")
	logrus.Info("")

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	logrus.Info(outStr)
	logrus.Infof(errStr)

	return nil
}
