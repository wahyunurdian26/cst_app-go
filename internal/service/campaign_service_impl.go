package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"gorm.io/gorm"
)

type campaignService struct {
	CampaignRepository repository.CampaignRepository
	Log                *logrus.Logger
	Validate           *validator.Validate
}

func NewCampaignService(campaignRepository repository.CampaignRepository, validator *validator.Validate, log *logrus.Logger) CampaignService {
	return &campaignService{
		CampaignRepository: campaignRepository,
		Validate:           validator,
		Log:                log,
	}
}

func (u *campaignService) GetAllOffer() ([]entity.Offer, error) {
	offers, err := u.CampaignRepository.GetAllOffer()
	if err != nil {
		u.Log.Warnf("Failed to retrieve offers: %v", err)
		return nil, err
	}
	return offers, nil
}

func (u *campaignService) GetAllSender() ([]entity.Sender, error) {
	senders, err := u.CampaignRepository.GetAllSender()
	if err != nil {
		u.Log.Warnf("Failed to retrieve senders: %v", err)
		return nil, err
	}
	return senders, nil
}

func (u *campaignService) GetAllProduct() ([]entity.Product, error) {
	products, err := u.CampaignRepository.GetAllProduct()
	if err != nil {
		u.Log.Warnf("Failed to retrieve products: %v", err)
		return nil, err
	}
	return products, nil
}

func (u *campaignService) GetAllBrand() ([]entity.Brand, error) {
	brands, err := u.CampaignRepository.GetAllBrand()
	if err != nil {
		u.Log.Warnf("Failed to retrieve brands: %v", err)
		return nil, err
	}
	return brands, nil
}

//Trx Campaign

func (u *campaignService) CreateCampaign(ctx context.Context, request *model.CampaignCreateRequest) (*entity.Campaign, error) {
	// Validate request
	if err := u.Validate.Struct(request); err != nil {
		u.Log.Warnf("Campaign creation failed: validation error: %v", err)
		return nil, fiber.ErrBadRequest
	}

	// Check if user already exists
	existingIdCmp, err := u.CampaignRepository.FindByIdCampaign(request.IDCampaign)
	if err != nil {
		u.Log.Warnf("Campaign creation failed: database error: %v", err)
		return nil, fmt.Errorf("failed to check existing campaign: %w", err)
	}

	if existingIdCmp != nil {
		u.Log.Warnf("Campaign creation failed: user already exists with id %s", request.IDCampaign)
		return nil, fiber.ErrConflict
	}

	// Create new campaign
	campaign := &entity.Campaign{
		IDCampaign:                request.IDCampaign,
		CampaignCode:              request.CampaignCode,
		CampaignName:              request.CampaignName,
		IDCampaignCategory:        request.IDCampaignCategory,
		IDBusinessGroup:           request.IDBusinessGroup,
		IDProductGroup:            request.IDProductGroup,
		IDGeneralObjective:        request.IDGeneralObjective,
		IDBrand:                   request.IDBrand,
		IDOfferingType:            request.IDOfferingType,
		StartDate:                 request.StartDate,
		EndDate:                   request.EndDate,
		BroadcastTime:             request.BroadcastTime,
		IDSender:                  request.IDSender,
		Wording:                   request.Wording,
		Remarks:                   request.Remarks,
		BonusDesc:                 request.BonusDesc,
		SubmissionTime:            request.SubmissionTime,
		CampaignApproveRemark:     request.CampaignApproveRemark,
		CampaignRejectReason:      request.CampaignRejectReason,
		CampaignConfigRemark:      request.CampaignConfigRemark,
		CampaignDiscontinueRemark: request.CampaignDiscontinueRemark,
		CampaignStatus:            request.CampaignStatus,
		EmailUser:                 request.EmailUser,
		WlFilename:                request.WlFilename,
		CampaignDescription:       request.CampaignDescription,
		WlType:                    request.WlType,
	}

	// Save user to database
	if err := u.CampaignRepository.CreateCampaign(campaign); err != nil {
		u.Log.Warnf("User creation failed: database error: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	u.Log.Infof("User created successfully: ID Campaign %s", campaign.IDCampaign)

	// Kembalikan entity.User sesuai dengan deklarasi interface
	return campaign, nil
}

func (u *campaignService) GetAll() ([]entity.Campaign, error) {
	campaign, err := u.CampaignRepository.GetAllCampaign()
	if err != nil {
		u.Log.Warnf("Failed to retrieve campaign: %v", err)
		return nil, err
	}
	return campaign, nil
}

func (u *campaignService) GetById(id_campaign string) (*entity.Campaign, error) {
	campaign, err := u.CampaignRepository.GetById(id_campaign)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		u.Log.Warnf("Failed to get campaign by ID %s: %v", id_campaign, err)
		return nil, err
	}
	return campaign, nil
}
