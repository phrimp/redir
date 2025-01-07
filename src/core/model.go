package core

type CoreModel interface {
	Create(map[string]string) (CoreModel, error)
	Update(CoreModel) (CoreModel, error)
}
