CREATE TABLE IF NOT EXISTS public.purchase_orders(
    id VARCHAR(36) PRIMARY KEY,
    total_amount decimal(19, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.purchase_order_items(
    id VARCHAR(36) PRIMARY KEY,
    purchase_order_id VARCHAR(36) NOT NULL,
    item_id VARCHAR(36) NOT NULL,
    item_price decimal(19, 2) NOT NULL,
    quantity integer NOT NULL,
    amount decimal(19, 2) NOT NULL
)