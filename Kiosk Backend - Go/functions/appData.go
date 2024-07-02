package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type appConfig struct {
	Timers struct {
		Popup          int `json:"popup"`
		Worning        int `json:"worning"`
		Success        int `json:"success"`
		Error          int `json:"error"`
		Print          int `json:"print"`
		FloorPage      int `json:"floorPage"`
		BookingsDetail int `json:"bookingsDetail"`
		Keyboard       int `json:"keyboard"`
	}
	DisplaySectorGroupBy string `json:"displaySectorGroupBy"`
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH") + "\\AppData" + "\\Local\\HCFILES"
		if home == "" {
			home = os.Getenv("USERPROFILE") + "\\AppData" + "\\Local\\HCFILES"
		}
		return home
	}
	return os.Getenv("HOME") + "\\AppData" + "\\Local\\HCFILES"
}

func CreateConfigFile() {
	// Create a CONFIG.json file if it does not exist in the user's home directory
	if _, err := os.Stat(UserHomeDir() + "\\CONFIG.json"); os.IsNotExist(err) {
		file, err := os.Create(UserHomeDir() + "\\CONFIG.json")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		// create a json object and write it to the file
		Config := appConfig{}
		Config.Timers.Popup = 20
		Config.Timers.Worning = 1
		Config.Timers.Success = 1
		Config.Timers.Error = 1
		Config.Timers.Print = 10
		Config.Timers.FloorPage = 90
		Config.Timers.BookingsDetail = 60
		Config.Timers.Keyboard = 15
		Config.DisplaySectorGroupBy = "FLOOR"

		jsonData, err := json.Marshal(Config)
		if err != nil {
			fmt.Println(err)
			return
		}
		file.Write(jsonData)
	}
}

func GetConfig() appConfig {
	// read the CONFIG.json file
	file, err := os.Open(UserHomeDir() + "\\CONFIG.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// decode the json file
	Config := appConfig{}
	json.NewDecoder(file).Decode(&Config)
	return Config
}

func SaveConfig(Config appConfig) {
	// create a json object and write it to the file
	jsonData, err := json.Marshal(Config)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile(UserHomeDir()+"\\CONFIG.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(jsonData)
}

func SECTOR_IMAGE_TO_BASE64(path string) string {
	// replace / to \ for windows
	path = strings.Replace(path, "/", "\\", -1)
	fmt.Println(UserHomeDir() + path)
	imageData, err := os.ReadFile(UserHomeDir() + path)
	if err != nil {
		return ""
	}
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	// Create a data URL
	mimeType := "image/jpeg" // Change this according to your image's MIME type
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)
	return dataURL
}
