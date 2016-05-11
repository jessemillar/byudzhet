package accessors

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// GetUserByEmail returns a user from the database by email
func (ag *AccessorGroup) GetUserByEmail(email string) (User, error) {
	user := &User{}
	err := ag.Database.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(&user.ID, &user.Email)

	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserByID returns a user from the database by userID
func (ag *AccessorGroup) GetUserByID(email string) (User, error) {
	user := &User{}
	err := ag.Database.QueryRow("SELECT * FROM users WHERE id=?", email).Scan(&user.ID, &user.Email)

	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserID returns a user from the database by userID
func (ag *AccessorGroup) GetUserID(email string) (int, error) {
	user, err := ag.GetUserByEmail(email)
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

	user, err := ag.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
