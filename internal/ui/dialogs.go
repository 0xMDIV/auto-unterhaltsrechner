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
	"fyne.io/fyne/v2/storage"
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
		dialog.ShowError(fmt.Errorf("berechnungsfehler"), a.window)
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
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
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

	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
	saveDialog.SetFileName("auto-unterhaltsrechner-export.csv")
	saveDialog.Show()
}

func (a *App) exportToJSON(profile *models.CarProfile) {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
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

	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	saveDialog.SetFileName("auto-unterhaltsrechner-profil.json")
	saveDialog.Show()
}

func (a *App) exportToPDF(calculation *models.CostCalculation) {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if writer == nil {
			return
		}
		defer writer.Close()

		// Get translations for PDF
		translations := a.getCurrentTranslations()

		// Create PDF with enhanced layout
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		// Define colors
		headerColor := []int{52, 73, 94}     // Dark blue-gray
		sectionColor := []int{236, 240, 241} // Light gray
		totalColor := []int{231, 76, 60}     // Red for totals

		// Header with background color
		pdf.SetFillColor(headerColor[0], headerColor[1], headerColor[2])
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFont("Arial", "B", 18)
		pdf.CellFormat(0, 15, translations.AppTitle, "0", 1, "C", true, 0, "")
		pdf.SetTextColor(0, 0, 0) // Reset text color
		pdf.Ln(5)

		// Subheader
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "Kostenaufstellung", "0", 1, "C", false, 0, "")
		pdf.Ln(10)

		// Profile info in a box
		pdf.SetFillColor(sectionColor[0], sectionColor[1], sectionColor[2])
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, 8, "Fahrzeugprofil", "1", 1, "L", true, 0, "")
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(0, 6, "Name: "+calculation.Profile.Name, "LR", 1, "L", false, 0, "")
		pdf.CellFormat(0, 6, "Erstellt am: "+time.Now().Format("02.01.2006 15:04"), "LRB", 1, "L", false, 0, "")
		pdf.Ln(8)

		// Helper function to create a section with table
		createSection := func(title string, data [][]string, showTotal bool, totalAmount float64) {
			// Section header
			pdf.SetFillColor(sectionColor[0], sectionColor[1], sectionColor[2])
			pdf.SetFont("Arial", "B", 12)
			pdf.CellFormat(0, 8, title, "1", 1, "L", true, 0, "")

			// Table rows
			pdf.SetFont("Arial", "", 10)
			for _, row := range data {
				if len(row) >= 2 {
					pdf.CellFormat(100, 6, row[0], "LR", 0, "L", false, 0, "")
					pdf.CellFormat(0, 6, row[1], "LR", 1, "R", false, 0, "")
				}
			}

			// Total row if needed
			if showTotal {
				pdf.SetFillColor(totalColor[0], totalColor[1], totalColor[2])
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFont("Arial", "B", 10)
				pdf.CellFormat(100, 8, "Gesamt:", "LR", 0, "L", true, 0, "")
				pdf.CellFormat(0, 8, FormatCurrencyPDF(totalAmount), "LR", 1, "R", true, 0, "")
				pdf.SetTextColor(0, 0, 0)
			}

			// Bottom border
			pdf.CellFormat(0, 0, "", "B", 1, "", false, 0, "")
			pdf.Ln(6)
		}

		// Monthly costs table
		monthlyData := [][]string{
			{translations.FuelCosts[:len(translations.FuelCosts)-2], FormatCurrencyPDF(calculation.MonthlyFuelCost)},
			{translations.ElectricityCosts[:len(translations.ElectricityCosts)-2], FormatCurrencyPDF(calculation.MonthlyElectricityCost)},
			{translations.TaxCosts[:len(translations.TaxCosts)-2], FormatCurrencyPDF(calculation.Profile.AnnualCarTax / 12)},
			{translations.InsuranceCosts[:len(translations.InsuranceCosts)-2], FormatCurrencyPDF(calculation.Profile.AnnualCarInsurance / 12)},
			{translations.FinancingCosts[:len(translations.FinancingCosts)-2], FormatCurrencyPDF(calculation.Profile.FinancingRate)},
		}
		createSection(translations.ResultsMonthlyCosts, monthlyData, true, calculation.MonthlyRunningCosts)

		// Annual costs table
		annualData := [][]string{
			{translations.FuelCosts[:len(translations.FuelCosts)-2], FormatCurrencyPDF(calculation.AnnualFuelCost)},
			{translations.ElectricityCosts[:len(translations.ElectricityCosts)-2], FormatCurrencyPDF(calculation.AnnualElectricityCost)},
			{translations.TaxCosts[:len(translations.TaxCosts)-2], FormatCurrencyPDF(calculation.Profile.AnnualCarTax)},
			{translations.InsuranceCosts[:len(translations.InsuranceCosts)-2], FormatCurrencyPDF(calculation.Profile.AnnualCarInsurance)},
			{translations.FinancingCosts[:len(translations.FinancingCosts)-2], FormatCurrencyPDF(calculation.Profile.FinancingRate * 12)},
		}
		createSection(translations.ResultsAnnualCosts, annualData, true, calculation.AnnualRunningCosts)

		// Key metrics table
		metricsData := [][]string{
			{translations.CostPerKilometer[:len(translations.CostPerKilometer)-2], FormatCurrencyPDF(calculation.CostPerKilometer)},
			{"Monatliche Fahrleistung:", FormatKilometers(calculation.Profile.MonthlyKilometers)},
			{translations.TotalOwnershipCost[:len(translations.TotalOwnershipCost)-2], FormatCurrencyPDF(calculation.TotalCostOfOwnership)},
		}
		createSection(translations.ResultsKeyMetrics, metricsData, false, 0)

		// Depreciation table
		depreciationData := [][]string{
			{"Kaufpreis:", FormatCurrencyPDF(calculation.Profile.PurchasePrice)},
			{"Besitzdauer:", fmt.Sprintf("%d Jahre", calculation.Profile.ExpectedYearsOfOwnership)},
			{"Jährlicher Wertverlust:", FormatCurrencyPDF(calculation.AnnualDepreciation)},
		}
		createSection(translations.ResultsDepreciation, depreciationData, false, 0)

		// Consumption section if applicable
		if calculation.Profile.FuelConsumption > 0 || calculation.Profile.ElectricConsumption > 0 {
			var consumptionData [][]string

			if calculation.Profile.FuelConsumption > 0 {
				monthlyFuel := (calculation.Profile.FuelConsumption * calculation.Profile.MonthlyKilometers) / 100
				annualFuel := monthlyFuel * 12
				consumptionData = append(consumptionData, []string{translations.MonthlyFuelAmount[:len(translations.MonthlyFuelAmount)-2], FormatLiters(monthlyFuel)})
				consumptionData = append(consumptionData, []string{translations.AnnualFuelAmount[:len(translations.AnnualFuelAmount)-2], FormatLiters(annualFuel)})

				if calculation.Profile.TankSize > 0 {
					tanksPerMonth := monthlyFuel / calculation.Profile.TankSize
					consumptionData = append(consumptionData, []string{translations.TanksPerMonth[:len(translations.TanksPerMonth)-2], FormatGermanNumber(tanksPerMonth, 1)})
				}
			}

			if calculation.Profile.ElectricConsumption > 0 {
				monthlyElectric := (calculation.Profile.ElectricConsumption * calculation.Profile.MonthlyKilometers) / 100
				annualElectric := monthlyElectric * 12
				consumptionData = append(consumptionData, []string{translations.MonthlyElectricAmount[:len(translations.MonthlyElectricAmount)-2], FormatKWh(monthlyElectric)})
				consumptionData = append(consumptionData, []string{translations.AnnualElectricAmount[:len(translations.AnnualElectricAmount)-2], FormatKWh(annualElectric)})

				if calculation.Profile.BatterySize > 0 {
					chargesPerMonth := monthlyElectric / calculation.Profile.BatterySize
					consumptionData = append(consumptionData, []string{translations.ChargesPerMonth[:len(translations.ChargesPerMonth)-2], FormatGermanNumber(chargesPerMonth, 1)})
				}
			}

			createSection(translations.ResultsConsumption, consumptionData, false, 0)
		}

		// Range information if applicable
		var rangeData [][]string
		if calculation.Profile.TankSize > 0 && calculation.Profile.FuelConsumption > 0 {
			fuelRange := (calculation.Profile.TankSize / calculation.Profile.FuelConsumption) * 100
			rangeData = append(rangeData, []string{translations.FuelRange[:len(translations.FuelRange)-2], FormatKilometers(fuelRange)})
		}
		if calculation.Profile.BatterySize > 0 && calculation.Profile.ElectricConsumption > 0 {
			electricRange := (calculation.Profile.BatterySize / calculation.Profile.ElectricConsumption) * 100
			rangeData = append(rangeData, []string{translations.ElectricRange[:len(translations.ElectricRange)-2], FormatKilometers(electricRange)})
		}
		if len(rangeData) > 0 {
			createSection(translations.ResultsRange, rangeData, false, 0)
		}

		// Footer
		pdf.SetY(-20)
		pdf.SetFont("Arial", "I", 8)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(0, 10, fmt.Sprintf("Erstellt mit %s", translations.AppTitle), "0", 0, "C", false, 0, "")

		// Save PDF
		err = pdf.Output(writer)
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}

		dialog.ShowInformation(translations.ExportSuccess, translations.ExportSuccessPDF, a.window)
	}, a.window)

	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	saveDialog.SetFileName("auto-unterhaltsrechner-kostenaufstellung.pdf")
	saveDialog.Show()
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
	translations := a.getCurrentTranslations()

	themeSelect := widget.NewSelect([]string{translations.ThemeLight, translations.ThemeDark}, nil)
	if a.settings.Theme == "dark" {
		themeSelect.SetSelected(translations.ThemeDark)
	} else {
		themeSelect.SetSelected(translations.ThemeLight)
	}

	languageSelect := widget.NewSelect([]string{translations.LanguageGerman, translations.LanguageEnglish}, nil)
	if a.settings.Language == "en" {
		languageSelect.SetSelected(translations.LanguageEnglish)
	} else {
		languageSelect.SetSelected(translations.LanguageGerman)
	}

	fuelPriceEntry := widget.NewEntry()
	fuelPriceEntry.SetText(FormatGermanNumber(a.settings.DefaultFuelPrice, 2))

	electricityPriceEntry := widget.NewEntry()
	electricityPriceEntry.SetText(FormatGermanNumber(a.settings.DefaultElectricityPrice, 2))

	form := widget.NewForm(
		widget.NewFormItem(translations.SettingsTheme, themeSelect),
		widget.NewFormItem(translations.SettingsLanguage, languageSelect),
		widget.NewFormItem(translations.SettingsDefaultFuel, fuelPriceEntry),
		widget.NewFormItem(translations.SettingsDefaultElec, electricityPriceEntry),
	)

	dialog.ShowCustomConfirm(translations.SettingsTitle, translations.SettingsSave, translations.SettingsCancel, form,
		func(confirmed bool) {
			if confirmed {
				// Update theme
				if themeSelect.Selected == translations.ThemeDark {
					a.settings.Theme = "dark"
					a.fyneApp.Settings().SetTheme(theme.DarkTheme())
				} else {
					a.settings.Theme = "light"
					a.fyneApp.Settings().SetTheme(theme.LightTheme())
				}

				// Update language
				if languageSelect.Selected == translations.LanguageEnglish {
					a.settings.Language = "en"
				} else {
					a.settings.Language = "de"
				}

				// Update window title and UI
				newTranslations := a.getCurrentTranslations()
				a.window.SetTitle(newTranslations.AppTitle)
				a.refreshUI()

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
