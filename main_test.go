package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetRisks(t *testing.T) {
    req, err := http.NewRequest("GET", "/v1/risks", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleRisks)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var risks []Risk
    if err := json.NewDecoder(rr.Body).Decode(&risks); err != nil {
        t.Errorf("handler returned unexpected body: %v", rr.Body.String())
    }
}

func TestCreateRisk(t *testing.T) {
    risk := Risk{
        State:       "open",
        Title:       "Test Risk",
        Description: "This is a test risk",
    }
    body, err := json.Marshal(risk)
    if err != nil {
        t.Fatal(err)
    }

    req, err := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleRisks)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    var createdRisk Risk
    if err := json.NewDecoder(rr.Body).Decode(&createdRisk); err != nil {
        t.Errorf("handler returned unexpected body: %v", rr.Body.String())
    }

    if createdRisk.ID == "" {
        t.Errorf("handler did not generate an ID for the risk")
    }
}

func TestGetRisk(t *testing.T) {
    risk := Risk{
        ID:          "test-id",
        State:       "open",
        Title:       "Test Risk",
        Description: "This is a test risk",
    }
    risks[risk.ID] = risk

    req, err := http.NewRequest("GET", "/v1/risks/test-id", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleRisk)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var retrievedRisk Risk
    if err := json.NewDecoder(rr.Body).Decode(&retrievedRisk); err != nil {
        t.Errorf("handler returned unexpected body: %v", rr.Body.String())
    }

    if retrievedRisk.ID != risk.ID {
        t.Errorf("handler returned wrong risk: got %v want %v", retrievedRisk.ID, risk.ID)
    }
}
 