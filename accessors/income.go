package accessors

import "github.com/labstack/echo"

type Income struct {
	ID     int    `json:"id"`
	User   int    `json:"user"`
	Time   string `json:"time"`
	Payer  string `json:"payer"`
	Amount int    `json:"amount,string"`
}

func (ag *AccessorGroup) LogIncome(c echo.Context, email string) (Income, error) {
	income := Income{}
	err := c.Bind(&income)
	if err != nil {
		return Income{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Income{}, err
	}

	income.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO income (user, payer, amount) VALUES (?,?,?)", income.User, income.Payer, income.Amount)
	if err != nil {
		return Income{}, err
	}

	return Income{}, nil
}

func (ag *AccessorGroup) GetIncome(c echo.Context, email string) ([]Income, error) {
	allIncome := []Income{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Income{}, err
	}

	allIncome, err = ag.GetIncomeByUserID(c, allIncome, userID)
	if err != nil {
		return []Income{}, err
	}

	return allIncome, nil
}

func (ag *AccessorGroup) GetIncomeByUserID(c echo.Context, allIncome []Income, id int) ([]Income, error) {
	rows, err := ag.Database.Query("SELECT * FROM income WHERE user=? AND MONTH(time) = MONTH(CURDATE())", id)
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

		allIncome = append(allIncome, income)
	}

	err = rows.Err()
	if err != nil {
		return []Income{}, err
	}

	return allIncome, nil
}
