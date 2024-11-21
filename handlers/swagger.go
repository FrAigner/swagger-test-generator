package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swagger-test-generator/metrics"
	"swagger-test-generator/swagger"
)

type SwaggerRequest struct {
	URL string `json:"url"`
}

func HandleGenerateTests(w http.ResponseWriter, r *http.Request) {
	var req SwaggerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		metrics.IncrementFailedRequests("400")
		http.Error(w, "Ungültige Anfrage", http.StatusBadRequest)
		return
	}

	swaggerDoc, err := swagger.FetchSwaggerDoc(req.URL)
	if err != nil {
		metrics.IncrementFailedRequests("500")
		http.Error(w, fmt.Sprintf("Fehler: %v", err), http.StatusInternalServerError)
		return
	}

	tests, err := swagger.GenerateTestsFromSwagger(swaggerDoc)
	if err != nil {
		metrics.IncrementFailedRequests("500")
		http.Error(w, fmt.Sprintf("Fehler bei der Testgenerierung: %v", err), http.StatusInternalServerError)
		return
	}

	metrics.IncrementSuccessfulRequests("200")
	w.Write([]byte(tests))
}
