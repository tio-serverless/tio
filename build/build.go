package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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

	err = tarDir([]string{
		b.Root + "/tio/bin/" + name, b.Root + "/tio/Dockerfile",
	}, b.Root+"/tio.tar")
	if err != nil {
		return err
	}

	version := "latest"
	if b.BuildInfo.Version != "" {
		version = b.BuildInfo.Version
	}

	return buildImage(fmt.Sprintf("%s-%s", name, version))
}

func createDockfile(name string) error {
	d := `FROM %s
COPY %s /%s
ENTRYPOINT ["/%s"]`

	content := fmt.Sprintf(d, baseImg, name, name, name)

	return ioutil.WriteFile(b.Root+"/tio/Dockerfile", []byte(content), 0777)
}

func tarDir(src []string, dst string) error {
	// 创建tar文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 通过fw创建一个tar.Writer
	tw := tar.NewWriter(fw)
	// 如果关闭失败会造成tar包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			log.Println(err)
		}
	}()

	for _, fileName := range src {
		fi, err := os.Stat(fileName)
		if err != nil {
			log.Println(err)
			continue
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		// 将tar的文件信息hdr写入到tw
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}

		// 将文件数据写入
		fs, err := os.Open(fileName)
		if err != nil {
			return err
		}
		if _, err = io.Copy(tw, fs); err != nil {
			return err
		}
		fs.Close()
	}
	return nil
}
