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

For help:

```bash
go run . -h
```

## Tools Version

```bash
go run . -v
```

## Horizontal Space Cleaning

```bash
go run . -c ./input.txt ./output.txt ./airport-lookup.csv
```

This command can use for trimming horizontal space. Without this command it will remain same as in input text file.

### How to check the codes:

First go to the input.txt file and you can put your testing material over there and then run the code by acceptable command
_go run . ./input.txt ./output.txt ./airport-lookup.csv_ then this will create new outcomes into the output.txt. Thats how you can get the testing result.

## Code to Airport and City name transformation check

"#" = IATA airport code → converted to full airport name

"##" = ICAO airport code → converted to full airport name

"\*#" = City name from IATA code → converted to city name

## 👤 Author

Rifat Bin Mahi  
Go CLI Project – Itinerary Prettifier
