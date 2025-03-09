# GOCLI

A Go cli tool that implements the popular UNIX commands `cat`, `wc`, and `ls`.

## Prerequisites

-   Go installed on your machine

# Usage

1️⃣ Build the modules and run them as executables

```bash
go build ./cmd/ls && go build ./cmd/cat && go build ./cmd/wc
```

**cat**
Display contents of file(s)

```bash
./cat filename.txt
```

**wc**
Count lines, words, and bytes of a file

```bash
./wc filename.txt
```

**ls**
List the contents of a directory

```bash
./ls filename.txt
```
