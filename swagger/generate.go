package swagger

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateTestsFromSwagger(swagger []byte) (string, error) {
	var swaggerDoc map[string]interface{}
	if err := json.Unmarshal(swagger, &swaggerDoc); err != nil {
		return "", fmt.Errorf("Fehler beim Parsen der Swagger-Dokumentation: %v", err)
	}

	paths := swaggerDoc["paths"].(map[string]interface{})
	var testCases strings.Builder

	for path, methods := range paths {
		for method := range methods.(map[string]interface{}) {
			testName := fmt.Sprintf("Test %s %s", strings.ToUpper(method), path)

			testCases.WriteString(fmt.Sprintf(`
test('%s', async ({ request }) => {
    const response = await request.%s('%s', {
        // Optional: Body, Headers, Query Parameters
    });
    expect(response.status()).toBe(200);
});
`, testName, method, path))
		}
	}

	playwrightTests := fmt.Sprintf(`
	
import { test, expect } from '@playwright/test';

test.describe('API Tests', () => {
%s
});
`, testCases.String())

	return playwrightTests, nil
}
