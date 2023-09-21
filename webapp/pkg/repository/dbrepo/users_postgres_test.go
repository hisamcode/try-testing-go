//go:build integration

package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/hisamcode/try-testing-go/webapp/pkg/data"
	"github.com/hisamcode/try-testing-go/webapp/pkg/repository"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB
var testRepo repository.DatabaseRepo

func TestMain(m *testing.M) {
	// connect to docker; fail if docker not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running ? %s", err)
	}
	pool = p

	// set up our docker options, specifying the image and so forth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	// start the image and wait until it's ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("error creating tables: %s", err)
	}

	// init db
	testRepo = &PostgresDBRepo{DB: testDB}

	// run tests
	code := m.Run()

	// clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("can't ping database")
	}
}

func Test_postgresDBRepo_insertUser(t *testing.T) {
	testUser := data.User{
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}

	if id != 1 {
		t.Errorf("insert user returned wrong id; expected 1, but got %d", id)
	}
}

func Test_postgresDBRepo_allUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("all users reports on error: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("all users reports wrong size; expected 1 but got %d", len(users))
	}

	testUser := data.User{
		FirstName: "Jackc",
		LastName:  "Smith",
		Email:     "jackc@smith.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("all users reports on error: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("all users reports wrong size after insert; expected 2 but got %d", len(users))
	}
}

func Test_postgresDBRepo_getUser(t *testing.T) {
	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("error getting user by id: %s", err)
	}

	if user.Email != "admin@example.com" {
		t.Errorf("wrong email returned by GetUser; expected admin@example.com but got %s", user.Email)
	}

	_, err = testRepo.GetUser(3)
	if err == nil {
		t.Error("no error reported when getting non existent user by id")
	}
}

func Test_postgresDBRepo_getUserByEmail(t *testing.T) {
	user, err := testRepo.GetUserByEmail("jackc@smith.com")
	if err != nil {
		t.Errorf("error getting user by email: %s", err)
	}

	if user.ID != 2 {
		t.Errorf("wrong id returned by GetUserByEmail; expected 2 but got %d", user.ID)
	}
}

func Test_postgresDBRepo_updateUser(t *testing.T) {
	user, _ := testRepo.GetUser(2)
	user.FirstName = "Jane"
	user.Email = "jane@smith.com"

	err := testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("error updating user %d: %s", 2, err)
	}

	user, _ = testRepo.GetUser(2)
	if user.FirstName != "Jane" || user.Email != "jane@smith.com" {
		t.Errorf("expected updated record to have firstname jane and email jane@smith.com but got %s %s", user.FirstName, user.Email)
	}
}

func Test_postgresDBRepo_deleteUser(t *testing.T) {
	err := testRepo.DeleteUser(2)
	if err != nil {
		t.Errorf("error deleting user id 2: %s", err)
	}

	_, err = testRepo.GetUser(2)
	if err == nil {
		t.Error("retrieved user id 2, who should have been deleted")
	}
}

func Test_postgresDBRepo_resetPassword(t *testing.T) {
	err := testRepo.ResetPassword(1, "password")
	if err != nil {
		t.Error("error reseting user's password ", err)
	}

	user, _ := testRepo.GetUser(1)
	matches, err := user.PasswordMatches("password")
	if err != nil {
		t.Error(err)
	}
	if !matches {
		t.Error("password should match 'password', but does not")
	}
}

func Test_postgresDBRepo_insertUserImage(t *testing.T) {
	var image data.UserImage
	image.UserID = 1
	image.FileName = "test.jpg"
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()

	newID, err := testRepo.InsertUserImage(image)
	if err != nil {
		t.Errorf("inserting user image failed: %s", err)
	}

	if newID != 1 {
		t.Error("got wrong id for image; should be 1, but got ", newID)
	}

	image.UserID = 100
	_, err = testRepo.InsertUserImage(image)
	if err == nil {
		t.Error("inserted a user image with non-existent user id")
	}

}
