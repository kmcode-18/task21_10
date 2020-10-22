# task21_10

## End Point(Image Get) - /image
### Request Method - GET
### Query Parameters - id=&offset=&limit=&sort_by=image_name&sort_order=asc
### Response - 
`{
    "result": [
        {
            "image_id": 2,
            "image_name": "go_cover.png",
            "created_at": "2020-10-21T19:44:56.892302Z"
        },
        {
            "image_id": 1,
            "image_name": "unnamed.jpg",
            "created_at": "2020-10-21T19:44:22.690508Z"
        }
    ]
}`

## End Point(Upload Get) - /image
### Request Method - POST
### Request Body - form-data with "file" key and image as a value
### Response - 
`{
    "message": "image uploaded successfully"
}`
