package accessors

import "github.com/labstack/echo"

type Bucket struct {
	ID     int     `json:"id,string"`
	User   int     `json:"user,string"`
	Amount float64 `json:"amount,string"`
	Spent  float64 `json:"spent,string"`
	Name   string  `json:"name"`
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

	_, err = ag.Database.Query("INSERT INTO buckets (user, amount, name) VALUES (?,?,?)", bucket.User, bucket.Amount*100, bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	return Bucket{}, nil
}

func (ag *AccessorGroup) GetBucket(c echo.Context, email string) ([]Bucket, error) {
	allBuckets := []Bucket{}

	userID, err := ag.GetUserID(email)
	if err != nil {
		return []Bucket{}, err
	}

	allBuckets, err = ag.GetBucketByUserID(c, allBuckets, userID)
	if err != nil {
		return []Bucket{}, err
	}

	return allBuckets, nil
}

func (ag *AccessorGroup) GetBucketSpent(userID int, bucketID int) (int, error) {
	spent := 0

	// Get the amount that's been spent
	err := ag.Database.QueryRow("SELECT COALESCE(SUM(amount),0) FROM expenses WHERE user=? AND bucket=?", userID, bucketID).Scan(&spent)
	if err != nil {
		return 0, err
	}

	return spent, nil
}

func (ag *AccessorGroup) GetBucketByUserID(c echo.Context, allBuckets []Bucket, userID int) ([]Bucket, error) {
	rows, err := ag.Database.Query("SELECT * FROM buckets WHERE user=? ORDER BY name", userID)
	if err != nil {
		return []Bucket{}, err
	}

	defer rows.Close()

	for rows.Next() {
		bucket := Bucket{}

		// Get general info about the bucket
		err := rows.Scan(&bucket.ID, &bucket.User, &bucket.Amount, &bucket.Name)
		if err != nil {
			return []Bucket{}, err
		}

		bucket.Amount = float64(bucket.Amount) / 100

		// Get the amount that's been spent
		spent, err := ag.GetBucketSpent(userID, bucket.ID)
		if err != nil {
			return []Bucket{}, err
		}

		bucket.Spent = float64(spent) / 100

		allBuckets = append(allBuckets, bucket)
	}

	err = rows.Err()
	if err != nil {
		return []Bucket{}, err
	}

	return allBuckets, nil
}

func (ag *AccessorGroup) GetBucketByName(c echo.Context, email string) (Bucket, error) {
	userID, err := ag.GetUserID(email)
	if err != nil {
		return Bucket{}, err
	}

	bucket, err := ag.GetBucketByNameAndUserID(c, c.Param("name"), userID)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}

func (ag *AccessorGroup) GetBucketByNameAndUserID(c echo.Context, name string, userID int) (Bucket, error) {
	bucket := Bucket{}

	// Get general info about the bucket
	err := ag.Database.QueryRow("SELECT * FROM buckets WHERE user=? AND name=?", userID, name).Scan(&bucket.ID, &bucket.User, &bucket.Amount, &bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	// Get the amount that's been spent
	spent, err := ag.GetBucketSpent(userID, bucket.ID)
	if err != nil {
		return Bucket{}, err
	}

	bucket.Spent = float64(spent) / 100

	return bucket, nil
}
