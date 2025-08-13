package ui

import (
	"auto-unterhaltsrechner/internal/models"
	"encoding/csv"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jung-kurt/gofpdf"
)

func (a *App) showLoadDialog() {
	profiles, err := a.storage.ListProfiles()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	if len(profiles) == 0 {
		dialog.ShowInformation("Keine Profile", "Es sind keine gespeicherten Profile vorhanden.", a.window)
		return
	}

	var profileList []string
	profileMap := make(map[string]*models.CarProfile)

	for _, profile := range profiles {
		displayName := fmt.Sprintf("%s (%s)", profile.Name, profile.ID)
		profileList = append(profileList, displayName)
		profileMap[displayName] = profile
	}

	profileSelect := widget.NewSelect(profileList, nil)

	dialog.ShowCustomConfirm("Profil laden", "Laden", "Abbrechen",
		container.NewVBox(
			widget.NewLabel("Wählen Sie ein Profil zum Laden:"),
			profileSelect,
		),
		func(confirmed bool) {
			if confirmed && profileSelect.Selected != "" {
				selectedProfile := profileMap[profileSelect.Selected]
				a.currentProfile = selectedProfile
				a.updateInputForm()
				a.updateResults()
			}
		}, a.window)
}

func (a *App) showExportDialog() {
	if a.currentProfile == nil {
		dialog.ShowInformation("Kein Profil", "Bitte wählen Sie zuerst ein Profil aus.", a.window)
		return
	}

	calculation := a.calculator.CalculateCosts(a.currentProfile)
	if calculation == nil {
		dialog.ShowError(fmt.Errorf("Berechnungsfehler"), a.window)
		return
	}

	exportOptions := []string{"CSV Export", "JSON Export", "PDF Export"}
	exportSelect := widget.NewSelect(exportOptions, nil)
	exportSelect.SetSelected("CSV Export")

	dialog.ShowCustomConfirm("Export", "Exportieren", "Abbrechen",
		container.NewVBox(
			widget.NewLabel("Export-Format wählen:"),
			exportSelect,
		),
		func(confirmed bool) {
			if confirmed {
				switch exportSelect.Selected {
				case "CSV Export":
					a.exportToCSV(calculation)
				case "JSON Export":
					a.exportToJSON(a.currentProfile)
				case "PDF Export":
					a.exportToPDF(calculation)
				}
			}
		}, a.window)
}

func (a *App) exportToCSV(calculation *models.CostCalculation) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if writer == nil {
			return
		}
		defer writer.Close()

		csvWriter := csv.NewWriter(writer)
		defer csvWriter.Flush()

		// Write header
		csvWriter.Write([]string{"Kategorie", "Beschreibung", "Monatlich (€)", "Jährlich (€)"})

		// Write data
		csvWriter.Write([]string{"Kraftstoff", "Kraftstoffkosten",
			FormatGermanNumber(calculation.MonthlyFuelCost, 2),
			FormatGermanNumber(calculation.AnnualFuelCost, 2)})

		csvWriter.Write([]string{"Strom", "Stromkosten",
			FormatGermanNumber(calculation.MonthlyElectricityCost, 2),
			FormatGermanNumber(calculation.AnnualElectricityCost, 2)})

		csvWriter.Write([]string{"Steuer", "KFZ-Steuer",
			FormatGermanNumber(calculation.Profile.AnnualCarTax/12, 2),
			FormatGermanNumber(calculation.Profile.AnnualCarTax, 2)})

		csvWriter.Write([]string{"Versicherung", "Versicherung",
			FormatGermanNumber(calculation.Profile.AnnualCarInsurance/12, 2),
			FormatGermanNumber(calculation.Profile.AnnualCarInsurance, 2)})

		csvWriter.Write([]string{"Finanzierung", "Finanzierung",
			FormatGermanNumber(calculation.Profile.FinancingRate, 2),
			FormatGermanNumber(calculation.Profile.FinancingRate*12, 2)})

		csvWriter.Write([]string{"Gesamt", "Gesamtkosten",
			FormatGermanNumber(calculation.MonthlyRunningCosts, 2),
			FormatGermanNumber(calculation.AnnualRunningCosts, 2)})

		csvWriter.Write([]string{"Wertverlust", "Jährlicher Wertverlust",
			FormatGermanNumber(calculation.AnnualDepreciation/12, 2),
			FormatGermanNumber(calculation.AnnualDepreciation, 2)})

		csvWriter.Write([]string{"Kennzahlen", "Kosten pro Kilometer",
			FormatGermanNumber(calculation.CostPerKilometer, 4), ""})

		dialog.ShowInformation("Export erfolgreich", "Die Daten wurden erfolgreich exportiert.", a.window)
	}, a.window)
}

