package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VidaCriativaCategory representa a categoria do item criativo.
type VidaCriativaCategory string

const (
	VidaCriativaCategoryDesign      VidaCriativaCategory = "Design"
	VidaCriativaCategoryMusica      VidaCriativaCategory = "Música"
	VidaCriativaCategoryEscrita     VidaCriativaCategory = "Escrita"
	VidaCriativaCategoryFotografia  VidaCriativaCategory = "Fotografia"
	VidaCriativaCategoryArte        VidaCriativaCategory = "Arte"
	VidaCriativaCategoryVideo       VidaCriativaCategory = "Vídeo"
	VidaCriativaCategoryOutro       VidaCriativaCategory = "Outro"
)

var ValidVidaCriativaCategories = map[VidaCriativaCategory]bool{
	VidaCriativaCategoryDesign:     true,
	VidaCriativaCategoryMusica:     true,
	VidaCriativaCategoryEscrita:    true,
	VidaCriativaCategoryFotografia: true,
	VidaCriativaCategoryArte:       true,
	VidaCriativaCategoryVideo:      true,
	VidaCriativaCategoryOutro:      true,
}

// VidaCriativaStatus representa o status de um item criativo.
type VidaCriativaStatus string

const (
	VidaCriativaStatusIdeia      VidaCriativaStatus = "Ideia"
	VidaCriativaStatusEmAndamento VidaCriativaStatus = "Em andamento"
	VidaCriativaStatusConcluido  VidaCriativaStatus = "Concluído"
	VidaCriativaStatusPausado    VidaCriativaStatus = "Pausado"
)

var ValidVidaCriativaStatuses = map[VidaCriativaStatus]bool{
	VidaCriativaStatusIdeia:       true,
	VidaCriativaStatusEmAndamento: true,
	VidaCriativaStatusConcluido:   true,
	VidaCriativaStatusPausado:     true,
}

// VidaCriativaItem representa um projeto ou obra da seção Vida Criativa.
type VidaCriativaItem struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"        json:"id"`
	Title       string               `bson:"title"                json:"title"`
	Description string               `bson:"description"          json:"description"`
	Category    VidaCriativaCategory `bson:"category"             json:"category"`
	Status      VidaCriativaStatus   `bson:"status"               json:"status"`
	Tags        []string             `bson:"tags"                 json:"tags"`
	ImageURL    string               `bson:"image_url,omitempty"  json:"imageUrl,omitempty"`
	URL         string               `bson:"url,omitempty"        json:"url,omitempty"`
	BannerColor string               `bson:"banner_color"         json:"bannerColor"`
	Active      bool                 `bson:"active"               json:"active"`
	Order       int                  `bson:"order"                json:"order"`
	CreatedAt   time.Time            `bson:"created_at"           json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updated_at"           json:"updatedAt"`
}
