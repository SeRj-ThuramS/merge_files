package main

const (
	helpF = `available commands:

  -h,  --help                  show this help message
  -v,  --version               show application version
  -p,  --path                  folder for packaging
  -o,  --out_file              path to the file where to pack
  -m,  --meta_file             selecting a meta file to work with
  -u,  --unpacking_path        select a folder for unpacking

usage:
  %[1]s -h
  %[1]s -v
  %[1]s -p <path_files> -o <out_file>
  %[1]s -m <meta_file> -u <out_path>

examples:
  %[1]s -p "path_files" -o "out_file.bin"
  %[1]s -p "./path_files" -o "out_file.bin"
  %[1]s -p "../path_files" -o "out_file.bin"
  %[1]s -p "/root/path_files" -o "out_file.bin"
  %[1]s -m "meta.json" -u "out_path"
  %[1]s -m "./meta.json" -u "out_path"
  %[1]s -m "../meta.json" -u "out_path"
  %[1]s -m "/root/meta.json" -u "out_path"
`
)
