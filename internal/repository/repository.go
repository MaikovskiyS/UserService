package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"practice/internal/models"
	"practice/internal/service"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type mySQL struct {
	log     *logrus.Logger // здесь не надо но пока не дойдем до логера можно
	db      *sql.DB        // должно быть не экспоритируемо db
	timeOut time.Duration
}

// NewRepository package constructor
func NewRepository(connString string, log *logrus.Logger, timeOut time.Duration) (service.Repository, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return &mySQL{
		db:      db,
		log:     log,
		timeOut: timeOut,
	}, nil
}

// Close ...
func (m *mySQL) Close() error {
	return m.db.Close()
}

func (m *mySQL) GetFriendsById(id uint) ([]string, error) {
	return nil, nil
}
func (m *mySQL) AddFriend(id1, id2 uint) (name1, name2 string, err error) {
	return
}
func (m *mySQL) DeleteUser(ctx context.Context, id *uint) error {
	fmt.Println("in deleteuserFunc  id:", *id)
	execCtx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()
	res, err := m.db.ExecContext(execCtx, "DELETE FROM user.users WHERE `idUsers`=?", id)
	if err != nil {
		m.log.Info("cant delete from db")
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	m.log.Printf("%d user deleted ", rows)
	return nil
}
func (m *mySQL) GetUserById(ctx context.Context, id *uint) (u *models.User, err error) {
	fmt.Println("in getUSerById    id:", *id)
	execCtx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()
	rows, err := m.db.QueryContext(execCtx, "SELECT `idUsers`,`name`,`age` FROM user.users WHERE idUsers=?", id)
	if err != nil {
		m.log.Info("cant select userbyid to db", err)
		cancel()
	}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			m.log.Info("cant write data from db to model user")
			return nil, err
		}
		u = &user
	}
	return u, nil
}
func (m *mySQL) CreateUser(ctx context.Context, u *models.User) (uint, error) {
	//TODO сделать нормальные id(не присваеваются освобожденные после удаления)
	execCtx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()

	res, err := m.db.ExecContext(execCtx, "INSERT user.users(name,age) VALUES (?,?)", u.Name, u.Age)
	if err != nil {
		m.log.Info("80", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	log.Printf("%d products created ", rows)
	id, err := res.LastInsertId()
	if err != nil {
		m.log.Info("cant return users id")
	}

	return uint(id), nil
}
func (m *mySQL) GetAllUsers(ctx context.Context) ([]models.User, error) {
	execCtx, cancel := context.WithTimeout(ctx, m.timeOut)
	defer cancel()
	rows, err := m.db.QueryContext(execCtx, "SELECT `idUsers`, `name`,`age` FROM user.users")
	if err != nil {
		m.log.Info("cant select request to db")
	}
	var arr []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			m.log.Info("cant parse data from db")
		}
		fmt.Println(user)
		arr = append(arr, user)

	}
	fmt.Println(arr)
	return arr, nil
}
