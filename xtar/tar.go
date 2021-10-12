package xtar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
)

var (
	ErrEmpty = errors.New("tar: file empty")
)

// TarFiles tar archive file list to dst file
func TarFiles(dst string, files []string) (err error) {
	return tarFiles(dst, files, false)
}

// TarGzipFiles tar archive file list to dst file with gzip
func TarGzipFiles(dst string, files []string) (err error) {
	return tarFiles(dst, files, true)
}

func tarFiles(dst string, files []string, useGzip bool) (err error) {
	if len(files) == 0 {
		return ErrEmpty
	}

	tarFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	var tarWriter *tar.Writer
	if useGzip {
		gzipWriter := gzip.NewWriter(tarFile)
		defer gzipWriter.Close()
		tarWriter = tar.NewWriter(gzipWriter)
	} else {
		tarWriter = tar.NewWriter(tarFile)
	}
	defer tarWriter.Close()

	for _, fileName := range files {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			panic(err)
		}

		header, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			panic(err)
		}
		header.Name = fileInfo.Name()

		if err = tarWriter.WriteHeader(header); err != nil {
			panic(err)
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()

		if _, err = io.Copy(tarWriter, fr); err != nil {
			return err
		}
	}

	return nil
}
