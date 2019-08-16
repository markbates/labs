package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedName | packages.NeedSyntax | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedTypesInfo}
	pkgs, err := packages.Load(cfg, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Print the names of the source files
	// for each package listed on the command line.
	for _, pkg := range pkgs {
		fmt.Println(pkg.ID)
		fmt.Println("ID", pkg.ID)
		fmt.Println("Name", pkg.Name)
		fmt.Println("PkgPath", pkg.PkgPath)
		// fmt.Println("Errors", pkg.Errors)
		fmt.Println("GoFiles", pkg.GoFiles)
		// fmt.Println("CompiledGoFiles", pkg.CompiledGoFiles)
		// fmt.Println("OtherFiles", pkg.OtherFiles)
		// fmt.Println("ExportFile", pkg.ExportFile)
		// fmt.Println("Imports", pkg.Imports)
		// fmt.Println("Types", pkg.Types)

		// s := pkg.Types.Scope()
		// fmt.Printf("### main.go:37 s (%T) -> %q %+v\n", s, s, s)
		// s.WriteTo(os.Stdout, 0, true)
		// for _, n := range s.Names() {
		// 	fmt.Printf("### _poc/main.go:38 n (%T) -> %q %+v\n", n, n, n)
		// 	o := s.Lookup(n)
		// 	fmt.Printf("### _poc/main.go:40 o.String() (%T) -> %q %+v\n", o.String(), o.String(), o.String())
		// }
		// fmt.Println("Fset", pkg.Fset)
		// fmt.Println("IllTyped", pkg.IllTyped)
		// fmt.Println("Syntax", pkg.Syntax)
		// fmt.Println("TypesInfo", pkg.TypesInfo)
		// fmt.Println("TypesSizes", pkg.TypesSizes)
		// pkg.Scope//

		// 	for _, v := range pkg.Syntax {
		// 		for _, d := range v.Decls {
		// 			switch t := d.(type) {
		// 			case *ast.GenDecl:
		// 				genDecl(t)
		// 			default:
		// 				// fmt.Println(t)
		// 			}
		// 		}
		// 		// fmt.Printf("### _poc/main.go:42 v (%T) -> %q %+v\n", v, v, v)
		// 	}
	}
}

// func genDecl(t *ast.GenDecl) {
// 	switch t.Tok {
// 	case token.TYPE:
// 		for _, s := range t.Specs {
// 			switch st := s.(type) {
// 			case *ast.TypeSpec:
// 			}
// 		}
// 	}
// }
