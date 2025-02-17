# 青训营 X MarsCode 技术训练营
### 一、项目概述
1. 项目愿景
希望同学们可以通过完成这个项目切实实践课程中(视频中)学到的知识点包括但是不限于Go 语言编程，常用框架、数据库、对象存储，服务治理，服务上云等内容，同时对开发工作有更多的深入了解与认识，长远讲能对大家的个人技术成长或视野有启发。
2. 项目目标
一句话做一个抖音商城。为用户提供便捷、优质的购物环境，满足用户多样化的购物需求，打造一个具有影响力的社交电商平台，提升抖音在电商领域的市场竞争力。
3. 技术栈
- Go - Hertz   -Kitex  -Consul   - OpenTelemetry   - Gorm   -cwgo   -Redis
使用其他语言以及其他语言对应的技术生态也可以，这里不做任何限制

### 二、技术需求
（一）注册中心集成
1. 服务注册与发现
  - 该服务能够与注册中心etcd进行集成，自动注册服务数据。
（二）身份认证
1. 登录认证
  - 可以使用第三方现成的登录验证框架CasBin，对请求进行身份验证
  - 可配置的认证白名单，对于某些不需要认证的接口或路径，允许直接访问
  - 可配置的黑名单，对于某些异常的用户，直接进行封禁处理
（三）可观测要求
1. 日志记录与监控
  - 对服务的运行状态和请求处理过程进行详细的日志记录，方便故障排查和性能分析。
  - 提供实时监控功能，能够及时发现和解决系统中的问题。
### 三、功能需求
认证中心
- 分发身份令牌
- 续期身份令牌
- 校验身份令牌

用户服务
- 创建用户
- 登录
- 用户登出
- 删除用户
- 更新用户
- 获取用户身份信息

商品服务
- 创建商品
- 修改商品信息
- 删除商品
- 查询商品信息（单个商品、批量商品）
购物车服务
- 创建购物车
- 清空购物车
- 获取购物车信息
订单服务
- 创建订单
- 修改订单信息
- 订单定时取消
结算
- 订单结算
支付
- 取消支付
- 定时取消支付
- 支付
### 四、考核方式
- 根据第二点技术需求设计一个合理且具有一定扩展性的系统架构
- 根据第三点功能需求设计出完整的库表结构
### 六、编码要求
- 在本机搭建运行环境或在云上进行开发都可，这里不做任何限制
侧重服务端实现，会提前定义好各个功能对应的接口（接口定义使用Protobuf），按说明实现接口即可在客户端中看到运行效果
服务端最基本的结构只需要服务端程序和数据库即可，服务端程序连接数据库，响应客户端请求完成对应功能。同时需要根据功能，设计合理的数据模型，并创建对应的数据表，其中日志文件等可以保存到本地，这里不做限制
为了数据库层面的安全考虑建议，建议使用提供ACL控制的云数据库，使用本地数据库也可，这里不做限制
数据库安装配置说明：MySQL 8.0 version +
对其他数据库或者其他中间件有了解的同学也可以根据实际情况选择，这里不做限制
十、接口文档
1) 认证服务
```protobuf
syntax="proto3";

package auth;

option go_package="/auth";

service AuthService {
    rpc DeliverTokenByRPC(DeliverTokenReq) returns (DeliveryResp) {}
    rpc VerifyTokenByRPC(VerifyTokenReq) returns (VerifyResp) {}
}

message DeliverTokenReq {
    int32  user_id= 0;
}

message VerifyTokenReq {
    string token = "emtp";
}

message DeliveryResp {
    string token = "emtp";
}

message VerifyResp {
    bool res = false;
}
```

