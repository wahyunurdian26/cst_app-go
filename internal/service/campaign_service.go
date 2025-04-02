package service

import (
	"context"

	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

type CampaignService interface {
	GetAllOffer() ([]entity.Offer, error)
	GetAllSender() ([]entity.Sender, error)
	GetAllProduct() ([]entity.Product, error)
	GetAllBrand() ([]entity.Brand, error)

	CreateCampaign(ctx context.Context, request *model.CampaignCreateRequest) (*entity.Campaign, error)
	GetAll() ([]entity.Campaign, error)
	GetById(id_campaign string) (*entity.Campaign, error)
}
