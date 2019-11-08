package filesystem

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//reference: https://blog.depado.eu/post/copy-files-and-directories-in-go
func CopyDir(src string, dst string) error {
	var (
		err     error
		fds     []os.FileInfo
		srcinfo os.FileInfo
	)

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				return err
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyFile(src, dst string) error {
	var (
		err     error
		srcfd   *os.File
		dstfd   *os.File
		srcinfo os.FileInfo
	)

	if srcfd, err = os.Open(src); err != nil {
		return err
	}

	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}

	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	return os.Chmod(dst, srcinfo.Mode())
}
