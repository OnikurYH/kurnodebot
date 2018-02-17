package openweathermap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"kurnode.com/kurnodebot/config"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var baseURL = "https://api.openweathermap.org/data/2.5/weather"

// WeatherMainParameters ...
type WeatherMainParameters struct {
	Temp      float64 `json:"temp"`
	Humidity  float64 `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

// CurrentWeather ...
type CurrentWeather struct {
	ID    int    `json:"id"`
	Dt    int64  `json:"dt"`
	Name  string `json:"name"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Sys struct {
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int64   `json:"sunrise"`
		Sunset  int64   `json:"sunset"`
	} `json:"sys"`
	Main WeatherMainParameters `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	}
	Clouds struct {
		All float64 `json:"all"`
	}
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
	Rain struct {
		Last3H int `json:"3h"`
	}
	Snow struct {
		Last3H int `json:"3h"`
	}
}

// GetCurrent ...
func GetCurrent(location string, units string) (*CurrentWeather, error) {
	querys := make([]string, 0)
	querys = append(querys, fmt.Sprintf("APPID=%s", config.Get().External.OpenWeatherMap.APIKey))
	querys = append(querys, fmt.Sprintf("q=%s", location))
	querys = append(querys, fmt.Sprintf("units=%s", units))

	url := fmt.Sprintf("%s?%s", baseURL, strings.Join(querys, "&"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	currentWeather := &CurrentWeather{}
	if err := json.Unmarshal(body, currentWeather); err != nil {
		return nil, err
	}
	return currentWeather, nil
}
