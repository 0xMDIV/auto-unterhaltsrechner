package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Test function to verify UI types compatibility
func testUITypes() {
	// Test that container.VBox can be assigned to fyne.Container
	vbox := container.NewVBox()

	// This should work without type errors
	var fyneContainer *fyne.Container = vbox

	fmt.Printf("VBox type: %T\n", vbox)
	fmt.Printf("Fyne Container type: %T\n", fyneContainer)

	// Test that both implement fyne.CanvasObject
	var canvasObj fyne.CanvasObject = vbox
	fmt.Printf("As CanvasObject: %T\n", canvasObj)

	fmt.Println("✅ All UI type conversions work correctly!")
}

func main() {
	fmt.Println("Testing UI type compatibility...")
	testUITypes()
}
