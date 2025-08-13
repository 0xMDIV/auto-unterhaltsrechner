package ui

// Translations for the application
type Translations struct {
	// Window title
	AppTitle string

	// Menu items
	MenuNew        string
	MenuSave       string
	MenuLoad       string
	MenuExport     string
	MenuComparison string
	MenuSettings   string

	// Profile section
	ProfileTitle      string
	ProfileSelect     string
	ProfileName       string
	ConsumptionTitle  string
	PricesTitle       string
	CapacityTitle     string
	UsageTitle        string
	CostsTitle        string
	FinancingTitle    string
	DepreciationTitle string

	// Input fields
	FuelConsumption     string
	ElectricConsumption string
	FuelPrice           string
	ElectricityPrice    string
	FuelType            string
	ElectricityType     string
	TankSize            string
	BatterySize         string
	MonthlyKilometers   string
	AnnualTax           string
	AnnualInsurance     string
	FinancingRate       string
	FinancingPeriod     string
	PurchasePrice       string
	OwnershipYears      string

	// Fuel types
	FuelTypeDiesel         string
	FuelTypeUltimate       string
	FuelTypeSuper          string
	FuelTypeSuperPlus      string
	FuelTypeUltimateDiesel string

	// Electricity types
	ElectricityTypeHome   string
	ElectricityTypePublic string

	// Results sections
	ResultsMonthlyCosts string
	ResultsAnnualCosts  string
	ResultsDepreciation string
	ResultsKeyMetrics   string
	ResultsConsumption  string
	ResultsRange        string

	// Result labels
	FuelCosts             string
	ElectricityCosts      string
	TaxCosts              string
	InsuranceCosts        string
	FinancingCosts        string
	TotalCosts            string
	CostPerKilometer      string
	TotalOwnershipCost    string
	MonthlyFuelAmount     string
	AnnualFuelAmount      string
	MonthlyElectricAmount string
	AnnualElectricAmount  string
	TanksPerMonth         string
	ChargesPerMonth       string
	FuelRange             string
	ElectricRange         string

	// Settings
	SettingsTitle       string
	SettingsTheme       string
	SettingsLanguage    string
	SettingsDefaultFuel string
	SettingsDefaultElec string
	SettingsSave        string
	SettingsCancel      string

	// Theme options
	ThemeLight string
	ThemeDark  string

	// Language options
	LanguageGerman  string
	LanguageEnglish string

	// Export
	ExportCSV         string
	ExportJSON        string
	ExportPDF         string
	ExportTitle       string
	ExportButton      string
	ExportCancel      string
	ExportSuccess     string
	ExportSuccessCSV  string
	ExportSuccessJSON string
	ExportSuccessPDF  string

	// Comparison
	ComparisonTitle       string
	ComparisonSelect      string
	ComparisonButton      string
	ComparisonNotEnough   string
	ComparisonTableTitle  string
	ComparisonChartsTitle string

	// Dialogs
	DialogNoProfile        string
	DialogSelectProfile    string
	DialogCalculationError string
	DialogSaved            string
	DialogLoad             string
	DialogCancel           string

	// Tooltips
	TooltipProfileName         string
	TooltipFuelConsumption     string
	TooltipElectricConsumption string
	TooltipFuelPrice           string
	TooltipElectricityPrice    string
	TooltipFuelType            string
	TooltipElectricityType     string
	TooltipTankSize            string
	TooltipBatterySize         string
	TooltipMonthlyKilometers   string
	TooltipAnnualTax           string
	TooltipAnnualInsurance     string
	TooltipFinancingRate       string
	TooltipFinancingPeriod     string
	TooltipPurchasePrice       string
	TooltipOwnershipYears      string
}

