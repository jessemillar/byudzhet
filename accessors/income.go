package accessors

import "github.com/labstack/echo"

type Income struct {
	ID     int     `json:"id,string"`
	User   int     `json:"user,string"`
	Time   string  `json:"time"`
	Payer  string  `json:"payer"`
	Amount float64 `json:"amount,string"`
}

func (ag *AccessorGroup) LogIncome(context echo.Context, email string) (Income, error) {
	income := Income{}
	err := context.Bind(&income)
	if err != nil {
		return Income{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Income{}, err
	}

	income.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO income (user, payer, amount) VALUES (?,?,?)", income.User, income.Payer, income.Amount*100)
	if err != nil {
		return Income{}, err
	}

	return Income{}, nil
}

func (ag *AccessorGroup) GetIncome(context echo.Context, email string) ([]Income, error) {
	allIncome := []Income{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Income{}, err
	}

	allIncome, err = ag.GetIncomeByUserID(context, allIncome, userID)
	if err != nil {
		return []Income{}, err
	}

	return allIncome, nil
}

func (ag *AccessorGroup) GetIncomeEarned(userID int) (float64, error) {
	var earned float64

	// Get the amount that's been earned
	err := ag.Database.QueryRow("SELECT COALESCE(SUM(amount),0) FROM income WHERE user=?", userID).Scan(&earned)
	if err != nil {
		return 0, err
	}

	return earned, nil
}

func (ag *AccessorGroup) GetIncomeByUserID(context echo.Context, allIncome []Income, userID int) ([]Income, error) {
	rows, err := ag.Database.Query("SELECT * FROM income WHERE user=? AND MONTH(time) = MONTH(CURDATE()) ORDER BY time DESC", userID)
	if err != nil {
		return []Income{}, err
	}

	defer rows.Close()

	for rows.Next() {
		income := Income{}

		err := rows.Scan(&income.ID, &income.User, &income.Time, &income.Payer, &income.Amount)
		if err != nil {
			return []Income{}, err
		}

		income.Amount = float64(income.Amount) / 100

		allIncome = append(allIncome, income)
	}

	err = rows.Err()
	if err != nil {
		return []Income{}, err
	}

	return allIncome, nil
}
