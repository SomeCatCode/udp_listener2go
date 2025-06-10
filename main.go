package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Port int `json:"port"`
}

func getConfig() (*Config, error) {
	// Ermittle den Pfad der ausführbaren Datei und das Verzeichnis, in dem sie sich befindet.
	exPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Fehler beim Ermitteln des Executable-Pfads: %v", err)
	}
	exDir := filepath.Dir(exPath)

	configFile := filepath.Join(exDir, "udp_listener2go.json")

	// Überprüfe, ob Config-Datei existiert
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Erstelle Config-Datei mit Standardwerten
		defaultConfig := Config{
			Port: 1234,
		}

		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			log.Fatalf("Fehler beim Erzeugen der Standard-Config: %v", err)
			return nil, err
		}

		// Schreibe die Standard-Konfiguration in die Datei.
		err = os.WriteFile(configFile, data, 0644)
		if err != nil {
			log.Fatalf("Fehler beim Schreiben der Standard-Config: %v", err)
			return nil, err
		}
	}

	// Lese Config-Datei
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Config-Datei: %v", err)
		return nil, err
	}

	// Parse die Konfigurationsdaten in die Config-Struktur.
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Config-Datei: %v", err)
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration: %v", err)
		return
	}

	// Starte Server
	addr := net.UDPAddr{
		Port: config.Port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening on", conn.LocalAddr().String())

	buffer := make([]byte, 1024)

	for {
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		data := buffer[:n]
		fmt.Printf("Received packet from %v at %v\n", src, time.Now())

		if isJSON(data) {
			var obj map[string]interface{}
			if err := json.Unmarshal(data, &obj); err != nil {
				fmt.Println("Error decoding JSON:", err)
			} else {
				fmt.Println("Decoded JSON:", obj)
			}
		} else {
			fmt.Println("Received:", string(data))
		}
	}
}

func isJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}
