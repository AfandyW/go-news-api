# News API

# Endpoints
### Header: 
- Content-Type: `application/json`
- Accept: `application/json`

### `[get]/api/tags` - List Tags  
- Response
    ```json
    {
        "code": 200,
        "status": "Success List Tags",
        "data": [
            {
                "id": 1,
                "name": "Crypto"
            }
        ]
    }
    ```
### `[post]/api/tags` - Create New Tags 
- Request Body
    - name: `string|required`
    ```json
    {
        "name": "Investment"
    }
    ```
- Response
    ```json
    {
        "code": 201,
        "status": "Success Create New Tags",
        "data": {
            "id": 1,
            "name": "Investment"
        }
    }
    ```
### `[put]/api/tags{tags_id}` - Update Tags by Id
- Request Body
    - name: `string|required`
    {
        "name": "Crypto"
    }
    ```
- Response
    ```json
    {
        "code": 201,
        "status": "Success Update Tags",
        "data": {
            "id": 1,
            "name": "Crypto"
        }
    }
    ```
### `[delete]/api/tags` - Delete Tags
- Response
    ```json
    {
        "code": 200,
        "status": "Success Deleted Tags"
    }
    ```
### `[post]/api/news/` - Create New News 
- Request Body
    - name: `string|required`
    - status: `string["draft","deleted","publish"]|required`
    - tags: `["arrayOfTags]`
    ```json
    {
        "name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
        "status": "draft",
        "tags": [
            "Investment", "Crypto"
        ]
    }
    ```
- Response
    ```json
    {
        "code": 201,
        "status": "Success Create New News",
        "data": {
            "Id": 1,
            "Name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
            "Status": "draft",
            "Tags": [
                {
                    "id": 1,
                    "name": "Investment"
                },
                {
                    "id": 2,
                    "name": "Crypto"
                }
            ]
        }
    }
    ```
### `[get]/api/news` - List News  
- Response
    ```json
    {
        "code": 200,
        "status": "Success List News",
        "data": [
            {
                "Id": 1,
                "Name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
                "Status": "draft",
                "Tags": [
                    {
                        "id": 1,
                        "name": "Investment"
                    },
                    {
                        "id": 2,
                        "name": "Crypto"
                    }
                ]
            }
        ]
    }
    ```
### `[put]/api/news/{news_id}` - Update News by Id
- Request Body
    - name: `string|required`
    - status: `string["draft","deleted","publish"]|required`
    - tags: `["arrayOfTags]`
    ```json
    {
        "name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
        "status": "publish",
        "tags": [
            "Investment", "Crypto"
        ]
    }
    ```
- Response
    ```json
    {
        "code": 201,
        "status": "Success Update News",
        "data": {
            "Id": 1,
            "Name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
            "Status": "publish",
            "Tags": [
                {
                    "id": 1,
                    "name": "Investment"
                },
                {
                    "id": 2,
                    "name": "Crypto"
                }
            ]
        }
    }
    ```
### `[delete]/api/{news_id}` - Delete News
- Response
    ```json
    {
        "code": 200,
        "status": "Success Deleted News"
    }
    ```

### `[get]/api/news?topic=` - Filter News by Topics/Tags
- Query Param:
    - topic: `string`
- Response
    ```json
    {
        "code": 200,
        "status": "Success List News By Topic",
        "data": [
            {
                "id": 1,
                "name": "Investment",
                "News": [
                    {
                        "Id": 1,
                        "Name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
                        "Status": "publish"
                    }
                ]
            }
        ]
    }
    ```
### `[get]/api/news?status=` - Filter News by Status News  
- Query Param:
    - status: `string`
- Response
    ```json
    {
        "code": 200,
        "status": "Success List News By Status",
        "data": [
            {
                "Id": 1,
                "Name": "Bukele and the volcano: El Salvador doubles-down on Bitcoin",
                "Status": "publish",
                "Tags": [
                    {
                        "id": 1,
                        "name": "Investment"
                    },
                    {
                        "id": 2,
                        "name": "Crypto"
                    }
                ]
            }
        ]
    }
    ```



