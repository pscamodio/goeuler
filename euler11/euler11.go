package euler11

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Grid struct {
	Data        [][]int
	Lines, Cols int
}

type Quad [4]int

func LoadGrid(path string) (g Grid, e error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\r\n")
	g.Data = make([][]int, 20)
	for l, line := range lines {
		g.Data[l] = make([]int, 20)
		vals := strings.Split(line, " ")
		for c, val := range vals {
			num, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				return
			}
			g.Data[l][c] = int(num)
		}
	}
	g.Lines = len(g.Data)
	g.Cols = len(g.Data[0])
	return
}

func (g *Grid) MaxQuad() int {
	quads := g.ExtractQuads()
	max := 0
	for _, quad := range quads {
		prod := 1
		for _, val := range quad {
			prod *= val
		}
		if prod > max {
			max = prod
		}
	}
	return max
}

func (g *Grid) ExtractQuads() []Quad {
	quads := make([]Quad, 0)
	for row := 0; row < g.Lines; row++ {
		for col := 0; col < g.Cols; col++ {
			q, ok := g.LeftQuadFromPoint(row, col)
			if ok == true {
				quads = append(quads, q)
			}
			q, ok = g.DownQuadFromPoint(row, col)
			if ok == true {
				quads = append(quads, q)
			}
			q, ok = g.DownLeftQuadFromPoint(row, col)
			if ok == true {
				quads = append(quads, q)
			}
			q, ok = g.UpLeftQuadFromPoint(row, col)
			if ok == true {
				quads = append(quads, q)
			}
		}
	}
	return quads
}

func (g *Grid) LeftQuadFromPoint(row, col int) (q Quad, ok bool) {
	if col+3 >= g.Cols {
		return q, false
	}
	q[0] = g.Data[row][col]
	q[1] = g.Data[row][col+1]
	q[2] = g.Data[row][col+2]
	q[3] = g.Data[row][col+3]
	return q, true
}

func (g *Grid) DownQuadFromPoint(row, col int) (q Quad, ok bool) {
	if row+3 >= g.Lines {
		return q, false
	}
	q[0] = g.Data[row][col]
	q[1] = g.Data[row+1][col]
	q[2] = g.Data[row+2][col]
	q[3] = g.Data[row+3][col]
	return q, true
}

func (g *Grid) UpLeftQuadFromPoint(row, col int) (q Quad, ok bool) {
	if (row-3 < 0) || (col+3 >= g.Cols) {
		return q, false
	}
	q[0] = g.Data[row][col]
	q[1] = g.Data[row-1][col+1]
	q[2] = g.Data[row-2][col+2]
	q[3] = g.Data[row-3][col+3]
	return q, true
}

func (g *Grid) DownLeftQuadFromPoint(row, col int) (q Quad, ok bool) {
	if (row+3 >= g.Lines) || (col+3 >= g.Cols) {
		return q, false
	}
	q[0] = g.Data[row][col]
	q[1] = g.Data[row+1][col+1]
	q[2] = g.Data[row+2][col+2]
	q[3] = g.Data[row+3][col+3]
	return q, true
}
