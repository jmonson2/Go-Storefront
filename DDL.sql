create table inventory
(
    sku   varchar not null
        constraint inventory_pk
            primary key,
    name  varchar not null,
    stock integer not null,
    price real    not null
);

create table product
(
    ID  integer not null
        constraint ID
            primary key autoincrement,
    SKU varchar
        references inventory (SKU)
        constraint product_inventory_SKU_fk
            references inventory (SKU)
);

CREATE TRIGGER update_inventory_stock_after_insert
    AFTER INSERT ON product
    FOR EACH ROW
BEGIN
    UPDATE inventory
    SET stock = stock + 1
    WHERE sku = NEW.sku;
END;

create table sold_products
(
    id  integer not null
        constraint sold_products_pk
            primary key,
    sku varchar not null
        constraint sold_products_inventory_sku_fk
            references inventory
);

CREATE TRIGGER update_inventory_stock_after_sale
    AFTER INSERT ON sold_products
    FOR EACH ROW
BEGIN
    UPDATE inventory
    SET stock = stock - 1
    WHERE sku = NEW.sku;
END;

create table storefront
(
    name    varchar not null,
    id      integer not null
        constraint storefront_pk
            primary key autoincrement,
    revenue float64 not null
);
