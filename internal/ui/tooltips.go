package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// createFormItemWithTooltip creates a form item with a tooltip icon
func createFormItemWithTooltip(label, tooltip string, object fyne.CanvasObject) *fyne.Container {
	labelWidget := widget.NewLabel(label)

	infoIcon := widget.NewIcon(theme.InfoIcon())
	infoIcon.Resize(fyne.NewSize(16, 16))

	infoButton := widget.NewButton("", func() {
		widget.ShowPopUp(
			widget.NewCard("Info", "", widget.NewLabel(tooltip)),
			fyne.CurrentApp().Driver().AllWindows()[0].Canvas(),
		)
	})
	infoButton.SetIcon(theme.InfoIcon())
	infoButton.Resize(fyne.NewSize(20, 20))

	headerContainer := container.NewHBox(
		labelWidget,
		infoButton,
	)

	return container.NewBorder(
		headerContainer,
		nil,
		nil,
		nil,
		object,
	)
}

// Tooltip texts in German
const (
	TooltipProfileName = "Eindeutiger Name für dieses Fahrzeugprofil"

	TooltipFuelConsumption = "Durchschnittlicher Kraftstoffverbrauch des Fahrzeugs in Litern pro 100 Kilometer. " +
		"Dieser Wert findet sich meist in den Fahrzeugpapieren oder kann über mehrere Tankfüllungen ermittelt werden."

	TooltipElectricConsumption = "Durchschnittlicher Stromverbrauch des Elektro-/Hybridfahrzeugs in kWh pro 100 Kilometer. " +
		"Dieser Wert wird im Bordcomputer angezeigt oder kann über mehrere Ladevorgänge ermittelt werden."

	TooltipFuelPrice = "Aktueller Preis für den gewählten Kraftstoff in Euro pro Liter. " +
		"Verwenden Sie einen Durchschnittspreis für bessere Kalkulationen."

	TooltipElectricityPrice = "Preis für Strom in Euro pro kWh. Bei Haushalten meist der Arbeitspreis aus der Stromrechnung. " +
		"An öffentlichen Ladestationen variiert der Preis je nach Anbieter."

	TooltipFuelType = "Art des verwendeten Kraftstoffs. Beeinflusst die Berechnung bei unterschiedlichen Preisen."

	TooltipElectricityType = "Art der Stromversorgung. Haushaltsstrom ist meist günstiger als öffentliche Ladestationen."

	TooltipTankSize = "Volumen des Kraftstofftanks in Litern. Wird für die Berechnung der Reichweite verwendet."

	TooltipBatterySize = "Kapazität der Fahrzeugbatterie in kWh. Wird für die Berechnung der elektrischen Reichweite verwendet."

	TooltipMonthlyKilometers = "Durchschnittlich gefahrene Kilometer pro Monat. " +
		"Grundlage für alle Kostenberechnungen. Kann aus dem Jahreskilometerstand ÷ 12 ermittelt werden."

	TooltipAnnualTax = "Jährliche KFZ-Steuer in Euro. Der Betrag steht im Steuerbescheid oder " +
		"kann online beim Kraftfahrt-Bundesamt berechnet werden."

	TooltipAnnualInsurance = "Jährliche Kosten für die Fahrzeugversicherung in Euro. " +
		"Umfasst Haftpflicht, Teil- oder Vollkasko je nach gewähltem Versicherungsschutz."

	TooltipFinancingRate = "Monatliche Rate für Finanzierung oder Leasing in Euro. " +
		"Bei Barkauf 0 eingeben."

	TooltipFinancingPeriod = "Laufzeit der Finanzierung oder des Leasings in Monaten. " +
		"Bei Barkauf 0 eingeben."

	TooltipPurchasePrice = "Kaufpreis des Fahrzeugs in Euro. Wird für die Wertverlustkalkulation verwendet. " +
		"Bei Gebrauchtwagen den tatsächlich gezahlten Preis eingeben."

	TooltipOwnershipYears = "Geplante Besitzdauer des Fahrzeugs in Jahren. " +
		"Beeinflusst die Berechnung des jährlichen Wertverlusts."
)

// Helper function to add tooltips to existing widgets
func (a *App) addTooltips() {
	// This would be called after creating the form to add tooltips to existing widgets
	// Implementation would depend on how tooltips are displayed in the final UI
}
