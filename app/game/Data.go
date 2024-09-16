package game

import (
	"encoding/json"
	"io"
	"os"
)

type GameData struct {
	HighScore int `json:"highScore"`
}

func (data *GameData) LoadData() {
	if err := os.MkdirAll(dataDirPath, os.ModePerm); err != nil {
		panic(err)
	}

	jsonFile, err := os.OpenFile(dataFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		panic(err)
	}

	data.SaveData()
}

func (g *GameData) SaveData() {
	if err := os.MkdirAll(dataDirPath, os.ModePerm); err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.OpenFile(dataFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	if _, err := jsonFile.Write(jsonBytes); err != nil {
		panic(err)
	}
}
