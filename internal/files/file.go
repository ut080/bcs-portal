package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/logging"
)

type File struct {
	dir    string
	base   string
	ext    string
	logger logging.Logger
}

func NewFile(path string, logger logging.Logger) (File, error) {
	dir, base, ext, err := DecomposePath(path)
	if err != nil {
		return File{}, errors.WithStack(err)
	}

	return File{
		dir:    dir,
		base:   base,
		ext:    ext,
		logger: logger,
	}, nil
}

func DecomposePath(path string) (string, string, string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", "", "", errors.WithStack(err)
	}

	dir := filepath.Dir(abs)

	base := strings.TrimSuffix(abs, filepath.Ext(abs))

	ext := filepath.Ext(abs)

	return dir, base, ext, nil
}

func (f File) Empty() bool {
	return f.dir == "" && f.base == "" && f.ext == ""
}

func (f File) Dir() string {
	return f.dir
}

func (f File) Base() string {
	return f.base
}

func (f File) Ext() string {
	return f.ext
}

func (f File) Name() string {
	return fmt.Sprintf("%s.%s", f.base, f.ext)
}

func (f File) FullPath() string {
	return filepath.Join(f.dir)
}

func (f File) Stat() (os.FileInfo, error) {
	return os.Stat(f.FullPath())
}

func (f File) Create() (*os.File, error) {
	return os.Create(f.FullPath())
}

func (f File) Open() (*os.File, error) {
	return os.Open(f.FullPath())
}

func (f File) Remove() error {
	return os.Remove(f.FullPath())
}

func (f File) Copy(destDir string) (File, error) {
	destPath := filepath.Join(destDir, f.Name())
	destFile, err := NewFile(destPath, f.logger)
	if err != nil {
		return File{}, errors.WithStack(err)
	}

	dest, err := destFile.Create()
	if err != nil {
		return File{}, errors.WithStack(err)
	}
	defer func(dest *os.File) {
		err := dest.Close()
		if err != nil {
			f.logCloseError(err)
		}
	}(dest)

	src, err := f.Open()
	if err != nil {
		return File{}, errors.WithStack(err)
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			f.logCloseError(err)
		}
	}(src)

	_, err = io.Copy(dest, src)
	if err != nil {
		return File{}, errors.WithStack(err)
	}

	return destFile, nil
}

func (f File) Move(destDir string) (File, error) {
	destFile, err := f.Copy(destDir)
	if err != nil {
		return File{}, errors.WithStack(err)
	}

	err = f.Remove()
	if err != nil {
		return File{}, errors.WithStack(err)
	}

	return destFile, nil
}

func (f File) ReadFile() ([]byte, error) {
	return os.ReadFile(f.FullPath())
}

func (f File) WriteBytes(content []byte) error {
	out, err := f.Create()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			f.logCloseError(err)
		}
	}(out)

	_, err = out.Write(content)
	if err != nil {
		return errors.WithStack(err)
	}

	err = out.Sync()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (f File) WriteString(content string) error {
	out, err := f.Create()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			f.logCloseError(err)
		}
	}(out)

	_, err = out.WriteString(content)
	if err != nil {
		return errors.WithStack(err)
	}

	err = out.Sync()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (f File) logCloseError(err error) {
	f.logger.Error().Err(err).Str("path", f.dir).Str("filename", f.Name()).Msg("failed to close file")
}
