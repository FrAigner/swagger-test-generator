package swagger

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func FetchSwaggerDoc(urlString string) ([]byte, error) {
	// Überprüfe, ob die URL ein vertrauenswürdiges Schema hat
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, fmt.Errorf("ungültige URL: %v", err)
	}

	if !strings.HasPrefix(parsedURL.Scheme, "http") {
		return nil, fmt.Errorf("unsichere URL, nur HTTP(S) ist erlaubt")
	}

	resp, err := http.Get(urlString)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Abrufen der Swagger-Dokumentation: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ungültiger Statuscode: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen des Swagger-Inhalts: %v", err)
	}

	return body, nil
}
