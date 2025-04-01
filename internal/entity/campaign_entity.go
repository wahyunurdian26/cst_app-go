package entity

import "time"

// Campaign represents the structure of the `trx_main_campaign` table.
type Campaign struct {
	IDCampaign                string    `gorm:"column:id_campaign;type:varchar(30);primary_key" json:"id_campaign"`
	CampaignCode              string    `gorm:"column:campaign_code;type:varchar(11);not null" json:"campaign_code"`
	CampaignName              string    `gorm:"column:campaign_name;type:varchar(100)" json:"campaign_name"`
	IDCampaignCategory        string    `gorm:"column:id_campaign_category;type:varchar(2)" json:"id_campaign_category"`
	IDBusinessGroup           string    `gorm:"column:id_business_group;type:varchar(2)" json:"id_business_group"`
	IDProductGroup            string    `gorm:"column:id_product_group;type:varchar(2)" json:"id_product_group"`
	IDGeneralObjective        string    `gorm:"column:id_general_objective;type:varchar(2)" json:"id_general_objective"`
	IDBrand                   string    `gorm:"column:id_brand;type:varchar(2)" json:"id_brand"`
	IDOfferingType            string    `gorm:"column:id_offering_type;type:varchar(2)" json:"id_offering_type"`
	StartDate                 time.Time `gorm:"column:start_date;type:date" json:"start_date"`
	EndDate                   time.Time `gorm:"column:end_date;type:date" json:"end_date"`
	BroadcastTime             time.Time `gorm:"column:broadcast_time;type:time" json:"broadcast_time"`
	IDSender                  string    `gorm:"column:id_sender;type:varchar(2)" json:"id_sender"`
	Wording                   string    `gorm:"column:wording;type:text" json:"wording"`
	Remarks                   string    `gorm:"column:remarks;type:text" json:"remarks"`
	BonusDesc                 string    `gorm:"column:bonus_desc;type:varchar(100)" json:"bonus_desc"`
	SubmissionTime            time.Time `gorm:"column:submission_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"submission_time"`
	CampaignApproveRemark     string    `gorm:"column:campaign_approve_remark;type:text" json:"campaign_approve_remark"`
	CampaignRejectReason      string    `gorm:"column:campaign_reject_reason;type:text" json:"campaign_reject_reason"`
	CampaignConfigRemark      string    `gorm:"column:campaign_config_remark;type:text" json:"campaign_config_remark"`
	CampaignDiscontinueRemark string    `gorm:"column:campaign_discontinue_remark;type:text" json:"campaign_discontinue_remark"`
	CampaignStatus            string    `gorm:"column:campaign_status;type:varchar(50)" json:"campaign_status"`
	EmailUser                 string    `gorm:"column:email_user;type:varchar(50)" json:"email_user"`
	WlFilename                string    `gorm:"column:wl_filename;type:varchar(100)" json:"wl_filename"`
	CampaignDescription       string    `gorm:"column:campaign_description;type:varchar(256)" json:"campaign_description"`
	WlType                    string    `gorm:"column:wl_type;type:varchar(8)" json:"wl_type"`
}

type Sender struct {
	ID         string `gorm:"column:id_sender;type:varchar(2);primaryKey" json:"id"`
	Name       string `gorm:"column:sender_name;type:varchar(50)" json:"name"`
	SenderType int    `gorm:"column:sender_type;type:int;default:0" json:"sender_type"`
}
type Product struct {
	ID        string `gorm:"column:id_product_group;type:varchar(2);primaryKey" json:"id"`
	Name      string `gorm:"column:product_group_name;type:varchar(50)" json:"name"`
	FlagUsage string `gorm:"column:flag_usage;type:varchar(1)" json:"flag_usage"`
	SearchATL string `gorm:"column:search_atl;type:text;not null" json:"search_atl"`
	SearchBTL string `gorm:"column:search_btl;type:text;not null" json:"search_btl"`
	EmailPIC  string `gorm:"column:email_pic;type:varchar(128)" json:"email_pic"`
}
type Brand struct {
	ID        string `gorm:"column:id_brand;type:varchar(2);primaryKey" json:"id"`
	Name      string `gorm:"column:brand_name;type:varchar(50)" json:"name"`
	SearchATL string `gorm:"column:search_atl;type:text;not null" json:"search_atl"`
	SearchBTL string `gorm:"column:search_btl;type:text" json:"search_btl"`
}

type Offer struct {
	ID   string `gorm:"column:id_offering_type;type:varchar(2);primaryKey" json:"id"`
	Name string `gorm:"column:offering_type_name;type:varchar(100)" json:"name"`
}
