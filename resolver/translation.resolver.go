package resolver

import (
	"aery-graphql/generated/gqlgen"
	"aery-graphql/guard"
	"aery-graphql/model"
	"aery-graphql/utility"
	"context"
)

/**
 * Generated Structs
 */

/**
 * Generated Funcs
 */

/**
 * Field Methods
 */

/**
 * Queries
 */
func (r *queryResolver) Translation(ctx context.Context, where *model.TranslationWhereInput, searchInKeyAndValue *string) (*model.Translation, error) {
	if where == nil {
		where = &model.TranslationWhereInput{}
	}
	if searchInKeyAndValue != nil {
		where.OR = utility.SearchParser(model.TranslationSearchFields, *searchInKeyAndValue)
	}

	item := model.Translation{}
	item.One(where)
	if item.Key == "" {
		return nil, nil
	}

	return &item, nil
}

func (r *queryResolver) Translations(
	ctx context.Context,
	where *model.TranslationWhereInput,
	searchInKeyAndValue *string,
	orderBy *model.TranslationOrderByInput,
	skip *int,
	limit *int,
) ([]*model.Translation, error) {
	if where == nil {
		where = &model.TranslationWhereInput{}
	}
	if searchInKeyAndValue != nil {
		where.OR = utility.SearchParser(model.PageSearchFields, *searchInKeyAndValue)
	}

	item := model.Translation{}
	items, err := item.List(where, orderBy, skip, limit)

	if err != nil {
		return nil, err
	}
	return items, nil
}

/**
 * Mutations
 */
func (r *mutationResolver) CreateTranslation(ctx context.Context, data gqlgen.TranslationCreateInput) (*model.Translation, error) {
	if err := guard.Auth([]guard.Role{guard.Editor, guard.Admin}, *r.Resolver.Token); err != nil {
		return nil, err
	}

	item := model.Translation{
		Locale: data.Locale,
		Key:    data.Key,
		Value:  data.Value,
	}

	if err := item.Create(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mutationResolver) UpdateTranslation(ctx context.Context, where model.TranslationWhereUniqueInput, data gqlgen.TranslationUpdateInput) (*model.Translation, error) {
	if err := guard.Auth([]guard.Role{guard.Editor, guard.Admin}, *r.Resolver.Token); err != nil {
		return nil, err
	}

	item := model.Translation{}
	item.ID = where.ID
	item.Locale = data.Locale
	item.Key = data.Key
	item.Value = data.Value

	if err := item.Update(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *mutationResolver) DeleteTranslation(ctx context.Context, where model.TranslationWhereUniqueInput) (*model.Translation, error) {
	if err := guard.Auth([]guard.Role{guard.Editor, guard.Admin}, *r.Resolver.Token); err != nil {
		return nil, err
	}

	item := model.Translation{ID: where.ID}

	if err := item.Delete(); err != nil {
		return nil, err
	}

	return &item, nil
}
