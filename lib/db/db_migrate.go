package db

func MigrateAll(repos ...IDbRepo) error {
	for _, repo := range repos {
		err := repo.AutoMigration()
		if err != nil {
			panic(err)
		}
	}

	return nil
}
