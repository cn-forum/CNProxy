<p align="center">
  <img src="5891a253da6faea29a4d326c6816e949a.png.0c8e73bdc5bcf40c017d036523a71d6a.png" alt="Logo" width="600"/>
</p>


## Überblick

Dieses Projekt ist Teil von CrimeNow, da wir großen Wert auf Transparenz in unseren Sicherheitsmaßnahmen legen. Mit der hier vorgestellten Methode haben wir ein vollständig integriertes Proxy-Netzwerk aufgebaut. Um es kurz zu erklären: Wir betreiben über 20 Server, von denen jeder drei Ports offen hat. Dieses System ist ausschließlich für die interne Kommunikation ausgelegt – externe Verbindungen werden konsequent abgelehnt.

Für zusätzliche Sicherheit empfehlen wir den Einsatz von SSH-Keys und das Konfigurieren von IPTables, um alle Verbindungen außer denen aus dem Proxy-Netzwerk zu blockieren. Idealerweise nutzt ihr einen Hauptserver mit <a href="https://github.com/41Baloo/balooProxy">BalooProxy</a>. Dieser fungiert als zentrale Anlaufstelle für HTTP-Anfragen, integriert eine Web Application Firewall (WAF) und kann auch als eigenständiger Proxy arbeiten. Von dort leitet ihr die Anfragen direkt in euer Netzwerk weiter.

Ein besonderer Vorteil ist, dass ihr auf einem einzigen Port mehrere Ziele definieren könnt, wodurch die Verbindung jedes Mal einen zufälligen Pfad nimmt. CNProxy selbst speichert keine IP-Adressen und nutzt keinen Proxy-Cache, was maximale Privatsphäre gewährleistet.

Wichtig: Beim Compilen benutzt <a href="https://github.com/burrowers/garble">garble</a> als Obfuscator und jagt die Output-Datei nochmal durch <a href="https://github.com/upx/upx">upx</a>, damit sie gepackt wird.

## Funktionen

- 🔒 IP-Whitelisting: Akzeptiert Connections nur von bestimmten IP-Adressen.
- 🌐 Wildcard-Unterstützung: Benutze `*`, um Verbindungen von allen IP-Adressen zu erlauben.
- 📡 Multi Proxy: Starte mehrere Proxy gleichzeitig auf unterschiedlichen Ports mit separaten Configs.
- 🎯 Load Balancing: Wählt für jede Connection zufällig einen der Target Server aus.

---

```go
ProxyConfig{
    LocalPort: "1337",
    Targets:   []string{"127.0.0.1:8080"}, // Targets angeben Port leitet auf eine Zufällige weiter.
    AllowedIPs: []string{
        "*", // Entweder * um Alle IP's zu whitelisten oder bestimmte IP's anzunehmen.
    },
}
```
