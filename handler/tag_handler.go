package handler

import (
	"strconv"

	"github.com/sinisaos/fiber-ent-admin/model"
	"github.com/sinisaos/fiber-ent-admin/service"

	"github.com/gofiber/fiber/v2"
)

type TagHandler struct {
	TagService service.TagService
}

func NewTagHandler(service service.TagService) *TagHandler {
	return &TagHandler{
		TagService: service,
	}
}

type TParams struct {
	Name  string `query:"title"`
	Sort  string `query:"_sort"`
	Order string `query:"_order"`
	Start int    `query:"_start"`
	End   int    `query:"_end"`
}

// All Tags
func (h TagHandler) GetAllTagsHandler(c *fiber.Ctx) error {
	p := new(TParams)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	// filtered result
	tags, err := h.TagService.GetAllTags(p.Start, p.End, p.Sort, p.Order, p.Name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// total number of filtered records for React Admin pagination
	if p.Name != "" {
		c.Set("X-Total-Count", strconv.Itoa(len(tags)))
	} else {
		// total number of records for React Admin pagination
		count, _ := h.TagService.CountTags()
		c.Set("X-Total-Count", strconv.Itoa(count))
	}
	return c.Status(200).JSON(tags)
}

// Single Tag
func (h TagHandler) GetTagHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	tag, err := h.TagService.GetTag(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(tag)
}

// New Tag
func (h TagHandler) CreateTagHandler(c *fiber.Ctx) error {
	payload := new(model.NewTagInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	tag, err := h.TagService.CreateTag(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": tag,
	})
}

// Update Tag
func (h TagHandler) UpdateTagHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	payload := new(model.UpdateTagInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	tag, err := h.TagService.UpdateTag(id, payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Tag successfully updated",
		"data":    tag,
	})
}

// Delete Tag
func (h TagHandler) DeleteTagHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	// Check if the record exists
	err := h.TagService.DeleteTag(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"message": "Tag successfully deleted",
	})
}

// Tag Questions
func (h TagHandler) GetTagQuestionHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	tag, err := h.TagService.GetTagQuestions(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(tag)
}
