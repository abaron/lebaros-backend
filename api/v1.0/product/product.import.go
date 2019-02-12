package product

import (
	"encoding/csv"
	"os"
	"regexp"
	"strings"
)

type CsvProduct struct {
	Sku   string
	Name  string
	Size  string
	Color string
	Stock string
}

func getCsvProduct(filename string) []CsvProduct {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	data := []CsvProduct{}

	// Loop through lines & turn into object
	for i, col := range lines {
		if i == 0 {
			continue
		}

		var rgx = regexp.MustCompile(`\((.*?)\)`)
		sc := rgx.FindStringSubmatch(col[1])
		sizeColor := strings.Split(sc[1], ",")

		rgx = regexp.MustCompile(`([^[\)]+)(?:$| \()`)
		name := rgx.FindStringSubmatch(col[1])

		row := CsvProduct{
			Sku:   col[0],
			Name:  name[1],
			Size:  sizeColor[0],
			Color: sizeColor[1],
			Stock: col[2],
		}

		data = append(data, row)
	}

	return data
}
