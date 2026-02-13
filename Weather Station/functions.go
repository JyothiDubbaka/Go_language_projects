package main

import "fmt"

// printState prints all sensors in order
func printState(state WeatherState) {
sensors := []struct {
name string
val  **float64
}{
{"airTemp", &state.airTemp},
{"airPressure", &state.airPressure},
{"precipitation", &state.precipitation},
{"windSpeed", &state.windSpeed},
{"windDirection", &state.windDirection},
{"humidity", &state.humidity},
{"dewPoint", &state.dewPoint},
{"soilMoisture", &state.soilMoisture},
{"cloudCover", &state.cloudCover},
}

for _, s := range sensors {
if *s.val == nil {
fmt.Println(s.name + ":NULL")
} else {
fmt.Printf("%s:%v\n", s.name, **s.val)
}
}
}

// clearState resets all sensors to nil
func clearState(state *WeatherState) {
state.airTemp = nil
state.airPressure = nil
state.precipitation = nil
state.windSpeed = nil
state.windDirection = nil
state.humidity = nil
state.dewPoint = nil
state.soilMoisture = nil
state.cloudCover = nil
}
