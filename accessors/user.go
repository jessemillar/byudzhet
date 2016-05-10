package accessors

import "database/sql"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// GetUser returns a user from the database by userID
func (ag *AccessorGroup) GetUser(email string) (User, error) {
	user1 := &User{}
	err1 := ag.Database.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(&user1.ID, &user1.Email)

	if err1 == sql.ErrNoRows { // If the user doesn't exist yet
		user2, err2 := ag.MakeUser(email)
		if err2 != nil {
			return User{}, err2
		}

		return user2, nil
	} else if err1 != nil {
		return User{}, err1
	}

	return *user1, nil
}

// GetUserID returns a user from the database by userID
func (ag *AccessorGroup) GetUserID(email string) (int, error) {
	user, err := ag.GetUser(email)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// MakeUser adds a user to the database
func (ag *AccessorGroup) MakeUser(email string) (User, error) {
	_, err := ag.Database.Query("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return User{}, err
	}

	user, err := ag.GetUser(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