func (a *App) exportToJSON(profile *models.CarProfile) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if writer == nil {
			return
		}
		defer writer.Close()

		err = a.storage.ExportProfileToJSON(profile, writer.URI().Path())
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}

		dialog.ShowInformation("Export erfolgreich", "Das Profil wurde erfolgreich exportiert.", a.window)
	}, a.window)
}

func (a *App) exportToPDF(calculation *models.CostCalculation) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if writer == nil {
			return
		}
		defer writer.Close()

		// Create PDF
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		// Set font
		pdf.SetFont("Arial", "B", 16)

		// Title
		pdf.Cell(0, 10, "Auto-Unterhaltsrechner - Kostenaufstellung")
		pdf.Ln(15)

		// Profile info
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 8, "Fahrzeugprofil: "+calculation.Profile.Name)
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 10)
		pdf.Cell(0, 6, "Erstellt am: "+time.Now().Format("02.01.2006 15:04"))
		pdf.Ln(10)

		// Monthly costs section
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 8, "Monatliche Kosten")
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 10)
		pdf.Cell(80, 6, "Kraftstoffkosten:")
		pdf.Cell(0, 6, FormatCurrency(calculation.MonthlyFuelCost))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Stromkosten:")
		pdf.Cell(0, 6, FormatCurrency(calculation.MonthlyElectricityCost))
		pdf.Ln(6)

		pdf.Cell(80, 6, "KFZ-Steuer:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.AnnualCarTax/12))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Versicherung:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.AnnualCarInsurance/12))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Finanzierung/Leasing:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.FinancingRate))
		pdf.Ln(6)

		pdf.SetFont("Arial", "B", 10)
		pdf.Cell(80, 6, "Gesamtkosten monatlich:")
		pdf.Cell(0, 6, FormatCurrency(calculation.MonthlyRunningCosts))
		pdf.Ln(12)

		// Annual costs section
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 8, "Jährliche Kosten")
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 10)
		pdf.Cell(80, 6, "Kraftstoffkosten:")
		pdf.Cell(0, 6, FormatCurrency(calculation.AnnualFuelCost))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Stromkosten:")
		pdf.Cell(0, 6, FormatCurrency(calculation.AnnualElectricityCost))
		pdf.Ln(6)

		pdf.Cell(80, 6, "KFZ-Steuer:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.AnnualCarTax))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Versicherung:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.AnnualCarInsurance))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Finanzierung/Leasing:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.FinancingRate*12))
		pdf.Ln(6)

		pdf.SetFont("Arial", "B", 10)
		pdf.Cell(80, 6, "Gesamtkosten jährlich:")
		pdf.Cell(0, 6, FormatCurrency(calculation.AnnualRunningCosts))
		pdf.Ln(12)

		// Depreciation section
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 8, "Wertverlust")
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 10)
		pdf.Cell(80, 6, "Kaufpreis:")
		pdf.Cell(0, 6, FormatCurrency(calculation.Profile.PurchasePrice))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Besitzdauer:")
		pdf.Cell(0, 6, fmt.Sprintf("%d Jahre", calculation.Profile.ExpectedYearsOfOwnership))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Jährlicher Wertverlust:")
		pdf.Cell(0, 6, FormatCurrency(calculation.AnnualDepreciation))
		pdf.Ln(12)

		// Key metrics section
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 8, "Kennzahlen")
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 10)
		pdf.Cell(80, 6, "Kosten pro Kilometer:")
		pdf.Cell(0, 6, FormatCurrency(calculation.CostPerKilometer))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Monatliche Fahrleistung:")
		pdf.Cell(0, 6, FormatKilometers(calculation.Profile.MonthlyKilometers))
		pdf.Ln(6)

		pdf.Cell(80, 6, "Gesamtkosten der Nutzung:")
		pdf.Cell(0, 6, FormatCurrency(calculation.TotalCostOfOwnership))
		pdf.Ln(12)

		// Consumption section if applicable
		if calculation.Profile.FuelConsumption > 0 || calculation.Profile.ElectricConsumption > 0 {
			pdf.SetFont("Arial", "B", 12)
			pdf.Cell(0, 8, "Verbrauchsdaten")
			pdf.Ln(8)

			pdf.SetFont("Arial", "", 10)
			if calculation.Profile.FuelConsumption > 0 {
				monthlyFuel := (calculation.Profile.FuelConsumption * calculation.Profile.MonthlyKilometers) / 100
				pdf.Cell(80, 6, "Kraftstoffverbrauch monatlich:")
				pdf.Cell(0, 6, FormatLiters(monthlyFuel))
				pdf.Ln(6)
			}

			if calculation.Profile.ElectricConsumption > 0 {
				monthlyElectric := (calculation.Profile.ElectricConsumption * calculation.Profile.MonthlyKilometers) / 100
				pdf.Cell(80, 6, "Stromverbrauch monatlich:")
				pdf.Cell(0, 6, FormatKWh(monthlyElectric))
				pdf.Ln(6)
			}
		}

		// Save PDF
		err = pdf.Output(writer)
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}

		dialog.ShowInformation("Export erfolgreich", "Das PDF wurde erfolgreich erstellt.", a.window)
	}, a.window)
}

