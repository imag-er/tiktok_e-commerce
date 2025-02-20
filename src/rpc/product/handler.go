package main

import (
	"context"
	"log"
	product "src/kitex_gen/product"

	"fmt"

	"gorm.io/gorm"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct {
	db *gorm.DB
}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	products := []*product.Product{}
	result := s.db.WithContext(ctx).Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list products: %w", result.Error)
	}

	log.Printf("list products: %v", products)
	
	return &product.ListProductsResp{
		Products: products,
	}, nil

}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	product_ := Product{}

	// search product by id
	result := s.db.WithContext(ctx).Where("id = ?", req.Id).First(&product_)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to find product: %w", result.Error)
	}


	log.Printf("get product: %v", product_)

	return &product.GetProductResp{
		Product: product_.ORM2RPC(),
	}, nil
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// search product by name or description
	products := []*Product{}
	result := s.db.WithContext(ctx).Where("name LIKE ? OR description LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%").Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to search products: %w", result.Error)
	}

	
	// convert orm to rpc
	productsresp := []*product.Product{}
	for _, p := range products {
		productsresp = append(productsresp, p.ORM2RPC())
	}
	return &product.SearchProductsResp{Results: productsresp}, nil
}

// CreateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// create product
	// 从idl定义的CreateProductReq结构体中取出字段
	// 创建一个orm的Product对象

	tmp := Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Picture:     req.Picture,
		Categories:  req.Categories,
	}

	result := s.db.WithContext(ctx).Create(&tmp)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create product: %w", result.Error)
	}

	log.Printf("create product: %v", tmp)
	return &product.CreateProductResp{Id: tmp.ID}, nil
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// delete product by id
	result := s.db.WithContext(ctx).Delete(&Product{}, req.Id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to delete product: %w", result.Error)
	}
	return &product.DeleteProductResp{}, nil

}
