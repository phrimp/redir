package core

type CoreModel interface {
	Create(map[string]interface{}) (CoreModel, error)
	Update(CoreModel) (CoreModel, error)
}
