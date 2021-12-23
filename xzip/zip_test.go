package xzip

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipFile(t *testing.T) {
	fileName := "testdata/a.txt"
	zipName := "testdata/test.zip"
	defer os.Remove(zipName)

	err := ZipFile(zipName, fileName)
	assert.Nil(t, err)
}

func TestZipFiles(t *testing.T) {
	files := []string{"testdata/a.txt", "testdata/b.txt"}
	zipName := "testdata/test.zip"
	defer os.Remove(zipName)

	err := ZipFiles(zipName, files)
	assert.Nil(t, err)
}

func TestZipDir(t *testing.T) {
	dir := "testdata"
	zipName := "test.zip"

	// 相对路径
	err := ZipDir(zipName, dir)
	assert.Nil(t, err)
	_ = os.Remove(zipName)

	// 绝对路径
	wd, _ := os.Getwd()
	err = ZipDir(filepath.Join(wd, zipName), filepath.Join(wd, dir))
	assert.Nil(t, err)
	_ = os.Remove(zipName)
}

func TestZipDirV2(t *testing.T) {
	dir := "testdata"
	zipName := "test.zip"

	// 相对路径
	err := ZipDirV2(zipName, dir)
	assert.Nil(t, err)
	_ = os.Remove(zipName)

	// 绝对路径
	wd, _ := os.Getwd()
	err = ZipDirV2(filepath.Join(wd, zipName), filepath.Join(wd, dir))
	assert.Nil(t, err)
	_ = os.Remove(zipName)
}

func TestUnzip(t *testing.T) {
	_ = ZipDir("test.zip", "testdata")
	defer os.Remove("test.zip")

	dir := "testdata"

	// 相对路径
	err := Unzip(dir, "test.zip")
	assert.Nil(t, err)
	_ = os.RemoveAll(filepath.Join(dir, "testdata"))

	// 绝对路径
	wd, _ := os.Getwd()
	err = Unzip(filepath.Join(wd, dir), filepath.Join(wd, "test.zip"))
	assert.Nil(t, err)
	_ = os.RemoveAll(filepath.Join(dir, "testdata"))
}
