# udp_listener2go

udp_listener2go ist ein kleiner UDP Listener für Debug zwecke. Sollte ein JSON content ankommen, so wird dieser geparsed ausgegeben.

## Verwendung
Nach dem Klonen des Repositories können Sie das Programm mit go run . ausführen oder eine ausführbare Datei mit go build erstellen.

## Konfiguration
Die Konfigurationsdatei config.json sollte im selben Verzeichnis wie die ausführbare Datei liegen. Sie sollte folgende Felder enthalten:

port: Port auf den gelauscht werden soll

Beispiel:

```json
{
"port": 2121
}
```

## Lizenz
Dieses Projekt steht unter der MIT-Lizenz - siehe die LICENSE.md Datei für weitere Informationen.

