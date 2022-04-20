package manager

type IManager interface {
	SetSuperior(IManager)
	Request(Request)
}

type Manager struct {
	Superior IManager
	Name     string
}

func (t *Manager) SetSuperior(manager IManager) {
	t.Superior = manager
}

func (t Manager) Request(request Request) {

}
