package repository

import "github.com/wahyunurdian26/cst_app_new/internal/entity"

type CampaignRepository interface {
	GetAllOffer() ([]entity.Offer, error)
	GetAllSender() ([]entity.Sender, error)
	GetAllProduct() ([]entity.Product, error)
	GetAllBrand() ([]entity.Brand, error)

	FindByIdCampaign(id_campaign string) (*entity.Campaign, error)
	CreateCampaign(campaign *entity.Campaign) error
	GetAllCampaign() ([]entity.Campaign, error)
}
