package assets

type Core struct {
	port        string
	resourcedir string
}

func (obj *Core) SetPort(port string) {
	obj.port = port
}

func (obj *Core) GetPort() string {
	return obj.port
}

func (obj *Core) SetResourceDir(resourceDir string) {
	obj.resourcedir = resourceDir
}

func (obj *Core) GetResourceDir() string {
	return obj.resourcedir
}
