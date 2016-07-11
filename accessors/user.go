package accessors

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// GetUserByEmail returns a user from the database by email
func (accessorGroup *AccessorGroup) GetUserByEmail(email string) (User, error) {
	user := &User{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(&user.ID, &user.Email)

	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserByID returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserByID(email string) (User, error) {
	user := &User{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM users WHERE id=?", email).Scan(&user.ID, &user.Email)

	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserID returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserID(email string) (int, error) {
	user, err := accessorGroup.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// MakeUser adds a user to the database
func (accessorGroup *AccessorGroup) MakeUser(email string) (User, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return User{}, err
	}

	user, err := accessorGroup.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
