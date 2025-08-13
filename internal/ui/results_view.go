package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createResultsView() *fyne.Container {
	return container.NewVBox(
		widget.NewCard("Berechnungsergebnisse", "", container.NewVBox(
			widget.NewLabel("Wählen Sie ein Profil aus oder erstellen Sie ein neues, um die Kosten zu berechnen."),
		)),
	)
}

func (a *App) updateResults() {
	if a.currentProfile == nil {
		a.resultsView.RemoveAll()
		a.resultsView.Add(widget.NewCard("Berechnungsergebnisse", "", container.NewVBox(
			widget.NewLabel("Wählen Sie ein Profil aus oder erstellen Sie ein neues, um die Kosten zu berechnen."),
		)))
		return
	}

	calculation := a.calculator.CalculateCosts(a.currentProfile)
	if calculation == nil {
		return
	}

	// Monthly costs section
	monthlyCostsContent := container.NewVBox(
		widget.NewLabel("Kraftstoffkosten: "+FormatCurrency(calculation.MonthlyFuelCost)),
		widget.NewLabel("Stromkosten: "+FormatCurrency(calculation.MonthlyElectricityCost)),
		widget.NewLabel("KFZ-Steuer: "+FormatCurrency(a.currentProfile.AnnualCarTax/12)),
		widget.NewLabel("Versicherung: "+FormatCurrency(a.currentProfile.AnnualCarInsurance/12)),
		widget.NewLabel("Finanzierung: "+FormatCurrency(a.currentProfile.FinancingRate)),
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Gesamt: "+FormatCurrency(calculation.MonthlyRunningCosts)+"**"),
	)

	// Annual costs section
	annualCostsContent := container.NewVBox(
		widget.NewLabel("Kraftstoffkosten: "+FormatCurrency(calculation.AnnualFuelCost)),
		widget.NewLabel("Stromkosten: "+FormatCurrency(calculation.AnnualElectricityCost)),
		widget.NewLabel("KFZ-Steuer: "+FormatCurrency(a.currentProfile.AnnualCarTax)),
		widget.NewLabel("Versicherung: "+FormatCurrency(a.currentProfile.AnnualCarInsurance)),
		widget.NewLabel("Finanzierung: "+FormatCurrency(a.currentProfile.FinancingRate*12)),
		widget.NewSeparator(),
		widget.NewRichTextFromMarkdown("**Gesamt: "+FormatCurrency(calculation.AnnualRunningCosts)+"**"),
	)

	// Depreciation section
	depreciationContent := container.NewVBox(
		widget.NewLabel("Gesamter Wertverlust: "+FormatCurrency(calculation.TotalDepreciation)),
		widget.NewLabel("Jährlicher Wertverlust: "+FormatCurrency(calculation.AnnualDepreciation)),
	)

	// Key metrics section
	keyMetricsContent := container.NewVBox(
		widget.NewLabel("Kosten pro Kilometer: "+FormatCurrency(calculation.CostPerKilometer)),
		widget.NewLabel("Gesamtkosten der Nutzung: "+FormatCurrency(calculation.TotalCostOfOwnership)),
	)

	// Consumption information
	consumptionContent := container.NewVBox()

	if a.currentProfile.FuelConsumption > 0 {
		monthlyFuelAmount := (a.currentProfile.FuelConsumption * a.currentProfile.MonthlyKilometers) / 100
		annualFuelAmount := monthlyFuelAmount * 12
		consumptionContent.Add(widget.NewLabel("Monatlicher Kraftstoffverbrauch: " + FormatLiters(monthlyFuelAmount)))
		consumptionContent.Add(widget.NewLabel("Jährlicher Kraftstoffverbrauch: " + FormatLiters(annualFuelAmount)))

		if a.currentProfile.TankSize > 0 {
			tanksPerMonth := monthlyFuelAmount / a.currentProfile.TankSize
			consumptionContent.Add(widget.NewLabel("Tankfüllungen pro Monat: " + FormatGermanNumber(tanksPerMonth, 1)))
		}
	}

	if a.currentProfile.ElectricConsumption > 0 {
		monthlyElectricAmount := (a.currentProfile.ElectricConsumption * a.currentProfile.MonthlyKilometers) / 100
		annualElectricAmount := monthlyElectricAmount * 12
		consumptionContent.Add(widget.NewLabel("Monatlicher Stromverbrauch: " + FormatKWh(monthlyElectricAmount)))
		consumptionContent.Add(widget.NewLabel("Jährlicher Stromverbrauch: " + FormatKWh(annualElectricAmount)))

		if a.currentProfile.BatterySize > 0 {
			chargesPerMonth := monthlyElectricAmount / a.currentProfile.BatterySize
			consumptionContent.Add(widget.NewLabel("Ladevorgänge pro Monat: " + FormatGermanNumber(chargesPerMonth, 1)))
		}
	}

	// Range information
	rangeContent := container.NewVBox()

	if a.currentProfile.TankSize > 0 && a.currentProfile.FuelConsumption > 0 {
		fuelRange := (a.currentProfile.TankSize / a.currentProfile.FuelConsumption) * 100
		rangeContent.Add(widget.NewLabel("Reichweite mit vollem Tank: " + FormatKilometers(fuelRange)))
	}

	if a.currentProfile.BatterySize > 0 && a.currentProfile.ElectricConsumption > 0 {
		electricRange := (a.currentProfile.BatterySize / a.currentProfile.ElectricConsumption) * 100
		rangeContent.Add(widget.NewLabel("Elektrische Reichweite: " + FormatKilometers(electricRange)))
	}

	// Update results view
	a.resultsView.RemoveAll()
	a.resultsView.Add(widget.NewCard("Monatliche Kosten", "", monthlyCostsContent))
	a.resultsView.Add(widget.NewCard("Jährliche Kosten", "", annualCostsContent))
	a.resultsView.Add(widget.NewCard("Wertverlust", "", depreciationContent))
	a.resultsView.Add(widget.NewCard("Kennzahlen", "", keyMetricsContent))

	if consumptionContent.Objects != nil && len(consumptionContent.Objects) > 0 {
		a.resultsView.Add(widget.NewCard("Verbrauch", "", consumptionContent))
	}

	if rangeContent.Objects != nil && len(rangeContent.Objects) > 0 {
		a.resultsView.Add(widget.NewCard("Reichweite", "", rangeContent))
	}
}
