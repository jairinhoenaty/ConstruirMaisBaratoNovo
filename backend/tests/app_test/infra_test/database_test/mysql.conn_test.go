package infra_test

import (
	pkgdatabase "construir_mais_barato/infra/database/mysql-db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionDB(t *testing.T) {
	db := pkgdatabase.ConnectionDB(nil)

	assert.NotEmpty(t, db)
}
