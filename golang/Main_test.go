package main
 
import (
	"fmt"
	"os"
	"testing"
)
 
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
 
func setup() {
	// Do something here.
 
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}
 
func teardown() {
	// Do something here.
 
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}