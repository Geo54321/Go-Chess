package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TwiN/go-color"
)

type pieceType int
type playerColor int
type boardColor bool

/* Enums */

const (
	Empty pieceType = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)

const (
	Blank playerColor = iota
	White
	Black
)

const (
	Board_White boardColor = true
	Board_Black boardColor = false
)

func (p pieceType) String() string {
	switch p {
	case Pawn:
		return "p"
	case Rook:
		return "r"
	case Knight:
		return "k"
	case Bishop:
		return "b"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Empty:
		return "."
	}
	return "??? :)"
}

func (c playerColor) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	case Blank:
		return ""
	}
	return "??? :)"
}

func (bc boardColor) String() string {
	switch bc {
	case Board_White:
		return "White"
	case Board_Black:
		return "Black"
	}
	return "??? :)"
}

/* Structs */
type Piece struct {
	pieceType pieceType
	player    playerColor
	rank      int
	file      int
	firstMove bool
}

func define_piece(t pieceType, p playerColor, r int, f int) Piece {
	var piece Piece
	piece.pieceType = t
	piece.player = p
	piece.rank = r
	piece.file = f
	piece.firstMove = true

	return piece
}

/* Functions */
func initialize_board() [8][8]Piece {
	board := [8][8]Piece{}
	for r := 0; r < len(board); r++ {
		for f := 0; f < len(board[r]); f++ {
			board[r][f] = Piece{}
		}
	}

	return board
}

func fill_board(board [8][8]Piece) [8][8]Piece {
	// White Player
	board[0][0] = define_piece(Rook, White, 0, 0)
	board[0][1] = define_piece(Knight, White, 0, 1)
	board[0][2] = define_piece(Bishop, White, 0, 2)
	board[0][3] = define_piece(King, White, 0, 3)
	board[0][4] = define_piece(Queen, White, 0, 4)
	board[0][5] = define_piece(Bishop, White, 0, 5)
	board[0][6] = define_piece(Knight, White, 0, 6)
	board[0][7] = define_piece(Rook, White, 0, 7)

	for p := 0; p < len(board[1]); p++ {
		board[1][p] = define_piece(Pawn, White, 1, p)
	}

	// Black Player
	board[7][0] = define_piece(Rook, Black, 7, 0)
	board[7][1] = define_piece(Knight, Black, 7, 1)
	board[7][2] = define_piece(Bishop, Black, 7, 2)
	board[7][3] = define_piece(King, Black, 7, 3)
	board[7][4] = define_piece(Queen, Black, 7, 4)
	board[7][5] = define_piece(Bishop, Black, 7, 5)
	board[7][6] = define_piece(Knight, Black, 7, 6)
	board[7][7] = define_piece(Rook, Black, 7, 7)

	for p := 0; p < len(board[6]); p++ {
		board[6][p] = define_piece(Pawn, Black, 6, p)
	}

	return board
}

func is_move(space [2]int, moves [][3]int) (bool, int) {
	isHighlight := false
	moveIndex := 0
	for s := 0; s < len(moves); s++ {
		if moves[s][0] == space[0] && moves[s][1] == space[1] {
			isHighlight = true
			moveIndex = s
		}
	}

	return isHighlight, moveIndex
}

func print_board(board [8][8]Piece, moves [][3]int, currentPiece Piece) {
	println("    A B C D E F G H")
	println("   ----------------")
	for r := len(board) - 1; r >= 0; r-- {
		print(r, " | ")
		for f := len(board[r]) - 1; f >= 0; f-- {
			isMove, moveIndex := is_move([2]int{r, f}, moves)
			if isMove {
				if currentPiece.player != board[r][f].player && board[r][f].player != Blank {
					// Attack Moves
					if moveIndex > 9 {
						print(color.InWhiteOverRed(moveIndex))
					} else {
						print(color.InWhiteOverRed(moveIndex) + " ")
					}
				} else {
					// Normal Moves
					if moveIndex > 9 {
						print(color.InCyanOverBlue(moveIndex))
					} else {
						print(color.InCyanOverBlue(moveIndex) + " ")
					}
				}
			} else if r == currentPiece.rank && f == currentPiece.file && currentPiece.pieceType != Empty {
				// Current Piece Highlight
				if currentPiece.player == White {
					print(color.InWhiteOverGreen(board[r][f].pieceType) + " ")
				} else {
					print(color.InCyanOverGreen(board[r][f].pieceType) + " ")
				}
			} else {
				// Other Pieces
				if board[r][f].player == White {
					print(color.InWhite(board[r][f].pieceType) + " ")
				} else {
					print(color.InCyan(board[r][f].pieceType) + " ")
				}
			}
		}
		println(" ")
	}
}

