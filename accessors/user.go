package accessors

import "database/sql"

type User struct {
	ID    int
	Email string
}

// MakeUser adds a user to the database
func (ag *AccessorGroup) MakeUser(email string) error {
	_, err := ag.Database.Query("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user from the database by userID
func (ag *AccessorGroup) GetUser(email string) (User, error) {
	user := &User{}
	err := ag.Database.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(&user.ID, &user.Email)

	if err == sql.ErrNoRows { // If the user doesn't exist yet
		err2 := ag.MakeUser(email)
		if err2 != nil {
			return User{}, err2
		}

		user2, err2 := ag.GetUser(email)
		if err2 != nil {
			return User{}, err2
		}

		return user2, nil
	} else if err != nil {
		return User{}, err
	}

	return *user, nil
}
