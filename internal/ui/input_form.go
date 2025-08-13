package ui

import (
	"auto-unterhaltsrechner/internal/models"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createInputForm() *fyne.Container {
	translations := a.getCurrentTranslations()
	// Profile selection
	a.profileSelect = widget.NewSelect([]string{}, func(value string) {
		a.loadSelectedProfile(value)
	})

	// Basic information
	a.nameEntry = widget.NewEntry()
	a.nameEntry.SetPlaceHolder("Profilname eingeben...")
	a.nameEntry.OnChanged = func(text string) {
		if a.currentProfile != nil {
			a.currentProfile.Name = text
			a.updateResults()
		}
	}

	// Fuel consumption
	a.fuelConsumptionEntry = widget.NewEntry()
	a.fuelConsumptionEntry.SetPlaceHolder("z.B. 6,5")
	a.fuelConsumptionEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "fuel_consumption")
	}

	// Electric consumption
	a.electricConsumptionEntry = widget.NewEntry()
	a.electricConsumptionEntry.SetPlaceHolder("z.B. 18,5")
	a.electricConsumptionEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "electric_consumption")
	}

	// Fuel price
	a.fuelPriceEntry = widget.NewEntry()
	a.fuelPriceEntry.SetPlaceHolder("z.B. 1,65")
	a.fuelPriceEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "fuel_price")
	}

	// Electricity price
	a.electricityPriceEntry = widget.NewEntry()
	a.electricityPriceEntry.SetPlaceHolder("z.B. 0,35")
	a.electricityPriceEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "electricity_price")
	}

	// Fuel type
	fuelTypes := a.getTranslatedFuelTypes()
	a.fuelTypeSelect = widget.NewSelect(fuelTypes, func(value string) {
		if a.currentProfile != nil {
			fuelTypeKey := a.getFuelTypeFromTranslation(value)
			a.currentProfile.FuelType = models.FuelType(fuelTypeKey)
			a.updateResults()
		}
	})

	// Electricity type
	electricityTypes := a.getTranslatedElectricityTypes()
	a.electricityTypeSelect = widget.NewSelect(electricityTypes, func(value string) {
		if a.currentProfile != nil {
			electricityTypeKey := a.getElectricityTypeFromTranslation(value)
			a.currentProfile.ElectricityType = models.ElectricityType(electricityTypeKey)
			a.updateResults()
		}
	})

	// Tank size
	a.tankSizeEntry = widget.NewEntry()
	a.tankSizeEntry.SetPlaceHolder("z.B. 50")
	a.tankSizeEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "tank_size")
	}

	// Battery size
	a.batterySizeEntry = widget.NewEntry()
	a.batterySizeEntry.SetPlaceHolder("z.B. 75")
	a.batterySizeEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "battery_size")
	}

	// Monthly kilometers
	a.monthlyKmEntry = widget.NewEntry()
	a.monthlyKmEntry.SetPlaceHolder("z.B. 1500")
	a.monthlyKmEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "monthly_km")
	}

	// Annual tax
	a.annualTaxEntry = widget.NewEntry()
	a.annualTaxEntry.SetPlaceHolder("z.B. 200")
	a.annualTaxEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "annual_tax")
	}

	// Annual insurance
	a.annualInsuranceEntry = widget.NewEntry()
	a.annualInsuranceEntry.SetPlaceHolder("z.B. 800")
	a.annualInsuranceEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "annual_insurance")
	}

	// Financing rate
	a.financingRateEntry = widget.NewEntry()
	a.financingRateEntry.SetPlaceHolder("z.B. 350")
	a.financingRateEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "financing_rate")
	}

	// Financing period
	a.financingPeriodEntry = widget.NewEntry()
	a.financingPeriodEntry.SetPlaceHolder("z.B. 60")
	a.financingPeriodEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "financing_period")
	}

	// Purchase price
	a.purchasePriceEntry = widget.NewEntry()
	a.purchasePriceEntry.SetPlaceHolder("z.B. 35000")
	a.purchasePriceEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "purchase_price")
	}

	// Ownership years
	a.ownershipYearsEntry = widget.NewEntry()
	a.ownershipYearsEntry.SetPlaceHolder("z.B. 5")
	a.ownershipYearsEntry.OnChanged = func(text string) {
		a.updateProfileFromEntry(text, "ownership_years")
	}

	// Create form sections
	profileForm := widget.NewForm(
		widget.NewFormItem(translations.ProfileSelect, a.profileSelect),
		widget.NewFormItem(translations.ProfileName, a.nameEntry),
	)
	profileSection := container.NewVBox(
		widget.NewCard(translations.ProfileTitle, "", profileForm),
	)

	consumptionForm := widget.NewForm(
		widget.NewFormItem(translations.FuelConsumption, a.fuelConsumptionEntry),
		widget.NewFormItem(translations.ElectricConsumption, a.electricConsumptionEntry),
	)
	consumptionSection := container.NewVBox(
		widget.NewCard(translations.ConsumptionTitle, "", consumptionForm),
	)

	pricesForm := widget.NewForm(
		widget.NewFormItem(translations.FuelPrice, a.fuelPriceEntry),
		widget.NewFormItem(translations.ElectricityPrice, a.electricityPriceEntry),
		widget.NewFormItem(translations.FuelType, a.fuelTypeSelect),
		widget.NewFormItem(translations.ElectricityType, a.electricityTypeSelect),
	)
	pricesSection := container.NewVBox(
		widget.NewCard(translations.PricesTitle, "", pricesForm),
	)

	capacityForm := widget.NewForm(
		widget.NewFormItem(translations.TankSize, a.tankSizeEntry),
		widget.NewFormItem(translations.BatterySize, a.batterySizeEntry),
	)
	capacitySection := container.NewVBox(
		widget.NewCard(translations.CapacityTitle, "", capacityForm),
	)

	usageForm := widget.NewForm(
		widget.NewFormItem(translations.MonthlyKilometers, a.monthlyKmEntry),
	)
	usageSection := container.NewVBox(
		widget.NewCard(translations.UsageTitle, "", usageForm),
	)

	costsForm := widget.NewForm(
		widget.NewFormItem(translations.AnnualTax, a.annualTaxEntry),
		widget.NewFormItem(translations.AnnualInsurance, a.annualInsuranceEntry),
	)
	costsSection := container.NewVBox(
		widget.NewCard(translations.CostsTitle, "", costsForm),
	)

	financingForm := widget.NewForm(
		widget.NewFormItem(translations.FinancingRate, a.financingRateEntry),
		widget.NewFormItem(translations.FinancingPeriod, a.financingPeriodEntry),
	)
	financingSection := container.NewVBox(
		widget.NewCard(translations.FinancingTitle, "", financingForm),
	)

	depreciationForm := widget.NewForm(
		widget.NewFormItem(translations.PurchasePrice, a.purchasePriceEntry),
		widget.NewFormItem(translations.OwnershipYears, a.ownershipYearsEntry),
	)
	depreciationSection := container.NewVBox(
		widget.NewCard(translations.DepreciationTitle, "", depreciationForm),
	)

	return container.NewVBox(
		profileSection,
		consumptionSection,
		pricesSection,
		capacitySection,
		usageSection,
		costsSection,
		financingSection,
		depreciationSection,
	)
}

