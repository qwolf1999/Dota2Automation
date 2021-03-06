package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const TEMPLATE = "gene_template.lua"

func GetGeneFromFile(filename string) []float64 {
	var gene []float64
	rawdata, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read file", err)
	}
	data := string(rawdata)
	lines := strings.Split(data, "\n")
	for i := 0; i < len(lines); i++ {
		numberStr := strings.Split(lines[i], "=")
		if len(numberStr) > 1 && !strings.Contains(numberStr[1], "{") {
			if strings.Contains(numberStr[1], ",") || strings.Contains(numberStr[1], "\r") {
				numberStr[1] = numberStr[1][:len(numberStr[1])-2]
			}
			if strings.Contains(numberStr[1], " ") {
				numberStr[1] = numberStr[1][1:]
			}

			number, err := strconv.ParseFloat(numberStr[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			gene = append(gene, number)
		}
	}
	return (gene)
}

func WriteGeneToFile(gene []float64, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	template, err := ioutil.ReadFile(TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(template), "\n")
	for i, g := 0, 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "%f") {
			fmt.Fprintln(file, fmt.Sprintf(lines[i], gene[g]))
			g++
		} else {
			fmt.Fprintln(file, lines[i])
		}
	}
	file.Close()
}

//ogremagi
//lich
//medusa
//chaosknight
//bane
func ConvertGeneData(ranking [5]Top5genes) [5][][]float64 {
	var result [5][][]float64
	for i := range ranking {
		result[i] = make([][]float64, 5)
		for j := 0; j < 5; j++ {
			result[i][j] = GetGeneFromFile(ranking[i].gene[j].fileName)
		}
	}
	return result
}

func WriteBestGenes(path string, result [5][][]float64) {
	bestGeneDir := path + "/best_genes"
	os.Mkdir(bestGeneDir, 0777)
	for i := 0; i < 5; i++ {
		WriteGeneToFile(result[i][0], bestGeneDir+"/gene_"+Roaster[i]+".lua")
	}
}
