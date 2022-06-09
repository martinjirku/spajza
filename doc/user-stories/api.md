# Data

- user
    - POST /user/register
    - POST /user/login
    - POST /user/logout
- storage
    - GET /amount-type      {id: ID, title: string, amountType: amountTypeId}
    - POST /amount-type - create new
    - GET /product-type     {id: ID, title: string, amountType: amountTypeId}
    - POST /product-type
    - GET /product-type/:id
    - GET /storage-item     {id: ID, title: string, amount: number, amountType: amountTypeId, productTypeId: ID}
    - POST /storage-item
    - GET /storage-item/:id