// German translations
var GermanTranslations = Translations{
	AppTitle: "Auto-Unterhaltsrechner v0.2",

	MenuNew:        "Neu",
	MenuSave:       "Speichern",
	MenuLoad:       "Laden",
	MenuExport:     "Export",
	MenuComparison: "Vergleich",
	MenuSettings:   "Einstellungen",

	ProfileTitle:      "Profil",
	ProfileSelect:     "Profil auswählen",
	ProfileName:       "Profilname",
	ConsumptionTitle:  "Verbrauch",
	PricesTitle:       "Preise",
	CapacityTitle:     "Kapazitäten",
	UsageTitle:        "Nutzung",
	CostsTitle:        "Fixkosten",
	FinancingTitle:    "Finanzierung",
	DepreciationTitle: "Wertverlust",

	FuelConsumption:     "Kraftstoffverbrauch (L/100km)",
	ElectricConsumption: "Stromverbrauch (kWh/100km)",
	FuelPrice:           "Kraftstoffpreis (€/L)",
	ElectricityPrice:    "Strompreis (€/kWh)",
	FuelType:            "Kraftstoffart",
	ElectricityType:     "Stromart",
	TankSize:            "Tankgröße (L)",
	BatterySize:         "Batteriegröße (kWh)",
	MonthlyKilometers:   "Monatliche Kilometer",
	AnnualTax:           "Jährliche KFZ-Steuer (€)",
	AnnualInsurance:     "Jährliche Versicherung (€)",
	FinancingRate:       "Finanzierungsrate (€/Monat)",
	FinancingPeriod:     "Finanzierungslaufzeit (Monate)",
	PurchasePrice:       "Kaufpreis (€)",
	OwnershipYears:      "Erwartete Besitzdauer (Jahre)",

	FuelTypeDiesel:         "Diesel",
	FuelTypeUltimate:       "Ultimate",
	FuelTypeSuper:          "Super",
	FuelTypeSuperPlus:      "Super Plus",
	FuelTypeUltimateDiesel: "Ultimate Diesel",

	ElectricityTypeHome:   "Haushaltsstrom",
	ElectricityTypePublic: "Öffentliche Ladestation",

	ResultsMonthlyCosts: "Monatliche Kosten",
	ResultsAnnualCosts:  "Jährliche Kosten",
	ResultsDepreciation: "Wertverlust",
	ResultsKeyMetrics:   "Kennzahlen",
	ResultsConsumption:  "Verbrauch",
	ResultsRange:        "Reichweite",

	FuelCosts:             "Kraftstoffkosten: ",
	ElectricityCosts:      "Stromkosten: ",
	TaxCosts:              "KFZ-Steuer: ",
	InsuranceCosts:        "Versicherung: ",
	FinancingCosts:        "Finanzierung: ",
	TotalCosts:            "Gesamt: ",
	CostPerKilometer:      "Kosten pro Kilometer: ",
	TotalOwnershipCost:    "Gesamtkosten der Nutzung: ",
	MonthlyFuelAmount:     "Monatlicher Kraftstoffverbrauch: ",
	AnnualFuelAmount:      "Jährlicher Kraftstoffverbrauch: ",
	MonthlyElectricAmount: "Monatlicher Stromverbrauch: ",
	AnnualElectricAmount:  "Jährlicher Stromverbrauch: ",
	TanksPerMonth:         "Tankfüllungen pro Monat: ",
	ChargesPerMonth:       "Ladevorgänge pro Monat: ",
	FuelRange:             "Reichweite mit vollem Tank: ",
	ElectricRange:         "Elektrische Reichweite: ",

	SettingsTitle:       "Einstellungen",
	SettingsTheme:       "Design",
	SettingsLanguage:    "Sprache",
	SettingsDefaultFuel: "Standard Kraftstoffpreis (€/L)",
	SettingsDefaultElec: "Standard Strompreis (€/kWh)",
	SettingsSave:        "Speichern",
	SettingsCancel:      "Abbrechen",

	ThemeLight: "Hell",
	ThemeDark:  "Dunkel",

	LanguageGerman:  "Deutsch",
	LanguageEnglish: "English",

	ExportCSV:         "CSV Export",
	ExportJSON:        "JSON Export",
	ExportPDF:         "PDF Export",
	ExportTitle:       "Export",
	ExportButton:      "Exportieren",
	ExportCancel:      "Abbrechen",
	ExportSuccess:     "Export erfolgreich",
	ExportSuccessCSV:  "Die Daten wurden erfolgreich exportiert.",
	ExportSuccessJSON: "Das Profil wurde erfolgreich exportiert.",
	ExportSuccessPDF:  "Das PDF wurde erfolgreich erstellt.",

	ComparisonTitle:       "Fahrzeugvergleich",
	ComparisonSelect:      "Wählen Sie bis zu 4 Profile für den Vergleich:",
	ComparisonButton:      "Vergleichen",
	ComparisonNotEnough:   "Für einen Vergleich werden mindestens 2 Profile benötigt.",
	ComparisonTableTitle:  "Tabelle",
	ComparisonChartsTitle: "Diagramme",

	DialogNoProfile:        "Kein Profil",
	DialogSelectProfile:    "Bitte wählen Sie zuerst ein Profil aus.",
	DialogCalculationError: "Berechnungsfehler",
	DialogSaved:            "Gespeichert",
	DialogLoad:             "Laden",
	DialogCancel:           "Abbrechen",

	TooltipProfileName:         "Eindeutiger Name für dieses Fahrzeugprofil",
	TooltipFuelConsumption:     "Durchschnittlicher Kraftstoffverbrauch des Fahrzeugs in Litern pro 100 Kilometer.",
	TooltipElectricConsumption: "Durchschnittlicher Stromverbrauch des Elektro-/Hybridfahrzeugs in kWh pro 100 Kilometer.",
	TooltipFuelPrice:           "Aktueller Preis für den gewählten Kraftstoff in Euro pro Liter.",
	TooltipElectricityPrice:    "Preis für Strom in Euro pro kWh.",
	TooltipFuelType:            "Art des verwendeten Kraftstoffs.",
	TooltipElectricityType:     "Art der Stromversorgung.",
	TooltipTankSize:            "Volumen des Kraftstofftanks in Litern.",
	TooltipBatterySize:         "Kapazität der Fahrzeugbatterie in kWh.",
	TooltipMonthlyKilometers:   "Durchschnittlich gefahrene Kilometer pro Monat.",
	TooltipAnnualTax:           "Jährliche KFZ-Steuer in Euro.",
	TooltipAnnualInsurance:     "Jährliche Kosten für die Fahrzeugversicherung in Euro.",
	TooltipFinancingRate:       "Monatliche Rate für Finanzierung oder Leasing in Euro.",
	TooltipFinancingPeriod:     "Laufzeit der Finanzierung oder des Leasings in Monaten.",
	TooltipPurchasePrice:       "Kaufpreis des Fahrzeugs in Euro.",
	TooltipOwnershipYears:      "Geplante Besitzdauer des Fahrzeugs in Jahren.",
}

