package entity

type Unit struct {
	Name       UnitName
	Symbol     string
	Quantity   QuantityType
	System     string
	PluralName string
	Names      []string
}

func (u *Unit) Validate() bool {
	return u.validateName()
}

func (u *Unit) validateName() bool {
	if !u.isUnitSupported() {
		return false
	}

	switch u.Quantity {
	case QuantityMass:
		return isOneOf(u.Name, massUnits)
	case QuantityLength:
		return isOneOf(u.Name, lengthUnits)
	case QuantityVolume:
		return isOneOf(u.Name, volummeUnits)
	case QuantityTemperature:
		return isOneOf(u.Name, temperatureUnits)
	case QuantityCount:
		return isOneOf(u.Name, countUnits)
	}
	return false
}

func (u *Unit) isUnitSupported() bool {
	for _, unit := range SupportedUnits {
		if u.Name == unit {
			return true
		}
	}
	return false
}

func isOneOf(unit UnitName, units []UnitName) bool {
	for _, u := range units {
		if u == unit {
			return true
		}
	}
	return false
}

type UnitName string

var (
	// mass
	UnitGram      UnitName = "gram"
	UnitMilligram UnitName = "milligram"
	UnitKilogram  UnitName = "kilogram"
	UnitDecagram  UnitName = "decagram"
	UnitPound     UnitName = "pound"
	UnitOunce     UnitName = "ounce"
	// length
	UnitMeter      UnitName = "meter"
	UnitCentimeter UnitName = "centimeter"
	UnitDecimeter  UnitName = "decimeter"
	UnitFoot       UnitName = "foot"
	UnitInch       UnitName = "inch"
	UnitKilometer  UnitName = "kilometer"
	UnitMile       UnitName = "mile"
	UnitYard       UnitName = "yard"
	// volume
	UnitGallon     UnitName = "gallon"
	UnitHectoliter UnitName = "hectoliter"
	UnitLiter      UnitName = "liter"
	UnitMilliliter UnitName = "milliliter"
	UnitPint       UnitName = "pint"
	// temperature
	UnitCelsius    UnitName = "celsius"
	UnitFahrenheit UnitName = "fahrenheit"
	UnitKelvin     UnitName = "kelvin"
	// time
	UnitCentury     UnitName = "century"
	UnitDay         UnitName = "day"
	UnitDecade      UnitName = "decade"
	UnitHour        UnitName = "hour"
	UnitMillisecond UnitName = "millisecond"
	UnitMinute      UnitName = "minute"
	UnitMonth       UnitName = "month"
	UnitYear        UnitName = "year"
	// count
	UnitCount UnitName = "count"
)

var SupportedUnits = []UnitName{
	UnitGram, UnitMilligram, UnitKilogram, UnitDecagram, UnitPound, UnitOunce, // mass
	UnitMeter, UnitCentimeter, UnitDecimeter, UnitFoot, UnitInch, UnitKilometer, UnitMile, UnitYard, // length
	UnitGallon, UnitHectoliter, UnitLiter, UnitMilliliter, UnitPint, // volume
	UnitCelsius, UnitFahrenheit, UnitKelvin, // temperature
	UnitCentury, UnitDay, UnitDecade, UnitHour, UnitMillisecond, UnitMinute, UnitMonth, UnitYear, // time
	UnitCount, // count
}

var (
	massUnits        = []UnitName{UnitGram, UnitMilligram, UnitKilogram, UnitDecagram, UnitPound, UnitOunce}
	lengthUnits      = []UnitName{UnitMeter, UnitCentimeter, UnitDecimeter, UnitFoot, UnitInch, UnitKilometer, UnitMile, UnitYard}
	volummeUnits     = []UnitName{UnitGallon, UnitHectoliter, UnitLiter, UnitMilliliter, UnitPint}
	temperatureUnits = []UnitName{UnitCentury, UnitDay, UnitDecade, UnitHour, UnitMillisecond, UnitMinute, UnitMonth, UnitYear}
	countUnits       = []UnitName{UnitCount}
)