func get_input(prompt string) int {
	println(prompt, ":")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		log.Fatal(err)
	}

	return choice
}

func get_space_format(space [2]int) string {
	var pos string
	switch space[1] {
	case 0:
		pos = "h"
	case 1:
		pos = "g"
	case 2:
		pos = "f"
	case 3:
		pos = "e"
	case 4:
		pos = "d"
	case 5:
		pos = "c"
	case 6:
		pos = "b"
	case 7:
		pos = "a"
	default:
		pos = "lmao"
	}
	pos += strconv.Itoa(space[0])
	return pos
}

func get_valid_pieces(board [8][8]Piece, player playerColor, lastMove [3]int, isInCheck bool, blockMoves [][2]int, enemyMoves [][3]int) []Piece {
	validPieces := make([]Piece, 0)
	if isInCheck {
		for r := len(board) - 1; r >= 0; r-- {
			for f := len(board[r]) - 1; f >= 0; f-- {
				if board[r][f].player == player {
					if board[r][f].pieceType == King {
						// King - Check if king can escape the check
						kingMoves := get_moves(board[r][f], board, lastMove, isInCheck)
						kingMoves = get_valid_moves(board, kingMoves, board[r][f], isInCheck, enemyMoves, blockMoves, lastMove)
						if len(kingMoves) > 0 {
							validPieces = append(validPieces, board[r][f])
						}
					} else {
						// Not King - Check if it can block the checking piece
						validPiece := false
						currMoves := get_moves(board[r][f], board, lastMove, isInCheck)
						currMoves = get_valid_moves(board, currMoves, board[r][f], isInCheck, enemyMoves, blockMoves, lastMove)
						for m := 0; m < len(currMoves); m++ {
							for b := 0; b < len(blockMoves); b++ {
								if currMoves[m][0] == blockMoves[b][0] && currMoves[m][1] == blockMoves[b][1] {
									validPiece = true
								}
							}
						}
						if validPiece {
							validPieces = append(validPieces, board[r][f])
						}
					}
				}
			}
		}

	} else {
		for r := len(board) - 1; r >= 0; r-- {
			for f := len(board[r]) - 1; f >= 0; f-- {
				if board[r][f].player == player && len(get_moves(board[r][f], board, lastMove, isInCheck)) != 0 {
					validPieces = append(validPieces, board[r][f])
				}
			}
		}
	}

	return validPieces
}

func select_piece(player playerColor, redo bool, pieces []Piece) (Piece, bool) {
	if !redo {
		println("\n  === Available Pieces === ")
		for p := 0; p < len(pieces); p++ {
			fmt.Printf("%v: \t%v @ %v \t\t", p, pieces[p].pieceType.String(), get_space_format([2]int{pieces[p].rank, pieces[p].file}))
			if p%2 == 1 {
				println()
			}
		}
		println()
	} else {
		println("Invalid piece. Please select another one.\n")
	}

	choice := get_input("Select a piece to move")

	if choice >= 0 && choice < len(pieces) {
		return pieces[choice], true
	}
	return Piece{}, false
}

func get_valid_moves(board [8][8]Piece, moves [][3]int, piece Piece, isInCheck bool, enemyMoves [][3]int, blockMoves [][2]int, lastMove [3]int) [][3]int {
	validMoves := make([][3]int, 0)
	for m := 0; m < len(moves); m++ {
		if (moves[m][0] >= 0 && moves[m][0] < 8) && (moves[m][1] >= 0 && moves[m][1] < 8) {
			if board[moves[m][0]][moves[m][1]].player != piece.player {
				if piece.pieceType != King {
					if isInCheck {
						for b := 0; b < len(blockMoves); b++ {
							if moves[m][0] == blockMoves[b][0] && moves[m][1] == blockMoves[b][1] {
								validMoves = append(validMoves, moves[m])
							}
						}
					} else {
						validMoves = append(validMoves, moves[m])
					}
				} else {
					validMoves = append(validMoves, moves[m])
				}
			}
		}
	}
	return validMoves
}

