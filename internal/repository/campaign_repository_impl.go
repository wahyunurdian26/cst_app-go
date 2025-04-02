package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{DB: db}
}

func (c *CampaignRepositoryImpl) GetAllOffer() ([]entity.Offer, error) {
	var offers []entity.Offer
	err := c.DB.Find(&offers).Error
	return offers, err
}

func (c *CampaignRepositoryImpl) GetAllSender() ([]entity.Sender, error) {
	var senders []entity.Sender
	err := c.DB.Where("sender_type = ?", 1).Find(&senders).Error
	return senders, err
}

func (c *CampaignRepositoryImpl) GetAllProduct() ([]entity.Product, error) {
	var products []entity.Product
	err := c.DB.Find(&products).Error
	return products, err
}

func (c *CampaignRepositoryImpl) GetAllBrand() ([]entity.Brand, error) {
	var brands []entity.Brand
	err := c.DB.Find(&brands).Error
	return brands, err
}

// Trx Campaign
func (c *CampaignRepositoryImpl) CreateCampaign(campaign *entity.Campaign) error {
	return c.DB.Create(campaign).Error
}

func (r *CampaignRepositoryImpl) FindByIdCampaign(id_campaign string) (*entity.Campaign, error) {
	var campaign entity.Campaign
	if err := r.DB.Where("id_campaign = ? ", id_campaign).First(&campaign).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &campaign, nil
}

func (r *CampaignRepositoryImpl) GetAllCampaign() ([]entity.Campaign, error) {
	var campaign []entity.Campaign
	err := r.DB.Find(&campaign).Error
	return campaign, err
}

func (r *CampaignRepositoryImpl) GetById(id_campaign string) (*entity.Campaign, error) {
	var campaign entity.Campaign
	err := r.DB.Where("id_campaign = ?", id_campaign).Take(&campaign).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &campaign, nil
}