func (a *App) showComparisonDialog() {
	profiles, err := a.storage.ListProfiles()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	if len(profiles) < 2 {
		dialog.ShowInformation("Nicht genügend Profile", "Für einen Vergleich werden mindestens 2 Profile benötigt.", a.window)
		return
	}

	// Create comparison window
	compWindow := a.fyneApp.NewWindow("Fahrzeugvergleich")
	compWindow.Resize(fyne.NewSize(1200, 800))

	var selectedProfiles []*models.CarProfile
	var checkboxes []*widget.Check

	content := container.NewVBox()
	content.Add(widget.NewLabel("Wählen Sie bis zu 4 Profile für den Vergleich:"))

	for _, profile := range profiles {
		p := profile // Capture loop variable
		check := widget.NewCheck(fmt.Sprintf("%s (%s)", p.Name, p.ID), func(checked bool) {
			if checked {
				if len(selectedProfiles) < 4 {
					selectedProfiles = append(selectedProfiles, p)
				} else {
					// Uncheck if more than 4 selected
					for _, cb := range checkboxes {
						if cb.Text == fmt.Sprintf("%s (%s)", p.Name, p.ID) {
							cb.SetChecked(false)
							break
						}
					}
				}
			} else {
				// Remove from selected profiles
				for i, sp := range selectedProfiles {
					if sp.ID == p.ID {
						selectedProfiles = append(selectedProfiles[:i], selectedProfiles[i+1:]...)
						break
					}
				}
			}
		})
		checkboxes = append(checkboxes, check)
		content.Add(check)
	}

	compareButton := widget.NewButton("Vergleichen", func() {
		if len(selectedProfiles) < 2 {
			dialog.ShowInformation("Auswahl unvollständig", "Bitte wählen Sie mindestens 2 Profile aus.", compWindow)
			return
		}
		a.showComparisonResults(selectedProfiles, compWindow)
	})

	content.Add(widget.NewSeparator())
	content.Add(compareButton)

	compWindow.SetContent(container.NewScroll(content))
	compWindow.Show()
}

func (a *App) showComparisonResults(profiles []*models.CarProfile, parentWindow fyne.Window) {
	resultsWindow := a.fyneApp.NewWindow("Vergleichsergebnisse")
	resultsWindow.Resize(fyne.NewSize(1400, 900))

	// Calculate costs for all profiles
	var calculations []*models.CostCalculation
	for _, profile := range profiles {
		calc := a.calculator.CalculateCosts(profile)
		calculations = append(calculations, calc)
	}

	// Create comparison table
	table := a.createComparisonTable(profiles, calculations)

	// Create charts (simplified version)
	chartsContent := a.createComparisonCharts(profiles, calculations)

	tabs := container.NewAppTabs(
		container.NewTabItem("Tabelle", container.NewScroll(table)),
		container.NewTabItem("Diagramme", container.NewScroll(chartsContent)),
	)

	resultsWindow.SetContent(tabs)
	resultsWindow.Show()
}

