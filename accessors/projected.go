package accessors

import "github.com/labstack/echo"

type ProjectedIncome struct {
	ID     int     `json:"id,string"`
	User   int     `json:"user,string"`
	Amount float64 `json:"amount,string"`
	Earned float64 `json:"earned,string"`
	Spent  float64 `json:"spent,string"`
}

func (ag *AccessorGroup) SetProjectedIncome(context echo.Context, email string) (ProjectedIncome, error) {
	projectedIncome := ProjectedIncome{}
	err := context.Bind(&projectedIncome)
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

func (ag *AccessorGroup) UpdateProjectedIncome(context echo.Context, email string) (ProjectedIncome, error) {
	projectedIncome := ProjectedIncome{}
	err := context.Bind(&projectedIncome)
	if err != nil {
		return ProjectedIncome{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("UPDATE projected SET amount=? WHERE user=?", projectedIncome.Amount*100, projectedIncome.User)
	if err != nil {
		return ProjectedIncome{}, err
	}

	return ProjectedIncome{}, nil
}

func (ag *AccessorGroup) GetProjectedIncome(context echo.Context, email string) (ProjectedIncome, error) {
	userID, err := ag.GetUserID(email)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome, err := ag.GetProjectedIncomeByUserID(context, userID)
	if err != nil {
		return ProjectedIncome{}, err
	}

	return projectedIncome, nil
}

func (ag *AccessorGroup) GetProjectedIncomeByUserID(context echo.Context, userID int) (ProjectedIncome, error) {
	projectedIncome := ProjectedIncome{}

	err := ag.Database.QueryRow("SELECT * FROM projected WHERE user=?", userID).Scan(&projectedIncome.ID, &projectedIncome.User, &projectedIncome.Amount)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.Amount = float64(projectedIncome.Amount) / 100

	earned, err := ag.GetIncomeEarned(userID)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.Earned = earned

	projectedIncome.Earned = float64(projectedIncome.Earned) / 100

	spent, err := ag.GetExpenseTotal(userID)
	if err != nil {
		return ProjectedIncome{}, err
	}

	projectedIncome.Spent = spent

	return projectedIncome, nil
}
