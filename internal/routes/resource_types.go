package routes

import "time"

// Структуры запросов
type GetWorkplacesRequest struct {
	Zone        string `json:"zone,omitempty"`
	Floor       int64  `json:"floor,omitempty"`
	Type        string `json:"type,omitempty"`
	Capacity    int64  `json:"capacity,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
	WithItems   bool   `json:"with_items,omitempty"`
	Page        int64  `json:"page,omitempty"`
}

type CreateWorkplaceRequest struct {
	Address           string `json:"address"`
	Zone              string `json:"zone"`
	Floor             int64  `json:"floor"`
	Number            int64  `json:"number"`
	Type              string `json:"type"`
	Capacity          int64  `json:"capacity"`
	Description       string `json:"description"`
	IsAvailable       bool   `json:"is_available"`
	MaintenanceStatus string `json:"maintenance_status"`
}

type UpdateWorkplaceRequest struct {
	Address           string `json:"address,omitempty"`
	Zone              string `json:"zone,omitempty"`
	Floor             int64  `json:"floor,omitempty"`
	Number            int64  `json:"number,omitempty"`
	Type              string `json:"type,omitempty"`
	Capacity          int64  `json:"capacity,omitempty"`
	Description       string `json:"description,omitempty"`
	IsAvailable       bool   `json:"is_available,omitempty"`
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
}

// Структуры ответов
type Item struct {
	ID          int64     `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Condition   string    `json:"condition"`
	WorkplaceID int64     `json:"workplace_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Workplace struct {
	ID                int64     `json:"id"`
	Address           string    `json:"address"`
	Zone              string    `json:"zone"`
	Floor             int64     `json:"floor"`
	Number            int64     `json:"number"`
	Type              string    `json:"type"`
	Capacity          int64     `json:"capacity"`
	Description       string    `json:"description"`
	IsAvailable       bool      `json:"is_available"`
	MaintenanceStatus string    `json:"maintenance_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Items             []Item    `json:"items,omitempty"`
	UniqueTag         string    `json:"unique_tag"`
}

type GetWorkplacesResponse struct {
	Workplaces []Workplace `json:"workplaces"`
	TotalCount int64       `json:"total_count"`
	Page       int64       `json:"page"`
	PageSize   int64       `json:"page_size"`
}

type DeleteWorkplaceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Структуры запросов
type GetParkingSpacesRequest struct {
	Address     string `json:"address,omitempty"`
	Zone        string `json:"zone,omitempty"`
	Type        string `json:"type,omitempty"`
	Number      int64  `json:"number,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
	Page        int64  `json:"page,omitempty"`
}

type CreateParkingSpaceRequest struct {
	Number      int64  `json:"number"`
	Address     string `json:"address"`
	Type        string `json:"type"`
	Zone        string `json:"zone"`
	IsAvailable bool   `json:"is_available"`
}

type UpdateParkingSpaceRequest struct {
	Number      int64  `json:"number,omitempty"`
	Address     string `json:"address,omitempty"`
	Zone        string `json:"zone,omitempty"`
	Type        string `json:"type,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
}

// Структуры ответов
type ParkingSpace struct {
	ID          int64     `json:"id"`
	Number      int64     `json:"number"`
	Address     string    `json:"address"`
	Zone        string    `json:"zone"`
	Type        string    `json:"type"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetParkingSpacesResponse struct {
	ParkingSpaces []ParkingSpace `json:"parking_spaces"`
	TotalCount    int64          `json:"total_count"`
	Page          int64          `json:"page"`
	PageSize      int64          `json:"page_size"`
}

type DeleteParkingSpaceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Структуры запросов
type GetItemsRequest struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	ConditionID int64  `json:"condition_id,omitempty"`
	WorkplaceID int64  `json:"workplace_id,omitempty"`
	Page        int64  `json:"page,omitempty"`
}

type CreateItemRequest struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	ConditionID int64  `json:"condition_id"`
	WorkplaceID int64  `json:"workplace_id,omitempty"`
}

type UpdateItemRequest struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	ConditionID int64  `json:"condition_id,omitempty"`
	WorkplaceID int64  `json:"workplace_id,omitempty"`
}

type AttachItemToWorkplaceRequest struct {
	ItemID      int64 `json:"item_id"`
	WorkplaceID int64 `json:"workplace_id"`
}

// Структуры ответов
type ItemCondition struct {
	ID          int64  `json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type GetItemsResponse struct {
	Items      []Item `json:"items"`
	TotalCount int64  `json:"total_count"`
	Page       int64  `json:"page"`
	PageSize   int64  `json:"page_size"`
}

type DeleteItemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
