package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var (
	Version    = "development"
	Name       = "Name"
	Namespace  = "Namespace"
	WORKDIR    string
	EXECUTABLE string
	mArgs      []IArgs
	pathFiles  string
	outFile    string
	outPath    string
	metaFile   string
)

func init() {
	w, err := os.Executable()

	if err != nil {
		fmt.Println("error get current directory")
		os.Exit(1)
	}

	WORKDIR = filepath.Dir(w)
	EXECUTABLE = filepath.Base(w)

	mArgs = makeArgs(os.Args[1:])

	if len(mArgs) == 0 || len(mArgs) > 2 {
		fmt.Printf(helpF, EXECUTABLE)
		os.Exit(0)
	}
}

func main() {
	// help
	if hasArg(0, mArgs, "-h", "--help") {
		fmt.Printf(helpF, EXECUTABLE)
		os.Exit(0)
	}

	// version
	if hasArg(0, mArgs, "-v", "--version") {
		fmt.Printf("Name: %s\nNamespace: %s\nVersion: %s\n", Name, Namespace, Version)
		os.Exit(0)
	}

	// -p || --path
	if hasArg(0, mArgs, "-p", "--path") {
		if pathFiles = valueArg(0, mArgs); len(pathFiles) == 0 {
			fmt.Printf(helpF, EXECUTABLE)
			os.Exit(0)
		}

		// -o || --out_file
		if hasArg(1, mArgs, "-o", "--out_file") {
			if outFile = valueArg(1, mArgs); len(outFile) == 0 {
				fmt.Printf(helpF, EXECUTABLE)
				os.Exit(0)
			}

			packing()
			os.Exit(0)
		}
	}

	// -m || --meta_file
	if hasArg(0, mArgs, "-m", "--meta_file") {
		if metaFile = valueArg(0, mArgs); len(metaFile) == 0 {
			fmt.Printf(helpF, EXECUTABLE)
			os.Exit(0)
		}

		// -u || --unpackick_path
		if hasArg(1, mArgs, "-u", "--unpacking_path") {
			if outPath = valueArg(1, mArgs); len(outPath) == 0 {
				fmt.Printf(helpF, EXECUTABLE)
				os.Exit(0)
			}

			unpacking()
			os.Exit(0)
		}
	}

	fmt.Printf(helpF, EXECUTABLE)
	os.Exit(0)
}

func packing() {
	fmt.Printf("=> (s) selected packaging folder: %s\n", pathFiles)
	fmt.Printf("=> (s) selected file for packaging: %s\n", outFile)

	if !isDir(pathFiles) {
		fmt.Printf("=> (e) the folder <pathFiles> does not exist or cannot be accessed.\n")
		os.Exit(1)
	}

	entries, err := os.ReadDir(pathFiles)
	if err != nil {
		fmt.Printf("=> (e) error reading directory: %v\n", err)
		os.Exit(1)
	}

	var outFiles = make([]string, 0)

	for _, entry := range entries {
		fullPath := filepath.Join(pathFiles, entry.Name())
		outFiles = append(outFiles, fullPath)
	}

	// meta file
	ext := filepath.Ext(outFile)
	metaFile = outFile[:len(outFile)-len(ext)] + ".json"

	// archive compilation
	err = pack(outFiles, outFile, metaFile)
	if err != nil {
		fmt.Printf("=> (e) error packing data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("=> (s) files are packed into <%s>\n", outFile)
}

func unpacking() {
	fmt.Printf("=> (s) selected metafile for unpacking: %s\n", metaFile)
	fmt.Printf("=> (s) selected packaging folder: %s\n", outPath)

	if !isDir(outPath) {
		fmt.Printf("=> (e) the folder <out_path> does not exist or cannot be accessed.\n")
		os.Exit(1)
	}

	if !isFile(metaFile) {
		fmt.Printf("=> (e) the file <meta_file> does not exist or cannot be accessed.\n")
		os.Exit(1)
	}

	mFile, err := os.Open(metaFile)
	if err != nil {
		fmt.Printf("=> (e) error opening meta file: %v\n", err)
		os.Exit(1)
	}

	var meta FilesMeta
	err = json.NewDecoder(mFile).Decode(&meta)
	if err != nil {
		fmt.Printf("=> (e) error parsing meta file: %v\n", err)
		os.Exit(1)
	}

	err = unpack(outPath, meta)
	if err != nil {
		fmt.Printf("=> (e) error unpacking data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("=> (s) files are unpacked into <%s>\n", outPath)
}
