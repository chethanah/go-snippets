//Package somepackage is used for demo
package somepackage

import "fmt"

/*
// For function to be exported use UPPERCASE name
// Below wont work
func printPkg(something string) {
*/
// Below will work

//PrintPkg will take input string and print string
func PrintPkg(something string) {
	fmt.Println(something)
}
