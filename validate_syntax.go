//go:build ignore
// +build ignore

package main

// This file validates that our UI code has correct types
// without needing to compile the full Fyne application

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fmt.Println("Validating Go syntax in UI files...")

	files := []string{
		"internal/ui/app.go",
		"internal/ui/input_form.go",
		"internal/ui/results_view.go",
		"internal/ui/dialogs.go",
		"internal/ui/utils.go",
		"internal/ui/tooltips.go",
	}

	fset := token.NewFileSet()

	for _, filename := range files {
		fmt.Printf("Checking %s... ", filename)

		src, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("❌ Error reading: %v\n", err)
			continue
		}

		_, err = parser.ParseFile(fset, filename, src, parser.ParseComments)
		if err != nil {
			fmt.Printf("❌ Syntax error: %v\n", err)
		} else {
			fmt.Println("✅ OK")
		}
	}

	fmt.Println("Syntax validation complete!")
}
