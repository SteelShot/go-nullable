module github.com/SteelShot/go-nullable

go 1.18.0

retract [v0.0.0-0, v0.1.0-retract] // Drop old version due to API change, lowercase module import

require gopkg.in/yaml.v3 v3.0.1
