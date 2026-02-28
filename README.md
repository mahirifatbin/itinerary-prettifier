# ✈️ Itinerary Prettifier

## 📌 Project Overview

Itinerary Prettifier is a Go-based Command Line Interface (CLI) application that transforms raw travel itinerary data into a clean, human-readable format.

The program reads an input file containing flight itinerary information, processes airport codes using a CSV lookup table, formats date-time values, cleans whitespace, and writes the formatted output to a new file.

This project demonstrates practical use of:

- File handling in Go
- CSV parsing
- Map-based lookup optimization
- String manipulation
- Date-time transformation
- CLI argument handling
- Modular project structure

---

## 🛠 Setup and Installation

### Requirements

- Go 1.20 or higher installed

Check your Go version:

```bash
go version
```

### Clone the Repository

```bash
git clone https://gitea.kood.tech/rifatbinmahi1/prettifier

cd itinerary-prettifier
```

### Run the Program

```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

### If incorrect arguments are provided:

```bash
go run . -h
```

### Project Structure

```itinerary-prettifier/
├── main.go
├── go.mod
├── README.md
├──airport-lookup.csv
├──input.txt
├──output.txt
│
├── internal/
│ ├── clean/ # Whitespace cleaning logic
│ ├── csvlookup/ # CSV validation and lookup map building
│ ├── filehandle/ # File reading and CSV opening
│ └── transform/ # Code and date transformations
```

## 👤 Author

Rifat Bin Mahi  
Go CLI Project – Itinerary Prettifier
