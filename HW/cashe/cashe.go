package cashe

type Cashe struct {
	cashMap map[string]interface{}
}

func New() *Cashe {
	return &Cashe{
		cashMap: make(map[string]interface{}),
	}
}

func (c *Cashe) Set(key string, value interface{}) {
	c.cashMap[key] = value
}

func (c *Cashe) Get(key string) interface{} {
	return c.cashMap[key]
}

func (c *Cashe) Delete(key string) {
	delete(c.cashMap, key)
}
