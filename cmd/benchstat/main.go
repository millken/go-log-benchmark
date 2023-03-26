package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Dataset struct {
	Name       string  // test name
	Iterations int     // iterations
	NsOp       float64 // ns/op
	BytesOp    float64 // B/op
	AllocsOp   float64 // allocs/op
}

type Datasets []Dataset

func (d Datasets) Len() int {
	return len(d)
}

func (d Datasets) Less(i, j int) bool {
	return d[i].NsOp < d[j].NsOp
}

func (d Datasets) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Datasets) CSVResult() string {
	csvHeader := "Name,ns/op,B/op,allocs/op\n"
	csvResult := csvHeader
	for _, dataset := range d {
		csvResult += fmt.Sprintf("%s,%.2f,%.0f,%.0f\n", dataset.Name[0:len(dataset.Name)-13], dataset.NsOp, dataset.BytesOp, dataset.AllocsOp)
	}
	return csvResult
}

func (d Datasets) MarkdownResult() string {
	markdownHeader := "| Name | ns/op | B/op | allocs/op |\n| --------- | --------- | --------- | --------- |\n"
	markdownResult := markdownHeader
	for _, dataset := range d {
		markdownResult += fmt.Sprintf("| %s | %.2f | %.0f | %.0f |\n", dataset.Name[0:len(dataset.Name)-13], dataset.NsOp, dataset.BytesOp, dataset.AllocsOp)
	}
	return markdownResult
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // skip header
	var datasets Datasets
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if dataset, ok := parseLine(line); ok {
			datasets = append(datasets, dataset)
		}
	}
	sort.Sort(datasets)
	var textNegative, jsonNegative, textPositive, jsonPositive Datasets
	for _, dataset := range datasets {
		lenth := len(dataset.Name)
		group := dataset.Name[lenth-12:]
		switch group {
		case "TextNegative":
			textNegative = append(textNegative, dataset)
		case "JSONNegative":
			jsonNegative = append(jsonNegative, dataset)
		case "TextPositive":
			textPositive = append(textPositive, dataset)
		case "JSONPositive":
			jsonPositive = append(jsonPositive, dataset)
		}
	}
	markdownResult := fmt.Sprintf("# Benchmark Results (%s)\n## TextNegative\n%s\n## JSONNegative\n%s\n## TextPositive\n%s\n## JSONPositive\n%s\n", time.Now().Format("2006-01-02"), textNegative.MarkdownResult(), jsonNegative.MarkdownResult(), textPositive.MarkdownResult(), jsonPositive.MarkdownResult())
	os.WriteFile("benchmark_result.md", []byte(markdownResult), 0644)

}

func parseLine(line string) (Dataset, bool) {
	dataset := Dataset{}
	// BenchmarkMemset-4 10000000 0.000000 ns/op 0 B/op 0 allocs/op
	f := strings.Fields(line)
	if len(f) < 4 {
		return dataset, false
	}
	name := f[0]
	if !strings.HasPrefix(name, "Benchmark") {
		return dataset, false
	}
	name = strings.TrimPrefix(name, "Benchmark")
	dataset.Iterations, _ = strconv.Atoi(f[1])

	dataset.Name = name[0 : len(name)-2]
	for i := 2; i+2 <= len(f); i += 2 {
		val, err := strconv.ParseFloat(f[i], 64)
		if err != nil {
			continue
		}
		unit := f[i+1]
		switch unit {
		case "ns/op":
			dataset.NsOp = val
		case "B/op":
			dataset.BytesOp = val
		case "allocs/op":
			dataset.AllocsOp = val
		}
	}
	// Name: BenchmarkMemset-4
	// Iterations: 10000000
	// ns/op: 0.000000
	// B/op: 0.000000
	// allocs/op: 0.000000
	return dataset, true
}
