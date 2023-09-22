package script

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strconv"
	"time"
)

// 格式支持更友好
func GenFiles(dataFileName string, templateFileName string, outputPath string) {
	dataFile, err := excelize.OpenFile(dataFileName)
	if err != nil {
		log.Fatalf("unable to open data file: %v", err)
	}

	dataSheet := dataFile.GetSheetMap()[dataFile.GetActiveSheetIndex()]
	if err != nil {
		log.Fatalf("unable to get avaliable sheet from template file: %v", err)
	}
	templateFile, err := excelize.OpenFile(templateFileName)
	if err != nil {
		log.Fatalf("unable to open template file: %v", err)
	}

	templateSheet := templateFile.GetSheetMap()[templateFile.GetActiveSheetIndex()]

	rows, err := dataFile.GetRows(dataSheet)
	if err != nil {
		log.Fatalf("unable to get rows: %v", err)
	}

	headers := rows[0]

	for rowIndex, row := range rows[1:] {
		newFile, err := excelize.OpenFile(templateFileName)
		if err != nil {
			log.Fatalf("unable to open template file: %v", err)
		}
		replacePlaceholders(newFile, templateSheet, headers, row)
		newFileName := fmt.Sprintf(outputPath+"output_%d.xlsx", rowIndex)
		if err := newFile.SaveAs(newFileName); err != nil {
			log.Fatalf("unable to save new file: %v", err)
		}
	}
}
func replacePlaceholders(file *excelize.File, sheet string, headers []string, row []string) {
	rows, err := file.GetRows(sheet)
	if err != nil {
		log.Fatalf("unable to get rows from template: %v", err)
	}

	for rowIndex, templateRow := range rows {
		for colIndex, cell := range templateRow {
			for i, header := range headers {
				placeholder := "{{" + header + "}}"
				if cell == placeholder {
					axis, err := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
					if err != nil {
						log.Fatalf("unable to convert coordinates to cell name: %v", err)
					}

					styleID, err := file.GetCellStyle(sheet, axis)
					if err != nil {
						log.Fatalf("unable to get cell style: %v", err)
					}

					// 基于原始数据的值格式设置单元格的值
					if val, err := strconv.Atoi(row[i]); err == nil {
						file.SetCellInt(sheet, axis, val)
					} else if val, err := strconv.ParseFloat(row[i], 64); err == nil {
						file.SetCellFloat(sheet, axis, val, -1, 64)
					} else if t, err := time.Parse("2006-01-02", row[i]); err == nil {
						file.SetCellDefault(sheet, axis, t.Format(time.RFC3339))
					} else if t, err := time.Parse("2006/01/02", row[i]); err == nil {
						file.SetCellDefault(sheet, axis, t.Format(time.RFC3339))
					} else if t, err := time.Parse("02-01-2006", row[i]); err == nil {
						file.SetCellDefault(sheet, axis, t.Format(time.RFC3339))
					} else if row[i] == "true" || row[i] == "false" {
						file.SetCellBool(sheet, axis, row[i] == "true")
					} else {
						file.SetCellValue(sheet, axis, row[i])
					}

					// 设置单元格的样式为模板中的样式
					file.SetCellStyle(sheet, axis, axis, styleID)
				}
			}
		}
	}
}
