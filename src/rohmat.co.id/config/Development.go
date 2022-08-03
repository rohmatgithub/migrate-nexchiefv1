package config

type DevelopmentConfig struct {
	Configuration
	Postgresql struct {
		Address           string `json:"address"`
		DefaultSchema     string `json:"default_schema"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql"`
	DirData DirPath `json:"directory_data"`
}

type DirPath struct {
	PathDir            string `json:"path_dir"`
	UserLevel1         string `json:"user_level_1"`
	UserLevel2         string `json:"user_level_2"`
	UserLevel3         string `json:"user_level_3"`
	UserLevel4         string `json:"user_level_4"`
	UserLevel5         string `json:"user_level_5"`
	Distributor        string `json:"distributor"`
	MappingDistributor string `json:"mapping_distributor"`
	Division           string `json:"division"`
	Vendor             string `json:"vendor"`
	Salesman           string `json:"salesman"`
	Customer           string `json:"customer"`
}

func (input DevelopmentConfig) GetPostgreSQLAddress() string {
	return input.Postgresql.Address
}
func (input DevelopmentConfig) GetPostgreSQLDefaultSchema() string {
	return input.Postgresql.DefaultSchema
}
func (input DevelopmentConfig) GetPostgreSQLMaxOpenConnection() int {
	return input.Postgresql.MaxOpenConnection
}
func (input DevelopmentConfig) GetPostgreSQLMaxIdleConnection() int {
	return input.Postgresql.MaxIdleConnection
}
func (input DevelopmentConfig) GetDirPath() DirPath {
	return input.DirData
}
