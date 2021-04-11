package accessor

import (
	"bogdanfloris-com/pkg/accessor"
	"bogdanfloris-com/pkg/logging"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var pgAccessor *accessor.PgAccessor

func TestMain(m *testing.M) {
	// Write code here to run before tests
	logging.InitLoggers()
	pgAccessor = accessor.NewPgAccessor()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests
	pgAccessor.Close()

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestPgAccessor_DatabaseName(t *testing.T) {
	assert.Equal(t, "bfdb", pgAccessor.DatabaseName(), "Wrong database name.")
}
