package vote

import (
	"Gopher-By-Example/graphql-example/model"
	"fmt"

	"github.com/graphql-go/graphql"
)

var Vote = graphql.NewObject(graphql.ObjectConfig{
	Name: "vote",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.CreatedAt, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.UpdatedAt, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Title, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Description, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"class": &graphql.Field{
			Type: Class,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Class, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"classString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.ClassString, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"options": &graphql.Field{
			Type: graphql.NewList(Options),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Options, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"deadline": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if vote, ok := p.Source.(model.VoteSerializer); ok {
					return vote.Deadline, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})

var OptionInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "optionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var Class = graphql.NewEnum(graphql.EnumConfig{
	Name: "classVote",
	Values: graphql.EnumValueConfigMap{
		"SINGLE": &graphql.EnumValueConfig{
			Value:       model.SINGLE,
			Description: model.ClassMap[model.SINGLE],
		},
		"MULTIPLE": &graphql.EnumValueConfig{
			Value:       model.MULTIPLE,
			Description: model.ClassMap[model.MULTIPLE],
		},
	},
})

var Options = graphql.NewObject(graphql.ObjectConfig{
	Name: "options",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if option, ok := p.Source.(model.OptionSerializer); ok {
					return option.Id, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if option, ok := p.Source.(model.OptionSerializer); ok {
					return option.CreatedAt, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if option, ok := p.Source.(model.OptionSerializer); ok {
					return option.UpdatedAt, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				if option, ok := p.Source.(model.OptionSerializer); ok {
					return option.Name, nil
				}
				return nil, fmt.Errorf("field not found")
			},
		},
	},
})
