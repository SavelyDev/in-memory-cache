package cache

func New() *cache {
	return &cache{
		storage: make(map[string]any),
	}
}
