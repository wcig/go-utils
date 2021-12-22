package xzip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipDirV1(t *testing.T) {
	_ = os.Mkdir("tmp", 0755)
	_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
	_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
	_ = os.Mkdir("tmp/c", 0755)
	_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)
	defer os.RemoveAll("tmp")

	{
		// 相对路径
		err := ZipDir("tmp.1.zip", "tmp")
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := ZipDir(filepath.Join(wd, "tmp.2.zip"), filepath.Join(wd, "tmp"))
		assert.Nil(t, err)
	}
}

func TestZipDirV2(t *testing.T) {
	_ = os.Mkdir("tmp", 0755)
	_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
	_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
	_ = os.Mkdir("tmp/c", 0755)
	_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)
	defer os.RemoveAll("tmp")

	{
		// 相对路径
		err := ZipDirV2("tmp.1.zip", "tmp")
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := ZipDirV2(filepath.Join(wd, "tmp.2.zip"), filepath.Join(wd, "tmp"))
		assert.Nil(t, err)
	}
}

func TestZipFiles(t *testing.T) {
	{
		_ = os.WriteFile("tmp.a.txt", []byte("aaaaa"), 0755)
		_ = os.WriteFile("tmp.b.txt", []byte("bbbbb"), 0755)
		_ = os.WriteFile("tmp.c.txt", []byte("ccccc"), 0755)
		err := ZipFiles("tmp.zip", []string{"tmp.a.txt", "tmp.b.txt", "tmp.c.txt"})
		assert.Nil(t, err)
	}
}

func TestUnzip(t *testing.T) {
	TestZipDirV1(t)

	{
		// 相对路径
		err := Unzip("", "tmp.zip")
		fmt.Println(err)
		assert.Nil(t, err)
	}
	{
		// 绝对路径
		wd, _ := os.Getwd()
		err := Unzip(filepath.Join(wd, "test"), filepath.Join(wd, "tmp.zip"))
		assert.Nil(t, err)
	}
}
