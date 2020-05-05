package taskman

type memoryDataStoreType struct{}

// InMemoryDataStore a data store that operates in memory only
var InMemoryDataStore = memoryDataStoreType{}

func (j memoryDataStoreType) Setup(dataPath string) error {
	return nil
}

func (j memoryDataStoreType) Save(data Data) error {
	return nil
}

func (j memoryDataStoreType) Load() (*Data, error) {
	return &defaultData, nil
}
