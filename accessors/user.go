package accessors

import "database/sql"

// MakeUser adds a user to the database
func (ag *AccessorGroup) MakeUser(ID string, username string) error {
	_, err := ag.Database.Query("INSERT INTO users (ID, username) VALUES (?,?)", ID, username)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user from the database by userID
func (ag *AccessorGroup) GetUser(ID int) (User, error) {
	user := &User{}
	err := ag.Database.QueryRow("SELECT * FROM users WHERE ID=?", ID).Scan(&user.ID, &user.Username)

	if err == sql.ErrNoRows { // If the user doesn't exist yet
		return User{}, nil // Return a blank user
	} else if err != nil {
		return User{}, err
	}

	return *user, nil
}
