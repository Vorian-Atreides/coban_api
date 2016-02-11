package common

import (
	"errors"
	"time"

	"coban/api/src/databases"
)

func typeOfMachine(value byte) bool {
	return value == byte(0x16)
}

func useCase(value byte) bool {
	return value == byte(0x01)
}

var payments = map[byte]string{
	0x00: "Cash",
	0x02: "View card",
	0x0B: "Pitapa card",
	0x0D: "Pasmo",
	0x3F: "Mobile suica",
}

func payment(value byte) (string, error) {
	if data, ok := payments[value]; ok {
		return data, nil
	}
	return "", errors.New("Payment not found.")
}

var methods = map[byte]string{
	0x01: "Enter",
	0x02: "Enter/Exit",
	0x03: "Season Ticket Enter/Exit",
	0x04: "Enter/Season Ticket Exit",
	0x0E: "Ticket Window Exit",
	0x0F: "Enter/Exit(Bus)",
}

func method(value byte) (string, error) {
	if data, ok := methods[value]; ok {
		return data, nil
	}
	return "", errors.New("Method not found.")
}

func date(value []byte) time.Time {
	// 7 bits
	year := value[0] >> 1
	// 4 bits
	month := value[0] & 0x01 << 3
	month |= value[1] >> 5
	// 5 bits
	day := value[1] & 0x1F

	restDate := time.Now().Year() / 100 * 100
	dateTime := time.Date(restDate+int(year), time.Month(month), int(day),
		0, 0, 0, 0, time.UTC)
	return dateTime
}

func station(value []byte) (databases.Station, error) {
	line := uint(value[0])
	station := uint(value[1])

	var data databases.Station
	databases.DB.Where(databases.Station{LineCode: line, StationCode: station}).
		First(&data)
	if data.ID == 0 {
		return data, errors.New("This station doesn't exist.")
	}

	return data, nil
}

func balance(values []byte) uint16 {
	var target uint16

	target = uint16(values[1])
	target <<= 8
	target += uint16(values[0])
	return target
}

// ParseTransportHistory parse an array of byte corresponding to a
// transport history from the card
func ParseTransportHistory(data []byte) (databases.TransportHistory, error) {
	if len(data) != 16 {
		return databases.TransportHistory{},
			errors.New("This data is corrupted.")
	}
	if !typeOfMachine(data[0]) {
		return databases.TransportHistory{},
			errors.New("This type of machine isn't supported.")
	}
	if !useCase(data[1]) {
		return databases.TransportHistory{},
			errors.New("This use case isn't supported.")
	}
	_, _ = payment(data[2])
	_, _ = method(data[3])
	_date := date(data[4:6])
	entrance, err := station(data[6:8])
	if err != nil {
		return databases.TransportHistory{},
			errors.New("There are an error with the entrance.")
	}
	exit, err := station(data[8:10])
	if err != nil {
		return databases.TransportHistory{},
			errors.New("There are an error with the exit.")
	}
	_balance := balance(data[10:12])

	transportHistory := databases.TransportHistory{
		Date:       _date,
		EntranceID: entrance.ID,
		ExitID:     exit.ID,
		Stock:      uint(_balance),
	}
	return transportHistory, nil
}
