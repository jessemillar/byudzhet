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

func (ag *AccessorGroup) UpdateProjectedIncome(c echo.Context, email string) (ProjectedIncome, error) {
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

	_, err = ag.Database.Query("UPDATE projected SET amount=? WHERE user=? (?,?)", projectedIncome.Amount*100, projectedIncome.User)
	if err != nil {
		return ProjectedIncome{}, err
	}

	return ProjectedIncome{}, nil
}

func (ag *AccessorGroup) GetProjectedIncome(c echo.Context, email string) (ProjectedIncome, error) {
	userID, err := ag.GetUserID(email)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome, err := ag.GetProjectedIncomeByUserID(c, userID)
	if err != nil {
		return ProjectedIncome{}, err
	}

	return projectedIncome, nil
}

func (ag *AccessorGroup) GetProjectedIncomeByUserID(c echo.Context, userId int) (ProjectedIncome, error) {
	projectedIncome := ProjectedIncome{}

	err := ag.Database.QueryRow("SELECT * FROM projected WHERE user=?", userId).Scan(&projectedIncome.ID, &projectedIncome.User, &projectedIncome.Amount)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.Amount = float64(projectedIncome.Amount) / 100

	return projectedIncome, nil
}
