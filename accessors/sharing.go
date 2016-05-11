package accessors

import "github.com/labstack/echo"

type Share struct {
	ID     int `json:"id"`
	User   int `json:"user"`
	Sharee int `json:"sharee"`
}

func (ag *AccessorGroup) Share(c echo.Context, email string) (Share, error) {
	share := Share{}
	err := c.Bind(&share)
	if err != nil {
		return Share{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Share{}, err
	}

	share.User = userID
	share.Sharee = share.Sharee

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO sharing (user, sharee) VALUES (?,?)", share.User, share.Sharee)
	if err != nil {
		return Share{}, err
	}

	return Share{}, nil
}

func (ag *AccessorGroup) GetSharing(c echo.Context, email string) ([]Share, error) {
	allShares := []Share{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Share{}, err
	}

	rows, err := ag.Database.Query("SELECT * FROM sharing WHERE user=? OR sharee=?", userID, userID)
	if err != nil {
		return []Share{}, err
	}

	defer rows.Close()

	for rows.Next() {
		share := Share{}

		err := rows.Scan(&share.ID, &share.User, &share.Sharee)
		if err != nil {
			return []Share{}, err
		}

		allShares = append(allShares, share)
	}

	err = rows.Err()
	if err != nil {
		return []Share{}, err
	}

	return allShares, nil
}