1) 用户服务
```protobuf
syntax="proto3";

package user;

option go_package="/user";

service UserService {
    rpc Register(RegisterReq) returns (RegisterResp) {}
    rpc Login(LoginReq) returns (LoginResp) {}
}

message RegisterReq {
    string email = 1;
    string password = 2;
    string confirm_password = 3;
}

message RegisterResp {
    int32 user_id = 1;
}

message LoginReq {
    string email= 1;
    string password = 2;
}

message LoginResp {
    int32 user_id = 1;
}
```
1) 商品服务
```protobuf
syntax = "proto3";

package product;

option go_package = "/product";

service ProductCatalogService {
  rpc ListProducts(ListProductsReq) returns (ListProductsResp) {}
  rpc GetProduct(GetProductReq) returns (GetProductResp) {}
  rpc SearchProducts(SearchProductsReq) returns (SearchProductsResp) {}
}

message ListProductsReq{
  int32 page = 1;
  int64 pageSize = 2;

  string categoryName = 3;
}

message Product {
  uint32 id = 1;
  string name = 2;
  string description = 3;
  string picture = 4;
  float price = 5;

  repeated string categories = 6;
}

message ListProductsResp {
  repeated Product products = 1;
}

message GetProductReq {
  uint32 id = 1;
}

message GetProductResp {
  Product product = 1;
}

message SearchProductsReq {
  string query = 1;
}

message SearchProductsResp {
  repeated Product results = 1;
}
```
1) 购物车服务
```protobuf
syntax = "proto3";

package cart;

option go_package = '/cart';

service CartService {
  rpc AddItem(AddItemReq) returns (AddItemResp) {}
  rpc GetCart(GetCartReq) returns (GetCartResp) {}
  rpc EmptyCart(EmptyCartReq) returns (EmptyCartResp) {}
}

message CartItem {
  uint32 product_id = 1;
  int32  quantity = 2;
}

message AddItemReq {
  uint32 user_id = 1;
  CartItem item = 2;
}

message AddItemResp {}

message EmptyCartReq {
  uint32 user_id = 1;
}

message GetCartReq {
  uint32 user_id = 1;
}

message GetCartResp {
  Cart cart = 1;
}

message Cart {
  uint32 user_id = 1;
  repeated CartItem items = 2;
}

message EmptyCartResp {}
```

1) 订单服务

```protobuf
syntax = "proto3";

package order;

import "cart.proto";

option go_package = "order";

service OrderService {
  rpc PlaceOrder(PlaceOrderReq) returns (PlaceOrderResp) {}
  rpc ListOrder(ListOrderReq) returns (ListOrderResp) {}
  rpc MarkOrderPaid(MarkOrderPaidReq) returns (MarkOrderPaidResp) {}
}

message Address {
  string street_address = 1;
  string city = 2;
  string state = 3;
  string country = 4;
  int32 zip_code = 5;
}

message PlaceOrderReq {
  uint32 user_id = 1;
  string user_currency = 2;

  Address address = 3;
  string email = 4;
  repeated OrderItem order_items = 5;
}

message OrderItem {
  cart.CartItem item = 1;
  float cost = 2;
}

message OrderResult {
  string order_id = 1;
}

message PlaceOrderResp {
  OrderResult order = 1;
}

message ListOrderReq {
  uint32 user_id = 1;
}

message Order {
  repeated OrderItem order_items = 1;
  string order_id = 2;
  uint32 user_id = 3;
  string user_currency = 4;
  Address address = 5;
  string email = 6;
  int32 created_at = 7;
}

message ListOrderResp {
  repeated Order orders = 1;
}

message MarkOrderPaidReq {
  uint32 user_id = 1;
  string order_id = 2;
}

message MarkOrderPaidResp {}
```

6) 结算服务
```protobuf
syntax = "proto3";

package  checkout;

import "payment.proto";

option go_package = "/checkout";

service CheckoutService {
  rpc Checkout(CheckoutReq) returns (CheckoutResp) {}
}

message Address {
  string street_address = 1;
  string city = 2;
  string state = 3;
  string country = 4;
  string zip_code = 5;
}

message CheckoutReq {
  uint32 user_id = 1;
  string firstname = 2;
  string lastname = 3;
  string email = 4;
  Address address = 5;
  payment.CreditCardInfo credit_card = 6;
}

message CheckoutResp {
  string order_id = 1;
  string transaction_id = 2;
}
1)  支付服务
syntax = "proto3";

package payment;

option go_package = "payment";


service PaymentService {
  rpc Charge(ChargeReq) returns (ChargeResp) {}
}

message CreditCardInfo {
  string credit_card_number = 1;
  int32 credit_card_cvv = 2;
  int32 credit_card_expiration_year = 3;
  int32 credit_card_expiration_month = 4;
}

message ChargeReq {
  float amount = 1;
  CreditCardInfo credit_card = 2;
  string order_id = 3;
  uint32 user_id = 4;
}

message ChargeResp {
  string transaction_id = 1;
}
```