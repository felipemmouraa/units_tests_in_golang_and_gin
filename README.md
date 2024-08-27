# Testes Unitários em Golang com Gin e Testify

Este repositório contém exemplos e exercícios relacionados a testes unitários em projetos Golang utilizando o framework Gin e a biblioteca Testify.

## Estrutura do Projeto

```plaintext
/units_tests_in_golang_and_gin
|-- /handlers
|   |-- handler.go
|   |-- handler_test.go
|
|-- /db
|   |-- database.go
|   |-- database_test.go
|
|-- main.go
|-- go.mod
|-- go.sum
|-- /tests
|   |-- main_test.go

![alt text](image.png)
```

### 1.TestMain e Conjunto de Testes

**TestMain:** Usado para configurar e desmontar recursos antes e depois dos testes.

``` golang
func TestMain(m *testing.M) {
    // Configurações globais
    code := m.Run()
    // Limpeza
    os.Exit(code)
}

```

**Conjuntos de Testes com Testify:** Agrupam testes logicamente, permitindo a configuração e desmontagem comuns.

```golang
type MySuite struct {
    suite.Suite
}

func (suite *MySuite) SetupTest() {
    // Código de configuração
}

func (suite *MySuite) TestSomething() {
    // Teste
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(MySuite))
}
```
### 2. Testes com PostgreSQL

**Configuração:** Conexão com um banco de dados PostgreSQL de teste.

**Migração Automática:** Auto-migração de tabelas para testes.

Exemplo de Teste com PostgreSQL:
``` golang
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

```

### 3. Conclusão:

Este projeto exemplifica a aplicação prática de testes unitários em Golang utilizando Testify, cobrindo desde a configuração até a execução de testes com PostgreSQL. 