// English translations
var EnglishTranslations = Translations{
	AppTitle: "Auto Maintenance Calculator v0.2",

	MenuNew:        "New",
	MenuSave:       "Save",
	MenuLoad:       "Load",
	MenuExport:     "Export",
	MenuComparison: "Comparison",
	MenuSettings:   "Settings",

	ProfileTitle:      "Profile",
	ProfileSelect:     "Select Profile",
	ProfileName:       "Profile Name",
	ConsumptionTitle:  "Consumption",
	PricesTitle:       "Prices",
	CapacityTitle:     "Capacities",
	UsageTitle:        "Usage",
	CostsTitle:        "Fixed Costs",
	FinancingTitle:    "Financing",
	DepreciationTitle: "Depreciation",

	FuelConsumption:     "Fuel Consumption (L/100km)",
	ElectricConsumption: "Electric Consumption (kWh/100km)",
	FuelPrice:           "Fuel Price (€/L)",
	ElectricityPrice:    "Electricity Price (€/kWh)",
	FuelType:            "Fuel Type",
	ElectricityType:     "Electricity Type",
	TankSize:            "Tank Size (L)",
	BatterySize:         "Battery Size (kWh)",
	MonthlyKilometers:   "Monthly Kilometers",
	AnnualTax:           "Annual Vehicle Tax (€)",
	AnnualInsurance:     "Annual Insurance (€)",
	FinancingRate:       "Financing Rate (€/Month)",
	FinancingPeriod:     "Financing Period (Months)",
	PurchasePrice:       "Purchase Price (€)",
	OwnershipYears:      "Expected Ownership Years",

	FuelTypeDiesel:         "Diesel",
	FuelTypeUltimate:       "Ultimate",
	FuelTypeSuper:          "Super",
	FuelTypeSuperPlus:      "Super Plus",
	FuelTypeUltimateDiesel: "Ultimate Diesel",

	ElectricityTypeHome:   "Home Electricity",
	ElectricityTypePublic: "Public Charging Station",

	ResultsMonthlyCosts: "Monthly Costs",
	ResultsAnnualCosts:  "Annual Costs",
	ResultsDepreciation: "Depreciation",
	ResultsKeyMetrics:   "Key Metrics",
	ResultsConsumption:  "Consumption",
	ResultsRange:        "Range",

	FuelCosts:             "Fuel costs: ",
	ElectricityCosts:      "Electricity costs: ",
	TaxCosts:              "Vehicle tax: ",
	InsuranceCosts:        "Insurance: ",
	FinancingCosts:        "Financing: ",
	TotalCosts:            "Total: ",
	CostPerKilometer:      "Cost per kilometer: ",
	TotalOwnershipCost:    "Total cost of ownership: ",
	MonthlyFuelAmount:     "Monthly fuel consumption: ",
	AnnualFuelAmount:      "Annual fuel consumption: ",
	MonthlyElectricAmount: "Monthly electricity consumption: ",
	AnnualElectricAmount:  "Annual electricity consumption: ",
	TanksPerMonth:         "Tank fills per month: ",
	ChargesPerMonth:       "Charging sessions per month: ",
	FuelRange:             "Range with full tank: ",
	ElectricRange:         "Electric range: ",

	SettingsTitle:       "Settings",
	SettingsTheme:       "Theme",
	SettingsLanguage:    "Language",
	SettingsDefaultFuel: "Default Fuel Price (€/L)",
	SettingsDefaultElec: "Default Electricity Price (€/kWh)",
	SettingsSave:        "Save",
	SettingsCancel:      "Cancel",

	ThemeLight: "Light",
	ThemeDark:  "Dark",

	LanguageGerman:  "Deutsch",
	LanguageEnglish: "English",

	ExportCSV:         "CSV Export",
	ExportJSON:        "JSON Export",
	ExportPDF:         "PDF Export",
	ExportTitle:       "Export",
	ExportButton:      "Export",
	ExportCancel:      "Cancel",
	ExportSuccess:     "Export successful",
	ExportSuccessCSV:  "Data was exported successfully.",
	ExportSuccessJSON: "Profile was exported successfully.",
	ExportSuccessPDF:  "PDF was created successfully.",

	ComparisonTitle:       "Vehicle Comparison",
	ComparisonSelect:      "Select up to 4 profiles for comparison:",
	ComparisonButton:      "Compare",
	ComparisonNotEnough:   "At least 2 profiles are required for comparison.",
	ComparisonTableTitle:  "Table",
	ComparisonChartsTitle: "Charts",

	DialogNoProfile:        "No Profile",
	DialogSelectProfile:    "Please select a profile first.",
	DialogCalculationError: "Calculation Error",
	DialogSaved:            "Saved",
	DialogLoad:             "Load",
	DialogCancel:           "Cancel",

	TooltipProfileName:         "Unique name for this vehicle profile",
	TooltipFuelConsumption:     "Average fuel consumption of the vehicle in liters per 100 kilometers.",
	TooltipElectricConsumption: "Average electricity consumption of the electric/hybrid vehicle in kWh per 100 kilometers.",
	TooltipFuelPrice:           "Current price for the selected fuel in euros per liter.",
	TooltipElectricityPrice:    "Price for electricity in euros per kWh.",
	TooltipFuelType:            "Type of fuel used.",
	TooltipElectricityType:     "Type of electricity supply.",
	TooltipTankSize:            "Volume of the fuel tank in liters.",
	TooltipBatterySize:         "Capacity of the vehicle battery in kWh.",
	TooltipMonthlyKilometers:   "Average kilometers driven per month.",
	TooltipAnnualTax:           "Annual vehicle tax in euros.",
	TooltipAnnualInsurance:     "Annual cost for vehicle insurance in euros.",
	TooltipFinancingRate:       "Monthly rate for financing or leasing in euros.",
	TooltipFinancingPeriod:     "Duration of financing or leasing in months.",
	TooltipPurchasePrice:       "Purchase price of the vehicle in euros.",
	TooltipOwnershipYears:      "Planned ownership duration of the vehicle in years.",
}

