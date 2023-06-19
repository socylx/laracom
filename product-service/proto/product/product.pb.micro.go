// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product/product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ProductService service

func NewProductServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductService service

type ProductService interface {
	Create(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	GetDetail(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) Create(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) Delete(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) Update(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) Get(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetDetail(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetDetail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	Create(context.Context, *Product, *Response) error
	Delete(context.Context, *Product, *Response) error
	Update(context.Context, *Product, *Response) error
	Get(context.Context, *Product, *Response) error
	GetDetail(context.Context, *Product, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		Create(ctx context.Context, in *Product, out *Response) error
		Delete(ctx context.Context, in *Product, out *Response) error
		Update(ctx context.Context, in *Product, out *Response) error
		Get(ctx context.Context, in *Product, out *Response) error
		GetDetail(ctx context.Context, in *Product, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) Create(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.Create(ctx, in, out)
}

func (h *productServiceHandler) Delete(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.Delete(ctx, in, out)
}

func (h *productServiceHandler) Update(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.Update(ctx, in, out)
}

func (h *productServiceHandler) Get(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.Get(ctx, in, out)
}

func (h *productServiceHandler) GetDetail(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.GetDetail(ctx, in, out)
}

func (h *productServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.ProductServiceHandler.GetAll(ctx, in, out)
}

// Api Endpoints for ImageService service

func NewImageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ImageService service

type ImageService interface {
	Create(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error)
	Delete(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error)
	Update(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error)
	Get(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error)
	GetByProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ImageResponse, error)
}

type imageService struct {
	c    client.Client
	name string
}

func NewImageService(name string, c client.Client) ImageService {
	return &imageService{
		c:    c,
		name: name,
	}
}

func (c *imageService) Create(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error) {
	req := c.c.NewRequest(c.name, "ImageService.Create", in)
	out := new(ImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Delete(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error) {
	req := c.c.NewRequest(c.name, "ImageService.Delete", in)
	out := new(ImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Update(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error) {
	req := c.c.NewRequest(c.name, "ImageService.Update", in)
	out := new(ImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Get(ctx context.Context, in *ProductImage, opts ...client.CallOption) (*ImageResponse, error) {
	req := c.c.NewRequest(c.name, "ImageService.Get", in)
	out := new(ImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) GetByProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*ImageResponse, error) {
	req := c.c.NewRequest(c.name, "ImageService.GetByProduct", in)
	out := new(ImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageService service

type ImageServiceHandler interface {
	Create(context.Context, *ProductImage, *ImageResponse) error
	Delete(context.Context, *ProductImage, *ImageResponse) error
	Update(context.Context, *ProductImage, *ImageResponse) error
	Get(context.Context, *ProductImage, *ImageResponse) error
	GetByProduct(context.Context, *Product, *ImageResponse) error
}

func RegisterImageServiceHandler(s server.Server, hdlr ImageServiceHandler, opts ...server.HandlerOption) error {
	type imageService interface {
		Create(ctx context.Context, in *ProductImage, out *ImageResponse) error
		Delete(ctx context.Context, in *ProductImage, out *ImageResponse) error
		Update(ctx context.Context, in *ProductImage, out *ImageResponse) error
		Get(ctx context.Context, in *ProductImage, out *ImageResponse) error
		GetByProduct(ctx context.Context, in *Product, out *ImageResponse) error
	}
	type ImageService struct {
		imageService
	}
	h := &imageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ImageService{h}, opts...))
}

type imageServiceHandler struct {
	ImageServiceHandler
}

func (h *imageServiceHandler) Create(ctx context.Context, in *ProductImage, out *ImageResponse) error {
	return h.ImageServiceHandler.Create(ctx, in, out)
}

func (h *imageServiceHandler) Delete(ctx context.Context, in *ProductImage, out *ImageResponse) error {
	return h.ImageServiceHandler.Delete(ctx, in, out)
}

func (h *imageServiceHandler) Update(ctx context.Context, in *ProductImage, out *ImageResponse) error {
	return h.ImageServiceHandler.Update(ctx, in, out)
}

func (h *imageServiceHandler) Get(ctx context.Context, in *ProductImage, out *ImageResponse) error {
	return h.ImageServiceHandler.Get(ctx, in, out)
}

func (h *imageServiceHandler) GetByProduct(ctx context.Context, in *Product, out *ImageResponse) error {
	return h.ImageServiceHandler.GetByProduct(ctx, in, out)
}

// Api Endpoints for BrandService service

func NewBrandServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BrandService service

type BrandService interface {
	Create(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error)
	Delete(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error)
	Update(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error)
	Get(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*BrandResponse, error)
	GetWithProducts(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error)
}

type brandService struct {
	c    client.Client
	name string
}

func NewBrandService(name string, c client.Client) BrandService {
	return &brandService{
		c:    c,
		name: name,
	}
}

func (c *brandService) Create(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.Create", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Delete(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.Delete", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Update(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.Update", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Get(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.Get", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.GetAll", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) GetWithProducts(ctx context.Context, in *Brand, opts ...client.CallOption) (*BrandResponse, error) {
	req := c.c.NewRequest(c.name, "BrandService.GetWithProducts", in)
	out := new(BrandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BrandService service

type BrandServiceHandler interface {
	Create(context.Context, *Brand, *BrandResponse) error
	Delete(context.Context, *Brand, *BrandResponse) error
	Update(context.Context, *Brand, *BrandResponse) error
	Get(context.Context, *Brand, *BrandResponse) error
	GetAll(context.Context, *Request, *BrandResponse) error
	GetWithProducts(context.Context, *Brand, *BrandResponse) error
}

func RegisterBrandServiceHandler(s server.Server, hdlr BrandServiceHandler, opts ...server.HandlerOption) error {
	type brandService interface {
		Create(ctx context.Context, in *Brand, out *BrandResponse) error
		Delete(ctx context.Context, in *Brand, out *BrandResponse) error
		Update(ctx context.Context, in *Brand, out *BrandResponse) error
		Get(ctx context.Context, in *Brand, out *BrandResponse) error
		GetAll(ctx context.Context, in *Request, out *BrandResponse) error
		GetWithProducts(ctx context.Context, in *Brand, out *BrandResponse) error
	}
	type BrandService struct {
		brandService
	}
	h := &brandServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BrandService{h}, opts...))
}

type brandServiceHandler struct {
	BrandServiceHandler
}

func (h *brandServiceHandler) Create(ctx context.Context, in *Brand, out *BrandResponse) error {
	return h.BrandServiceHandler.Create(ctx, in, out)
}

func (h *brandServiceHandler) Delete(ctx context.Context, in *Brand, out *BrandResponse) error {
	return h.BrandServiceHandler.Delete(ctx, in, out)
}

func (h *brandServiceHandler) Update(ctx context.Context, in *Brand, out *BrandResponse) error {
	return h.BrandServiceHandler.Update(ctx, in, out)
}

func (h *brandServiceHandler) Get(ctx context.Context, in *Brand, out *BrandResponse) error {
	return h.BrandServiceHandler.Get(ctx, in, out)
}

func (h *brandServiceHandler) GetAll(ctx context.Context, in *Request, out *BrandResponse) error {
	return h.BrandServiceHandler.GetAll(ctx, in, out)
}

func (h *brandServiceHandler) GetWithProducts(ctx context.Context, in *Brand, out *BrandResponse) error {
	return h.BrandServiceHandler.GetWithProducts(ctx, in, out)
}

// Api Endpoints for CategoryService service

func NewCategoryServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CategoryService service

type CategoryService interface {
	Create(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Delete(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Update(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Get(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*CategoryResponse, error)
	GetWithProducts(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) Create(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Create", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Delete(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Delete", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Update(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Update", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Get(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Get", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.GetAll", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) GetWithProducts(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.GetWithProducts", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CategoryService service

type CategoryServiceHandler interface {
	Create(context.Context, *Category, *CategoryResponse) error
	Delete(context.Context, *Category, *CategoryResponse) error
	Update(context.Context, *Category, *CategoryResponse) error
	Get(context.Context, *Category, *CategoryResponse) error
	GetAll(context.Context, *Request, *CategoryResponse) error
	GetWithProducts(context.Context, *Category, *CategoryResponse) error
}

func RegisterCategoryServiceHandler(s server.Server, hdlr CategoryServiceHandler, opts ...server.HandlerOption) error {
	type categoryService interface {
		Create(ctx context.Context, in *Category, out *CategoryResponse) error
		Delete(ctx context.Context, in *Category, out *CategoryResponse) error
		Update(ctx context.Context, in *Category, out *CategoryResponse) error
		Get(ctx context.Context, in *Category, out *CategoryResponse) error
		GetAll(ctx context.Context, in *Request, out *CategoryResponse) error
		GetWithProducts(ctx context.Context, in *Category, out *CategoryResponse) error
	}
	type CategoryService struct {
		categoryService
	}
	h := &categoryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CategoryService{h}, opts...))
}

type categoryServiceHandler struct {
	CategoryServiceHandler
}

func (h *categoryServiceHandler) Create(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Create(ctx, in, out)
}

func (h *categoryServiceHandler) Delete(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Delete(ctx, in, out)
}

func (h *categoryServiceHandler) Update(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Update(ctx, in, out)
}

func (h *categoryServiceHandler) Get(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Get(ctx, in, out)
}

func (h *categoryServiceHandler) GetAll(ctx context.Context, in *Request, out *CategoryResponse) error {
	return h.CategoryServiceHandler.GetAll(ctx, in, out)
}

func (h *categoryServiceHandler) GetWithProducts(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.GetWithProducts(ctx, in, out)
}

// Api Endpoints for AttributeService service

func NewAttributeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttributeService service

type AttributeService interface {
	CreateAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error)
	DeleteAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error)
	UpdateAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error)
	CreateValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error)
	DeleteValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error)
	UpdateValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error)
	CreateProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error)
	DeleteProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error)
	UpdateProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error)
	GetAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error)
	GetAttributes(ctx context.Context, in *Request, opts ...client.CallOption) (*AttributeResponse, error)
	GetValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error)
	GetValues(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeValueResponse, error)
	GetProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error)
	GetProductAttributes(ctx context.Context, in *Product, opts ...client.CallOption) (*ProductAttributeResponse, error)
}

type attributeService struct {
	c    client.Client
	name string
}

func NewAttributeService(name string, c client.Client) AttributeService {
	return &attributeService{
		c:    c,
		name: name,
	}
}

func (c *attributeService) CreateAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.CreateAttribute", in)
	out := new(AttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) DeleteAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.DeleteAttribute", in)
	out := new(AttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) UpdateAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.UpdateAttribute", in)
	out := new(AttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) CreateValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.CreateValue", in)
	out := new(AttributeValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) DeleteValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.DeleteValue", in)
	out := new(AttributeValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) UpdateValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.UpdateValue", in)
	out := new(AttributeValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) CreateProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.CreateProductAttribute", in)
	out := new(ProductAttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) DeleteProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.DeleteProductAttribute", in)
	out := new(ProductAttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) UpdateProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.UpdateProductAttribute", in)
	out := new(ProductAttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetAttribute(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetAttribute", in)
	out := new(AttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetAttributes(ctx context.Context, in *Request, opts ...client.CallOption) (*AttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetAttributes", in)
	out := new(AttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetValue(ctx context.Context, in *AttributeValue, opts ...client.CallOption) (*AttributeValueResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetValue", in)
	out := new(AttributeValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetValues(ctx context.Context, in *Attribute, opts ...client.CallOption) (*AttributeValueResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetValues", in)
	out := new(AttributeValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetProductAttribute(ctx context.Context, in *ProductAttribute, opts ...client.CallOption) (*ProductAttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetProductAttribute", in)
	out := new(ProductAttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeService) GetProductAttributes(ctx context.Context, in *Product, opts ...client.CallOption) (*ProductAttributeResponse, error) {
	req := c.c.NewRequest(c.name, "AttributeService.GetProductAttributes", in)
	out := new(ProductAttributeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttributeService service

type AttributeServiceHandler interface {
	CreateAttribute(context.Context, *Attribute, *AttributeResponse) error
	DeleteAttribute(context.Context, *Attribute, *AttributeResponse) error
	UpdateAttribute(context.Context, *Attribute, *AttributeResponse) error
	CreateValue(context.Context, *AttributeValue, *AttributeValueResponse) error
	DeleteValue(context.Context, *AttributeValue, *AttributeValueResponse) error
	UpdateValue(context.Context, *AttributeValue, *AttributeValueResponse) error
	CreateProductAttribute(context.Context, *ProductAttribute, *ProductAttributeResponse) error
	DeleteProductAttribute(context.Context, *ProductAttribute, *ProductAttributeResponse) error
	UpdateProductAttribute(context.Context, *ProductAttribute, *ProductAttributeResponse) error
	GetAttribute(context.Context, *Attribute, *AttributeResponse) error
	GetAttributes(context.Context, *Request, *AttributeResponse) error
	GetValue(context.Context, *AttributeValue, *AttributeValueResponse) error
	GetValues(context.Context, *Attribute, *AttributeValueResponse) error
	GetProductAttribute(context.Context, *ProductAttribute, *ProductAttributeResponse) error
	GetProductAttributes(context.Context, *Product, *ProductAttributeResponse) error
}

func RegisterAttributeServiceHandler(s server.Server, hdlr AttributeServiceHandler, opts ...server.HandlerOption) error {
	type attributeService interface {
		CreateAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error
		DeleteAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error
		UpdateAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error
		CreateValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error
		DeleteValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error
		UpdateValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error
		CreateProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error
		DeleteProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error
		UpdateProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error
		GetAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error
		GetAttributes(ctx context.Context, in *Request, out *AttributeResponse) error
		GetValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error
		GetValues(ctx context.Context, in *Attribute, out *AttributeValueResponse) error
		GetProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error
		GetProductAttributes(ctx context.Context, in *Product, out *ProductAttributeResponse) error
	}
	type AttributeService struct {
		attributeService
	}
	h := &attributeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttributeService{h}, opts...))
}

type attributeServiceHandler struct {
	AttributeServiceHandler
}

func (h *attributeServiceHandler) CreateAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error {
	return h.AttributeServiceHandler.CreateAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) DeleteAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error {
	return h.AttributeServiceHandler.DeleteAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) UpdateAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error {
	return h.AttributeServiceHandler.UpdateAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) CreateValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error {
	return h.AttributeServiceHandler.CreateValue(ctx, in, out)
}

func (h *attributeServiceHandler) DeleteValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error {
	return h.AttributeServiceHandler.DeleteValue(ctx, in, out)
}

func (h *attributeServiceHandler) UpdateValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error {
	return h.AttributeServiceHandler.UpdateValue(ctx, in, out)
}

func (h *attributeServiceHandler) CreateProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error {
	return h.AttributeServiceHandler.CreateProductAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) DeleteProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error {
	return h.AttributeServiceHandler.DeleteProductAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) UpdateProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error {
	return h.AttributeServiceHandler.UpdateProductAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) GetAttribute(ctx context.Context, in *Attribute, out *AttributeResponse) error {
	return h.AttributeServiceHandler.GetAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) GetAttributes(ctx context.Context, in *Request, out *AttributeResponse) error {
	return h.AttributeServiceHandler.GetAttributes(ctx, in, out)
}

func (h *attributeServiceHandler) GetValue(ctx context.Context, in *AttributeValue, out *AttributeValueResponse) error {
	return h.AttributeServiceHandler.GetValue(ctx, in, out)
}

func (h *attributeServiceHandler) GetValues(ctx context.Context, in *Attribute, out *AttributeValueResponse) error {
	return h.AttributeServiceHandler.GetValues(ctx, in, out)
}

func (h *attributeServiceHandler) GetProductAttribute(ctx context.Context, in *ProductAttribute, out *ProductAttributeResponse) error {
	return h.AttributeServiceHandler.GetProductAttribute(ctx, in, out)
}

func (h *attributeServiceHandler) GetProductAttributes(ctx context.Context, in *Product, out *ProductAttributeResponse) error {
	return h.AttributeServiceHandler.GetProductAttributes(ctx, in, out)
}
