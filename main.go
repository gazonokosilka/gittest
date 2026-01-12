package main

import (
	"gitprc/feature"
	simpleconnection "gitprc/featurePostgres/simpleConnection"
)

func main() {
	feature.DaetSalam()
	simpleconnection.CheckConnection()
}
