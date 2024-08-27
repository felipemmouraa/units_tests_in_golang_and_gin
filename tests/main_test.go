package tests

import (
    "testing"

    "units_tests_in_golang_and_gin/handlers"
)

func TestMain(m *testing.M) {
    // Configurações globais
    code := m.Run()
    // Limpeza
    os.Exit(code)
}