func (a *App) getCurrentTranslations() Translations {
	if a.settings.Language == "en" {
		return EnglishTranslations
	}
	return GermanTranslations
}

// Helper functions to translate enum values
func (a *App) translateFuelType(fuelType string) string {
	translations := a.getCurrentTranslations()
	switch fuelType {
	case "diesel":
		return translations.FuelTypeDiesel
	case "ultimate":
		return translations.FuelTypeUltimate
	case "super":
		return translations.FuelTypeSuper
	case "super_plus":
		return translations.FuelTypeSuperPlus
	case "ultimate_diesel":
		return translations.FuelTypeUltimateDiesel
	default:
		return fuelType
	}
}

func (a *App) translateElectricityType(electricityType string) string {
	translations := a.getCurrentTranslations()
	switch electricityType {
	case "home_socket":
		return translations.ElectricityTypeHome
	case "public_charging_station":
		return translations.ElectricityTypePublic
	default:
		return electricityType
	}
}

func (a *App) getTranslatedFuelTypes() []string {
	translations := a.getCurrentTranslations()
	return []string{
		translations.FuelTypeDiesel,
		translations.FuelTypeUltimate,
		translations.FuelTypeSuper,
		translations.FuelTypeSuperPlus,
		translations.FuelTypeUltimateDiesel,
	}
}

func (a *App) getTranslatedElectricityTypes() []string {
	translations := a.getCurrentTranslations()
	return []string{
		translations.ElectricityTypeHome,
		translations.ElectricityTypePublic,
	}
}

func (a *App) getFuelTypeFromTranslation(translation string) string {
	translations := a.getCurrentTranslations()
	switch translation {
	case translations.FuelTypeDiesel:
		return "diesel"
	case translations.FuelTypeUltimate:
		return "ultimate"
	case translations.FuelTypeSuper:
		return "super"
	case translations.FuelTypeSuperPlus:
		return "super_plus"
	case translations.FuelTypeUltimateDiesel:
		return "ultimate_diesel"
	default:
		return translation
	}
}

func (a *App) getElectricityTypeFromTranslation(translation string) string {
	translations := a.getCurrentTranslations()
	switch translation {
	case translations.ElectricityTypeHome:
		return "home_socket"
	case translations.ElectricityTypePublic:
		return "public_charging_station"
	default:
		return translation
	}
}