func select_promotion(piece Piece) (Piece, bool) {
	isValid, newPiece := true, Piece{}
	println("=== Available Promotions  ===")
	println("1: \tKnight \t\t 2:\tBishop")
	println("3: \tRook \t\t 4: \tQueen")

	choice := get_input("Select what to promote the pawn to")
	println(choice)

	switch choice {
	case 1:
		newPiece = define_piece(Knight, piece.player, piece.rank, piece.file)
	case 2:
		newPiece = define_piece(Bishop, piece.player, piece.rank, piece.file)
	case 3:
		newPiece = define_piece(Rook, piece.player, piece.rank, piece.file)
	case 4:
		newPiece = define_piece(Queen, piece.player, piece.rank, piece.file)
	default:
		newPiece = Piece{}
		isValid = false
	}

	return newPiece, isValid
}

func get_moves(piece Piece, board [8][8]Piece, lastMove [3]int, isInCheck bool) [][3]int {
	moves := make([][3]int, 0)
	switch piece.pieceType {
	case Pawn:
		// Pawn Movement
		if piece.player == White {
			// White Forward Moves
			if board[piece.rank+1][piece.file].player == Blank {
				if piece.rank+1 == 7 {
					// Pawn promotion
					moves = append(moves, [3]int{piece.rank + 1, piece.file, 4})
				} else {
					moves = append(moves, [3]int{piece.rank + 1, piece.file, 0})
				}
			}
			if piece.firstMove && board[piece.rank+1][piece.file].player == Blank && board[piece.rank+2][piece.file].player == Blank {
				// First move check
				moves = append(moves, [3]int{piece.rank + 2, piece.file, 0})
			}

			// White Attack Moves
			if piece.file+1 < 8 {
				if board[piece.rank+1][piece.file+1].player == Black {
					if piece.rank+1 == 7 {
						// Pawn promotion
						moves = append(moves, [3]int{piece.rank + 1, piece.file + 1, 4})
					} else {
						moves = append(moves, [3]int{piece.rank + 1, piece.file + 1, 0})
					}
				}
			}
			if piece.file-1 >= 0 {
				if board[piece.rank+1][piece.file-1].player == Black {
					if piece.rank+1 == 7 {
						// Pawn promotion
						moves = append(moves, [3]int{piece.rank + 1, piece.file - 1, 4})
					} else {
						moves = append(moves, [3]int{piece.rank + 1, piece.file - 1, 0})
					}
				}
			}
		} else if piece.player == Black {
			// Black Forward Moves
			if board[piece.rank-1][piece.file].player == Blank {
				if piece.rank-1 == 0 {
					// Pawn promotion
					moves = append(moves, [3]int{piece.rank - 1, piece.file, 4})
				} else {
					moves = append(moves, [3]int{piece.rank - 1, piece.file, 0})
				}
			}
			if piece.firstMove && board[piece.rank-1][piece.file].player == Blank && board[piece.rank-2][piece.file].player == Blank {
				// First move check
				moves = append(moves, [3]int{piece.rank - 2, piece.file, 0})
			}

			// Black Attack Moves
			if piece.file+1 < 8 {
				if board[piece.rank-1][piece.file+1].player == White {
					if piece.rank-1 == 0 {
						// Pawn promotion
						moves = append(moves, [3]int{piece.rank - 1, piece.file + 1, 4})
					} else {
						moves = append(moves, [3]int{piece.rank - 1, piece.file + 1, 0})
					}
				}
			}
			if piece.file-1 >= 0 {
				if board[piece.rank-1][piece.file-1].player == White {
					if piece.rank-1 == 0 {
						// Pawn promotion
						moves = append(moves, [3]int{piece.rank - 1, piece.file - 1, 4})
					} else {
						moves = append(moves, [3]int{piece.rank - 1, piece.file - 1, 0})
					}
				}
			}
		}
		// En passant
		if piece.player == White && lastMove[0] == 4 && board[lastMove[0]][lastMove[1]].pieceType == Pawn && piece.rank == 4 && (piece.file == lastMove[1]+1 || piece.file == lastMove[1]-1) {
			moves = append(moves, [3]int{lastMove[0] + 1, lastMove[1], 1})
		}
		if piece.player == Black && lastMove[0] == 3 && board[lastMove[0]][lastMove[1]].pieceType == Pawn && piece.rank == 3 && (piece.file == lastMove[1]+1 || piece.file == lastMove[1]-1) {
			moves = append(moves, [3]int{lastMove[0] - 1, lastMove[1], 1})
		}
	case Rook:
		// Rook Movement
		for up := piece.rank + 1; up < len(board); up++ {
			moves = append(moves, [3]int{up, piece.file, 0})
			if board[up][piece.file].player != Blank {
				break
			}
		}
		for down := piece.rank - 1; down >= 0; down-- {
			moves = append(moves, [3]int{down, piece.file, 0})
			if board[down][piece.file].player != Blank {
				break
			}
		}
		for left := piece.file + 1; left < len(board[piece.rank]); left++ {
			moves = append(moves, [3]int{piece.rank, left, 0})
			if board[piece.rank][left].player != Blank {
				break
			}
		}
		for right := piece.file - 1; right >= 0; right-- {
			moves = append(moves, [3]int{piece.rank, right, 0})
			if board[piece.rank][right].player != Blank {
				break
			}
		}
	case Knight:
		// Knight Movement
		moves = append(moves, [3]int{piece.rank + 1, piece.file + 2, 0})
		moves = append(moves, [3]int{piece.rank + 2, piece.file + 1, 0})
		moves = append(moves, [3]int{piece.rank + 1, piece.file - 2, 0})
		moves = append(moves, [3]int{piece.rank + 2, piece.file - 1, 0})
		moves = append(moves, [3]int{piece.rank - 1, piece.file + 2, 0})
		moves = append(moves, [3]int{piece.rank - 2, piece.file + 1, 0})
		moves = append(moves, [3]int{piece.rank - 1, piece.file - 2, 0})
		moves = append(moves, [3]int{piece.rank - 2, piece.file - 1, 0})
	case Bishop:
		// Bishop Movement
		upleft, upright, downleft, downright := true, true, true, true
		for count := 1; count < len(board)-1; count++ {
			if upleft {
				trank, tfile := piece.rank+count, piece.file+count
				if trank < 8 && tfile < 8 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						upleft = false
					}
				} else {
					upleft = false
				}
			}
			if upright {
				trank, tfile := piece.rank-count, piece.file+count
				if trank >= 0 && tfile < 8 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						upright = false
					}
				} else {
					upright = false
				}
			}
			if downleft {
				trank, tfile := piece.rank+count, piece.file-count
				if (trank) < 8 && (tfile) >= 0 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						downleft = false
					}
				} else {
					downleft = false
				}
			}
			if downright {
				trank, tfile := piece.rank-count, piece.file-count
				if (trank) >= 0 && (tfile) >= 0 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, piece.file - count, 0})
					} else {
						moves = append(moves, [3]int{trank, piece.file - count, 0})
						downright = false
					}
				} else {
					downright = false
				}
			}
		}
	case Queen:
		// Queen Movement
		// Rook like movement
		for up := piece.rank + 1; up < len(board); up++ {
			moves = append(moves, [3]int{up, piece.file, 0})
			if board[up][piece.file].player != Blank {
				break
			}
		}
		for down := piece.rank - 1; down >= 0; down-- {
			moves = append(moves, [3]int{down, piece.file, 0})
			if board[down][piece.file].player != Blank {
				break
			}
		}
		for left := piece.file + 1; left < len(board[piece.rank]); left++ {
			moves = append(moves, [3]int{piece.rank, left, 0})
			if board[piece.rank][left].player != Blank {
				break
			}
		}
		for right := piece.file - 1; right >= 0; right-- {
			moves = append(moves, [3]int{piece.rank, right, 0})
			if board[piece.rank][right].player != Blank {
				break
			}
		}

		// Bishop like movement
		upleft, upright, downleft, downright := true, true, true, true
		for count := 1; count < len(board)-1; count++ {
			if upleft {
				trank, tfile := piece.rank+count, piece.file+count
				if trank < 8 && tfile < 8 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						upleft = false
					}
				} else {
					upleft = false
				}
			}
			if upright {
				trank, tfile := piece.rank-count, piece.file+count
				if trank >= 0 && tfile < 8 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						upright = false
					}
				} else {
					upright = false
				}
			}
			if downleft {
				trank, tfile := piece.rank+count, piece.file-count
				if (trank) < 8 && (tfile) >= 0 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, tfile, 0})
					} else {
						moves = append(moves, [3]int{trank, tfile, 0})
						downleft = false
					}
				} else {
					downleft = false
				}
			}
			if downright {
				trank, tfile := piece.rank-count, piece.file-count
				if (trank) >= 0 && (tfile) >= 0 {
					if board[trank][tfile].player == Blank {
						moves = append(moves, [3]int{trank, piece.file - count, 0})
					} else {
						moves = append(moves, [3]int{trank, piece.file - count, 0})
						downright = false
					}
				} else {
					downright = false
				}
			}
		}
	case King:
		// King Movement
		kingMoves := make([][3]int, 0)
		kingMoves = append(kingMoves, [3]int{piece.rank + 1, piece.file, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank + 1, piece.file + 1, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank, piece.file + 1, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank - 1, piece.file + 1, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank - 1, piece.file, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank - 1, piece.file - 1, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank, piece.file - 1, 0})
		kingMoves = append(kingMoves, [3]int{piece.rank + 1, piece.file - 1, 0})

		// Castle Movement - King Side
		if piece.firstMove && board[piece.rank][7].firstMove && !isInCheck {
			if board[piece.rank][piece.file+1].player == Blank && board[piece.rank][piece.file+2].player == Blank {
				kingMoves = append(moves, [3]int{piece.rank, piece.file + 2, 2})
			}
		}

		// Castle Movement - Queen Side
		if piece.firstMove && board[piece.rank][0].firstMove && !isInCheck {
			if board[piece.rank][piece.file-1].player == Blank && board[piece.rank][piece.file-2].player == Blank && board[piece.rank][piece.file-3].player == Blank {
				kingMoves = append(moves, [3]int{piece.rank, piece.file - 2, 3})
			}
		}

		// Confirm the move doesn't leave the king in check
		for k := 0; k < len(kingMoves); k++ {
			tempBoard := board
			enemyColor := Blank
			if piece.player == White {
				enemyColor = Black
			} else {
				enemyColor = White
			}

			if kingMoves[k][0] < 8 && kingMoves[k][0] >= 0 && kingMoves[k][1] < 8 && kingMoves[k][1] >= 0 {
				tempBoard = move_piece(tempBoard, piece, kingMoves[k])
				stillInCheck, _ := check_check(tempBoard, enemyColor)

				if !stillInCheck {
					moves = append(moves, kingMoves[k])
				}
			}
		}
	}
	return moves
}

