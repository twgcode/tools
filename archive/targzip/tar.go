/**
@Author: wei-g
@Email: guantingwei@sixents.com
@Date: 2022/8/1 15:01
@Description: 开箱即用的 tar + gzip 进行压缩和解压文件
*/

package targzip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Compress 打包压缩
func Compress(source []string, dest string) (err error) {
	files := make([]*os.File, 0, len(source))
	for _, v := range source {
		var (
			f *os.File
		)
		if f, err = os.Open(v); err != nil {
			err = fmt.Errorf("open file %q failed, err: %w", v, err)
			return
		}
		files = append(files, f)
	}
	var d *os.File
	if d, err = os.Create(dest); err != nil {
		err = fmt.Errorf("create tar package file %q failed, err: %w", dest, err)
		return
	}
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err = compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) (err error) {
	var (
		info      os.FileInfo
		header    *tar.Header
		fileInfos []os.FileInfo
	)
	info, err = file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if header, err = tar.FileInfoHeader(info, ""); err != nil {
			return err
		}
		header.Name = filepath.Join(prefix, header.Name)
		if err = tw.WriteHeader(header); err != nil {
			return err
		}

		prefix = filepath.Join(prefix, info.Name())
		if fileInfos, err = file.Readdir(-1); err != nil {
			return err
		}
		for _, fi := range fileInfos {
			var f *os.File
			f, err = os.Open(filepath.Join(file.Name(), fi.Name()))
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}

		if err = file.Close(); err != nil {
			return err
		}

	} else {
		header, err = tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		//header.Format = tar.FormatGNU
		header.Name = filepath.Join(prefix, header.Name)

		if err = tw.WriteHeader(header); err != nil {
			return err
		}

		if _, err = io.Copy(tw, file); err != nil {
			return
		}

		if err = file.Close(); err != nil {
			return
		}
	}
	return nil
}

// Decompress 解压
func Decompress(src, dst string) (err error) {
	var srcFile *os.File
	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()
	dst = filepath.Clean(dst)
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}
	err = decompress(srcFile, dst)
	return
}

// decompress will decompress a tar.gz archive, moving all files and folders
// within the archive (parameter 1) to an output directory (parameter 2).
func decompress(src io.Reader, dst string) (err error) {
	// TODO 未保留 解压文件/目录的 最后修改时间
	var gr *gzip.Reader
	gr, err = gzip.NewReader(src)
	if err != nil {
		return fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		var header *tar.Header
		header, err = tr.Next()

		switch {
		// If no more files are found return
		case err == io.EOF:
			return nil

		// Return any other error
		case err != nil:
			return err

		// If the header is nil, skip it
		case header == nil:
			continue
		}

		// The target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)
		fi := header.FileInfo()
		if fi.IsDir() {
			_ = os.MkdirAll(target, os.ModePerm)
			continue
		}

		// 预防 目标文件的 父目录不存在
		if dst != filepath.Dir(target) {
			if err = os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
				return err
			}
		}
		var fd *os.File
		fd, err = os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fi.Mode())
		if err != nil {
			return err
		}

		// NOTE: We use looped CopyN() not Copy() to avoid gosec G110 (CWE-409):
		// Potential DoS vulnerability via decompression bomb.
		for {
			_, err = io.CopyN(fd, tr, 1024)
			if err != nil {
				if err == io.EOF {
					err = nil
					break
				}
				return err
			}
		}
		_ = fd.Close()
	}
}
