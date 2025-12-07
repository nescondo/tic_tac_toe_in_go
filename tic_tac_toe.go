package main

import (
        "fmt"
        "math/rand"
)

// Prints the current status of the board.
// Takes in a parameter 'board' that is an array of 9 strings.
// Returns nothing.
func printBoard(board [9]string) {
    fmt.Println(board[0], "|", board[1], "|", board[2]) // First row
    fmt.Println(board[3], "|", board[4], "|", board[5]) // Second row
    fmt.Println(board[6], "|", board[7], "|", board[8]) // Third row
}

// Handles player move functionality.
// Takes in the parameter 'board' as a pointer to the board and the 'player' string representation of the player's symbol.
// Returns nothing.
func playerMove(board *[9]string, player string) {
    var boardPos int
    for true { // No 'while' loops in Go - will use 'for true' and break when needed.
        fmt.Print("Enter a board position (0-8): ")
        fmt.Scanln(&boardPos)
        if board[boardPos] != " " { // Checks if the board is empty, if not, ask for another input.
            fmt.Print("Position already filled, please choose another one: ")
        } else { // If empty, break from the loop.
            break
        }
    }
    
    board[boardPos] = player // Assign the 'player' character (X) to the position on the board.
}

// Handles the 'computer' move functionality.
// Takes in the parameter 'board' as a pointer to the board and the 'computer' string representation of the computer's symbol.
// Returns nothing.
func computerMove(board *[9]string, computer string) {
    var boardPos = rand.Intn(9) // Randomly select a number between 0-8.
    for true {
        if board[boardPos] != " " { // Check if the board position randomly selected is not empty.
            boardPos = rand.Intn(9)
        } else { // If the board position randomly selected is empty, assign the 'computer' character (O) to the position on the board.
            board[boardPos] = computer
            break
        } 
    }
}

// Function to detect whether the player has won or lost.
// Takes in the parameter 'board' as a pointer to the board, the 'player' string representation of the player's symbol, and the 'computer' string representation of the computer's symbol.
// Returns a string to represent if the player (X), computer (O), or no one (" ") has won.
func detectWinLoss(board [9]string, player string, computer string) string {
    winningBoards := [][]int { // Represents all possible winning combinations.
        {0, 1, 2}, {3, 4, 5}, {6, 7, 8},
        {0, 3, 6}, {1, 4, 7}, {2, 5, 8},
        {0, 4, 8}, {2, 4, 6},
    }
    
    for _, winningBoard := range winningBoards { // Iterate through every winning combination.
        if board[winningBoard[0]] != " " && board[winningBoard[0]] == board[winningBoard[1]] && board[winningBoard[1]] == board[winningBoard[2]] { // Check each position within the current winning board combination for a match in the current board.
            return board[winningBoard[0]] // Return the winning symbol within the current board (X or O).
        }
    }
    return " " // Return an empty string if no winner.
}

// Function to detect a tie.
// Takes in the parameter 'board' as an array with 9 strings.
// Returns a boolean representing whether there is or is not a tie.
func detectTie(board [9]string) bool {
    numFilled := 0 // Keeps track of how many positions have been filled on the board.
    for _, symbol := range board { // Check each position on the board.
        if symbol != " " {
            numFilled++ // If a position is filled, increment by 1.
        }
    }
    if numFilled == 9 { // If all positions are filled, return true.
        return true
    }
    return false
}

func main() {
    var board [9]string = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}    

    var player string = "X" 
    computer := "O"
    
    for true {
        // Player move.
        playerMove(&board, player)
        printBoard(board)
       
        // Check if a win/loss or tie has occured.
        if detectWinLoss(board, player, computer) != " " || detectTie(board) {
            break
        }
        
        // Computer move.
        fmt.Println("Computer's turn...")
        computerMove(&board, computer)
        printBoard(board)
        
        // Check if a win/loss or tie has occured.
        if detectWinLoss(board, player, computer) != " " || detectTie(board) {
            break
        } 
    }
   
    // Once a win/loss or tie has occured, print the corresponding statement and exit the program. 
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
