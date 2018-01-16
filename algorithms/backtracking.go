package algorithms

import (
	"strings"
	"log"
	"strconv"
)

type QueenStruct struct{
	value int
	columnS int
	leftD int
	rightD int
}

var lDiagonal = [][]int{}
var rDiagonal = [][]int{}
var column = [][]int{}

func Solve8QueenProblem(n int) int{
	lDiagonal = buildLeftDiagonal(n)
	rDiagonal = buildRightDiagonal(n)
	column = buildColumn(n)
	board := buildEmptyBoard(n)
	return solve8Queen(0, board, 0, n)
}

func solve8Queen(solution int, board [][]QueenStruct, x int, n int) int {
	if x == n {
		printBoard(board, n)
		solution++
		return solution
	}
	for y:=0; y<n; y++{
		if !isPositionValid(board, n, x, y) {
			continue
		}
		q := QueenStruct{value:1, columnS:column[x][y], leftD: lDiagonal[x][y], rightD: rDiagonal[x][y]}
		board[x][y] = q
		solution = solve8Queen(solution, board, x+1, n)
		board[x][y].value = 0
		board[x][y].rightD = -1
		board[x][y].leftD = -1
		board[x][y].columnS = -1
	}
	return solution
}

func isPositionValid(board [][]QueenStruct, n int, x int, y int) bool{
    for i:=0; i<=x ;i++{
    	for j:=0; j<n ;j++{
    		if isEmpty(board[i][j]){
    			continue
			}else {
				q := board[i][j]
				if q.leftD == lDiagonal[x][y] || q.rightD == rDiagonal[x][y] || q.columnS == column[x][y] {
					return false
				}
			}
		}
	}
	return true
}

func isEmpty(q QueenStruct) bool{
	if q.value == 0 {
		return true
	}
	return false
}

func printBoard(board [][]QueenStruct, n int){
	strArr := []string{}
	log.Printf("[")
	for i:=0; i<n; i++{
		for j:=0; j<n; j++{
			strArr = append(strArr, strconv.Itoa(board[i][j].value))
		}
		log.Printf("%s", strings.Join(strArr,", "))
		strArr = []string{}
	}
	log.Printf("]")
}

func buildEmptyBoard(n int) [][]QueenStruct{
	column := make([][]QueenStruct , n)
	for i :=range column{
		column[i] = make([]QueenStruct, n)
		for j:=0; j<n; j++{
			column[i][j].value = 0
			column[i][j].rightD = -1
			column[i][j].leftD = -1
			column[i][j].columnS = -1
		}
	}
	return column
}

func buildColumn (n int) [][]int{
	column := make([][]int , n)
	for i :=range column{
		column[i] = make([]int, n)
	}
	for i:=0; i<n; i++{
		for j:=0; j<n; j++{
			column[i][j] = j
		}
	}
	return column
}

func buildLeftDiagonal(n int) [][]int{
	var leftDiagonal = make([][]int, n)
	for i :=range leftDiagonal{
		leftDiagonal[i] = make([]int, n)
	}
	offset := 0
	for i:=0; i<n; i++{
		for j:=0; j<n; j++{
			leftDiagonal[i][j] = j+offset
		}
		offset++
	}
	//log.Printf("%v", leftDiagonal)
	return leftDiagonal
}

func buildRightDiagonal(n int) [][]int{
	var rDiagonal = make([][]int, n)
	for i :=range rDiagonal{
		rDiagonal[i] = make([]int, n)
	}
	offset := 0
	for i:=n-1; i>=0; i--{
		for j:=0; j<n; j++{
			rDiagonal[i][j] = j+offset
		}
		offset++
	}
	//log.Printf("+v", rDiagonal)
	return rDiagonal
}