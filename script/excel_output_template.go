package script

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func GenFiles(dataFileName string, templateFileName string, outputPath string) {
	templateFile, err := xlsx.OpenFile(templateFileName)
	if err != nil {
		log.Fatalf("unable to open template file: %v", err)
	}

	dataFile, err := xlsx.OpenFile(dataFileName)
	if err != nil {
		log.Fatalf("unable to open data file: %v", err)
	}

	dataSheet := dataFile.Sheets[0]

	headerRow := dataSheet.Row(0)

	headers := getHeaders(headerRow)

	for rowIndex, row := range dataSheet.Rows[1:] {
		replacePlaceholders(templateFile, headers, row)
		newFileName := fmt.Sprintf(outputPath+"output_%d.xlsx", rowIndex)
		err := templateFile.Save(newFileName)
		if err != nil {
			log.Fatalf("unable to save new file: %v", err)
		}
	}
}

func getHeaders(row *xlsx.Row) []string {
	var headers []string
	for _, cell := range row.Cells {
		header := cell.String()
		headers = append(headers, header)
	}
	return headers
}

func replacePlaceholders(file *xlsx.File, headers []string, row *xlsx.Row) {
	sheet := file.Sheets[0]
	for i, header := range headers {
		cellValue := row.Cells[i].String()
		placeholder := "{{" + header + "}}"
		for _, templateRow := range sheet.Rows {
			for _, cell := range templateRow.Cells {
				if cell.Value == placeholder {
					cell.Value = cellValue
				}
			}
		}
	}
}
