package weather

func NewWeatherMock() *WeatherMock {
	return &WeatherMock{
		FuncGetTemperature: func(city string) (degrees float64, err error) {
			return
		},
	}
}

type WeatherMock struct {
	FuncGetTemperature func(city string) (degrees float64, err error)

	//Spy
	Call struct {
		// GetTemperature is the number of calls to the GetTemperature method.
		GetTemperature int
		// CurrentParamCity is the last city parameter passed to the GetTemperature method.
		CurrentParamCity string
	}
}

func (w *WeatherMock) GetTemperature(city string) (degrees float64, err error) {
	w.Call.GetTemperature++
	w.Call.CurrentParamCity = city

	degrees, err = w.FuncGetTemperature(city)
	return
}
