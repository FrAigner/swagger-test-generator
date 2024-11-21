
# Swagger Test Generator

## Beschreibung

Die Anwendung dient dazu, Playwright-Tests aus einer Swagger-Dokumentation zu generieren. Die API empfängt eine URL, die auf eine Swagger-Dokumentation verweist, und gibt Playwright-Tests zurück, die die APIs testen.

## Installation

### Abhängigkeiten installieren

Stelle sicher, dass Go und die Abhängigkeiten installiert sind:

```bash
go mod tidy
```

### Swagger-Dokumentation erstellen (optional)

Wenn du die Swagger-Dokumentation generieren möchtest, führe den folgenden Befehl aus:

```bash
swag init
```

Dies erzeugt die Swagger-Dokumentation in der `docs`-Ordnerstruktur.

### Server starten

Starte den Server im HTTPS-Modus:

```bash
go run main.go
```

Der Server ist nun unter https://localhost:8080 erreichbar.

### HTTP Umleitung (optional)

Wenn du möchtest, dass HTTP-Anfragen zu HTTPS umgeleitet werden, stelle sicher, dass der HTTP-Server auf Port 80 läuft. Dies ist im Code als Go-Routine implementiert und leitet alle HTTP-Anfragen automatisch um.
