package handler

import (
	"time"

	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/domain"
	"github.com/lolibrary/lolibrary/service.category/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateCategory(req typhon.Request) typhon.Response {
	body := &categoryproto.POSTCreateCategoryRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Slug == "":
		return typhon.Response{Error: validation.ErrMissingParam("slug")}
	case body.Name == "":
		return typhon.Response{Error: validation.ErrMissingParam("name")}
	case body.Slug != strcase.ToKebab(body.Slug):
		return typhon.Response{Error: validation.ErrBadParam("slug", "slug should be in kebab-case")}
	}

	id, err := idgen.New()
	if err != nil {
		slog.Error(req, "Error generating ID: %v", err)
		return typhon.Response{Error: err}
	}

	category := &domain.Category{
		ID:          id,
		Slug:        body.Slug,
		Name:        body.Name,
		CreatedAt:   time.Now().UTC(),
	}

	if err := dao.CreateCategory(category); err != nil {
		slog.Error(req, "Failed to create category entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&categoryproto.POSTCreateCategoryResponse{
		Category: marshaling.CategoryToProto(category),
	})
}