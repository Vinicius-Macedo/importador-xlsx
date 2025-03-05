CREATE TABLE
  IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS partners (
    id SERIAL PRIMARY KEY,
    partner_key VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255),
    mpn_id VARCHAR(50)
  );

CREATE TABLE
  IF NOT EXISTS customers (
    id serial PRIMARY KEY,
    customer_key VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    domain_name VARCHAR(255),
    country VARCHAR(2),
    tier_to_mpn_id VARCHAR(50)
  );

CREATE TABLE
  IF NOT EXISTS skus (
    id SERIAL PRIMARY KEY,
    sku_key VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50),
    availability_id VARCHAR(50)
  );

CREATE TABLE
  IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_key VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255),
    sku_id VARCHAR(50) REFERENCES skus (sku_key)
  );

CREATE TABLE
  IF NOT EXISTS publishers (
    id SERIAL PRIMARY KEY,
    publisher_key VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255)
  );

CREATE TABLE
  IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    subscription_key VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255)
  );

CREATE TABLE
  IF NOT EXISTS meters (
    id SERIAL PRIMARY KEY,
    meters_key VARCHAR(255) UNIQUE NOT NULL,
    type VARCHAR(255),
    category VARCHAR(255),
    sub_category VARCHAR(255),
    name VARCHAR(255),
    region VARCHAR(255),
    unit VARCHAR(50)
  );

CREATE TABLE
  IF NOT EXISTS resources (
    id SERIAL PRIMARY KEY,
    uri TEXT UNIQUE NOT NULL,
    location VARCHAR(255),
    consumed_service VARCHAR(255),
    resource_group VARCHAR(255),
    info1 TEXT,
    info2 TEXT,
    tags TEXT,
    additional_info TEXT
  );

CREATE TABLE
  IF NOT EXISTS entitlements (
    id SERIAL PRIMARY KEY,
    entitlement_key varchar(255) UNIQUE NOT NULL,
    description VARCHAR(255)
  );

CREATE TABLE
  IF NOT EXISTS benefits (
    id SERIAL PRIMARY KEY,
    benefit_key VARCHAR(255) UNIQUE NOT NULL,
    benefit_order_id VARCHAR(255),
    benefit_type VARCHAR(255)
  );

CREATE TABLE
  IF NOT EXISTS billings (
    id SERIAL PRIMARY KEY,
    partner_key VARCHAR(50) REFERENCES partners (partner_key),
    customer_key VARCHAR(255) REFERENCES customers (customer_key),
    product_key VARCHAR(50) REFERENCES products (product_key),
    publisher_key VARCHAR(50) REFERENCES publishers (publisher_key),
    subscription_key VARCHAR(255) REFERENCES subscriptions (subscription_key),
    meters_key VARCHAR(255) REFERENCES meters (meters_key),
    resource_uri TEXT REFERENCES resources (uri),
    entitlement_key VARCHAR(255) REFERENCES entitlements (entitlement_key),
    benefit_key VARCHAR(255) REFERENCES benefits (benefit_key),
    invoice VARCHAR(255),
    unit_price DECIMAL(18, 8),
    quantity DECIMAL(18, 8),
    unit_type VARCHAR(50),
    billing_pre_tax_total DECIMAL(18, 8),
    billing_currency VARCHAR(50),
    pricing_pre_tax_total DECIMAL(18, 8),
    pricing_currency VARCHAR(50),
    effective_unit_price DECIMAL(18, 8),
    pc_to_bc_exchange_rate DECIMAL(18, 8),
    pc_to_bc_exchange_rate_date DATE,
    charge_start_date DATE,
    charge_end_date DATE,
    usage_date DATE,
    charge_type varchar(50)
  );

CREATE INDEX idx_billing_customer_key ON billings (customer_key);
CREATE INDEX idx_billing_meter_key ON billings (meters_key);
CREATE INDEX idx_billing_resource_uri ON billings (resource_uri);
CREATE INDEX idx_billing_charge_start_date ON billings (charge_start_date);
CREATE INDEX idx_meter_category ON meters (category);