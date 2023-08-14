-- +goose Up
-- +goose StatementBegin
create table if not exists products (
    id int not null auto_increment primary key,
    title varchar(255) not null,
    description text not null,
    price int not null
);
-- +goose StatementEnd

-- +goose StatementBegin
create table if not exists carts (
    id int not null auto_increment primary key,
    session_id varchar(255) not null
);
-- +goose StatementEnd

-- +goose StatementBegin
create table if not exists cart_products (
    product_id int not null,
    cart_id int not null,

    quantity int not null,
    price int not null,

    primary key (product_id, cart_id),
    constraint fk_product_id foreign key (product_id) references products(id),
    constraint fk_cart_id foreign key (cart_id) references carts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table cart_products;
-- +goose StatementEnd

-- +goose StatementBegin
drop table carts;
-- +goose StatementEnd

-- +goose StatementBegin
drop table products;
-- +goose StatementEnd
