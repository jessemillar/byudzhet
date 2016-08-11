package accessors

import "github.com/labstack/echo"

type Expense struct {
	ID        int     `json:"id,string"`
	User      int     `json:"user,string"`
	Time      string  `json:"time"`
	Bucket    int     `json:"bucket,string"`
	Amount    float64 `json:"amount,string"`
	Recipient string  `json:"recipient"`
	Note      string  `json:"note"`
}

func (accessorGroup *AccessorGroup) LogExpense(context echo.Context, email string) (Expense, error) {
	expense := Expense{}
	err := context.Bind(&expense)
	if err != nil {
		return Expense{}, err
	}

	userID, err := accessorGroup.GetUserID(email)
	if err != nil {
		return Expense{}, err
	}

	expense.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = accessorGroup.Database.Query("INSERT INTO expenses (user, bucket, amount, recipient, note) VALUES (?,?,?,?,?)", expense.User, expense.Bucket, expense.Amount*100, expense.Recipient, expense.Note)
	if err != nil {
		return Expense{}, err
	}

	return Expense{}, nil
}

func (accessorGroup *AccessorGroup) GetExpense(context echo.Context, email string) ([]Expense, error) {
	allExpenses := []Expense{}

	userID, err := accessorGroup.GetUserID(email)
	if err != nil {
		return []Expense{}, err
	}

	allExpenses, err = accessorGroup.GetExpenseByUserID(allExpenses, userID)
	if err != nil {
		return []Expense{}, err
	}

	return allExpenses, nil
}

func (accessorGroup *AccessorGroup) GetExpenseByUserID(expenses []Expense, userID int) ([]Expense, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM expenses WHERE user=? AND MONTH(time) = MONTH(CURDATE()) ORDER BY time DESC", userID)
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

		expense.Amount = float64(expense.Amount) / 100

		expenses = append(expenses, expense)
	}

	err = rows.Err()
	if err != nil {
		return []Expense{}, err
	}

	return expenses, nil
}

func (accessorGroup *AccessorGroup) GetExpenseTotal(userID int) (float64, error) {
	var total float64

	err := accessorGroup.Database.QueryRow("SELECT COALESCE(SUM(amount),0) FROM expenses WHERE user=? AND MONTH(time) = MONTH(CURDATE())", userID).Scan(&total)
	if err != nil {
		return 0, err
	}

	total = total / 100

	return total, nil
}
