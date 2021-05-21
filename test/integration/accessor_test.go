package accessor

import (
	"bogdanfloris-com/internal/accessor"
	"bogdanfloris-com/internal/logging"
	"bogdanfloris-com/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const USERNAME = "TEST"
const PASSWORD = "TEST"

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

func TestPgAccessor_UserAddDelete(t *testing.T) {
	// Add user
	_, err := pgAccessor.AddUser(USERNAME, PASSWORD)
	assert.NoError(t, err)

	// Get user
	user, err := pgAccessor.User(USERNAME)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, USERNAME)
	assert.True(t, utils.CheckPasswordHash(PASSWORD, user.PasswordHash))

	// Remove user
	err = pgAccessor.RemoveUser(USERNAME)
	assert.NoError(t, err)
}