func select_move(moves [][3]int) ([3]int, bool) {
	println("\n  === Available Moves ===")
	for m := 0; m < len(moves); m++ {
		print(m, ": to ", get_space_format([2]int{moves[m][0], moves[m][1]}), "\t")
		if m%2 == 1 {
			println()
		}
	}
	move := get_input("\nSelect a move to make")
	println()
	if move >= 0 && move < len(moves) {
		return moves[move], true
	}
	return [3]int{}, false
}

func move_piece(board [8][8]Piece, piece Piece, move [3]int) [8][8]Piece {
	if move[2] == 1 {
		// En Passat pawn removal
		board[piece.rank][move[1]] = Piece{}
	}
	if move[2] == 2 {
		// Castle - King Side
		board[piece.rank][7] = Piece{}
		board[piece.rank][piece.file+1] = define_piece(Rook, piece.player, piece.rank, piece.file+1)
		board[piece.rank][piece.file+1].firstMove = false
	}
	if move[2] == 3 {
		// Castle - Queen Side
		board[piece.rank][0] = Piece{}
		board[piece.rank][piece.file-1] = define_piece(Rook, piece.player, piece.rank, piece.file-1)
		board[piece.rank][piece.file-1].firstMove = false
	}
	if move[2] == 4 {
		// Pawn promotion
		isValidPromotion, flag, newPiece := false, false, Piece{}
		for {
			if isValidPromotion {
				break
			}
			if flag {
				println("ERROR: Invalid promotion. Please choose another promotion.\n")
			}
			newPiece, isValidPromotion = select_promotion(piece)
			flag = true
		}
		piece = newPiece
	}
	board[piece.rank][piece.file] = Piece{}
	piece.rank = move[0]
	piece.file = move[1]

	if piece.firstMove {
		piece.firstMove = false
	}

	board[move[0]][move[1]] = piece
	return board
}

