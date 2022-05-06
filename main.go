package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type sudoku struct {
	field    [][]string
	sizeArea int
	space    int
}

var (
	solved    = errors.New("sudoku is solved")
	notSolved = errors.New("sudoku is not solved")
)

func (s *sudoku) loadFromFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0x444)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	r := csv.NewReader(f)
	r.Comma = ' '
	s.field, err = r.ReadAll()
	if err != nil {
		return err
	}
	s.fillFields()

	return nil
}

func (s *sudoku) solve() error {
	row, col, isNotFound := s.getEmpty()
	if isNotFound {
		return solved
	}
	for d := range s.field {
		digit := strconv.Itoa(d + 1)
		if s.validate(row, col, digit) {
			s.field[row][col] = digit
			if err := s.solve(); errors.Is(err, solved) {
				return solved
			}
			s.field[row][col] = "."
		}
	}
	return notSolved
}

func (s *sudoku) getEmpty() (int, int, bool) {
	for i, col := range s.field {
		for j, cell := range col {
			if cell == "." {
				return i, j, false
			}
		}
	}
	return 0, 0, true
}

func (s *sudoku) validate(row, col int, digit string) bool {
	// rows
	for i := range s.field {
		if i != row && s.field[i][col] == digit {
			return false
		}
	}
	// cols
	for j := range s.field[0] {
		if j != col && s.field[row][j] == digit {
			return false
		}
	}
	// area
	rowBegin := row / s.sizeArea * s.sizeArea
	colBegin := col / s.sizeArea * s.sizeArea
	for i := rowBegin; i < rowBegin+s.sizeArea; i++ {
		for j := colBegin; j < colBegin+s.sizeArea; j++ {
			if i != row && j != col && s.field[i][j] == digit {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) fillFields() {
	s.sizeArea = int(math.Sqrt(float64(len(s.field))))
	s.space = len(strconv.Itoa(len(s.field)))
}

func (s sudoku) String() string {
	lines := make([]string, 0)
	for i := range s.field {
		if i%s.sizeArea == 0 {
			lines = append(lines, s.rowSep())
		}
		lines = append(lines, s.colSep(i))
	}
	lines = append(lines, s.rowSep())
	return strings.Join(lines, "\n")
}

func (s sudoku) rowSep() string {
	line := strings.Builder{}
	line.WriteString("+")
	for i := 0; i < len(s.field)/s.sizeArea; i++ {
		line.WriteString(strings.Repeat("–", s.space*s.sizeArea*2-1) + "+")
	}
	return line.String()
}

func (s sudoku) colSep(row int) string {
	line := strings.Builder{}
	for j, cell := range s.field[row] {
		if j%s.sizeArea == 0 {
			line.WriteString("|")
		} else {
			line.WriteString(" ")
		}
		line.WriteString(fmt.Sprintf(fmt.Sprintf("%%%ds", s.space), cell))
	}
	line.WriteString("|")
	return line.String()
}

func (s *sudoku) saveToFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	w := csv.NewWriter(f)
	w.Comma = ' '
	return w.WriteAll(s.field)
}

func main() {
	var (
		s                     sudoku
		err                   error
		inputFile, outputFile string
	)
	flag.StringVar(&inputFile, "i", "matrix.csv", "input csv-file")
	flag.StringVar(&outputFile, "o", "result.csv", "output csv-file")
	flag.Parse()
	err = s.loadFromFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s)
	if err = s.solve(); errors.Is(err, notSolved) {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(err)
	fmt.Println(s)
	err = s.saveToFile(outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
