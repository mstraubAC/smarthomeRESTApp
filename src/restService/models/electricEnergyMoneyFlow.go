package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ElectricEnergyMoneyFlow struct {
	Logdate                      time.Time       `json:"logdate" db:"logdate"`
	UtilitiesBoughInclVat        zeronull.Float8 `json:"utilitiesBoughInclVat" db:"vnbbuyinclvat"`
	UtilitiesSoldInclVat         zeronull.Float8 `json:"utilitiesSoldInclVat" db:"vnbsellinclvat"`
	PvProductionSoldInclVat      zeronull.Float8 `json:"pvProductionSoldInclVat" db:"pvproductionsellinclvat"`
	VatToPayForDirectConsumption zeronull.Float8 `json:"vatToPayForDirectConsumption" db:"vatforpvdirectconsumption"`
	SavedUtilitiesBuy            zeronull.Float8 `json:"savedUtilitiesBuy" db:"savedbypvdirectuse"`
	MoneyFlowOut                 zeronull.Float8 `json:"moneyFlowOut" db:"moneyflowout"`
	MoneyFlowInAndSavings        zeronull.Float8 `json:"moneyFlowInAndSavings" db:"moneyflowinandsavings"`
}
