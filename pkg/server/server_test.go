package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/cepriego/test/pkg/postgres"
	"github.com/gin-gonic/gin"
)

const pgaddr string = "postgresql://user:password123@localhost:5432/shore_test"

func TestRunServer(t *testing.T) {

	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)

	handlers := NewHandlers(*pgserrep)
	server := NewShoreServer(handlers)
	server.Init()

	if err != nil {
		fmt.Println("exiting with error", err)
	}
}

//Punto 1
func TestGetUsers(t *testing.T) {
	//Test to get all users
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/getUsers", handlers.getUsers)

	payload := `{}` //Ignored
	req, _ := http.NewRequest("GET", "/getUsers", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(p))
}

//Punto 2
func TestGetUser(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/getUser", handlers.getUser)

	payload := `{"Id":5}`
	req, _ := http.NewRequest("GET", "/getUser", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(p))
}

//Punto 3
func TestCreateUser(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/newUser", handlers.CreateUser)

	payload := `{"Id":5,"FechaCreacion":"2021-01-15T15:32:00Z","Nombre":"Santiago Priego Zurita","Eliminado":false}`
	req, _ := http.NewRequest("POST", "/newUser", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || !strings.Contains(string(p), "user created") {
		t.Fail()
	}
}

func TestCreateGroup(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/newGroup", handlers.CreateGroup)

	payload := `{"Id":4,"FechaCreacion":"2021-01-15T15:32:00Z","Nombre":".Net Developer","Eliminado":false}`
	req, _ := http.NewRequest("POST", "/newGroup", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || !strings.Contains(string(p), "group created") {
		t.Fail()
	}
}

//Punto 5
func TestGetGroupById(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/getGroup", handlers.getGroup)

	payload := `{"Id":2}`
	req, _ := http.NewRequest("GET", "/getGroup", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(p))
}

// Punto 7
func TestAssignUserToGroup(t *testing.T) {
	//Test to assign a user to a specific group
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/assignToGroup", handlers.AssignUserToGroup)

	payload := `{"IdUsuario":2,"IdGrupo":4}`
	req, _ := http.NewRequest("POST", "/assignToGroup", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || !strings.Contains(string(p), "user assigned to group") {
		t.Fail()
	}
}

//Punto 4
func TestGetGroupAndUsers(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/getGroupAndUsers", handlers.GetGroupsUsers)

	payload := `{"Id":2}`
	req, _ := http.NewRequest("GET", "/getGroupAndUsers", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(p))
}

// Punto 9
func TestDeleteUser(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/deleteUser", handlers.DeleteUser)

	payload := `{"Id":5}`
	req, _ := http.NewRequest("POST", "/deleteUser", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || !strings.Contains(string(p), "user created") {
		t.Fail()
	}
}

//Punto 8
func TestDeleteUserFromGroup(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/deleteUserFromGroup", handlers.DeleteUserFromGroup)

	payload := `{"IdUsuario":3, "IdGrupo":2}`
	req, _ := http.NewRequest("POST", "/deleteUserFromGroup", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}

func TestDeleteGroup(t *testing.T) {
	pgrep := postgres.New()
	ctx := context.Background()

	pgCloser, err := pgrep.Connect(ctx, pgaddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer pgCloser()
	pgserrep := NewPostgresRepo(pgrep)
	handlers := NewHandlers(*pgserrep)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/deleteGroup", handlers.DeleteGroup)

	payload := `{"Id":4}`
	req, _ := http.NewRequest("POST", "/deleteGroup", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || !strings.Contains(string(p), "user created") {
		t.Fail()
	}
}
