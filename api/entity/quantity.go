package entity

import goUnits "github.com/bcicen/go-units"

type Quantity struct {
	Value float64
	Unit  UnitName
}

func (q Quantity) ToUnit(u UnitName) (Quantity, error) {
	if u == q.Unit {
		return Quantity{q.Value, u}, nil
	}
	toUnit, err := goUnits.Find(string(u))
	if err != nil {
		return q, ErrInvalidParameter
	}
	fromUnit, err := goUnits.Find(string(q.Unit))
	if err != nil {
		return q, ErrInvalidEntity
	}
	valueToAdd, err := goUnits.ConvertFloat(q.Value, fromUnit, toUnit)
	if err != nil {
		return q, ErrInvalidParameter
	}
	return Quantity{valueToAdd.Float(), u}, nil
}

func (q Quantity) Add(quantity Quantity) (Quantity, error) {
	valueToAdd, err := quantity.ToUnit(q.Unit)
	if err != nil {
		return q, err
	}
	q.Value += valueToAdd.Value
	return q, nil
}

func (q Quantity) Subtract(quantity Quantity) (Quantity, error) {
	valueToAdd, err := quantity.ToUnit(q.Unit)
	if err != nil {
		return q, err
	}
	q.Value -= valueToAdd.Value
	return q, nil
}

func (q Quantity) Verify() error {
	_, err := goUnits.Find(string(q.Unit))
	if err != nil {
		return ErrInvalidEntity
	}
	return nil
}
