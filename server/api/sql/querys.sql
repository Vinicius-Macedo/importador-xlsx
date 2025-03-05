-- name: CreateUser :exec
INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CountUserByUsername :one
SELECT COUNT(*) FROM users WHERE username = $1;

-- name: CountUserByEmail :one
SELECT COUNT(*) FROM users WHERE email = $1;

-- name: UpdateUserPasswordByID :exec
UPDATE users SET password = $1 WHERE id = $2;


-- name: CreatePartner :copyfrom
INSERT INTO partners (partner_key, name, mpn_id) VALUES ($1, $2, $3);

-- name: CreateCustomer :copyfrom
INSERT INTO customers (customer_key, name, domain_name, country, tier_to_mpn_id) VALUES ($1, $2, $3, $4, $5);

-- name: CreateProduct :copyfrom
INSERT INTO products (product_key, name, sku_id) VALUES ($1, $2, $3);

-- name: CreateSku :copyfrom
INSERT INTO skus (sku_key, name, availability_id) VALUES ($1, $2, $3);

-- name: CreatePublisher :copyfrom
INSERT INTO publishers (publisher_key, name) VALUES ($1, $2);

-- name: CreateSubscription :copyfrom
INSERT INTO subscriptions (subscription_key, description) VALUES ($1, $2);

-- name: CreateMeter :copyfrom
INSERT INTO meters (meters_key, type, category, sub_category, name, region, unit) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateResource :copyfrom
INSERT INTO resources (uri, location, consumed_service, resource_group, info1, info2, tags, additional_info) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: CreateEntitlement :copyfrom
INSERT INTO entitlements (entitlement_key, description) VALUES ($1, $2);

-- name: CreateBenefit :copyfrom
INSERT INTO benefits (benefit_key, benefit_order_id, benefit_type) VALUES ($1, $2, $3);

-- name: CreateBilling :copyfrom
INSERT INTO billings (partner_key, customer_key, product_key, publisher_key, subscription_key, meters_key, resource_uri, entitlement_key, benefit_key,invoice, unit_price, quantity, unit_type, billing_pre_tax_total, billing_currency, pricing_pre_tax_total, pricing_currency, effective_unit_price, pc_to_bc_exchange_rate, pc_to_bc_exchange_rate_date, charge_start_date, charge_end_date, usage_date, charge_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24);


-- name: GetAllCustomers :many
SELECT
  customer_key,
  name,
  domain_name,
  country,
  tier_to_mpn_id
FROM
  customers
ORDER BY
  name;

-- name: GetChargeMonths :many
SELECT
  DISTINCT EXTRACT(MONTH FROM charge_start_date) AS month,
  EXTRACT(YEAR FROM charge_start_date) AS year
FROM
  billings
ORDER BY
  year DESC,
  month DESC;




-- name: GetAllResources :many
SELECT
  id,
  uri,
  location,
  consumed_service,
  resource_group
FROM
  resources
ORDER BY
  resource_group;

-- name: GetCategories :many
SELECT meters_key, type, category, sub_category, name, region, unit FROM meters ORDER BY category, sub_category, name;