func (a *App) updateProfileFromEntry(text, field string) {
	if a.currentProfile == nil {
		return
	}

	value, err := ParseGermanNumber(text)
	if err != nil && text != "" {
		return // Invalid number, skip update
	}

	switch field {
	case "fuel_consumption":
		a.currentProfile.FuelConsumption = value
	case "electric_consumption":
		a.currentProfile.ElectricConsumption = value
	case "fuel_price":
		a.currentProfile.FuelPrice = value
	case "electricity_price":
		a.currentProfile.ElectricityPrice = value
	case "tank_size":
		a.currentProfile.TankSize = value
	case "battery_size":
		a.currentProfile.BatterySize = value
	case "monthly_km":
		a.currentProfile.MonthlyKilometers = value
	case "annual_tax":
		a.currentProfile.AnnualCarTax = value
	case "annual_insurance":
		a.currentProfile.AnnualCarInsurance = value
	case "financing_rate":
		a.currentProfile.FinancingRate = value
	case "purchase_price":
		a.currentProfile.PurchasePrice = value
	case "financing_period":
		a.currentProfile.FinancingPeriod = int(value)
	case "ownership_years":
		a.currentProfile.ExpectedYearsOfOwnership = int(value)
	}

	a.updateResults()
}

func (a *App) updateInputForm() {
	if a.currentProfile == nil {
		return
	}

	a.nameEntry.SetText(a.currentProfile.Name)
	a.fuelConsumptionEntry.SetText(FormatGermanNumber(a.currentProfile.FuelConsumption, 1))
	a.electricConsumptionEntry.SetText(FormatGermanNumber(a.currentProfile.ElectricConsumption, 1))
	a.fuelPriceEntry.SetText(FormatGermanNumber(a.currentProfile.FuelPrice, 2))
	a.electricityPriceEntry.SetText(FormatGermanNumber(a.currentProfile.ElectricityPrice, 2))
	a.fuelTypeSelect.SetSelected(a.translateFuelType(string(a.currentProfile.FuelType)))
	a.electricityTypeSelect.SetSelected(a.translateElectricityType(string(a.currentProfile.ElectricityType)))
	a.tankSizeEntry.SetText(FormatGermanNumber(a.currentProfile.TankSize, 0))
	a.batterySizeEntry.SetText(FormatGermanNumber(a.currentProfile.BatterySize, 0))
	a.monthlyKmEntry.SetText(FormatGermanNumber(a.currentProfile.MonthlyKilometers, 0))
	a.annualTaxEntry.SetText(FormatGermanNumber(a.currentProfile.AnnualCarTax, 0))
	a.annualInsuranceEntry.SetText(FormatGermanNumber(a.currentProfile.AnnualCarInsurance, 0))
	a.financingRateEntry.SetText(FormatGermanNumber(a.currentProfile.FinancingRate, 0))
	a.financingPeriodEntry.SetText(fmt.Sprintf("%d", a.currentProfile.FinancingPeriod))
	a.purchasePriceEntry.SetText(FormatGermanNumber(a.currentProfile.PurchasePrice, 0))
	a.ownershipYearsEntry.SetText(fmt.Sprintf("%d", a.currentProfile.ExpectedYearsOfOwnership))
}

