package User

import (
	"context"
	"testing"
	"waterguy/database"
	"waterguy/models"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/gofr"

	"github.com/stretchr/testify/assert"
)
func initializeTest(t *testing.T) *gofr.Context {
	app := gofr.New()

	// initializing the seeder
	database.InitDB()

	seeder := datastore.NewSeeder(&app.DataStore, "../../database")
	seeder.RefreshMongoCollections(t, "user_goal")

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	return ctx
}
func TestModel_Create(t *testing.T) {
	tests := []struct {
		desc     string
		user string
		err      error
	}{
		{"create success", `{"userID":"testuser1","value":42}`, nil},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		var c models.UserEntry

		err := store.CreateUserEntry(ctx, c)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}