func do_turn(board [8][8]Piece, player playerColor, lastMove [3]int, isInCheck bool, blockMoves [][2]int, enemyMoves [][3]int) ([8][8]Piece, bool, [3]int, bool, [][2]int) {
	isValid, flag := false, false
	var pieceChoice Piece
	var moveChoice [3]int

	// Select Piece
	println("\n\n===", player.String(), "Turn ===\n\n")

	if isInCheck {
		println("!!! YOU ARE IN CHECK !!! \n")
	}

	print_board(board, make([][3]int, 0), Piece{})
	for {
		if isValid {
			break
		}
		if flag {
			println("ERROR: Invalid piece. Please choose another piece.\n")
		}
		pieceOptions := get_valid_pieces(board, player, lastMove, isInCheck, blockMoves, enemyMoves)
		pieceChoice, isValid = select_piece(player, false, pieceOptions)
		flag = true
	}

	// Display Moves or Redo turn
	moveOptions := get_moves(pieceChoice, board, lastMove, isInCheck)
	moveOptions = get_valid_moves(board, moveOptions, pieceChoice, isInCheck, enemyMoves, blockMoves, lastMove)
	if len(moveOptions) != 0 {
		print_board(board, moveOptions, pieceChoice)

		// Select Move
		isValid, flag = false, false
		for {
			if isValid {
				break
			}
			if flag {
				println("ERROR: Invalid move. Please choose another move.\n")
			}
			moveChoice, isValid = select_move(moveOptions)
			flag = true
		}

		// Move Piece
		board = move_piece(board, pieceChoice, moveChoice)
		lastMove = moveChoice
		print_board(board, make([][3]int, 0), pieceChoice)

		isCheck, blockMoves := check_check(board, pieceChoice.player)

		return board, true, lastMove, isCheck, blockMoves
	}

	return board, false, lastMove, false, blockMoves
}