func (a *App) createComparisonTable(profiles []*models.CarProfile, calculations []*models.CostCalculation) *widget.Table {
	headers := []string{"Kategorie"}
	for _, profile := range profiles {
		headers = append(headers, profile.Name)
	}

	rows := [][]string{
		append([]string{"Monatliche Kraftstoffkosten"}, a.getCalculationValues(calculations, "monthly_fuel")...),
		append([]string{"Monatliche Stromkosten"}, a.getCalculationValues(calculations, "monthly_electric")...),
		append([]string{"Monatliche Gesamtkosten"}, a.getCalculationValues(calculations, "monthly_total")...),
		append([]string{"Jährliche Gesamtkosten"}, a.getCalculationValues(calculations, "annual_total")...),
		append([]string{"Kosten pro Kilometer"}, a.getCalculationValues(calculations, "cost_per_km")...),
		append([]string{"Jährlicher Wertverlust"}, a.getCalculationValues(calculations, "annual_depreciation")...),
		append([]string{"Gesamtkosten der Nutzung"}, a.getCalculationValues(calculations, "total_ownership")...),
	}

	allRows := append([][]string{headers}, rows...)

	table := widget.NewTable(
		func() (int, int) {
			return len(allRows), len(headers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row < len(allRows) && id.Col < len(allRows[id.Row]) {
				label.SetText(allRows[id.Row][id.Col])
				if id.Row == 0 {
					label.TextStyle.Bold = true
				}
			}
		},
	)

	// Set column widths
	table.SetColumnWidth(0, 200)
	for i := 1; i < len(headers); i++ {
		table.SetColumnWidth(i, 150)
	}

	return table
}

func (a *App) getCalculationValues(calculations []*models.CostCalculation, valueType string) []string {
	var values []string
	for _, calc := range calculations {
		var value float64
		switch valueType {
		case "monthly_fuel":
			value = calc.MonthlyFuelCost
		case "monthly_electric":
			value = calc.MonthlyElectricityCost
		case "monthly_total":
			value = calc.MonthlyRunningCosts
		case "annual_total":
			value = calc.AnnualRunningCosts
		case "cost_per_km":
			value = calc.CostPerKilometer
		case "annual_depreciation":
			value = calc.AnnualDepreciation
		case "total_ownership":
			value = calc.TotalCostOfOwnership
		}
		values = append(values, FormatCurrency(value))
	}
	return values
}

func (a *App) createComparisonCharts(profiles []*models.CarProfile, calculations []*models.CostCalculation) *fyne.Container {
	// Simplified chart representation using text
	content := container.NewVBox()

	content.Add(widget.NewCard("Monatliche Kosten Vergleich", "", a.createTextChart(profiles, calculations, "monthly")))
	content.Add(widget.NewCard("Jährliche Kosten Vergleich", "", a.createTextChart(profiles, calculations, "annual")))
	content.Add(widget.NewCard("Kosten pro Kilometer", "", a.createTextChart(profiles, calculations, "per_km")))

	return content
}

func (a *App) createTextChart(profiles []*models.CarProfile, calculations []*models.CostCalculation, chartType string) *fyne.Container {
	content := container.NewVBox()

	var maxValue float64
	var values []float64

	for _, calc := range calculations {
		var value float64
		switch chartType {
		case "monthly":
			value = calc.MonthlyRunningCosts
		case "annual":
			value = calc.AnnualRunningCosts
		case "per_km":
			value = calc.CostPerKilometer
		}
		values = append(values, value)
		if value > maxValue {
			maxValue = value
		}
	}

	for i, profile := range profiles {
		percentage := int((values[i] / maxValue) * 100)
		bar := strings.Repeat("█", percentage/2)
		content.Add(widget.NewLabel(fmt.Sprintf("%s: %s %s",
			profile.Name,
			FormatCurrency(values[i]),
			bar)))
	}

	return content
}

func (a *App) showSettingsDialog() {
	themeSelect := widget.NewSelect([]string{"Hell", "Dunkel"}, nil)
	if a.settings.Theme == "dark" {
		themeSelect.SetSelected("Dunkel")
	} else {
		themeSelect.SetSelected("Hell")
	}

	fuelPriceEntry := widget.NewEntry()
	fuelPriceEntry.SetText(FormatGermanNumber(a.settings.DefaultFuelPrice, 2))

	electricityPriceEntry := widget.NewEntry()
	electricityPriceEntry.SetText(FormatGermanNumber(a.settings.DefaultElectricityPrice, 2))

	form := widget.NewForm(
		widget.NewFormItem("Design", themeSelect),
		widget.NewFormItem("Standard Kraftstoffpreis (€/L)", fuelPriceEntry),
		widget.NewFormItem("Standard Strompreis (€/kWh)", electricityPriceEntry),
	)

	dialog.ShowCustomConfirm("Einstellungen", "Speichern", "Abbrechen", form,
		func(confirmed bool) {
			if confirmed {
				// Update theme
				if themeSelect.Selected == "Dunkel" {
					a.settings.Theme = "dark"
					a.fyneApp.Settings().SetTheme(theme.DarkTheme())
				} else {
					a.settings.Theme = "light"
					a.fyneApp.Settings().SetTheme(theme.LightTheme())
				}

				// Update default prices
				if val, err := ParseGermanNumber(fuelPriceEntry.Text); err == nil {
					a.settings.DefaultFuelPrice = val
				}
				if val, err := ParseGermanNumber(electricityPriceEntry.Text); err == nil {
					a.settings.DefaultElectricityPrice = val
				}

				// Save settings
				err := a.storage.SaveSettings(a.settings)
				if err != nil {
					dialog.ShowError(err, a.window)
				}
			}
		}, a.window)
}
