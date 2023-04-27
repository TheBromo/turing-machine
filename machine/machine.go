package machine

type Machine struct {
	Name string
}

func (m *Machine) GetTestString() string {
	return m.Name

}