func check_check(board [8][8]Piece, player playerColor) (bool, [][2]int) {
	isCheck := false
	kingSpace := [2]int{}
	allMoves := [][3]int{}
	movePieces := make([]Piece, 0)
	checkerPiece := Piece{}
	blockSpaces := make([][2]int, 0)

	for r := len(board) - 1; r >= 0; r-- {
		for f := len(board[r]) - 1; f >= 0; f-- {
			if board[r][f].pieceType == King && board[r][f].player != player {
				// Found enemy king
				kingSpace = [2]int{r, f}
			} else if board[r][f].player == player && board[r][f].pieceType != King {
				// Add moves to catalog of current available moves
				tempMoves := get_moves(board[r][f], board, [3]int{0, 0, 0}, false)
				for m := 0; m < len(tempMoves); m++ {
					allMoves = append(allMoves, tempMoves[m])
					movePieces = append(movePieces, board[r][f])
				}

			}
		}
	}

	for m := 0; m < len(allMoves); m++ {
		if allMoves[m][0] == kingSpace[0] && allMoves[m][1] == kingSpace[1] {
			// Checks all moves against kings location for check
			isCheck = true
			checkerPiece = movePieces[m]
		}
	}

	switch checkerPiece.pieceType {
	case Pawn, Knight:
		blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank, checkerPiece.file})
	case Bishop:
		count := checkerPiece.rank - kingSpace[0]
		if count < 0 {
			count *= -1
		}
		if checkerPiece.rank < kingSpace[0] && checkerPiece.file < kingSpace[1] {
			for s := 0; s <= count; s++ {
				blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file + s})
			}
		} else if checkerPiece.rank < kingSpace[0] && checkerPiece.file > kingSpace[1] {
			for s := 0; s <= count; s++ {
				blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file - s})
			}
		} else if checkerPiece.rank > kingSpace[0] && checkerPiece.file < kingSpace[1] {
			for s := 0; s <= count; s++ {
				blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file + s})
			}
		} else if checkerPiece.rank > kingSpace[0] && checkerPiece.file > kingSpace[1] {
			for s := 0; s <= count; s++ {
				blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file - s})
			}
		}
	case Rook:
		if checkerPiece.rank == kingSpace[0] {
			count := checkerPiece.file - kingSpace[1]
			if count < 0 {
				count *= -1
			}
			if checkerPiece.file < kingSpace[1] {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank, checkerPiece.file + s})
				}
			} else {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank, checkerPiece.file - s})
				}
			}
		} else {
			count := checkerPiece.rank - kingSpace[0]
			if count < 0 {
				count *= -1
			}
			if checkerPiece.rank <= kingSpace[0] {
				for s := 0; s < count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file})
				}
			} else {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file})
				}
			}
		}
	case Queen:
		if checkerPiece.rank != kingSpace[0] && checkerPiece.file != kingSpace[1] {
			// Bishop Moves
			count := checkerPiece.rank - kingSpace[0]
			if count < 0 {
				count *= -1
			}
			if checkerPiece.rank < kingSpace[0] && checkerPiece.file < kingSpace[1] {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file + s})
				}
			} else if checkerPiece.rank < kingSpace[0] && checkerPiece.file > kingSpace[1] {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file - s})
				}
			} else if checkerPiece.rank > kingSpace[0] && checkerPiece.file < kingSpace[1] {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file + s})
				}
			} else if checkerPiece.rank > kingSpace[0] && checkerPiece.file > kingSpace[1] {
				for s := 0; s <= count; s++ {
					blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file - s})
				}
			}
		} else {
			// Rook Moves
			if checkerPiece.rank == kingSpace[0] {
				count := checkerPiece.file - kingSpace[1]
				if count < 0 {
					count *= -1
				}
				if checkerPiece.file < kingSpace[1] {
					for s := 0; s <= count; s++ {
						blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank, checkerPiece.file + s})
					}
				} else {
					for s := 0; s <= count; s++ {
						blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank, checkerPiece.file - s})
					}
				}
			} else {
				count := checkerPiece.rank - kingSpace[0]
				if count < 0 {
					count *= -1
				}
				if checkerPiece.rank < kingSpace[0] {
					for s := 0; s <= count; s++ {
						blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank + s, checkerPiece.file})
					}
				} else {
					for s := 0; s <= count; s++ {
						blockSpaces = append(blockSpaces, [2]int{checkerPiece.rank - s, checkerPiece.file})
					}
				}
			}
		}
	default:

	}

	return isCheck, blockSpaces
}

