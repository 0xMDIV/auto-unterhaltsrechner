package calculator

import (
	"auto-unterhaltsrechner/internal/models"
	"math"
)

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) CalculateCosts(profile *models.CarProfile) *models.CostCalculation {
	if profile == nil {
		return nil
	}

	calc := &models.CostCalculation{
		Profile: profile,
	}

	// Calculate fuel costs
	calc.MonthlyFuelCost = c.calculateMonthlyFuelCost(profile)
	calc.AnnualFuelCost = calc.MonthlyFuelCost * 12

	// Calculate electricity costs
	calc.MonthlyElectricityCost = c.calculateMonthlyElectricityCost(profile)
	calc.AnnualElectricityCost = calc.MonthlyElectricityCost * 12

	// Calculate total running costs
	monthlyTax := profile.AnnualCarTax / 12
	monthlyInsurance := profile.AnnualCarInsurance / 12
	calc.MonthlyRunningCosts = calc.MonthlyFuelCost + calc.MonthlyElectricityCost +
		monthlyTax + monthlyInsurance + profile.FinancingRate
	calc.AnnualRunningCosts = calc.MonthlyRunningCosts * 12

	// Calculate depreciation
	calc.TotalDepreciation = c.calculateDepreciation(profile)
	calc.AnnualDepreciation = calc.TotalDepreciation / float64(profile.ExpectedYearsOfOwnership)

	// Calculate cost per kilometer
	annualKm := profile.MonthlyKilometers * 12
	if annualKm > 0 {
		calc.CostPerKilometer = (calc.AnnualRunningCosts + calc.AnnualDepreciation) / annualKm
	}

	// Calculate total cost of ownership
	totalYears := float64(profile.ExpectedYearsOfOwnership)
	calc.TotalCostOfOwnership = (calc.AnnualRunningCosts * totalYears) + calc.TotalDepreciation

	return calc
}

func (c *Calculator) calculateMonthlyFuelCost(profile *models.CarProfile) float64 {
	if profile.FuelConsumption <= 0 || profile.MonthlyKilometers <= 0 {
		return 0
	}

	// Fuel consumption per 100km * monthly km / 100 * fuel price
	return (profile.FuelConsumption * profile.MonthlyKilometers / 100) * profile.FuelPrice
}

func (c *Calculator) calculateMonthlyElectricityCost(profile *models.CarProfile) float64 {
	if profile.ElectricConsumption <= 0 || profile.MonthlyKilometers <= 0 {
		return 0
	}

	// Electric consumption per 100km * monthly km / 100 * electricity price
	return (profile.ElectricConsumption * profile.MonthlyKilometers / 100) * profile.ElectricityPrice
}

func (c *Calculator) calculateDepreciation(profile *models.CarProfile) float64 {
	if profile.PurchasePrice <= 0 || profile.ExpectedYearsOfOwnership <= 0 {
		return 0
	}

	// Simple linear depreciation - assumes 20% residual value after ownership period
	residualValuePercentage := 0.20

	// For cars older than 10 years, depreciation slows down
	if profile.ExpectedYearsOfOwnership > 10 {
		residualValuePercentage = 0.10
	}

	residualValue := profile.PurchasePrice * residualValuePercentage
	return profile.PurchasePrice - residualValue
}

func (c *Calculator) CalculateBreakEven(electricProfile, combustionProfile *models.CarProfile) *BreakEvenAnalysis {
	if electricProfile == nil || combustionProfile == nil {
		return nil
	}

	electricCalc := c.CalculateCosts(electricProfile)
	combustionCalc := c.CalculateCosts(combustionProfile)

	analysis := &BreakEvenAnalysis{
		ElectricProfile:   electricProfile,
		CombustionProfile: combustionProfile,
		ElectricCosts:     electricCalc,
		CombustionCosts:   combustionCalc,
	}

	// Calculate break-even point
	priceDifference := electricProfile.PurchasePrice - combustionProfile.PurchasePrice
	monthlySavings := combustionCalc.MonthlyRunningCosts - electricCalc.MonthlyRunningCosts

	if monthlySavings > 0 {
		analysis.BreakEvenMonths = int(math.Ceil(priceDifference / monthlySavings))
		analysis.BreakEvenKilometers = float64(analysis.BreakEvenMonths) * electricProfile.MonthlyKilometers
	} else {
		analysis.BreakEvenMonths = -1 // Never breaks even
	}

	// Calculate total savings over ownership period
	totalMonths := float64(electricProfile.ExpectedYearsOfOwnership * 12)
	if monthlySavings > 0 {
		analysis.TotalSavings = (monthlySavings * totalMonths) - priceDifference
	} else {
		analysis.TotalSavings = (monthlySavings * totalMonths) - priceDifference
	}

	return analysis
}

type BreakEvenAnalysis struct {
	ElectricProfile     *models.CarProfile      `json:"electric_profile"`
	CombustionProfile   *models.CarProfile      `json:"combustion_profile"`
	ElectricCosts       *models.CostCalculation `json:"electric_costs"`
	CombustionCosts     *models.CostCalculation `json:"combustion_costs"`
	BreakEvenMonths     int                     `json:"break_even_months"`
	BreakEvenKilometers float64                 `json:"break_even_kilometers"`
	TotalSavings        float64                 `json:"total_savings"`
}

func (c *Calculator) ValidateProfile(profile *models.CarProfile) []string {
	var errors []string

	if profile.Name == "" {
		errors = append(errors, "Profilname ist erforderlich")
	}

	if profile.MonthlyKilometers < 0 {
		errors = append(errors, "Monatliche Kilometer müssen >= 0 sein")
	}

	if profile.FuelConsumption < 0 {
		errors = append(errors, "Kraftstoffverbrauch muss >= 0 sein")
	}

	if profile.ElectricConsumption < 0 {
		errors = append(errors, "Stromverbrauch muss >= 0 sein")
	}

	if profile.FuelPrice < 0 {
		errors = append(errors, "Kraftstoffpreis muss >= 0 sein")
	}

	if profile.ElectricityPrice < 0 {
		errors = append(errors, "Strompreis muss >= 0 sein")
	}

	if profile.TankSize < 0 {
		errors = append(errors, "Tankgröße muss >= 0 sein")
	}

	if profile.BatterySize < 0 {
		errors = append(errors, "Batteriegröße muss >= 0 sein")
	}

	if profile.AnnualCarTax < 0 {
		errors = append(errors, "Jährliche KFZ-Steuer muss >= 0 sein")
	}

	if profile.AnnualCarInsurance < 0 {
		errors = append(errors, "Jährliche Versicherung muss >= 0 sein")
	}

	if profile.FinancingRate < 0 {
		errors = append(errors, "Finanzierungsrate muss >= 0 sein")
	}

	if profile.FinancingPeriod < 0 {
		errors = append(errors, "Finanzierungslaufzeit muss >= 0 sein")
	}

	if profile.PurchasePrice < 0 {
		errors = append(errors, "Kaufpreis muss >= 0 sein")
	}

	if profile.ExpectedYearsOfOwnership <= 0 {
		errors = append(errors, "Erwartete Besitzdauer muss > 0 sein")
	}

	return errors
}
