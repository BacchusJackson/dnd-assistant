package repositories

// Client interface
// An generic abstraction layer that can be implemented using any data store
type Client interface {
	Append(key string, value string) error
	Update(key string, field string, value string) error
	Write(key string, value map[string]string) error
	Read(key string) (map[string]string, error)
	Ping() error
	Clean() error
}
