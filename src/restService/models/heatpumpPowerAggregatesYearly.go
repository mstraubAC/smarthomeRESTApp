package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type HeatpumpPowerAggregatesYearlyType struct {
	Logdate                                  time.Time       `json:"logdate" db:"logdate"`
	YearlyWorkCoefficient                    zeronull.Float8 `json:"yearlyWorkCoefficient" db:"jahresarbeitszahlinclcontrolandpumps"`
	YearlyWorkCoefficientHeatSourceOnly      zeronull.Float8 `json:"yearlyWorkCoefficientHeatSourceOnly" db:"jahresarbeitszahlsolvis"`
	YearlyWorkCoefficientIncludingControl    zeronull.Float8 `json:"yearlyWorkCoefficientIncludingControl" db:"jahresarbeitszahlfullelectricmeasurement"`
	TotalElectricEnergyInkWh                 zeronull.Float8 `json:"totalElectricEnergyInkWh" db:"totalelectricenergy"`
	TotalElectricEnergyIncludingControlInkWh zeronull.Float8 `json:"totalElectricEnergyIncludingControlInkWh" db:"totalelectricenergyfullmeasurement"`
	TotalElectricEnergyHeatSourceOnlyInkWh   zeronull.Float8 `json:"totalElectricEnergyHeatSourceOnlyInkWh" db:"totalelectricenergysolvismeasurement"`
	TotalThermalEnergyInkWh                  zeronull.Float8 `json:"totalThermalEnergyInkWh" db:"totalheatingenergy"`
	HeatpumpThermalEnergyInkWh               zeronull.Float8 `json:"heatpumpThermalEnergyInkWh" db:"heatpumpthermalpowerenergy"`
	ResistiveHeatingThermalEnergyInkWh       zeronull.Float8 `json:"resistiveHeatingThermalEnergyInkWh" db:"heatpumpresistanceheatingenergy"`
	AverageOutsideTemperaturInCelsius        zeronull.Float8 `json:"averageOutsideTemperaturInCelsius" db:"outsidetemperatureavg"`
	AverageHeatingFlowTemperatureInCelsius   zeronull.Float8 `json:"averageHeatingFlowTemperatureInCelsius" db:"flowtemperaturecircuit1avg"`
}
