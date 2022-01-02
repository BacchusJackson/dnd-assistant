package repositories

type Mappable interface {
	Map() map[string]string
	GetId() string
}
