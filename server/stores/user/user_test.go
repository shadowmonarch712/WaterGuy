package User

import (
	"context"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/gofr"
	"testing"
	"waterguy/database"
	"waterguy/models"

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
		desc string
		user string
		err  error
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

func TestCustomer_Fetch(t *testing.T) {
	tests := []struct {
		desc  string
		value string
		resp  models.UserEntry
		err   error
	}{
		{
			desc:  "get single entity",
			value: "Messi",
			resp:  models.UserEntry{UserID: "Messi", Value: 32},
			err:   nil,
		},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		resp, err := store.FetchUserEntry(ctx, tc.value)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestModel_Delete(t *testing.T) {
	tests := []struct {
		desc   string
		userID string
		count  int
		err    error
	}{
		{"delete non existent entity", "Alex", 0, nil},
		{"delete single entity", "Thomas", 1, nil},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		count, err := store.DeleteUserEntry(ctx, tc.userID)
		//count := result.(int)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.count, count, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
func TestModel_Update(t *testing.T) {
	tests := []struct {
		desc string
		user models.UserEntry
		resp interface{}
		err  error
	}{
		{"update non existent entity", models.UserEntry{UserID: "Alex", Value: 50}, "No document found with the given userID to update", nil},
		{"update existing entity", models.UserEntry{UserID: "Thomas", Value: 100}, 100, nil},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		resp, err := store.UpdateUserEntry(ctx, tc.user)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
