# udp_listener2go

udp_listener2go ist ein kleiner UDP Listener für Debug zwecke. Sollte ein JSON content ankommen, so wird dieser formatiert ausgegeben.

## Download
Fertige Binaries für Windows, Linux und macOS liegen unter [Releases](https://github.com/SomeCatCode/udp_listener2go/releases). Die Binaries haben keine Abhängigkeiten.

## Release erstellen
Ein Tag `v*` startet den GitHub-Workflow. Der Workflow baut alle Binaries und erstellt das Release.

```bash
git tag v1.2.0
git push origin v1.2.0
```

## Verwendung
Nach dem Klonen des Repositories können Sie das Programm mit go run . ausführen oder eine ausführbare Datei mit go build erstellen.

## Konfiguration
Die Konfigurationsdatei `udp_listener2go.json` sollte im selben Verzeichnis wie die ausführbare Datei liegen. Sie sollte folgende Felder enthalten:

port: Port auf den gelauscht werden soll

Beispiel:

```json
{
"port": 1234
}
```

## Lizenz
Dieses Projekt steht unter der MIT-Lizenz - siehe die LICENSE Datei für weitere Informationen.

