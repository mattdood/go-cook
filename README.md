<div align="center">
    <img src="https://github.com/mattdood/go-cook/raw/master/assets/ein-ed-cook.gif" alt="Gif of Spike Spiegel from Cowboy Bebop in space"/>
</div>

**Note:** I do not make any claims to the [Cowboy Bebop](https://en.wikipedia.org/wiki/Cowboy_Bebop) assets, names, or trademarks.

# Go Cook
A Go CLI to create recipe and cooking tip files from templates. The goal of this
project is to organize and standardize cooking information that can be parsed
to individual pages on a website.

<img src="https://img.shields.io/github/issues/mattdood/go-cook"
    target="https://github.com/mattdood/go-cook/issues"
    alt="Badge for GitHub issues."/>
<img src="https://img.shields.io/github/forks/mattdood/go-cook"
    target="https://github.com/mattdood/go-cook/forks"
    alt="Badge for GitHub forks."/>
<img src="https://img.shields.io/github/stars/mattdood/go-cook"
    alt="Badge for GitHub stars."/>
<img src="https://img.shields.io/github/license/mattdood/go-cook"
    target="https://github.com/mattdood/go-cook/raw/master/LICENSE"
    alt="Badge for GitHub license, MIT."/>
<img src="https://img.shields.io/twitter/url?url=https%3A%2F%2Fgithub.com%2Fmattdood%2Fgo-cook"
    target="https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fmattdood%2Fgo-cook"
    alt="Badge for sharable Twitter link."/>
<img src="https://img.shields.io/github/go-mod/go-version/mattdood/go-cook"
    alt="Badge for Golang version." />

## Installation
Ensure that you have a Golang version `>= 1.19`.

```bash
# If git cloning
make install

# If using go
go install https://github.com/mattdood/go-cook
```

## Usage
The CLI is meant to generate template files for recipes and tips, use the following:

```bash
go-cook --help

# Create a template
go-cook create -title "some recipe" -category "dinner" -tags "beef pasta tomatoes" -type "recipe"
```

### File placement
Files are created in the `~/cook/` directory, all `git` wrapper commands point
to this directory exclusively.

## Concepts
This repository's purpose is to give users a defined, parseable format for recipes
to avoid having content that has **endless** story above the actual implementation
of the dish.

An ancillary goal of this project is to define a standard format that encompasses
most recipes/tips/etc. that may find a home in a cook book. This would create a
set of validation criteria that can be used to ensure each recipe adhears to some
agreed upon format.
