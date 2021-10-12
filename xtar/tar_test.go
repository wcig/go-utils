package xtar

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarFiles(t *testing.T) {
	{
		_ = os.WriteFile("tmp.a.txt", []byte("aaaaa"), 0755)
		_ = os.WriteFile("tmp.b.txt", []byte("bbbbb"), 0755)
		_ = os.WriteFile("tmp.c.txt", []byte("ccccc"), 0755)
		err := TarFiles("tmp.1.tar", []string{"tmp.a.txt", "tmp.b.txt", "tmp.c.txt"})
		assert.Nil(t, err)

		defer func() {
			_ = os.Remove("tmp.a.txt")
			_ = os.Remove("tmp.b.txt")
			_ = os.Remove("tmp.c.txt")
			_ = os.Remove("tmp.1.tar")
		}()
	}
	{
		dir, err := os.MkdirTemp("", "tmp")
		assert.Nil(t, err)
		fmt.Println("tmp dir:", dir)

		_ = os.WriteFile(filepath.Join(dir, "tmp.a.txt"), []byte("aaaaa"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.b.txt"), []byte("bbbbb"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.c.txt"), []byte("ccccc"), 0755)
		err = TarFiles(filepath.Join(dir, "tmp.1.tar"),
			[]string{
				filepath.Join(dir, "tmp.a.txt"),
				filepath.Join(dir, "tmp.b.txt"),
				filepath.Join(dir, "tmp.c.txt"),
			})
		assert.Nil(t, err)

		defer func() {
			_ = os.RemoveAll(dir)
		}()
	}
}

func TestTarGzipFiles(t *testing.T) {
	{
		_ = os.WriteFile("tmp.a.txt", []byte("aaaaa"), 0755)
		_ = os.WriteFile("tmp.b.txt", []byte("bbbbb"), 0755)
		_ = os.WriteFile("tmp.c.txt", []byte("ccccc"), 0755)
		err := TarGzipFiles("tmp.1.tar.gz", []string{"tmp.a.txt", "tmp.b.txt", "tmp.c.txt"})
		assert.Nil(t, err)

		defer func() {
			_ = os.Remove("tmp.a.txt")
			_ = os.Remove("tmp.b.txt")
			_ = os.Remove("tmp.c.txt")
			_ = os.Remove("tmp.1.tar.gz")
		}()
	}
	{
		dir, err := os.MkdirTemp("", "tmp")
		assert.Nil(t, err)
		fmt.Println("tmp dir:", dir)

		_ = os.WriteFile(filepath.Join(dir, "tmp.a.txt"), []byte("aaaaa"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.b.txt"), []byte("bbbbb"), 0755)
		_ = os.WriteFile(filepath.Join(dir, "tmp.c.txt"), []byte("ccccc"), 0755)
		err = TarGzipFiles(filepath.Join(dir, "tmp.1.tar.gz"),
			[]string{
				filepath.Join(dir, "tmp.a.txt"),
				filepath.Join(dir, "tmp.b.txt"),
				filepath.Join(dir, "tmp.c.txt"),
			})
		assert.Nil(t, err)

		defer func() {
			_ = os.RemoveAll(dir)
		}()
	}
}
