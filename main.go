package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

var version = "dev"

type Config struct {
	Port int `json:"port"`
}

func getConfig() (*Config, error) {
	exPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Executable-Pfad nicht ermittelbar: %w", err)
	}
	configFile := filepath.Join(filepath.Dir(exPath), "udp_listener2go.json")

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		data, err := json.MarshalIndent(Config{Port: 1234}, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("Standard-Config nicht erzeugbar: %w", err)
		}
		if err := os.WriteFile(configFile, data, 0644); err != nil {
			return nil, fmt.Errorf("Standard-Config nicht schreibbar: %w", err)
		}
		log.Printf("Standard-Config erstellt: %s", configFile)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("Config-Datei nicht lesbar: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Config-Datei %s ungültig: %w", configFile, err)
	}

	return &config, nil
}

func main() {
	log.Printf("udp_listener2go %s", version)

	config, err := getConfig()
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration: %v", err)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: config.Port})
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer conn.Close()

	log.Printf("Listening on %s", conn.LocalAddr())

	// max. UDP-Payload
	buffer := make([]byte, 65535)

	for {
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading: %v", err)
			continue
		}
		data := bytes.TrimSpace(buffer[:n])
		log.Printf("Received %d bytes from %v", n, src)

		var pretty bytes.Buffer
		if json.Valid(data) && json.Indent(&pretty, data, "", "  ") == nil {
			fmt.Println(pretty.String())
		} else {
			fmt.Println(string(data))
		}
	}
}
