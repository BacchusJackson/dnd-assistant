package repositories

// Client interface
// An generic abstraction layer that can be implemented using any data store
type Client interface {
	Append(key string, value string) error
	Set(key string, field string, value string) error
	SetMap(key string, value map[string]string) error
	Get(key string) (map[string]string, error)
	Ping() error
	Clean() error
}
