package main

type WeatherState struct {
    airTemp       *float64
    airPressure   *float64
    precipitation *float64
    windSpeed     *float64
    windDirection *float64
    humidity      *float64
    dewPoint      *float64
    soilMoisture  *float64
    cloudCover    *float64
}

func NewSensorMap(state *WeatherState) map[int]**float64 {
    return map[int]**float64{
        1:  &state.airTemp,
        2:  &state.airPressure,
        7:  &state.precipitation,
        11: &state.windSpeed,
        12: &state.windDirection,
        13: &state.humidity,
        14: &state.dewPoint,
        15: &state.soilMoisture,
        22: &state.cloudCover,
    }
}
