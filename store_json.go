package taskman

import (
	"encoding/json"
	"os"
	"path"
)

type jsonDataStoreType struct {
	fileName string
	filePath string
}

// JSONDataStore a data store that saves the data in a JSON file
var JSONDataStore = jsonDataStoreType{
	fileName: "taskman_config.json",
}

func (j jsonDataStoreType) Setup(dataPath string) error {
	j.filePath = path.Join(dataPath, j.fileName)

	if _, err := j.Load(); err != nil {
		return err
	}

	return nil
}

func (j jsonDataStoreType) Save(data Data) error {
	filePath := path.Join(j.filePath, j.fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error("Error writing JSON data store at path '%s': %s", filePath, err.Error())
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(&data); err != nil {
		log.Error("Error writing JSON data store at path '%s': %s", filePath, err.Error())
		return err
	}

	return nil
}

func (j jsonDataStoreType) Load() (*Data, error) {
	filePath := path.Join(j.filePath, j.fileName)
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error("Error reading JSON data store at path '%s': %s", filePath, err.Error())
		return nil, err
	}
	defer f.Close()

	data := defaultData
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		log.Error("Error reading JSON data store at path '%s': %s", filePath, err.Error())
		return nil, err
	}
	return &data, nil
}
