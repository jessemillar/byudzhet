package accessors

import "github.com/labstack/echo"

type Bucket struct {
	ID     int
	User   int
	Amount int `json:",string"`
	Name   string
}

func (ag *AccessorGroup) MakeBucket(c echo.Context, email string) (Bucket, error) {
	bucket := Bucket{}
	err := c.Bind(&bucket)
	if err != nil {
		return Bucket{}, err
	}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Bucket{}, err
	}

	bucket.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = ag.Database.Query("INSERT INTO buckets (user, amount, name) VALUES (?,?,?)", bucket.User, bucket.Amount, bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	return Bucket{}, nil
}

// func (ag *AccessorGroup) GetBucket()
