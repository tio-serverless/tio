package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

func zipAndparser(file string) error {
	sources, err := unzip(file, b.Root+"/tio")
	if err != nil {
		return err
	}

	var config string

	for _, s := range sources {
		if strings.HasSuffix(s, ".tio.toml") {
			config = s
			break
		}
	}

	if config == "" {
		return errors.New("Not found .tio.toml in this zip. ")
	}

	err = parserSource()
	if err != nil {
		return err
	}

	return nil
}

func parserSource() error {
	_, err := toml.DecodeFile(b.Root+"/tio/.tio.toml", b)
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadFile(b.Root + "/tio/.tio.toml")
	raw = string(data)

	return nil
}

func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	return filenames, nil
}
