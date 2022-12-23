package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ElectricEnergyFlow struct {
	Logdate                 time.Time       `json:"logdate" db:"logdate"`
	ElectricGridConsumption zeronull.Float8 `json:"electricGridConsumption" db:"electricconsumevalue"`
	ElectricGridFeedIn      zeronull.Float8 `json:"electricGridFeedIn" db:"electricgridfeedinvalue"`
	PvGeneration            zeronull.Float8 `json:"pvGeneration" db:"electricgridpvgenerationvalue"`
	HeatingConusmption      zeronull.Float8 `json:"heatingConsumption" db:"electricheatingconsumevalue"`
	ItConsumption           zeronull.Float8 `json:"itConsumption" db:"electricitconsumevalue"`
	WallboxConsumption      zeronull.Float8 `json:"wallboxConsumption" db:"electricwallboxconsumevalue"`
}
