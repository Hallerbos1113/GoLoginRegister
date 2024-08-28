package model

import(
	"fmt"
	"example/utils"
)

func RegUser(user *User) (*User, error) {
	user.Password = utils.HashPassword(user.Password)
	user.Uuid =  utils.GenerateUserID(user.UserName)
	query := `INSERT INTO tb_test ("Uuid", "Username", "Password", "Api_user_id", "Is_admin", "Email") VALUES ($1,$2,$3,$4,$5,$6)`
	res, err := DB.Exec(query, user.Uuid, user.UserName, user.Password, user.ApiUserID, user.IsAdmin, user.Email)
	if err != nil || res == nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func LoginUser(email string) (*User, error) {
	var Uuid string
	var Username string
	var Password string
	var Email string
	var Api_user_id string
	var Is_admin bool
	query := `SELECT "Uuid", "Username", "Password", "Email", "Api_user_id", "Is_admin" FROM tb_test WHERE "Email"=$1`
	row := DB.QueryRow(query, email)
	err := row.Scan(&Uuid, &Username, &Password, &Email, &Api_user_id, &Is_admin)
	// fmt.Println(err)
	// fmt.Println(Uuid, Username, Password, Email, Api_user_id, Is_admin)
	if err != nil {
		return nil, err
	}
	return &User{
		Uuid: Uuid, Password: Password, Email: Email, ApiUserID: Api_user_id, IsAdmin: Is_admin, UserName: Username,
	}, nil
}

func GetUsers(userID string) ([]User, error) {
	var query string
	var ret []User

	if userID == "" { // get all user
		query = `SELECT "Uuid", "Username", "Password", "Email", "Api_user_id", "Is_admin" FROM tb_test`
	} else { // get one user
		query = "SELECT \"Uuid\", \"Username\", \"Password\", \"Email\", \"Api_user_id\", \"Is_admin\" FROM tb_test WHERE \"Uuid\"='%s'"
		query = fmt.Sprintf(query, userID)
		fmt.Printf(query)
	}
	rows, _ := DB.Query(query)
	defer rows.Close()
	
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Uuid, &u.UserName, &u.Password, &u.Email, &u.ApiUserID, &u.IsAdmin); err != nil {
			return ret, err
		}
		ret = append(ret, u)
	}
	if err := rows.Err(); err != nil {
		return ret, err
	}
	return ret, nil
}

func DelUserUID(userID string) bool {
	query := `DELETE FROM "tb_test" WHERE "Uuid"=$1`
	DB.Exec(query, userID)
	return true
}

func UpdateUserAdmin(userID string, user *User) bool {
	query := `UPDATE tb_test SET "Username"=$1, "Email"=$2, "Api_user_id"=$3 WHERE "Uuid"=$4`
	// Execute the update query
	result, err := DB.Exec(query, user.UserName, user.Email, user.ApiUserID, userID)
	if err != nil {
		return false
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println(rowsAffected)
	if err != nil || rowsAffected < 1 {
		return false
	}
	return true
}
func UpdateUser(user *User) bool {
	query := `UPDATE tb_test SET "Username"=$1, "Email"=$2, "Password"=$3 WHERE "Uuid"=$4`
	fmt.Printf(query, user.UserName, user.Email, user.Password, user.Uuid)
	user.Password = utils.HashPassword(user.Password)
	// Execute the update query
	result, err := DB.Exec(query, user.UserName, user.Email, user.Password, user.Uuid)
	if err != nil {
		return false
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	fmt.Println(rowsAffected)
	if err != nil || rowsAffected < 1 {
		return false
	}
	return true
}