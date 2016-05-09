package accessors

import "github.com/labstack/echo"

type Expense struct {
	ID        int
	User      int
	Timestamp string
	Bucket    int `json:",string"`
	Amount    int `json:",string"`
	Recipient string
	Note      string
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

// func (ag *AccessorGroup) GetExpense()
