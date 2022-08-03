package temperature

type Temperature interface {
	ConvertTo(arg float64) float64
	GetName() string
}

// 1) kelvin to fahrenheit
type KelvinToFahrenheit struct {
	Name string
}

func (kt KelvinToFahrenheit) ConvertTo(arg float64) float64 {
	return 1.8*(arg-273.15) + 32.0
}

func (kt KelvinToFahrenheit) GetName() string {
	return kt.Name
}

// 2) Celcius to fahrenheit
type CelciusToFahrenheit struct {
	Name string
}

func (cf CelciusToFahrenheit) ConvertTo(arg float64) float64 {
	return 1.8*arg + 32
}

func (cf CelciusToFahrenheit) GetName() string {
	return cf.Name
}

// 3 Fahrenheit to kelvin
type FahrenheitToKelvin struct {
	Name string
}

func (fk FahrenheitToKelvin) ConvertTo(arg float64) float64 {
	return (arg-32)/1.8 + 273.15
}

func (fk FahrenheitToKelvin) GetName() string {
	return fk.Name
}

// 4) Celcius to Kelvin
type CelciusToKelvin struct {
	Name string
}

func (fk CelciusToKelvin) ConvertTo(arg float64) float64 {
	return arg + 273.15
}

func (fk CelciusToKelvin) GetName() string {
	return fk.Name
}

// 5) Fahrenheit to Celcius
type FahrenheitToCelcius struct {
	Name string
}

func (fc FahrenheitToCelcius) ConvertTo(arg float64) float64 {
	return (arg - 32) / 1.8
}

func (fc FahrenheitToCelcius) GetName() string {
	return fc.Name
}

// 6) Kelvin to celcius
type KelvinToCelcius struct {
	Name string
}

func (kc KelvinToCelcius) ConvertTo(arg float64) float64 {
	return arg - 273.15
}

func (kc KelvinToCelcius) GetName() string {
	return kc.Name
}
