package impl

import (
	"context"
	"errors"
	"log"
	"profile-service/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ProfileRepository struct {
	db neo4j.DriverWithContext
}

func NewProfileRepository(db neo4j.DriverWithContext) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (p *ProfileRepository) GetProfileById(c *gin.Context, id string) (*domain.Profile, error) {
	session := p.db.NewSession(c, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})

	if session != nil {
		log.Println("Session created successfully")
	} else {
		log.Println("Failed to create session")
		return nil, errors.New("session in get user profile repository is nil")
	}

	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Println("Error closing session:", err)
		}
	}(session, c)

	result, err := session.ExecuteRead(c, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MATCH (p:Profile {userId: $userId})
					RETURN p.userId as userId,
						p.first_name as firstName,
						p.last_name as lastName,
						p.display_name as displayName,
						p.bio as bio,
						p.avatar_url as avatarURL,
						p.birth_date as birthDate `
		params := map[string]any{
			"userId": id,
		}

		record, err := tx.Run(c, query, params)
		if err != nil {
			return nil, err
		}

		if record.Next(c) {
			r := record.Record()
			userId, _ := r.Get("userId")
			firstName, _ := r.Get("firstName")
			lastName, _ := r.Get("lastName")
			displayName, _ := r.Get("displayName")
			bio, _ := r.Get("bio")
			avatarURL, _ := r.Get("avatarURL")

			// the time.Time type is available here, we need to cast to time.Time type instead of string
			birthDate, _ := r.Get("birthDate")

			if err != nil {
				return nil, err
			}

			profile := &domain.Profile{
				UserID:      userId.(string),
				FirstName:   firstName.(string),
				LastName:    lastName.(string),
				DisplayName: displayName.(string),
				Bio:         bio.(string),
				AvatarURL:   avatarURL.(string),
				BirthDate:   birthDate.(time.Time),
			}

			return profile, nil
		}
		return nil, errors.New("profile not found")
	})

	if err != nil {
		return nil, err
	}

	return result.(*domain.Profile), nil
}

func (p *ProfileRepository) UpdateProfile(c *gin.Context, profile *domain.Profile, userId string) (*domain.Profile, error) {

	session := p.db.NewSession(c, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	if session != nil {
		log.Println("Session created successfully")
	} else {
		log.Println("Failed to create session")
		return nil, errors.New("session in update user profile repository is nil")
	}

	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Println("Error closing session:", err)
		}
	}(session, c)

	_, err := session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MERGE (p:Profile {userId: $userId})
		Set p.userId = $userId, 
			p.first_name = $firstName,
			p.last_name = $lastName,
			p.display_name = $displayName,
			p.bio = $bio,
			p.avatar_url = $avatarURL,
			p.birth_date = $birthDate
		`
		params := map[string]any{
			"userId":      userId,
			"firstName":   profile.FirstName,
			"lastName":    profile.LastName,
			"displayName": profile.DisplayName,
			"bio":         profile.Bio,
			"avatarURL":   profile.AvatarURL,
			"birthDate":   profile.BirthDate,
		}

		_, err := tx.Run(c, query, params)
		return nil, err
	})

	if err != nil {
		return nil, err
	}
	return nil, nil
}
