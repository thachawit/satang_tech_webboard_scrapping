package entity

type DataBaseColumn struct {
	WebboardCategory string `json:"webboatdcategory" bson:"webboatd_category" `
	Topic            string `json:"topic" bson:"topic" `
	CoverImage       string `json:"cover_image" bson:"cover_image" `
	CreatedBy        string `json:"created_by" bson:"created_by" `
	CreatedDate      string `json:"created_date" bson:"created_date" `
	ReplyTotalNumber string `json:"reply_total_number" bson:"reply_total_number" `
	Content          string `json:"content" bson:"content" `
}
