# Go Cook
A Go CLI to create recipe and cooking tip files from templates. The goal of this
project is to organize and standardize cooking information that can be parsed
to individual pages on a website.

## Installation
Ensure that you have a Golang version `>= 1.19`.

```bash
make install
```

## Usage
The CLI is meant to generate template files for recipes and tips, use the following:

```bash
go-cook --help
```

### File placement
Files are created in the `~/cook/` directory, all `git` wrapper commands point
to this directory exclusively.