func (a *App) updateProfileFromForm() {
	if a.currentProfile == nil {
		return
	}

	a.currentProfile.Name = a.nameEntry.Text

	if val, err := ParseGermanNumber(a.fuelConsumptionEntry.Text); err == nil {
		a.currentProfile.FuelConsumption = val
	}
	if val, err := ParseGermanNumber(a.electricConsumptionEntry.Text); err == nil {
		a.currentProfile.ElectricConsumption = val
	}
	if val, err := ParseGermanNumber(a.fuelPriceEntry.Text); err == nil {
		a.currentProfile.FuelPrice = val
	}
	if val, err := ParseGermanNumber(a.electricityPriceEntry.Text); err == nil {
		a.currentProfile.ElectricityPrice = val
	}
	if val, err := ParseGermanNumber(a.tankSizeEntry.Text); err == nil {
		a.currentProfile.TankSize = val
	}
	if val, err := ParseGermanNumber(a.batterySizeEntry.Text); err == nil {
		a.currentProfile.BatterySize = val
	}
	if val, err := ParseGermanNumber(a.monthlyKmEntry.Text); err == nil {
		a.currentProfile.MonthlyKilometers = val
	}
	if val, err := ParseGermanNumber(a.annualTaxEntry.Text); err == nil {
		a.currentProfile.AnnualCarTax = val
	}
	if val, err := ParseGermanNumber(a.annualInsuranceEntry.Text); err == nil {
		a.currentProfile.AnnualCarInsurance = val
	}
	if val, err := ParseGermanNumber(a.financingRateEntry.Text); err == nil {
		a.currentProfile.FinancingRate = val
	}
	if val, err := strconv.Atoi(a.financingPeriodEntry.Text); err == nil {
		a.currentProfile.FinancingPeriod = val
	}
	if val, err := ParseGermanNumber(a.purchasePriceEntry.Text); err == nil {
		a.currentProfile.PurchasePrice = val
	}
	if val, err := strconv.Atoi(a.ownershipYearsEntry.Text); err == nil {
		a.currentProfile.ExpectedYearsOfOwnership = val
	}

	a.currentProfile.FuelType = models.FuelType(a.getFuelTypeFromTranslation(a.fuelTypeSelect.Selected))
	a.currentProfile.ElectricityType = models.ElectricityType(a.getElectricityTypeFromTranslation(a.electricityTypeSelect.Selected))
}

func (a *App) loadSelectedProfile(value string) {
	if value == "" {
		return
	}

	// Extract ID from selection (format: "Name (ID)")
	parts := strings.Split(value, " (")
	if len(parts) < 2 {
		return
	}

	id := strings.TrimSuffix(parts[1], ")")

	profile, err := a.storage.LoadProfile(id)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	a.currentProfile = profile
	a.updateInputForm()
	a.updateResults()
}
