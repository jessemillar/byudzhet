package accessors

import "github.com/labstack/echo"

type ProjectedIncome struct {
	ID     int     `json:"id,string"`
	User   int     `json:"user,string"`
	Amount float64 `json:"amount,string"`
}

func (ag *AccessorGroup) SetProjectedIncome(c echo.Context, email string) (ProjectedIncome, error) {
	projectedIncome := ProjectedIncome{}
	err := c.Bind(&projectedIncome)
	if err != nil {
		return ProjectedIncome{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO projected (user, amount) VALUES (?,?)", projectedIncome.User, projectedIncome.Amount*100)
	if err != nil {
		return ProjectedIncome{}, err
	}

	return ProjectedIncome{}, nil
}

func (ag *AccessorGroup) GetProjectedIncome(c echo.Context, email string) ([]ProjectedIncome, error) {
	allProjectedIncome := []ProjectedIncome{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []ProjectedIncome{}, err
	}

	allProjectedIncome, err = ag.GetProjectedIncomeByUserID(c, allProjectedIncome, userID)
	if err != nil {
		return []ProjectedIncome{}, err
	}

	return allProjectedIncome, nil
}

func (ag *AccessorGroup) GetProjectedIncomeByUserID(c echo.Context, allProjectedIncome []ProjectedIncome, userId int) ([]ProjectedIncome, error) {
	rows, err := ag.Database.Query("SELECT * FROM projected WHERE user=?", userId)
	if err != nil {
		return []ProjectedIncome{}, err
	}

	defer rows.Close()

	for rows.Next() {
		projectedIncome := ProjectedIncome{}

		err := rows.Scan(&projectedIncome.ID, &projectedIncome.User, &projectedIncome.Amount)
		if err != nil {
			return []ProjectedIncome{}, err
		}

		projectedIncome.Amount = float64(projectedIncome.Amount) / 100

		allProjectedIncome = append(allProjectedIncome, projectedIncome)
	}

	err = rows.Err()
	if err != nil {
		return []ProjectedIncome{}, err
	}

	return allProjectedIncome, nil
}
