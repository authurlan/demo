package main

import (
    "flag"
    "fmt"
    "strconv"
)

func human_to_machine(rows *IntVec99) {
    for i, row := range rows {
        for j, v := range row {
            if v == 0 {
                rows[i][j] = 0x01FF
                continue
            }

            rows[i][j] = 0x0000
            for _, c := range strconv.Itoa(v) {
                switch c {
                case '1':
                    rows[i][j] |= 0x0001
                case '2':
                    rows[i][j] |= 0x0002
                case '3':
                    rows[i][j] |= 0x0004
                case '4':
                    rows[i][j] |= 0x0008
                case '5':
                    rows[i][j] |= 0x0010
                case '6':
                    rows[i][j] |= 0x0020
                case '7':
                    rows[i][j] |= 0x0040
                case '8':
                    rows[i][j] |= 0x0080
                case '9':
                    rows[i][j] |= 0x0100
                }
            }
        }
    }
}

func machine_to_human(rows *IntVec99) {
    for i, row := range rows {
        for j, v := range row {
            t := ""
            for k := 0; k < 9; k++ {
                if (v >> k) & 0x0001 == 0x0001 {
                    t += strconv.Itoa(k + 1)
                }
            }

            rows[i][j], _ = strconv.Atoi(t)
        }
    }
}

func sudoku_print(rows *IntVec99) {
    fmt.Println("-------------------------------------------------------------------------")
    for i, row := range rows {
        fmt.Print("| ")
        for j, v := range row {
            fmt.Print(v)
            if j % 3 == 2 {
                fmt.Print("\t| ")
            } else {
                fmt.Print("\t")
            }
        }

        fmt.Println()
        if i % 3 == 2 {
            fmt.Println("-------------------------------------------------------------------------")
        }
    }
}

func sudoku_confirm(v int) bool {
    switch v {
    case 0x0001:
        fallthrough
    case 0x0002:
        fallthrough
    case 0x0004:
        fallthrough
    case 0x0008:
        fallthrough
    case 0x0010:
        fallthrough
    case 0x0020:
        fallthrough
    case 0x0040:
        fallthrough
    case 0x0080:
        fallthrough
    case 0x0100:
        return true
    default:
        return false
    }
}

func sudoku_bitwise(v int, rows *IntVec99, i, j int) {
    if sudoku_confirm(v) {
        rows[i][j] &= ^v
    }
}

func sudoku_assign(v int, rows *IntVec99, i, j int) {
    if sudoku_confirm(v) {
        rows[i][j] = v
    }
}

func main() {
    inputPath := flag.String("i", "", "input path")
    outputPath := flag.String("o", "", "output path")
    flag.Parse()

    var rows IntVec99
    rows.Load(*inputPath)
    sudoku_print(&rows)
    human_to_machine(&rows)

    for i, row := range rows {
        for j, v := range row {
            if sudoku_confirm(v) {
                continue
            }

            x, y, z := v, v, v
            for k := 0; k < 9; k++ {
                if j != k {
                    sudoku_bitwise(row[k], &rows, i, j)
                    x &= rows[i][k] ^ x
                }

                if i != k {
                    sudoku_bitwise(rows[k][j], &rows, i, j)
                    y &= rows[k][j] ^ y
                }

                p, q := i / 3 * 3 + k / 3, j / 3 * 3 + k % 3
                if p != i || q != j {
                    sudoku_bitwise(rows[p][q], &rows, i, j)
                    z &= rows[p][q] ^ z
                }
            }

            sudoku_assign(x, &rows, i, j)
            sudoku_assign(y, &rows, i, j)
            sudoku_assign(z, &rows, i, j)
        }
    }

    machine_to_human(&rows)
    sudoku_print(&rows)
    rows.Dump(*outputPath)
}
