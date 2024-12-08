/*
The United States uses the imperial system of weights and measures, which means that
there are many different, seemingly arbitrary units to measure distance. There are
12 inches in a foot, 3 feet in a yard, 22 yards in a chain, and so on.

Create a data structure that can efficiently convert a certain quantity of one unit
to the correct amount of any other unit. You should also allow for additional units
to be added to the system.
* */

package usunitconverter

import (
	"errors"
	"testing"
)

// UnitCoverter stores conversion factor relative to a base unit
type UnitCoverter struct {
	conversionFactors map[string]float64
	baseUnit          string
}

// NewUnitConverter creates a new UnitConverter
func NewUnitConverter(baseUnit string) *UnitCoverter {
	return &UnitCoverter{
		conversionFactors: map[string]float64{baseUnit: 1.0},
		baseUnit:          baseUnit,
	}
}

// AddUnit adds a new unit and its conversion factor relative to the base unit
func (uc *UnitCoverter) AddUnit(unit string, toBaseFactor float64) {
	uc.conversionFactors[unit] = toBaseFactor
}

func (uc *UnitCoverter) Convert(quantity float64, from, to string) (float64, error) {
	framFactor, fromExists := uc.conversionFactors[from]
	toFactor, toExists := uc.conversionFactors[to]

	if !fromExists {
		return 0, errors.New("unknown unit " + from)
	}

	if !toExists {
		return 0, errors.New("unknown unit " + to)
	}

	// convert quantity to base unit then target unit
	quantityInBase := quantity * framFactor
	convertedQuanitity := quantityInBase / toFactor

	return convertedQuanitity, nil
}

func TestMain(t *testing.T) {
	uc := NewUnitConverter("inch")

	uc.AddUnit("foot", 12.0)
	uc.AddUnit("yard", 36.0)
	uc.AddUnit("chain", 792.0)

	quantity, err := uc.Convert(10, "foot", "inch")
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	} else {
		t.Logf("10 feet is %.2f inches\n", quantity)
	}

	quantity, err = uc.Convert(15, "yard", "foot")
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	} else {
		t.Logf("15 yard is %.2f feet\n", quantity)
	}
	quantity, err = uc.Convert(5, "chain", "yard")
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	} else {
		t.Logf("10 chain is %.2f yard\n", quantity)
	}
}