func main() {
	// Game Setup
	Board := initialize_board() // TODO : This can be refactors to remove the space coloring if i don't end up using space colors
	Board = fill_board(Board)
	isTurnValid, flag, isCheck := false, false, false
	player := White
	lastMove := [3]int{}
	blockMoves := make([][2]int, 0)
	enemyMoves := make([][3]int, 0)

	// Game Loop
	for {
		isTurnValid, flag = false, false
		for {
			if isTurnValid {
				break
			}
			if flag {
				println("ERROR: Invalid turn. Please choose another piece.\n")
			}
			Board, isTurnValid, lastMove, isCheck, blockMoves = do_turn(Board, player, lastMove, isCheck, blockMoves, enemyMoves)
			flag = true
		}

		if player == White {
			player = Black
		} else {
			player = White
		}

		// Index all enemy moves
		for r := len(Board) - 1; r >= 0; r-- {
			for f := len(Board[r]) - 1; f >= 0; f-- {
				if Board[r][f].player != player && Board[r][f].player != Blank {
					tempMoves := get_moves(Board[r][f], Board, lastMove, false)
					enemyMoves = append(enemyMoves, get_valid_moves(Board, tempMoves, Board[r][f], false, make([][3]int, 0), make([][2]int, 0), lastMove)...)
				}
			}
		}

		if len(get_valid_pieces(Board, player, lastMove, isCheck, blockMoves, enemyMoves)) == 0 {
			// Current player has won if this is reached
			break
		}
	}

	if player == White {
		player = Black
	} else {
		player = White
	}
	println("\n\n\n\n\n=== CONGRATS ON THE WIN: ", player.String(), "===")

	// TODO : Fix weird issues with Get_valid_pieces() letting no valid pieces through
	// TODO : Fix castle moves not being available when they should be
	// TODO : Add ability to cancel piece selection
	// TODO : Add ability to choose pieces and mvoes by space instead of the index
}
