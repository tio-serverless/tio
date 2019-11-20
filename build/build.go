package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func build(name string) error {

	cmd := exec.Command("go", "build", "-x", "-o", fmt.Sprintf("bin/%s", name))
	cmd.Dir = fmt.Sprintf("%s/tio", b.Root)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	logrus.Infof("Work Dir: %s", cmd.Dir)
	logrus.Infof("Command: %s %v", cmd.Path, cmd.Args)
	logrus.Info("===========Build Log===========")
	logrus.Info("")

	err := cmd.Run()
	if err != nil {
		return err
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	logrus.Info(outStr)
	logrus.Infof(errStr)

	err = createDockfile(name)
	if err != nil {
		return err
	}

	return buildImage(name)
}

func createDockfile(name string) error {
	d := `FROM %s
COPY bin/%s /%s
ENTRYPOINT ["/%s"]`

	content := fmt.Sprintf(d, baseImg, name, name, name)

	return ioutil.WriteFile(b.Root+"/tio/Dockerfile", []byte(content), 0777)
}
