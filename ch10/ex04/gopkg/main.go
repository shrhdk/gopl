package main

import (
	"bytes"
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
	var depth int
	walkImports(os.Args[1], func(pkg *build.Package) {
		fmt.Println(indent(depth) + pkg.ImportPath)
		depth++
	}, func(pkg *build.Package) {
		depth--
	})
}

func indent(depth int) string {
	var buf bytes.Buffer
	for i := 0; i < depth; i++ {
		buf.WriteString("  ")
	}
	return buf.String()
}

func walkImports(importPath string, start, end func(*build.Package)) {
	seen := make(map[string]bool)
	var walk func(string, func(*build.Package), func(*build.Package))
	walk = func(importPath string, start, end func(*build.Package)) {
		if seen[importPath] {
			return
		}

		pkginfo, err := readPkg(importPath)
		if err != nil {
			return
		}

		seen[importPath] = true

		start(pkginfo)
		for _, dep := range pkginfo.Imports {
			walk(dep, start, end)
		}
		end(pkginfo)
	}
	walk(importPath, start, end)
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
