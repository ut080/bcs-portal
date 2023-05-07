package files

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Copy is a basic file copy utility function, because this doesn't exist in the std library for... reasons.
func Copy(srcPath, destPath string) (err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	return nil
}

func Move(srcPath, destPath string) (err error) {
	err = Copy(srcPath, destPath)
	if err != nil {
		return err
	}

	err = os.Remove(srcPath)
	if err != nil {
		return err
	}

	return nil
}

func Write(destPath, content string) (err error) {
	dest, err := os.Create(destPath)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	defer dest.Close()

	_, err = io.WriteString(dest, content)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	err = dest.Sync()
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	return nil
}

func Remove(path string) error {
	return os.Remove(path)
}

func SplitPath(path string) (dir string, file string, err error) {
	dir, file = filepath.Split(path)

	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return "", "", errors.WithStack(err)
		}
	}

	return dir, file, nil
}
