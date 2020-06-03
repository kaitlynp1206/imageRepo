package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type UsersManager struct {
	db *sql.DB
}

func NewUsersManager(db *sql.DB) *UsersManager {
	return &UsersManager{
		db: db,
	}
}

//User Handler
func (m *UsersManager) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		m.GetUserHandler(w, r)
	case "POST":
		m.AddUserHandler(w, r)
	case "DELETE":
		m.RemoveUserHandler(w, r)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (m *UsersManager) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var usr User

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &User{}
	//Get by ID or username
	if usr.UserID != 0 {
		user, err = m.GetUserByUserID(r.Context(), usr.UserID)
	} else if usr.Username != "" {
		user, err = m.GetUserByUsername(r.Context(), usr.Username)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.UserID == 0 {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(user)

	fmt.Fprintf(w, string(resp))
	return

}

func (m *UsersManager) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var usr User

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Add Image to DB
	err = m.AddUser(r.Context(), usr.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Consideration - return added user userID
	fmt.Fprintf(w, fmt.Sprintf("Successfully added user as %s", usr.Username))
	return

}

func (m *UsersManager) RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
	var usr User

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.DeleteUser(r.Context(), usr.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Successfully removed user %d", usr.UserID))
	return
}

func (m *UsersManager) AddUser(ctx context.Context, name string) error {
	//TODO: Validate input
	insForm, err := m.db.Prepare("INSERT INTO users(username) VALUES(?)")
	if err != nil {
		return err
	}
	insForm.Exec(name)
	//TODO: Return userId
	return nil
}

func (m *UsersManager) DeleteUser(ctx context.Context, usrID int) error {
	//TODO: Validate input - ID exists
	delForm, err := m.db.Prepare("DELETE FROM users WHERE user_id=?")
	if err != nil {
		return err
	}
	delForm.Exec(usrID)
	//TODO: Validate change
	return nil
}

func (m *UsersManager) GetUserByUserID(ctx context.Context, usrID int) (*User, error) {
	selDB, err := m.db.Query("SELECT user_id, username FROM users WHERE user_id=?", usrID)
	if err != nil {
		return nil, err
	}
	var usr User
	for selDB.Next() {
		err = selDB.Scan(&usr.UserID, &usr.Username)
	}

	return &usr, nil
}

func (m *UsersManager) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	selDB, err := m.db.Query("SELECT user_id, username FROM users WHERE username=?", username)
	if err != nil {
		return nil, err
	}
	var usr User
	for selDB.Next() {
		err = selDB.Scan(&usr.UserID, &usr.Username)
	}
	return &usr, nil
}
