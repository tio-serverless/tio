package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zipAndparser(file string) error {
	sources, err := unzip(file+".zip", b.Root+"/"+file)
	if err != nil {
		return err
	}

	var gofile string

	for _, s := range sources {
		if strings.HasSuffix(s, ".go") {
			gofile = s
			break
		}
	}

	if gofile == "" {
		return errors.New("Not found go source code in this zip. ")
	}

	err = parserSource(gofile)
	if err != nil {
		return err
	}

	return nil
}

func parserSource(file string) error {
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)

	if err != nil {
		return err
	}

	for _, c := range f.Comments {
		cc := strings.TrimSpace(c.Text())
		if strings.HasPrefix(cc, "tio-name:") {
			n := strings.Split(cc, ":")
			if len(n) < 2 {
				continue
			}

			b.Name = strings.TrimSpace(n[1])
			continue
		}

		if strings.HasPrefix(cc, "tio-api:") {
			n := strings.Split(cc, ":")
			if len(n) < 2 {
				continue
			}

			b.API = strings.TrimSpace(n[1])
			continue
		}
	}

	if b.Name == "" {
		b.Name = zipName
	}

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
