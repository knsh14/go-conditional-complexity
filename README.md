# go-conditional-complexity

# Install

```
$ go get github.com/knsh14/go-conditional-complexity/cmd/gocc
```

# Usage

```
$ gocc -max 16 -exclude testdata -top 10 -avg ./
```

## -max N
threshold to show. command shows more complicated functions than this parameter.   
default value is 12. it is same score to flake8.

## -exclude PATTERN
command will not checks any files to match input pattern.  
pattern is complied to go regular expression.

## -top N
output most complicated functions

## -avg
output average complexity for all functions and filtered functions if this flag is set.

## ARG
path to check complexity.  
checks all go files under input directory recursive.

if no path is passed. checks current directory.
