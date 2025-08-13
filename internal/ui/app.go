package ui

import (
	"auto-unterhaltsrechner/internal/calculator"
	"auto-unterhaltsrechner/internal/models"
	"auto-unterhaltsrechner/internal/storage"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	fyneApp    fyne.App
	window     fyne.Window
	calculator *calculator.Calculator
	storage    *storage.Storage
	settings   *models.AppSettings

	// Current profile being edited
	currentProfile *models.CarProfile

	// UI components
	profileSelect *widget.Select
	inputForm     *fyne.Container
	resultsView   *fyne.Container
	mainContent   *container.Split

	// Input widgets
	nameEntry                *widget.Entry
	fuelConsumptionEntry     *widget.Entry
	electricConsumptionEntry *widget.Entry
	fuelPriceEntry           *widget.Entry
	electricityPriceEntry    *widget.Entry
	fuelTypeSelect           *widget.Select
	electricityTypeSelect    *widget.Select
	tankSizeEntry            *widget.Entry
	batterySizeEntry         *widget.Entry
	monthlyKmEntry           *widget.Entry
	annualTaxEntry           *widget.Entry
	annualInsuranceEntry     *widget.Entry
	financingRateEntry       *widget.Entry
	financingPeriodEntry     *widget.Entry
	purchasePriceEntry       *widget.Entry
	ownershipYearsEntry      *widget.Entry
}

func NewApp() *App {
	fyneApp := app.NewWithID("auto-unterhaltsrechner")
	fyneApp.SetIcon(theme.ComputerIcon())

	// Title will be set after settings are loaded
	window := fyneApp.NewWindow("Auto-Unterhaltsrechner v0.1")

	// Set window to reasonable size (75% of common 1920x1080 screen)
	window.Resize(fyne.NewSize(1440, 810))
	window.SetMaster()
	window.CenterOnScreen()

	// Set close handler to properly exit the application
	window.SetCloseIntercept(func() {
		fyneApp.Quit()
	})

	calc := calculator.New()
	store := storage.New()

	settings, _ := store.LoadSettings()
	if settings == nil {
		settings = &models.AppSettings{
			Theme:                   "light",
			Language:                "de",
			DefaultFuelPrice:        1.65,
			DefaultElectricityPrice: 0.35,
		}
	}

	appInstance := &App{
		fyneApp:    fyneApp,
		window:     window,
		calculator: calc,
		storage:    store,
		settings:   settings,
	}

	appInstance.setupUI()
	appInstance.addTooltips()
	appInstance.loadProfiles()

	return appInstance
}

func (a *App) refreshUI() {
	// Recreate the entire UI
	a.setupUI()

	// Restore current profile data
	if a.currentProfile != nil {
		a.updateInputForm()
		a.updateResults()
	}
}

func (a *App) setupUI() {
	// Apply theme
	if a.settings.Theme == "dark" {
		a.fyneApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		a.fyneApp.Settings().SetTheme(theme.LightTheme())
	}

	// Set window title based on language
	translations := a.getCurrentTranslations()
	a.window.SetTitle(translations.AppTitle)

	// Create toolbar
	toolbar := a.createToolbar()

	// Create main content
	a.mainContent = a.createMainContent()

	// Layout
	borderContainer := container.NewBorder(
		toolbar,       // top
		nil,           // bottom
		nil,           // left
		nil,           // right
		a.mainContent, // center
	)

	a.window.SetContent(borderContainer)
}

func (a *App) createToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar()

	toolbar.Append(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		a.newProfile()
	}))

	toolbar.Append(widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		a.saveCurrentProfile()
	}))

	toolbar.Append(widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		a.showLoadDialog()
	}))

	toolbar.Append(widget.NewToolbarSeparator())

	toolbar.Append(widget.NewToolbarAction(theme.DocumentIcon(), func() {
		a.showExportDialog()
	}))

	toolbar.Append(widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
		a.showComparisonDialog()
	}))

	toolbar.Append(widget.NewToolbarSeparator())

	toolbar.Append(widget.NewToolbarAction(theme.SettingsIcon(), func() {
		a.showSettingsDialog()
	}))

	return toolbar
}

func (a *App) createMainContent() *container.Split {
	// Left side - input form
	a.inputForm = a.createInputForm()

	// Right side - results
	a.resultsView = a.createResultsView()

	// Use Border containers to provide stable sizing
	leftPanel := container.NewBorder(nil, nil, nil, nil,
		container.NewScroll(a.inputForm))

	rightPanel := container.NewBorder(nil, nil, nil, nil,
		container.NewScroll(a.resultsView))

	split := container.NewHSplit(leftPanel, rightPanel)
	split.SetOffset(0.5)

	return split
}

func (a *App) Run() {
	a.window.ShowAndRun()
}

func (a *App) newProfile() {
	a.currentProfile = models.NewCarProfile()
	if a.settings.Language == "en" {
		a.currentProfile.Name = "New Profile"
	} else {
		a.currentProfile.Name = "Neues Profil"
	}
	a.currentProfile.FuelPrice = a.settings.DefaultFuelPrice
	a.currentProfile.ElectricityPrice = a.settings.DefaultElectricityPrice
	a.updateInputForm()
	a.updateResults()
}

func (a *App) saveCurrentProfile() {
	if a.currentProfile == nil {
		return
	}

	// Update profile from form
	a.updateProfileFromForm()

	// Validate profile
	errors := a.calculator.ValidateProfile(a.currentProfile)
	if len(errors) > 0 {
		dialog.ShowError(
			fmt.Errorf("Validierungsfehler:\n%s", strings.Join(errors, "\n")),
			a.window,
		)
		return
	}

	// Save profile
	a.currentProfile.UpdatedAt = time.Now()
	err := a.storage.SaveProfile(a.currentProfile)
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	// Reload profiles list
	a.loadProfiles()

	dialog.ShowInformation("Gespeichert", "Profil wurde erfolgreich gespeichert", a.window)
}

func (a *App) loadProfiles() {
	profiles, err := a.storage.ListProfiles()
	if err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	var options []string
	for _, profile := range profiles {
		options = append(options, profile.Name+" ("+profile.ID+")")
	}

	if a.profileSelect != nil {
		a.profileSelect.Options = options
		a.profileSelect.Refresh()
	}
}

// Note: Methods updateInputForm(), updateResults(), and updateProfileFromForm()
// are implemented in input_form.go and results_view.go respectively
