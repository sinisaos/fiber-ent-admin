package service

import (
	"context"

	"github.com/sinisaos/fiber-ent-admin/ent"
	"github.com/sinisaos/fiber-ent-admin/ent/predicate"
	"github.com/sinisaos/fiber-ent-admin/ent/tag"
	"github.com/sinisaos/fiber-ent-admin/pkg/model"
)

type TagService struct {
	Client *ent.Client
}

func NewTagService(client *ent.Client) *TagService {
	return &TagService{
		Client: client,
	}
}

func (s TagService) GetAllTags(start int, end int, sort string, order string, name string) ([]*ent.Tag, error) {
	var orderFunc []tag.OrderOption
	var predicateTag []predicate.Tag

	if name != "" {
		predicateTag = append(predicateTag, tag.NameContainsFold(name))
	}

	if order != "" {
		if order == "ASC" {
			orderFunc = append(orderFunc, ent.Asc(sort))
		} else {
			orderFunc = append(orderFunc, ent.Desc(sort))
		}
	}

	tags, _ := s.Client.Tag.Query().
		Where(predicateTag...).
		// Aggregate(
		// 	ent.Count(),
		// ).
		WithQuestions().
		Limit(end - start).
		Offset(start).
		Order(orderFunc...).
		All(context.Background())

	return tags, nil
}

func (s TagService) CountTags() (int, error) {
	count, _ := s.Client.Tag.Query().
		Aggregate(
			ent.Count(),
		).
		Int(context.Background())
	return count, nil
}

func (s TagService) GetTag(id int) (*ent.Tag, error) {
	tag, err := s.Client.Tag.Query().
		Where(tag.IDEQ(id)).
		WithQuestions().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagService) CreateTag(payload *model.NewTagInput) (*ent.Tag, error) {
	tag, err := s.Client.Tag.Create().
		SetName(payload.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagService) UpdateTag(id int, payload *model.UpdateTagInput) (*ent.Tag, error) {
	tag, err := s.Client.Tag.UpdateOneID(id).
		SetName(payload.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagService) DeleteTag(id int) error {
	err := s.Client.Tag.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s TagService) GetTagQuestions(id int) (*ent.Tag, error) {
	tag, _ := s.Client.Tag.Query().
		Where(tag.IDEQ(id)).
		WithQuestions().
		Only(context.Background())

	return tag, nil
}
