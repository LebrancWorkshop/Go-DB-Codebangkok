package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
	Displayname string
}

var db *sql.DB;

func main() {
	// Open SQL Database. 
	var err error;
	db, err = sql.Open("mysql", "root:<SECRET_PASSWORD>@tcp(0.0.0.0:3306)/codebangkok");
	if err != nil {
		fmt.Println("[ERROR] Open SQL error");
		panic(err);
	}

	// Implement Add User.
	// newUser := User{"alarm9", "w3ddin@", "Joe Bongo"};
	// err = AddUser(newUser);
	// if err != nil {
	// 	fmt.Println("[ERROR] Cannot Add User");
	// 	fmt.Println(err);
	// }

	// Implement Update User. 
	// updateUser := User{"alarm9", "w3ddin@", "Joe Nat"};
	// err = UpdateUser(updateUser, 3);
	// if err != nil {
	// 	fmt.Println("[ERROR] Cannot Update User");
	// 	fmt.Println(err);
	// }

	// Implement Delete User. 
	err = DeleteUser(3);
	if err != nil {
		fmt.Println("[ERROR] Cannot Delete User");
		fmt.Println(err); 
	}

	// Implement Get Users.
	users, err := GetUsers();
	if err != nil {
		fmt.Println("[ERROR] Get Users error");
		fmt.Println(err);
	}

	for _, user := range users {
		fmt.Println(user); 
	}

	// Implement Get User By ID.
	// user, err := GetUserByID(2);
	// if err != nil {
	// 	fmt.Println("[ERROR] Get User error")
	// 	fmt.Println(err);
	// }
	// fmt.Println(user); 
}

// Get Users. 
func GetUsers() ([]User, error) {
	// Ping Database Server. 
	err := db.Ping(); 
	if err != nil {
		return nil, err; 
	}

	// Query Database. 
	query := "SELECT ID, username, password, displayname FROM lingoquest_users;"; 
	rows, err := db.Query(query);
	if err != nil {
		return nil, err;
	}
	defer rows.Close(); 

	// Scan Data from Database. 
	users := []User{}; 
	for rows.Next() {
		user := User{};
		ID := 0;
		err := rows.Scan(&ID, &user.Username, &user.Password, &user.Displayname);
		if err != nil {
			return nil, err;
		}

		users = append(users, user); 		
	}

	return users, nil; 
}

func GetUserByID(id int) (*User, error) {
	// Ping Database Server. 
	err := db.Ping(); 
	if err != nil {
		return nil, err; 
	}

	// Query Database. 
	query := "SELECT ID, username, password, displayname FROM lingoquest_users WHERE id=?"; 
	row := db.QueryRow(query, id);
	if err != nil {
		return nil, err;
	}
	user := User{};
	ID := 0;
	err = row.Scan(&ID, &user.Username, &user.Password, &user.Displayname);
	if err != nil {
		return nil, err; 
	}

	return &user, nil; 
}

// Add New User. 
func AddUser(user User) error {
	query := "INSERT INTO lingoquest_users (username, password, displayname) VALUES (?, ?, ?)";
	result, err := db.Exec(query, user.Username, user.Password, user.Displayname);
	if err != nil {
		return err; 
	}

	affected, err := result.RowsAffected();
	if err != nil {
		return err;
	}
	if affected <= 0 {
		return errors.New("[ERROR] Cannot Insert Data");
	}

	return nil; 
}

// Update New User. 
func UpdateUser(user User, id int) error {
	query := "UPDATE lingoquest_users SET username=?, password=?, displayname=? WHERE id=?";
	result, err := db.Exec(query, user.Username, user.Password, user.Displayname, id);
	if err != nil {
		return err; 
	}

	affected, err := result.RowsAffected();
	if err != nil {
		return err;
	}
	if affected <= 0 {
		return errors.New("[ERROR] Cannot Update Data");
	}

	return nil; 
}

// Delete New User. 
func DeleteUser(id int) error {
	query := "DELETE FROM lingoquest_users WHERE id=?";
	result, err := db.Exec(query, id);
	if err != nil {
		return err; 
	}

	affected, err := result.RowsAffected();
	if err != nil {
		return err;
	}
	if affected <= 0 {
		return errors.New("[ERROR] Cannot Delete Data");
	}

	return nil; 
}