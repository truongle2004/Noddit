package config

import (
	"context"
	"fmt"
	"log"
	"profile-service/internal/environment"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitNeo4j() (neo4j.DriverWithContext, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		environment.Neo4jUri,
		neo4j.BasicAuth(environment.Neo4jUsername, environment.Neo4jPassword, ""))

	if err != nil {
		return nil, fmt.Errorf("failed to create Neo4j driver: %w", err)
	}

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, fmt.Errorf("failed to verify Neo4j connectivity: %w", err)
	}

	log.Println(" Connection established.")
	return driver, nil
}
