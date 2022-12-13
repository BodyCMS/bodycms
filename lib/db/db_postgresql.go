package db

type DbPostgreSql struct {
	DbOptions
}

func (d *DbPostgreSql) ConnectionString() string {
	return d.GetConnectionString()
}

func (d *DbPostgreSql) Connect() error {
	return nil
}

func (d *DbPostgreSql) Close() error {
	return nil
}
