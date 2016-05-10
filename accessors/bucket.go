package accessors

import "github.com/labstack/echo"

type Bucket struct {
	ID     int    `json:"id"`
	User   int    `json:"user"`
	Amount int    `json:"amount,string"`
	Name   string `json:"name"`
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

func (ag *AccessorGroup) GetBucket(c echo.Context, email string) ([]Bucket, error) {
	buckets := []Bucket{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Bucket{}, err
	}

	rows, err := ag.Database.Query("SELECT * FROM buckets WHERE user=?", userID)
	if err != nil {
		return []Bucket{}, err
	}

	defer rows.Close()

	for rows.Next() {
		bucket := Bucket{}

		err := rows.Scan(&bucket.ID, &bucket.User, &bucket.Amount, &bucket.Name)
		if err != nil {
			return []Bucket{}, err
		}

		buckets = append(buckets, bucket)
	}

	err = rows.Err()
	if err != nil {
		return []Bucket{}, err
	}

	return buckets, nil
}

func (ag *AccessorGroup) GetBucketByName(c echo.Context, email string) (Bucket, error) {
	bucket := Bucket{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return Bucket{}, err
	}

	err = ag.Database.QueryRow("SELECT * FROM buckets WHERE user=? AND name=?", userID, c.Param("bucket")).Scan(&bucket.ID, &bucket.User, &bucket.Amount, &bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}
