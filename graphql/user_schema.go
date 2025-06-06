package graphql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	repository "github.com/lava15/timeslots-booking-golang/internal/repositories"
)

func resolveUserByID(ur *repository.UserRepository) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		idStr, _ := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, errors.New("invalid user ID")
		}
		return ur.GetUserByID(id)
	}
}

func NewSchema(ur *repository.UserRepository) (*graphql.Schema, error) {
	userType := graphql.NewObject(graphql.ObjectConfig{})
	queryType := graphql.NewObject(graphql.ObjectConfig{})
	mutationType := graphql.NewObject(graphql.ObjectConfig{})
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
