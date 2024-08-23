package main
import (
    "testing"
    "github.com/stretchr/testify/suite"
)

// TestMain sets up and tears down resources for all tests.
func TestMain(m *testing.M) {
    // Additional setup code here

    // Run all tests
    code := m.Run()

    // Additional teardown code here

    // Exit with the test result code
    os.Exit(code)
}