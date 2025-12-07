package main

import (
        "fmt"
        "math/rand"
)

func printBoard(board [9]string) {
    fmt.Println(board[0], "|", board[1], "|", board[2])
    fmt.Println(board[3], "|", board[4], "|", board[5])
    fmt.Println(board[6], "|", board[7], "|", board[8])
}

func playerMove(board *[9]string, player string) {
    var boardPos int
    for true {
        fmt.Print("Enter a board position (0-8): ")
        fmt.Scanln(&boardPos)
        if board[boardPos] != " " {
            fmt.Print("Position already filled, please choose another one: ")
        } else {
            break
        }
    }
    
    board[boardPos] = player 
}

func computerMove(board *[9]string, computer string) {
    var boardPos = rand.Intn(9)
    for true {
        if board[boardPos] != " " {
            boardPos = rand.Intn(9)
        } else {
            board[boardPos] = computer
            break
        } 
    }
}

func detectWinLoss(board [9]string, player string, computer string) string {
    winningBoards := [][]int {
        {0, 1, 2}, {3, 4, 5}, {6, 7, 8},
        {0, 3, 6}, {1, 4, 7}, {2, 5, 8},
        {0, 4, 8}, {2, 4, 6},
    }
    
    for _, winningBoard := range winningBoards {
        if board[winningBoard[0]] != " " && board[winningBoard[0]] == board[winningBoard[1]] && board[winningBoard[1]] == board[winningBoard[2]] {
            return board[winningBoard[0]]
        }
    }
    return " "
}

func detectTie(board [9]string) bool {
    numFilled := 0
    for _, symbol := range board {
        if symbol != " " {
            numFilled++
        }
    }
    if numFilled == 9 {
        return true
    }
    return false
}

func main() {
    var board [9]string = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
    var player = "X"
    var computer = "O"
    
    for true {
        playerMove(&board, player)
        printBoard(board)
        
        if detectWinLoss(board, player, computer) != " " || detectTie(board) {
            break
        }
        
        fmt.Println("Computer's turn...")
        computerMove(&board, computer)
        printBoard(board)
        
        if detectWinLoss(board, player, computer) != " " || detectTie(board) {
            break
        } 
    }
    
    if detectWinLoss(board, player, computer) != " " {
        winner := detectWinLoss(board, player, computer)
        if winner == "X" {
            fmt.Println("You won!!!")
        } else {
            fmt.Println("You Lost :(")
        }
        return
    }
    
    if detectTie(board) {
        fmt.Println("It was a tie.")
        return
    }
}
