-- +goose Up
-- +goose StatementBegin
create table products (
    id int not null auto_increment primary key,
    title varchar(255) not null,
    description text not null,
    price int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table products;
-- +goose StatementEnd
