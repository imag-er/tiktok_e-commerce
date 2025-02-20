package main
import (
	"time"
	"gorm.io/gorm"
	"github.com/lib/pq"
	product "src/kitex_gen/product"

)



type Product struct {
	ID          uint32         `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Name        string   `gorm:"size:255;not null"`
	Description string   
	Picture     string   
	Price       float32  `gorm:"not null"`
	Categories  pq.StringArray `gorm:"type:text[]"` 

}

func (this *Product) ORM2RPC() *product.Product {
	return &product.Product{
		Id:          this.ID,
		Name:        this.Name,
		Description: this.Description,
		Picture:     this.Picture,
		Price:       this.Price,
		Categories:  this.Categories,
	}
}