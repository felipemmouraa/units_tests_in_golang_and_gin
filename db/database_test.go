package db

import (
    "fmt"
    "os"
    "testing"

    "github.com/stretchr/testify/suite"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// User represents a simple user model.
type User struct {
    ID   uint
    Name string
}

// DatabaseTestSuite is the test suite.
type DatabaseTestSuite struct {
    suite.Suite
    db *gorm.DB
}

// SetupSuite is called once before the test suite runs.
func (suite *DatabaseTestSuite) SetupSuite() {
    // Set up a PostgreSQL database for testing
    dsn := "user=testuser password=testpassword dbname=testdb sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    suite.Require().NoError(err, "Error connecting to the test database")

    // Enable logging for Gorm during tests
    suite.db = db.Debug()

    // Auto-migrate tables
    err = suite.db.AutoMigrate(&User{})
    suite.Require().NoError(err, "Error auto-migrating database tables")
}

// TestUserInsertion tests inserting a user record.
func (suite *DatabaseTestSuite) TestUserInsertion() {
    // Create a user
    user := User{Name: "John Doe"}
    err := suite.db.Create(&user).Error
    suite.Require().NoError(err, "Error creating user record")

    // Retrieve the inserted user
    var retrievedUser User
    err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
    suite.Require().NoError(err, "Error retrieving user record")

    // Verify that the retrieved user matches the inserted user
    suite.Equal(user.Name, retrievedUser.Name, "Names should match")
}

// TearDownSuite is called once after the test suite runs.
func (suite *DatabaseTestSuite) TearDownSuite() {
    // Clean up: Close the database connection
    err := suite.db.Exec("DROP TABLE users;").Error
    suite.Require().NoError(err, "Error dropping test table")

    err = suite.db.Close()
    suite.Require().NoError(err, "Error closing the test database")
}

// TestSuite runs the test suite.
func TestSuite(t *testing.T) {
    // Skip the tests if the PostgreSQL connection details are not provided
    if os.Getenv("POSTGRES_DSN") == "" {
        t.Skip("Skipping PostgreSQL tests; provide POSTGRES_DSN environment variable.")
    }

    suite.Run(t, new(DatabaseTestSuite))
}