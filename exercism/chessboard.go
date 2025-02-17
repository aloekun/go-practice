package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	// count in column
	for _, exist := range cb[file] {
		if exist {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	const RankMin = 1
	const RankMax = 8
	if rank < RankMin || RankMax < rank {
		return 0
	}

	count := 0
	// count in row
	for _, fileColumn := range cb {
		if fileColumn[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	countFile := len(cb)

	countRank := 0
	for _, file := range cb {
		countRank = len(file)
		break
	}

	return countFile * countRank
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for key, _ := range cb {
		count += CountInFile(cb, key)
	}
	return count
}
