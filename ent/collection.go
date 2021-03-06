// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AccountQuery) CollectFields(ctx context.Context, satisfies ...string) *AccountQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		a = a.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return a
}

func (a *AccountQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AccountQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "organization":
			a = a.WithOrganization(func(query *OrganizationQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return a
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CampaignQuery) CollectFields(ctx context.Context, satisfies ...string) *CampaignQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CampaignQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CampaignQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (o *OrganizationQuery) CollectFields(ctx context.Context, satisfies ...string) *OrganizationQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		o = o.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return o
}

func (o *OrganizationQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *OrganizationQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "accounts":
			o = o.WithAccounts(func(query *AccountQuery) {
				query.collectField(ctx, field)
			})
		case "campaigns":
			o = o.WithCampaigns(func(query *CampaignQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return o
}
