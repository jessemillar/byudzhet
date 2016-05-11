package accessors

import "github.com/labstack/echo"

type Expense struct {
	ID        int    `json:"id"`
	User      int    `json:"user"`
	Time      string `json:"time"`
	Bucket    int    `json:"bucket,string"`
	Amount    int    `json:"amount,string"`
	Recipient string `json:"recipient"`
	Note      string `json:"note"`
}

func (ag *AccessorGroup) LogExpense(c echo.Context, email string) (Expense, error) {
	expense := Expense{}
	err := c.Bind(&expense)
	if err != nil {
		return Expense{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Expense{}, err
	}

	expense.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO expenses (user, bucket, amount, recipient, note) VALUES (?,?,?,?,?)", expense.User, expense.Bucket, expense.Amount, expense.Recipient, expense.Note)
	if err != nil {
		return Expense{}, err
	}

	return Expense{}, nil
}

func (ag *AccessorGroup) GetExpense(c echo.Context, email string) ([]Expense, error) {
	allExpenses := []Expense{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Expense{}, err
	}

	allExpenses, err = ag.GetExpenseByUserID(c, allExpenses, userID)
	if err != nil {
		return []Expense{}, err
	}

	allShares, err := ag.GetSharing(c, email)
	if err != nil {
		return []Expense{}, err
	}

	for i := range allShares {
		if allShares[i].User != userID {
			allExpenses, err = ag.GetExpenseByUserID(c, allExpenses, allShares[i].User)
			if err != nil {
				return []Expense{}, err
			}
		} else if allShares[i].Sharee != userID {
			allExpenses, err = ag.GetExpenseByUserID(c, allExpenses, allShares[i].Sharee)
			if err != nil {
				return []Expense{}, err
			}
		}
	}

	return allExpenses, nil
}

func (ag *AccessorGroup) GetExpenseByUserID(c echo.Context, expenses []Expense, id int) ([]Expense, error) {
	rows, err := ag.Database.Query("SELECT * FROM expenses WHERE user=? AND MONTH(time) = MONTH(CURDATE())", id)
	if err != nil {
		return []Expense{}, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := Expense{}

		err := rows.Scan(&expense.ID, &expense.User, &expense.Time, &expense.Bucket, &expense.Amount, &expense.Recipient, &expense.Note)
		if err != nil {
			return []Expense{}, err
		}

		expenses = append(expenses, expense)
	}

	err = rows.Err()
	if err != nil {
		return []Expense{}, err
	}

	return expenses, nil
}
