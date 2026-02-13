package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

//Handling the input 
func RunWeatherStation() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("--- Weather Station ---")
    state := WeatherState{}
    sensorMap := NewSensorMap(&state)

    for {
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }
        input = strings.TrimSpace(input)

        switch {
        case input == "get":
            printState(state)

        case input == "clear":
            clearState(&state)

        case input == "exit":
            fmt.Println("Exiting...")
            return

        default:
            handleSensorInput(input, sensorMap)
        }
    }
}

//Handling the input if we give sensor values
func handleSensorInput(input string, sensorMap map[int]**float64) {
    parts := strings.Split(input, ",")
    if len(parts) != 2 {
        return
    }

    idStr := strings.TrimSpace(parts[0])
    valueStr := strings.TrimSpace(parts[1])

    id, err := strconv.Atoi(idStr)
    if err != nil {
        return
    }

    ptr, ok := sensorMap[id]
    if !ok {
        return
    }

    if strings.ToUpper(valueStr) == "NULL" {
        *ptr = nil
        return
    }

    val, err := strconv.ParseFloat(valueStr, 64)
    if err != nil {
        return
    }

    *ptr = &val
}
