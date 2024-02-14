package weather

// WeatherStub is a WeatherAPI that stubs the GetTemperature method.
func NewWeatherStub() *WeatherStub {
	return &WeatherStub{}
}

type WeatherStub struct {
	FuncGetTemperature func(city string) (degrees float64, err error)
}

func (w *WeatherStub) GetTemperature(city string) (degrees float64, err error) {
	degrees, err = w.FuncGetTemperature(city)
	return
}
