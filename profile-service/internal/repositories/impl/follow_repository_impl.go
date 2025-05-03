package impl

import (
	"context"
	"fmt"
	"log"
	"profile-service/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowRepositoryImpl struct {
	db neo4j.DriverWithContext
}

func NewFollowRepository(db neo4j.DriverWithContext) *FollowRepositoryImpl {
	return &FollowRepositoryImpl{
		db: db,
	}
}

func (f *FollowRepositoryImpl) Follow(c *gin.Context, followerId, followeeId string) error {
	session := f.db.NewSession(c, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Println("Error closing session:", err)
		}
	}(session, c)

	_, err := session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
			MATCH (follower:Profile {userId: $followerId})
			MATCH (followee:Profile {userId: $followeeId})
			WHERE follower.userId <> followee.userId  
			AND NOT (follower)-[:FOLLOWS]->(followee) 
			MERGE (follower)-[:FOLLOWS]->(followee)
			RETURN follower.userId as followerId, followee.userId as followeeId
		`

		params := map[string]interface{}{
			"followerId": followerId, "followeeId": followeeId,
		}

		result, err := tx.Run(c, query, params)

		if err != nil {
			log.Println("Error creating follow relationship:", err)
			return nil, err
		}

		if result.Next(c) {
			values := result.Record().Values
			log.Println("Follow relationship created between:", values)
		} else if err := result.Err(); err != nil {
			log.Println("Error consuming result:", err)
			return nil, err
		} else {
			// No record -> match failed -> nothing created
			return nil, fmt.Errorf("failed to create follow relationship")
		}

		return nil, err
	})

	if err != nil {
		log.Println("Error creating follow relationship:", err)
		return err
	}

	return nil
}

func (f *FollowRepositoryImpl) UnFollow(c *gin.Context, followerId, followeeId string) error {
	session := f.db.NewSession(c, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Println("Error closing session:", err)
		}
	}(session, c)

	_, err := session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
			MATCH (follower:Profile {userId: $followerId})-[r:FOLLOWS]->(followee:Profile {userId: $followeeId})
			DELETE r
		`

		params := map[string]interface{}{
			"followerId": followerId,
			"followeeId": followeeId,
		}

		_, err := tx.Run(c, query, params)
		return nil, err
	})

	return err
}

func (f *FollowRepositoryImpl) GetFollows(c *gin.Context, userId string) ([]domain.Profile, error) {
	session := f.db.NewSession(c, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Println("Error closing session:", err)
		}
	}(session, c)

	result, err := session.ExecuteRead(c, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
			MATCH (follower:Profile)-[:FOLLOWS]->(user:Profile {userId: $userId})
			RETURN follower.userId AS userId, follower.display_name AS display_name, follower.avatar_url AS avatar_url
		`
		params := map[string]interface{}{
			"userId": userId,
		}

		res, err := tx.Run(c, query, params)
		if err != nil {
			return nil, err
		}

		var profiles []domain.Profile

		for res.Next(c) {
			record := res.Record()

			userIdVal, _ := record.Get("userId")
			displayName, _ := record.Get("display_name")
			avatarUrl, _ := record.Get("avatar_url")

			profile := domain.Profile{
				UserID:      userIdVal.(string),
				DisplayName: displayName.(string),
				AvatarURL:   avatarUrl.(string),
			}
			profiles = append(profiles, profile)
		}

		if err = res.Err(); err != nil {
			return nil, err
		}

		return profiles, nil
	})

	if err != nil {
		log.Println("Error getting follows:", err)
		return nil, err
	}

	// Safe type assertion
	profiles, ok := result.([]domain.Profile)
	if !ok {

		return nil, fmt.Errorf("unexpected result type")

	}

	return profiles, nil
}
