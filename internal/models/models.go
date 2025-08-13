package models

import "time"

type FuelType string

const (
	Diesel         FuelType = "Diesel"
	Ultimate       FuelType = "Ultimate"
	Super          FuelType = "Super"
	SuperPlus      FuelType = "SuperPlus"
	UltimateDiesel FuelType = "Ultimate Diesel"
)

type ElectricityType string

const (
	HomeSocket            ElectricityType = "Home socket"
	PublicChargingStation ElectricityType = "Public charging station"
)

type CarProfile struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	FuelConsumption          float64         `json:"fuel_consumption"`     // L/100km
	ElectricConsumption      float64         `json:"electric_consumption"` // kWh/100km
	FuelPrice                float64         `json:"fuel_price"`           // €/L
	ElectricityPrice         float64         `json:"electricity_price"`    // €/kWh
	FuelType                 FuelType        `json:"fuel_type"`
	ElectricityType          ElectricityType `json:"electricity_type"`
	TankSize                 float64         `json:"tank_size"`    // L
	BatterySize              float64         `json:"battery_size"` // kWh
	MonthlyKilometers        float64         `json:"monthly_kilometers"`
	AnnualCarTax             float64         `json:"annual_car_tax"`       // €
	AnnualCarInsurance       float64         `json:"annual_car_insurance"` // €
	FinancingRate            float64         `json:"financing_rate"`       // €/month
	FinancingPeriod          int             `json:"financing_period"`     // months
	PurchasePrice            float64         `json:"purchase_price"`       // €
	ExpectedYearsOfOwnership int             `json:"expected_years_of_ownership"`
	CreatedAt                time.Time       `json:"created_at"`
	UpdatedAt                time.Time       `json:"updated_at"`
}

type CostCalculation struct {
	Profile                *CarProfile `json:"profile"`
	MonthlyFuelCost        float64     `json:"monthly_fuel_cost"`
	AnnualFuelCost         float64     `json:"annual_fuel_cost"`
	MonthlyElectricityCost float64     `json:"monthly_electricity_cost"`
	AnnualElectricityCost  float64     `json:"annual_electricity_cost"`
	MonthlyRunningCosts    float64     `json:"monthly_running_costs"`
	AnnualRunningCosts     float64     `json:"annual_running_costs"`
	TotalDepreciation      float64     `json:"total_depreciation"`
	AnnualDepreciation     float64     `json:"annual_depreciation"`
	CostPerKilometer       float64     `json:"cost_per_kilometer"`
	TotalCostOfOwnership   float64     `json:"total_cost_of_ownership"`
}

type ComparisonResult struct {
	Profiles     []*CarProfile      `json:"profiles"`
	Calculations []*CostCalculation `json:"calculations"`
	CreatedAt    time.Time          `json:"created_at"`
}

type AppSettings struct {
	Theme                   string  `json:"theme"` // "light" or "dark"
	DefaultFuelPrice        float64 `json:"default_fuel_price"`
	DefaultElectricityPrice float64 `json:"default_electricity_price"`
	LastProfilesDir         string  `json:"last_profiles_dir"`
	LastExportDir           string  `json:"last_export_dir"`
}

func GetFuelTypes() []FuelType {
	return []FuelType{Diesel, Ultimate, Super, SuperPlus, UltimateDiesel}
}

func GetElectricityTypes() []ElectricityType {
	return []ElectricityType{HomeSocket, PublicChargingStation}
}

func NewCarProfile() *CarProfile {
	now := time.Now()
	return &CarProfile{
		ID:        generateID(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
