package taskman

// DataInterface describes the interface for saving and loading taskman data
type DataInterface interface {
	Setup(dataPath string) error
	Save(data Data) error
	Load() (*Data, error)
}
