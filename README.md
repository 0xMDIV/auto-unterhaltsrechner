# Auto-Unterhaltsrechner

Ein umfassender Fahrzeugkostenrechner für Windows Desktop, entwickelt in Go mit Fyne GUI.

## Features

### Eingabefelder
- Kraftstoffverbrauch pro 100km
- Stromverbrauch pro 100km (für Hybrid/Elektro)
- Kraftstoff- und Strompreise
- Kraftstoffart (Diesel, Ultimate, Super, SuperPlus, Ultimate Diesel)
- Stromart (Haushaltssteckdose, Öffentliche Ladestation)
- Tankgröße in Litern
- Batteriegröße in kWh
- Monatliche Kilometer
- Jährliche KFZ-Steuer
- Jährliche Versicherung
- Finanzierungs-/Leasingrate pro Monat
- Finanzierungs-/Leasinglaufzeit in Monaten
- Kaufpreis (für Wertverlustkalkulation)
- Erwartete Besitzdauer in Jahren

### Berechnungen
- Monatliche und jährliche Kraftstoffkosten
- Monatliche und jährliche Stromkosten
- Gesamte monatliche Betriebskosten
- Gesamte jährliche Betriebskosten
- Wertverlust (separate Berechnung)
- Kosten pro Kilometer
- Break-Even-Analyse für Elektro vs. Verbrenner

### Funktionen
- Profile speichern/laden für verschiedene Fahrzeuge
- Ergebnisse als PDF oder CSV exportieren
- Vergleichsmodus für mehrere Fahrzeuge
- Moderner Dark/Light Theme Toggle
- Diagramme zur Kostenvisualisierung
- Responsive Layout
- Deutsche Zahlenformatierung (Komma als Dezimaltrennzeichen)
- Tooltips für alle Eingaben mit Erklärungen

## Installation

### Voraussetzungen
- Go 1.21 oder höher
- CGO-fähiger C-Compiler (für Windows: TDM-GCC oder Microsoft Visual Studio)

### Windows Build mit TDM-GCC

1. TDM-GCC installieren:
   - Downloaden von: https://jmeubank.github.io/tdm-gcc/
   - 64-bit Version installieren

2. Umgebungsvariablen setzen:
   ```cmd
   set CGO_ENABLED=1
   set CC=gcc
   ```

3. Dependencies installieren:
   ```cmd
   go mod tidy
   ```

4. Anwendung kompilieren:
   ```cmd
   go build -o auto-unterhaltsrechner.exe ./cmd/app
   ```

### Alternative: Cross-Platform Build
Für einfachere Builds können Sie auch fyne package verwenden:

```cmd
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -o auto-unterhaltsrechner.exe ./cmd/app
```

## Struktur

```
auto-unterhaltsrechner/
├── cmd/app/                 # Einstiegspunkt der Anwendung
│   └── main.go
├── internal/
│   ├── calculator/          # Alle Berechnungslogik
│   │   └── calculator.go
│   ├── ui/                  # GUI-Komponenten
│   │   ├── app.go          # Haupt-App-Struktur
│   │   ├── input_form.go   # Eingabeformular
│   │   ├── results_view.go # Ergebnisanzeige
│   │   ├── dialogs.go      # Dialoge (Export, Vergleich, etc.)
│   │   └── utils.go        # Deutsche Zahlenformatierung
│   ├── models/              # Datenstrukturen
│   │   └── models.go
│   └── storage/             # Speicher-/Ladefunktionen
│       └── storage.go
├── go.mod
├── go.sum
└── README.md
```

## Verwendung

1. Anwendung starten: `./auto-unterhaltsrechner.exe`

2. **Neues Profil erstellen:**
   - Auf "Neu" klicken
   - Fahrzeugdaten eingeben
   - Profil speichern

3. **Profil laden:**
   - Auf "Öffnen" klicken
   - Aus der Liste wählen

4. **Vergleich:**
   - Mehrere Profile erstellen
   - Auf "Vergleich" klicken
   - Profile auswählen

5. **Export:**
   - Profil auswählen
   - Auf "Export" klicken
   - Format wählen (CSV/JSON)

## Berechnungslogik

### Kraftstoffkosten
```
Monatliche Kosten = (Verbrauch L/100km × Monatliche km ÷ 100) × Kraftstoffpreis €/L
```

### Stromkosten
```
Monatliche Kosten = (Verbrauch kWh/100km × Monatliche km ÷ 100) × Strompreis €/kWh
```

### Wertverlust
- Linearer Wertverlust mit 20% Restwert nach Besitzdauer
- Für Fahrzeuge älter als 10 Jahre: 10% Restwert

### Gesamtkosten
```
Monatliche Gesamtkosten = Kraftstoff + Strom + KFZ-Steuer/12 + Versicherung/12 + Finanzierung
```

## Einstellungen

Die Anwendung speichert Einstellungen und Profile in:
- Windows: `%USERPROFILE%\.auto-unterhaltsrechner\`

### Konfigurierbare Einstellungen:
- Theme (Hell/Dunkel)
- Standard-Kraftstoffpreis
- Standard-Strompreis

## Problembehandlung

### Build-Probleme
1. **CGO-Fehler:** Stellen Sie sicher, dass ein C-Compiler installiert ist
2. **GL-Fehler:** Installieren Sie TDM-GCC oder Visual Studio
3. **Import-Fehler:** Führen Sie `go mod tidy` aus

### Runtime-Probleme
1. **GUI startet nicht:** Stellen Sie sicher, dass alle DLLs verfügbar sind
2. **Profile werden nicht gespeichert:** Überprüfen Sie Schreibrechte im Benutzerverzeichnis

## Mitwirken

Beiträge sind willkommen! Bitte:
1. Fork des Repositories erstellen
2. Feature-Branch erstellen
3. Änderungen committen
4. Pull Request erstellen

## Lizenz

Dieses Projekt steht unter der MIT-Lizenz. Siehe LICENSE-Datei für Details.

## Systemanforderungen

- Windows 10 oder höher
- 512 MB RAM
- 50 MB freier Speicherplatz
- OpenGL 2.1 oder höher

## Changelog

### Version 1.0.0
- Initiale Version
- Vollständige Kostenberechnung
- Profile speichern/laden
- Export-Funktionalität
- Vergleichsmodus
- Deutsche Lokalisierung