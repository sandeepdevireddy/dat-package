package models

// Company represents a company document in MongoDB
type PackageComponent struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}
