# 📦 merge_files

A simple CLI tool for packing multiple files from a directory into a single binary file.

---

## 🚀 Features

- Pack multiple files into a single file
- Unpacking files by passing meta.json
- Stores metadata (offset + size) for each file
- Supports input directory selection
- Configurable output file path
- Built-in version and help commands

---

## 🛠 Build (Linux)

Run the build script:

```bash
./build_linux
```

After build:

```bash
./merge_files
```

---

## 📖 Usage

```bash
merge_files -h
merge_files -v
merge_files -p <path_files> -o <out_file>
merge_files -m <meta_file> -u <out_path>
```

---

## 📌 Flags

| Flag | Long form          | Description                           |
|------|--------------------|---------------------------------------|
| `-h` | `--help`           | Show help message                     |
| `-v` | `--version`        | Show application version              |
| `-p` | `--path`           | Path to the folder with files to pack |
| `-o` | `--out_file`       | Path to the output binary file        |
| `-m` | `--meta_file`      | Selecting a meta file to work with    |
| `-u` | `--unpacking_path` | Select a folder for unpacking         |

---

## 📂 Examples

```bash
merge_files -p "path_files" -o "out_file.bin"
merge_files -p "./path_files" -o "out_file.bin"
merge_files -p "../path_files" -o "out_file.bin"
merge_files -p "/root/path_files" -o "out_file.bin"
merge_files -m "meta.json" -u "out_path"
merge_files -m "./meta.json" -u "out_path"
merge_files -m "../meta.json" -u "out_path"
merge_files -m "/root/meta.json" -u "out_path"
```

---

## ℹ️ Built-in Help Output

```text
available commands:

  -h,  --help                  show this help message
  -v,  --version               show application version
  -p,  --path                  folder for packaging
  -o,  --out_file              path to the file where to pack
  -m,  --meta_file             selecting a meta file to work with
  -u,  --unpacking_path        select a folder for unpacking

usage:
  merge_files -h
  merge_files -v
  merge_files -p <path_files> -o <out_file>
  merge_files -m <meta_file> -u <out_path>

examples:
  merge_files -p "path_files" -o "out_file.bin"
  merge_files -p "./path_files" -o "out_file.bin"
  merge_files -p "../path_files" -o "out_file.bin"
  merge_files -p "/root/path_files" -o "out_file.bin"
  merge_files -m "meta.json" -u "out_path"
  merge_files -m "./meta.json" -u "out_path"
  merge_files -m "../meta.json" -u "out_path"
  merge_files -m "/root/meta.json" -u "out_path"

```

---

## 📦 How it works

The tool:

- Reads all files from the specified directory
- Packs them into a single binary file
- Unpacking files by passing meta data
- Generates metadata (file offset + size)

---

## ⚡ Roadmap

- [x] unpacking files
- [ ] recursive addition of folders and files
- [ ] gzip compression support
- [ ] SHA256 integrity check
      
