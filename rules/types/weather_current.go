package types

import (
	"strconv"
	"strings"

	"kurnode.com/kurnodebot/config"
	"kurnode.com/kurnodebot/external/openweathermap"
	"kurnode.com/kurnodebot/irc"
)

// RuleTypeCurrentWeather ...
type RuleTypeCurrentWeather struct{}

// Do ...
func (rt *RuleTypeCurrentWeather) Do(ro config.RuleOutput) string {
	return rt.injectWeather(ro, ro.Content, ro.DefaultLocation)
}

// DoWithMessage ...
func (rt *RuleTypeCurrentWeather) DoWithMessage(msg irc.Message, ro config.RuleOutput) string {
	location := ro.DefaultLocation
	if len(msg.ContentArgs) >= 1 {
		location = strings.Join(msg.ContentArgs, " ")
	}
	return rt.injectWeather(ro, ro.Content, location)
}

func (rt *RuleTypeCurrentWeather) injectWeather(ro config.RuleOutput, output string, location string) string {
	w, err := openweathermap.GetCurrent(location, ro.Units)
	if err != nil {
		return "Cannot get weather"
	}

	output = strings.Replace(output, "{{temp}}", strconv.FormatFloat(w.Main.Temp, 'f', -1, 32), -1)
	output = strings.Replace(output, "{{temp_min}}", strconv.FormatFloat(w.Main.TempMin, 'f', -1, 32), -1)
	output = strings.Replace(output, "{{temp_max}}", strconv.FormatFloat(w.Main.TempMax, 'f', -1, 32), -1)
	output = strings.Replace(output, "{{name}}", w.Name, -1)
	var weathers []string
	for _, we := range w.Weather {
		weathers = append(weathers, we.Description)
	}
	output = strings.Replace(output, "{{weathers}}", strings.Join(weathers, ", "), -1)
	return output
}
