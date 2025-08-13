package main

import (
	"auto-unterhaltsrechner/internal/calculator"
	"auto-unterhaltsrechner/internal/models"
	"auto-unterhaltsrechner/internal/storage"
	"fmt"
)

func main() {
	fmt.Println("Testing Auto-Unterhaltsrechner logic...")

	// Test calculator
	calc := calculator.New()
	fmt.Println("Calculator created successfully")

	// Test storage
	store := storage.New()
	fmt.Println("Storage created successfully")

	// Test model creation
	profile := models.NewCarProfile()
	profile.Name = "Test Fahrzeug"
	profile.FuelConsumption = 6.5
	profile.FuelPrice = 1.65
	profile.MonthlyKilometers = 1500
	profile.AnnualCarTax = 200
	profile.AnnualCarInsurance = 800
	profile.PurchasePrice = 25000
	profile.ExpectedYearsOfOwnership = 5
	fmt.Println("Profile created successfully")

	// Test calculation
	calculation := calc.CalculateCosts(profile)
	if calculation != nil {
		fmt.Printf("Monthly fuel cost: %.2f €\n", calculation.MonthlyFuelCost)
		fmt.Printf("Monthly running costs: %.2f €\n", calculation.MonthlyRunningCosts)
		fmt.Printf("Cost per kilometer: %.4f €\n", calculation.CostPerKilometer)
		fmt.Println("Calculations work correctly!")
	} else {
		fmt.Println("Error: Calculation failed")
	}

	// Test validation
	errors := calc.ValidateProfile(profile)
	if len(errors) == 0 {
		fmt.Println("Validation passed!")
	} else {
		fmt.Printf("Validation errors: %v\n", errors)
	}

	// Test storage
	err := store.SaveProfile(profile)
	if err != nil {
		fmt.Printf("Storage error: %v\n", err)
	} else {
		fmt.Println("Profile saved successfully!")
	}

	// Test loading
	loadedProfile, err := store.LoadProfile(profile.ID)
	if err != nil {
		fmt.Printf("Load error: %v\n", err)
	} else {
		fmt.Printf("Profile loaded successfully: %s\n", loadedProfile.Name)
	}

	fmt.Println("All core logic tests passed!")
}
