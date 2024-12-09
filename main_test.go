package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCalculateSum(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"zero and positive", 0, 1, 1},
        {"negative numbers", -1, -2, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calculateSum(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("calculateSum(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestSumHandler(t *testing.T) {
    tests := []struct {
        name           string
        url            string
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "valid parameters",
            url:            "/sum?a=5&b=3",
            expectedStatus: http.StatusOK,
            expectedBody:   "Sum of 5 and 3 is: 8",
        },
        {
            name:           "invalid parameter a",
            url:            "/sum?a=invalid&b=3",
            expectedStatus: http.StatusBadRequest,
            expectedBody:   "Parameter 'a' must be an integer\n",
        },
        {
            name:           "invalid parameter b",
            url:            "/sum?a=5&b=invalid",
            expectedStatus: http.StatusBadRequest,
            expectedBody:   "Parameter 'b' must be an integer\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, err := http.NewRequest("GET", tt.url, nil)
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(sumHandler)
            handler.ServeHTTP(rr, req)

            if status := rr.Code; status != tt.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v",
                    status, tt.expectedStatus)
            }

            if rr.Body.String() != tt.expectedBody {
                t.Errorf("handler returned unexpected body: got %v want %v",
                    rr.Body.String(), tt.expectedBody)
            }
        })
    }
}
