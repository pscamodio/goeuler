package main

import (
    "fmt"
    "strings"
)

const (
    gridSize int = 1001
)

type Row []int
type Grid []Row
type Point [2]int

func (g Grid) String() string {
    lines := make([]string, gridSize)
    for idx, row := range g {
        vals := make([]string, gridSize)
        for jdx, val := range row {
            vals[jdx] = fmt.Sprintf("%5d", val)
        }
        lines[idx] = strings.Join(vals, "")
    }
    return strings.Join(lines, "\n")
}

func (g Grid) SetPoint(p Point, val int) {
    g[p[0]][p[1]] = val
}

func (g Grid) GetPoint(p Point) int {
    return g[p[0]][p[1]]
}

func main() {
    g := make(Grid, gridSize)
    for idx := range g {
        g[idx] = make(Row, gridSize)
    }
    loops := (gridSize - 1) / 2
    p := Point{loops, loops}
    val := 1
    g.SetPoint(p, val)
    for loop := 0; loop < loops; loop++ {
        //Right
        for idx := 0; idx < 1+2*loop; idx++ {
            val++
            p[1] += 1
            g.SetPoint(p, val)
        }
        //Down
        for idx := 0; idx < 1+2*loop; idx++ {
            val++
            p[0] += 1
            g.SetPoint(p, val)
        }
        //Left
        for idx := 0; idx < 2+2*loop; idx++ {
            val++
            p[1] -= 1
            g.SetPoint(p, val)
        }
        //Up
        for idx := 0; idx < 2+2*loop; idx++ {
            val++
            p[0] -= 1
            g.SetPoint(p, val)
        }
    }
    for idx := 0; idx < 2*loops; idx++ {
        val++
        p[1] += 1
        g.SetPoint(p, val)
    }
    tot := 0
    for idx := 0; idx < gridSize; idx++ {
        tot += g.GetPoint(Point{idx, idx})
        tot += g.GetPoint(Point{idx, gridSize - idx - 1})
    }
    tot -= g.GetPoint(Point{loops, loops})
    fmt.Println(tot)
}
