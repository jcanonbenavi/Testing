package service_test

import (
	"doubles/internal/city"
	"doubles/internal/city/repository"
	"doubles/internal/city/service"
	"doubles/internal/logger"
	"doubles/internal/weather"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDefaultAddCity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		rp := repository.NewStub()
		rp.FuncSaveCity = func(c *city.City) (err error) {
			c.ID = 1
			return
		}

		weather_stub := weather.NewWeatherStub()
		weather_stub.FuncGetTemperature = func(city string) (degrees float64, err error) {
			degrees = 10
			return
		}
		logger := logger.NewDummy()
		service := service.NewDefault(rp, weather_stub, logger)

		// act
		c, err := service.AddCity("Paris", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expected := city.City{
			ID:          1,
			Name:        "Paris",
			Country:     "France",
			Population:  2_200_000,
			Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 10,
		}
		require.Equal(t, expected, c)
		require.Nil(t, err)
	})

	t.Run("error: invalid City", func(t *testing.T) {

		rp := repository.NewStub()
		rp.FuncSaveCity = func(c *city.City) (err error) {
			c.ID = 1
			return
		}
		//arrange
		weather_stub := weather.NewWeatherStub()
		weather_stub.FuncGetTemperature = func(city string) (degrees float64, err error) {
			degrees = 10
			return
		}
		logger := logger.NewDummy()
		sv := service.NewDefault(rp, weather_stub, logger)

		// act
		_, err := sv.AddCity("Paris", "", 0, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		expectedError := service.ErrCityInvalid
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("success 02: add other city", func(t *testing.T) {
		//arrange
		rp := repository.NewStub()
		rp.FuncSaveCity = func(c *city.City) (err error) {
			c.ID = 1
			return
		}

		weather_mock := weather.NewWeatherMock()
		weather_mock.FuncGetTemperature = func(city string) (degrees float64, err error) {
			degrees = 10
			return
		}
		logger := logger.NewDummy()
		service := service.NewDefault(rp, weather_mock, logger)

		// act
		c, err := service.AddCity("Paris", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expected := city.City{
			ID:          1,
			Name:        "Paris",
			Country:     "France",
			Population:  2_200_000,
			Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 10,
		}
		require.Equal(t, expected, c)
		require.NoError(t, err)
		require.Equal(t, 1, weather_mock.Call.GetTemperature)
		require.Equal(t, "Paris", weather_mock.Call.CurrentParamCity)
	})
}

func TestDefaultAddCity_Testify(t *testing.T) {
	t.Run("success - add city", func(t *testing.T) {
		lg := logger.NewDummy()
		wa := weather.NewWeatherStub()
		wa.FuncGetTemperature = func(city string) (degrees float64, err error) {
			degrees = 10
			return
		}
		rp := repository.NewMock()
		// SaveCity is the method of the mock
		//city.City is the input of the method
		rp.On("SaveCity", &city.City{
			ID:          0,
			Name:        "Paris",
			Country:     "France",
			Population:  2_200_000,
			Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 10,
		}).Return(nil) // return nil error
		rp.FuncSaveCity = func(c *city.City) {
			c.ID = 1
		}
		// - service
		sv := service.NewDefault(rp, wa, lg)

		// act
		c, err := sv.AddCity("Paris", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expectedCity := city.City{
			ID:          1,
			Name:        "Paris",
			Country:     "France",
			Population:  2_200_000,
			Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 10,
		}
		require.Equal(t, expectedCity, c)
		require.NoError(t, err)
		rp.AssertExpectations(t)
	})
}
