package store

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func startSession(driver neo4j.Driver, dbName string) neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{DatabaseName: dbName, AccessMode: neo4j.AccessModeWrite})
}

func endSession(session neo4j.Session) {
	err := session.Close()
	if err != nil {
		log.Println(err)
	}
}
