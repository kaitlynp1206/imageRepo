package user

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type UsersManager struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *UsersManager {
	return &UsersManager{
		db: db,
	}
}

func (m *UsersManager) AddHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *UsersManager) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *UsersManager) AddUser(ctx context.Context, name string, password string) error {
	//TODO: Validate input
	insForm, err := m.db.Prepare("INSERT INTO users(username, password) VALUES(?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(name, password)
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

func (m *UsersManager) GetUserByID(ctx context.Context, usrID int) (User, error) {
	selDB, err := m.db.Query("SELECT user_id, username, password FROM users WHERE user_id=?", usrID)
	if err != nil {
		return nil, err
	}
	var usr User
	for selDB.Next() {
		err = selDB.Scan(&usr.UserID, &usr.Username, &usr.Password)
	}
	//TODO: Validate change
	return usr, nil
}

func (m *UsersManager) GetUserByUserID(ctx context.Context, usrID int) (User, error) {
	selDB, err := m.db.Query("SELECT user_id, username, password FROM users WHERE user_id=?", usrID)
	if err != nil {
		return nil, err
	}
	var usr User
	for selDB.Next() {
		err = selDB.Scan(&usr.UserID, &usr.Username, &usr.Password)
	}
	return usr, nil
}

func (m *UsersManager) GetUserByUsername(ctx context.Context, username string) (User, error) {
	selDB, err := m.db.Query("SELECT user_id, username, password FROM users WHERE username=?", username)
	if err != nil {
		return nil, err
	}
	var usr User
	for selDB.Next() {
		err = selDB.Scan(&usr.UserID, &usr.Username, &usr.Password)
	}
	return usr, nil
}
