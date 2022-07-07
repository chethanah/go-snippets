package somepackage

import "fmt"

/*
// For function to be exported use UPPERCASE name
// Below wont work
func printPkg(something string) {
*/
// This will work
func PrintPkg(something string) {
	fmt.Println(something)
}
