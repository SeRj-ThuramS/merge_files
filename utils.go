package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type IArgs struct {
	Command string
	Args    string
}

func makeArgs(args []string) []IArgs {
	params := make([]IArgs, 0)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if len(arg) > 0 && arg[0] == '-' {
			a := IArgs{
				Command: arg,
			}

			if i+1 < len(args) && args[i+1][0] != '-' {
				a.Args = args[i+1]
				i++
			} else {
				a.Args = ""
			}

			params = append(params, a)
		}
	}

	return params
}

func hasArg(index int, m []IArgs, keys ...string) bool {
	for i := 0; i < len(m); i++ {
		s := m[i].Command
		if i == index {
			for _, k := range keys {
				if s == k {
					return true
				}
			}
		}
	}

	return false
}

func valueArg(index int, m []IArgs) string {
	if index >= len(m) {
		return ""
	}

	return m[index].Args
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func pack(files []string, dataFile, metaFile string) error {
	out, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer out.Close()

	meta := FilesMeta{
		Filename: dataFile,
		Files:    make([]FileMeta, 0),
	}

	var offset int64 = 0

	i := 0
	lF := len(files)
	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		stat, _ := f.Stat()

		n, err := io.Copy(out, f)
		if err != nil {
			return err
		}

		meta.Files = append(meta.Files, FileMeta{
			ID:     int64(i + 1),
			Name:   filepath.Base(path),
			Offset: offset,
			Size:   n,
		})

		offset += n
		f.Close()

		i++
		fmt.Printf("=> (%d - %d) added %s (%d bytes)\n", i, lF, path, stat.Size())
	}

	metaOut, err := os.Create(metaFile)
	if err != nil {
		return err
	}
	defer metaOut.Close()

	return json.NewEncoder(metaOut).Encode(meta)
}

func unpack(outPath string, meta FilesMeta) error {
	data, err := os.Open(meta.Filename)
	if err != nil {
		return err
	}

	defer data.Close()

	if !isFile(meta.Filename) {
		return err
	}

	lF := len(meta.Files)
	i := 0
	for _, target := range meta.Files {
		buf := make([]byte, target.Size)
		_, err := data.ReadAt(buf, target.Offset)
		if err != nil {
			return err
		}

		// create file
		nameFile := filepath.Join(outPath, target.Name)
		out, err := os.Create(nameFile)
		if err != nil {
			return err
		}

		// write file
		_, err = out.Write(buf)
		if err != nil {
			return err
		}

		// close file
		err = out.Close()
		if err != nil {
			return err
		}

		i++
		fmt.Printf("=> (%d - %d) saved %s\n", i, lF, nameFile)
	}
	return nil
}
