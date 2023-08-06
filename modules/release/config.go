package release

// Config for a releaseService
type Config struct {
	Enabled        bool   `mapstructure:"enabled"`
	DatabaseDriver string `mapstructure:"databaseDriver"`
	DatabaseUrl    string `mapstructure:"databaseUrl"`
}
