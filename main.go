package main

import (
    "encoding/json"
    "github.com/google/uuid"
    "log"
    "net/http"
)

type Risk struct {
    ID          string `json:"id"`
    State       string `json:"state"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

var risks = make(map[string]Risk)

func main() {
    http.HandleFunc("/v1/risks", handleRisks)
    http.HandleFunc("/v1/risks/", handleRisk)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRisks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getRisks(w, r)
    case "POST":
        createRisk(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleRisk(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/v1/risks/"):]
    switch r.Method {
    case "GET":
        getRisk(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getRisks(w http.ResponseWriter, r *http.Request) {
    var riskList []Risk
    for _, risk := range risks {
        riskList = append(riskList, risk)
    }
    json.NewEncoder(w).Encode(riskList)
}

func createRisk(w http.ResponseWriter, r *http.Request) {
    var risk Risk
    if err := json.NewDecoder(r.Body).Decode(&risk); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    risk.ID = uuid.New().String()
    risks[risk.ID] = risk
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(risk)
}

func getRisk(w http.ResponseWriter, r *http.Request, id string) {
    risk, exists := risks[id]
    if !exists {
        http.Error(w, "Risk not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(risk)
}

