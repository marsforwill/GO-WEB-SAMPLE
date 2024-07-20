# Product API Documentation

## Endpoints

### Get Product By ID

**URL**: `/products/:id`

**Method**: `GET`

**Description**: Retrieves a product by its ID.

**URL Parameters**:
- `id` (required): The ID of the product to retrieve.

**Response**:
- `200 OK`:
  ```json
  {
    "product": {
      "ID": 1,
      "Name": "Wireless Mouse",
      "Category": "Electronics",
      "Price": 25.99,
      "Description": "Ergonomic wireless mouse with long battery life",
      "StockQuantity": 150,
      "CreatedAt": "2023-01-01T00:00:00Z",
      "UpdatedAt": "2023-01-01T00:00:00Z",
      "DeletedAt": null
    }
  }

  ### Get Products

**URL**: `/products`

**Method**: `GET`

**Description**: Retrieves a list of products with optional filters for category and price range, along with pagination and sorting options.

**Query Parameters**:
- `category` (optional): Filter products by category.
- `price_min` (optional): Minimum price of products to retrieve.
- `price_max` (optional): Maximum price of products to retrieve.
- `page` (optional): Page number for pagination (default is 1).
- `page_size` (optional): Number of items per page (default is 10).
- `sort_order` (optional): Sort order for the results, either `asc` or `desc` (default is `desc`).

**Response**:
- `200 OK`:
  ```json
  {
    "total_count": 100,
    "page": 1,
    "page_size": 10,
    "products": [
      {
        "ID": 1,
        "Name": "Wireless Mouse",
        "Category": "Electronics",
        "Price": 25.99,
        "Description": "Ergonomic wireless mouse with long battery life",
        "StockQuantity": 150,
        "CreatedAt": "2023-01-01T00:00:00Z",
        "UpdatedAt": "2023-01-01T00:00:00Z",
        "DeletedAt": null
      },
      {
        "ID": 2,
        "Name": "Bluetooth Speaker",
        "Category": "Electronics",
        "Price": 49.99,
        "Description": "Portable Bluetooth speaker with high-quality sound",
        "StockQuantity": 85,
        "CreatedAt": "2023-01-02T00:00:00Z",
        "UpdatedAt": "2023-01-02T00:00:00Z",
        "DeletedAt": null
      }
      // More products...
    ]
  }
