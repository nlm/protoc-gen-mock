package main

type ServerNamer struct {
	BaseName string
}

func (sn ServerNamer) BaseServerName() string {
	return sn.BaseName + "Server"
}

func (sn ServerNamer) MockServerName() string {
	return "Mock" + sn.BaseServerName()
}

func (sn ServerNamer) MockServerDefaultsName() string {
	return sn.MockServerName() + "Defaults"
}
