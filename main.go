package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func calculateSum(a, b int) int {
    return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Go Web App!")
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
    // Get query parameters
    query := r.URL.Query()

    // Get a and b parameters
    aStr := query.Get("a")
    bStr := query.Get("b")

    // Convert string to int
    a, err := strconv.Atoi(aStr)
    if err != nil {
        http.Error(w, "Parameter 'a' must be an integer", http.StatusBadRequest)
        return
    }

    b, err := strconv.Atoi(bStr)
    if err != nil {
        http.Error(w, "Parameter 'b' must be an integer", http.StatusBadRequest)
        return
    }

    result := calculateSum(a, b)

    fmt.Fprintf(w, "Sum of %d and %d is: %d", a, b, result)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/sum", sumHandler)
    fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
