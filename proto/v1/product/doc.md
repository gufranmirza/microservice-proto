# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [product.proto](#product.proto)
    - [Product](#v1.product.Product)
    - [Products](#v1.product.Products)
  
    - [ProductCategory](#v1.product.ProductCategory)
  
- [Scalar Value Types](#scalar-value-types)



<a name="product.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product.proto



<a name="v1.product.Product"></a>

### Product
Product message structure


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id of the product |
| name | [string](#string) |  | name of the product |
| description | [string](#string) |  | product description |
| manufacturer | [string](#string) |  | product manufacturer |
| price | [string](#string) |  | product price |
| in_stock | [bool](#bool) |  | product is in stock or not |
| category | [ProductCategory](#v1.product.ProductCategory) |  | category of the product it belogs to |






<a name="v1.product.Products"></a>

### Products
Products contains the list of products


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#v1.product.Product) | repeated |  |





 


<a name="v1.product.ProductCategory"></a>

### ProductCategory
ProductCategory is the category a product belogs to

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Product has not any category unspecified |
| BOOKS_AND_LITERATURE | 1 | Books and literature |
| FOOD_AND_DRINK | 2 | Food and drink |
| COMPUTERS_AND_ELECTONICS | 3 | Computers and electronic |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

