package accessors

import "github.com/labstack/echo"

type Bucket struct {
	ID     int     `json:"id,string"`
	User   int     `json:"user,string"`
	Amount float64 `json:"amount,string"`
	Spent  float64 `json:"spent,string"`
	Name   string  `json:"name"`
}

func (accessorGroup *AccessorGroup) MakeBucket(context echo.Context, email string) (Bucket, error) {
	bucket := Bucket{}
	err := context.Bind(&bucket)
	if err != nil {
		return Bucket{}, err
	}

	userID, err := accessorGroup.GetUserID(email)
	if err != nil {
		return Bucket{}, err
	}

	bucket.User = userID

	// TODO: Make sure the information passed in is complete and don't submit if it's not

	_, err = accessorGroup.Database.Query("INSERT INTO buckets (user, amount, name) VALUES (?,?,?)", bucket.User, bucket.Amount*100, bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	return Bucket{}, nil
}

func (accessorGroup *AccessorGroup) GetBucket(context echo.Context, email string) ([]Bucket, error) {
	allBuckets := []Bucket{}

	userID, err := accessorGroup.GetUserID(email)
	if err != nil {
		return []Bucket{}, err
	}

	allBuckets, err = accessorGroup.GetBucketByUserID(context, allBuckets, userID)
	if err != nil {
		return []Bucket{}, err
	}

	return allBuckets, nil
}

func (accessorGroup *AccessorGroup) GetBucketSpent(userID int, bucketID int) (int, error) {
	spent := 0

	// Get the amount that's been spent
	err := accessorGroup.Database.QueryRow("SELECT COALESCE(SUM(amount),0) FROM expenses WHERE user=? AND bucket=? AND MONTH(time) = MONTH(CURDATE())", userID, bucketID).Scan(&spent)
	if err != nil {
		return 0, err
	}

	return spent, nil
}

func (accessorGroup *AccessorGroup) GetBucketByUserID(context echo.Context, allBuckets []Bucket, userID int) ([]Bucket, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM buckets WHERE user=? ORDER BY name", userID)
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
		spent, err := accessorGroup.GetBucketSpent(userID, bucket.ID)
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

func (accessorGroup *AccessorGroup) GetBucketByName(context echo.Context, email string) (Bucket, error) {
	userID, err := accessorGroup.GetUserID(email)
	if err != nil {
		return Bucket{}, err
	}

	bucket, err := accessorGroup.GetBucketByNameAndUserID(context, context.Param("name"), userID)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}

func (accessorGroup *AccessorGroup) GetBucketByNameAndUserID(context echo.Context, name string, userID int) (Bucket, error) {
	bucket := Bucket{}

	// Get general info about the bucket
	err := accessorGroup.Database.QueryRow("SELECT * FROM buckets WHERE user=? AND name=?", userID, name).Scan(&bucket.ID, &bucket.User, &bucket.Amount, &bucket.Name)
	if err != nil {
		return Bucket{}, err
	}

	// Get the amount that's been spent
	spent, err := accessorGroup.GetBucketSpent(userID, bucket.ID)
	if err != nil {
		return Bucket{}, err
	}

	bucket.Spent = float64(spent) / 100

	return bucket, nil
}
