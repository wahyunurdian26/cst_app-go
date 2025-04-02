package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/helper"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/service"
)

type CampaignController struct {
	CampaignService service.CampaignService
}

func NewCampaignController(campaignService service.CampaignService) *CampaignController {
	return &CampaignController{CampaignService: campaignService}
}

func (h *CampaignController) GetAllOffer(ctx *fiber.Ctx) error {
	offers, err := h.CampaignService.GetAllOffer()
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve offers")
	}

	return helper.JSONResponse(ctx, fiber.StatusOK, "Offers retrieved successfully", offers)
}

func (c *CampaignController) GetAllSender(ctx *fiber.Ctx) error {
	senders, err := c.CampaignService.GetAllSender()
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve senders")
	}

	return helper.JSONResponse(ctx, fiber.StatusOK, "Senders retrieved successfully", senders)
}

func (h *CampaignController) GetAllProduct(ctx *fiber.Ctx) error {
	offers, err := h.CampaignService.GetAllProduct()
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve offers")
	}

	return helper.JSONResponse(ctx, fiber.StatusOK, "Offers retrieved successfully", offers)
}

func (c *CampaignController) GetAllBrand(ctx *fiber.Ctx) error {
	senders, err := c.CampaignService.GetAllBrand()
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve senders")
	}

	return helper.JSONResponse(ctx, fiber.StatusOK, "Senders retrieved successfully", senders)
}

func (h *CampaignController) CreateCampaign(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body is required")
	}

	request := new(model.CampaignCreateRequest)
	if err := c.BodyParser(request); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := h.CampaignService.CreateCampaign(c.Context(), request)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.JSONResponse(c, fiber.StatusCreated, "Campaign successfully created", response)
}

func (h *CampaignController) GetAllCampaign(c *fiber.Ctx) error {
	campaigns, err := h.CampaignService.GetAll()
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve campaigns")
	}

	return helper.JSONResponse(c, fiber.StatusOK, "Campaigns retrieved successfully", campaigns)
}

func (h *CampaignController) GetCampaignByID(ctx *fiber.Ctx) error {
	id_campaign := ctx.Params("id_campaign")

	campaign, err := h.CampaignService.GetById(string(id_campaign))
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return helper.ErrorResponse(ctx, fiberErr.Code, fiberErr.Message)
		}
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, "Something went wrong")
	}

	if campaign == nil {
		return helper.ErrorResponse(ctx, fiber.StatusNotFound, "campaign not found")
	}

	return helper.JSONResponse(ctx, fiber.StatusOK, "campaign retrieved successfully", campaign)
}
