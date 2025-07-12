package registration

// Personal information of the shopkeeper
type PersonalInfo struct {
	FullName        string `json:"full_name" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// Shop-related details
type ShopInfo struct {
	ShopName      string `json:"shop_name" validate:"required"`
	ShopCategory  string `json:"shop_category" validate:"required"`
	ShopAddress   string `json:"shop_address" validate:"required"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
	Pincode       string `json:"pincode" validate:"required"`
	OpeningTime   string `json:"opening_time"` // e.g. "09:00"
	ClosingTime   string `json:"closing_time"` // e.g. "21:00"
	LicenseNumber string `json:"license_number,omitempty"`
}

// Bank or payment information
type PaymentInfo struct {
	AccountHolder string `json:"account_holder,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	IFSCCode      string `json:"ifsc_code,omitempty"`
	UPIID         string `json:"upi_id,omitempty"`
}

// Document uploads (file URLs or filenames)
type DocumentInfo struct {
	ShopPhotoURL        string `json:"shop_photo_url,omitempty"`
	OwnerIDProofURL     string `json:"owner_id_proof_url,omitempty"`
	RegistrationCertURL string `json:"registration_cert_url,omitempty"`
}

// GPS location of the shop
type LocationInfo struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// Main wrapper request for shopkeeper registration
type ShopkeeperRegisterRequest struct {
	PersonalInfo PersonalInfo `json:"personal_info"`
	ShopInfo     ShopInfo     `json:"shop_info"`
	PaymentInfo  PaymentInfo  `json:"payment_info,omitempty"`
	DocumentInfo DocumentInfo `json:"document_info,omitempty"`
	LocationInfo LocationInfo `json:"location_info,omitempty"`
}
