package service

type WeatherServicer interface {
}

// WeatherService 天气服务
type WeatherService struct {
}

func NewWeatherService() WeatherServicer {
	s := new(WeatherService)
	return s
}
