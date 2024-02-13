package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"sort"
	"time"
)

// type Ticket struct {
// 	Entries [3][9]int
// }

func GenerateTicket() [][]int {
	ticket := make([][]int, 3)
	//Initializing ticket
	for i := range ticket {
		ticket[i] = make([]int, 9)
	}

	rand.Seed(time.Now().UnixNano())

	//Inserting the ticket Numbers

	for col := 0; col < 9; col++ {
		colmNum := rand.Perm(10)
		for i := 0; i < 3; i++ {
			ticket[i][col] = col*10 + colmNum[i] + 1
		}
	}

	//shuffle the number in each column

	for col := 0; col < 9; col++ {
		rand.Shuffle(3, func(i, j int) {
			ticket[i][col], ticket[j][col] = ticket[j][col], ticket[i][col]
		})

	}
	for row := 0; row < 3; row++ {
		for i := 0; i < 4; i++ {
			col := rand.Intn(9) // Randomly select a column
			ticket[row][(col+row)%9] = 0
		}

	}

	SortRows(ticket)

	return ticket
}
func SortRows(ticket [][]int) {
	sort.Slice(ticket, func(i, j int) bool {
		return ticket[i][0] < ticket[j][0]
	})
}

func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

// func UniqueValues(slice []int) bool {
// 	seen := make(map[int]struct{})
// 	for _, value := range slice {
// 		if _, exists := seen[value]; exists {
// 			return false // Duplicate value found
// 		}
// 		seen[value] = struct{}{}
// 	}
// 	return true
// }

// func IsvalidTicket(ticket domain.Ticket) bool {

// 	// Check Entries of array is empty or not

// 	if len(ticket.Entries) == 0 {
// 		return false
// 	}

// 	for _, row := range ticket.Entries {
// 		if !UniqueValues(row) {
// 			return false
// 		}
// 	}
// 	return true
// }

// ValidateTambolaTicket validates a single Tambola ticket
func ValidateTambolaTicket(ticket [][]int) error {
	// Rule: Each row must have exactly 5 numbers in it
	for _, row := range ticket {
		if countNonZero(row) != 5 {
			return errors.New("each row must have exactly 5 numbers")
		}
	}

	// Rule: In a specific column, numbers must be arranged in ascending order from top to bottom
	for col := 0; col < 9; col++ {
		for row := 1; row < 3; row++ {
			if ticket[row-1][col] > ticket[row][col] {
				return errors.New("numbers in each column must be in ascending order from top to bottom")
			}
		}
	}

	// Rule: Each column must have at least 1 number
	// for col := 0; col < 9; col++ {
	// 	if countNonZeroColumn(ticket, col) == 0 {
	// 		return errors.New("each column must have at least 1 number")
	// 	}
	// }

	// Rule: All the numbers 1 to 90 are used only once in each set of 6 tickets
	// usedNumbers := make(map[int]bool)
	// for _, row := range ticket {
	// 	for _, num := range row {
	// 		if num != 0 {
	// 			if usedNumbers[num] {
	// 				return errors.New("all the numbers 1 to 90 must be used only once in each set of 6 tickets")
	// 			}
	// 			usedNumbers[num] = true
	// 		}
	// 	}
	// }
	seen := make(map[int]bool)

	for _, row := range ticket {
		for _, value := range row {
			if value != 0 {
				if _, exists := seen[value]; exists {
					return errors.New("all the numbers 1 to 90 must be used only once in each set of 6 tickets") // Duplicate value found
				}
				seen[value] = true
			}
		}
	}

	return nil
}

// Helper function to count non-zero elements in a slice
func countNonZero(slice []int) int {
	count := 0
	for _, value := range slice {
		if value != 0 {
			count++
		}
	}
	return count
}

// Helper function to count non-zero elements in a specific column
func countNonZeroColumn(ticket [][]int, col int) int {
	count := 0
	for _, row := range ticket {
		if row[col] != 0 {
			count++
		}
	}
	return count
}

