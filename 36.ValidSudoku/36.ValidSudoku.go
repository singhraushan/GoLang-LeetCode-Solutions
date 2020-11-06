package main

import "fmt"

func main(){
	/*arr:= [][]byte{{'8','3','.','.','7','.','.','.','.'},
	{'6','.','.','1','9','5','.','.','.'},
	{'.','9','8','.','.','.','.','6','.'},
	{'8','.','.','.','6','.','.','.','3'},
	{'4','.','.','8','.','3','.','.','1'},
	{'7','.','.','.','2','.','.','.','6'},
	{'.','6','.','.','.','.','2','8','.'},
	{'.','.','.','4','1','9','.','.','5'},
	{'.','.','.','.','8','.','.','7','9'}}*/

	arr:= [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
                 {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                 {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                 {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                 {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                 {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                 {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                 {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				 {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
				
	res := isValidSudoku(arr);
	fmt.Println("res: ",res)
}

func isValidSudoku(board [][]byte) bool {
     return isRowValid(board) && isColumnValid(board) && is3X3BoxValidValid(board)
}
func isRowValid(board [][]byte) bool{
	seen := make(map[byte]interface{});
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if !isValidSudokuBoxValue(board[r][c],seen) {
				return false;
			}
		}
		clear(seen);
	}
	return true;
}

func isColumnValid(board [][]byte) bool{
	seen := make(map[byte]interface{});
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			if !isValidSudokuBoxValue(board[r][c],seen) {
				return false;
			}
		}
		clear(seen);
	}
	return true;
}

func is3X3BoxValidValid(board [][]byte) bool{
	seen := make(map[byte]interface{});
	 i,j := 0, 0
        for i != 9 && j != 9 {
            j = 0;
            //3X3 box
            for r := i ; r < 3 + i; r++ {
                for c := j ; c < 3 + j; c++ {
					if !isValidSudokuBoxValue(board[r][c],seen) {
						return false;
					}
                }
            }
            clear(seen);

            j += 3;
            //3X3 box
            for r := i ; r < 3 + i; r++ {
                for c := j ; c < 3 + j; c++ {
                    if !isValidSudokuBoxValue(board[r][c],seen) {
						return false;
					}
                }
            }
            clear(seen);


            j += 3;
            //3X3 box
            for r := i ; r < 3 + i; r++ {
                for  c := j ; c < 3 + j; c++ {
                    if !isValidSudokuBoxValue(board[r][c],seen) {
						return false;
					}
                }
            }
            clear(seen);
            i += 3;

        }

        return true;
}

func isValidSudokuBoxValue(v byte,seen map[byte]interface{}) bool{
	if (v == '.') {
		return true
	} else if (v > 57 || v < 49) {
		return false
	} else {
		_,okay := seen[v]
		if (okay) {
			return false
		}
		seen[v]=true
	}
	return true
}

func clear(seen map[byte]interface{}){
	for e := range seen{
		delete(seen,e)
	}
}
