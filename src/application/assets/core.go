package assets

type Core struct {
	port           string
	com_port       string
	resourcedir    string
	awake_interval int64
}

func (obj *Core) SetPort(port string) {
	if len(port) == 0 || port == "" {
		obj.port = "8080"
		return
	}
	obj.port = port
}

func (obj *Core) GetPort() string {
	return obj.port
}

func (obj *Core) SetResourceDir(resourceDir string) {
	if len(resourceDir) == 0 || resourceDir == "" {
		obj.resourcedir = ""
		return
	}
	obj.resourcedir = resourceDir
}

func (obj *Core) GetResourceDir() string {
	return obj.resourcedir
}

func (obj *Core) SetComPort(comport string) {
	if len(comport) == 0 || comport == "" {
		obj.com_port = "COM5"
		return
	}
	obj.com_port = comport
}

func (obj *Core) GetComPort() string {
	return obj.com_port
}

func (obj *Core) SetAwakeInterval(interval int64) {
	if interval < 1 {
		interval = 1
	}
	obj.awake_interval = interval
}

func (obj *Core) GetAwakeInterval() int64 {
	return obj.awake_interval
}