// func GenerateUniqueID(digits int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	min := int64(10000)
// 	max := int64(99999)
// 	randomID := min + rand.Int63n(max-min+1)
// 	return fmt.Sprintf("%0*d", digits, randomID)
// }

func MatrixToString(matrix [][]int) (string, error) {
	data, err := json.Marshal(matrix)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// StringToMatrix converts a JSON string to a [][]int
func StringToMatrix(data string) ([][]int, error) {
	var matrix [][]int
	err := json.Unmarshal([]byte(data), &matrix)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}

// func CheckDuplicates(matrix [][]int) map[int]bool {
// 	seen := make(map[int]bool)

// 	for _, row := range matrix {
// 		for _, value := range row {
// 			// Check if the value already exists in the map
// 			if _, exists := seen[value]; exists {
// 				// Duplicate value found
// 				//fmt.Printf("Duplicate value found: %d\n", value)
// 				seen[value] =
// 			} else {
// 				// Insert the value into the map
// 				seen[value] = true
// 			}
// 		}
// 	}

// 	return seen
// }

func HandleDuplicates(matrix [][]int, valueMap map[int]bool, checkMap map[int][]int) [][]int {

	for i := range matrix {
		for j := range matrix[i] {
			value := matrix[i][j]
			// minValue := (i - 1) * 10
			// maxValue := (i * 10) + 1
			//temp := value
			if value == 0 {
				continue
			}
			checkValues, ok := checkMap[j]
			if !ok {
				// If the key is not found in checkMap, skip the replacement
				continue
			}

			if !contains(checkMap[j], value) {
				// Insert the new value into the array
				checkMap[j] = append(checkMap[j], value)

			}

			//temp := value

			// Check for duplicates
			// for valueMap[value] {
			// 	//Replace duplicate with a number between 20 and 30
			// 	// value = rand.Intn(11) + ((value / 10) * 10)
			// 	// if value == temp && value > j*10 || value < j*10 {
			// 	// 	value = rand.Intn(11) + ((value / 10) * 10)
			// 	// }
			// 	// minValue := (i-1) * 10
			// 	// maxValue := i * 10
			// 	// value = rand.Intn(11) + ((value / 10) * 10)

			// 	// // Ensure the value is not greater than 90
			// 	// if value > 90 {
			// 	// 	value = rand.Intn(11) + ((value / 10) * 10)
			// 	// }
			// 	// Replace duplicate with a number based on checkMap range
			// 	value = rand.Intn(11) + (checkValues[0] / 10 * 10)
			// 	checkValues, ok = value

			// 	// Ensure the value is not greater than checkMap range
			// 	if value > checkValues[j] {
			// 		value = rand.Intn(11) + (checkValues[0] / 10 * 10)
			// 	}

			// }
			for valueMap[value] || isValueRepeated(matrix, i, j, value) {
				// Replace duplicate with a number based on checkMap range
				value = FindValuesNotInRange(matrix, j, checkValues[0], checkValues[1])
			}
			if !contains(checkMap[j], value) {
				// Insert the new value into the array
				checkMap[j] = append(checkMap[j], value)

			}

			// Insert value into map
			valueMap[value] = true

			// Update matrix with the modified value
			matrix[i][j] = value
		}
	}
	result := matrix
	return result

}

// Helper function to get the next unique value within the specified range (excluding 0)
func getNextUniqueValue(valueMap map[int]bool, minValue, maxValue int) int {
	for {
		value := rand.Intn(maxValue-minValue+1) + minValue
		if value != 0 && !valueMap[value] {
			return value
		}
	}
}

// Helper function to check if a value is repeated in the matrix within the column's range
func isValueRepeated(matrix [][]int, row, col, value int) bool {
	for i := 0; i < row; i++ {
		// Check if the value is within the column's range
		if matrix[i][col] == value {
			return true
		}
	}
	return false
}
func FindValuesNotInRange(matrix [][]int, col, minValue, maxValue int) int {
	var notInRange []int
	for _, row := range matrix {
		value := row[col]
		if value != 0 && (value < minValue || value > maxValue) {
			notInRange = append(notInRange, value)
		}
	}
	return notInRange[0]
}
