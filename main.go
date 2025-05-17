package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type InputData struct {
    Numbers []int `json:"numbers"`
}

type OutputData struct {
    Sum    int     `json:"sum"`
    Avg    float64 `json:"average"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
    var input InputData
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    sum := 0
    for _, num := range input.Numbers {
        sum += num
    }

    avg := 0.0
    if len(input.Numbers) > 0 {
        avg = float64(sum) / float64(len(input.Numbers))
    }

    result := OutputData{Sum: sum, Avg: avg}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/calculate", calculateHandler)

    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
