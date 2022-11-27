package files

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

// Copy is a basic file copy utility function, because this doesn't exist in the std library for... reasons.
func Copy(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Move(srcPath, destPath string) error {
	err := Copy(srcPath, destPath)
	if err != nil {
		return err
	}

	err = os.Remove(srcPath)
	if err != nil {
		return err
	}

	return nil
}

func Write(destPath, content string) error {
	dest, err := os.Create(destPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer dest.Close()

	_, err = io.WriteString(dest, content)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dest.Sync()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
