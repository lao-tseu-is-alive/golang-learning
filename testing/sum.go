package math_operation

/*
Code based on the article from  Jon Calhoun :
https://www.calhoun.io/how-to-test-with-go/
*/
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
