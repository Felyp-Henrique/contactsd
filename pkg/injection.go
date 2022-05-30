package pkg

type Injection struct {
	objects map[string]interface{}
}

func NewInjection() *Injection {
	return &Injection{
		objects: map[string]interface{}{},
	}
}

func (i *Injection) Get(tag string) interface{} {
	return i.objects[tag]
}

func (i *Injection) Put(tag string, object interface{}) {
	i.objects[tag] = object
}
