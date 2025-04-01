package model

import "time"

type CampaignCreateRequest struct {
	IDCampaign                string    `json:"id_campaign" validate:"required,max=30"`
	CampaignCode              string    `json:"campaign_code" validate:"required,max=11"`
	CampaignName              string    `json:"campaign_name" validate:"required,max=100"`
	IDCampaignCategory        string    `json:"id_campaign_category" validate:"required,max=2"`
	IDBusinessGroup           string    `json:"id_business_group" validate:"required,max=2"`
	IDProductGroup            string    `json:"id_product_group" validate:"required,max=2"`
	IDGeneralObjective        string    `json:"id_general_objective" validate:"required,max=2"`
	IDBrand                   string    `json:"id_brand" validate:"required,max=2"`
	IDOfferingType            string    `json:"id_offering_type" validate:"required,max=2"`
	StartDate                 time.Time `json:"start_date" validate:"required"`
	EndDate                   time.Time `json:"end_date" validate:"required,gtfield=StartDate"`
	BroadcastTime             time.Time `json:"broadcast_time" validate:"required"`
	IDSender                  string    `json:"id_sender" validate:"required,max=2"`
	Wording                   string    `json:"wording" validate:"required"`
	Remarks                   string    `json:"remarks"`
	BonusDesc                 string    `json:"bonus_desc" validate:"max=100"`
	SubmissionTime            time.Time `json:"submission_time" validate:"required"`
	CampaignApproveRemark     string    `json:"campaign_approve_remark"`
	CampaignRejectReason      string    `json:"campaign_reject_reason"`
	CampaignConfigRemark      string    `json:"campaign_config_remark"`
	CampaignDiscontinueRemark string    `json:"campaign_discontinue_remark"`
	CampaignStatus            string    `json:"campaign_status" validate:"required,max=50"`
	EmailUser                 string    `json:"email_user" validate:"required,email,max=50"`
	WlFilename                string    `json:"wl_filename" validate:"max=100"`
	CampaignDescription       string    `json:"campaign_description" validate:"max=256"`
	WlType                    string    `json:"wl_type" validate:"max=8"`
}
