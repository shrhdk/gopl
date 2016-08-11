package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
)

var ctxt build.Context

func init() {
	ctxt = build.Default
}

func main() {
	walkImports(os.Args[1], make(map[string]bool), 0, func(importPath string, depth int) {
		for i := 0; i < depth; i++ {
			fmt.Print("  ")
		}
		fmt.Println(importPath)
	})
}

func walkImports(importPath string, seen map[string]bool, depth int, f func(string, int)) {
	if seen[importPath] {
		return
	}

	pkginfo, err := readPkg(importPath)
	if err != nil {
		return
	}

	seen[importPath] = true

	f(pkginfo.ImportPath, depth)
	for _, dep := range pkginfo.Imports {
		walkImports(dep, seen, depth+1, f)
	}
}

func readPkg(importPath string) (*build.Package, error) {
	pkginfo, err := readPkgFrom(ctxt.GOROOT, importPath)
	if err == nil {
		return pkginfo, nil
	}

	pkginfo, err = readPkgFrom(ctxt.GOPATH, importPath)
	if err == nil {
		return pkginfo, nil
	}

	return nil, fmt.Errorf("%s is not found", importPath)
}

func readPkgFrom(base, importPath string) (*build.Package, error) {
	path := filepath.Join(base, "src", importPath)
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	return ctxt.ImportDir(path, 0)
}
