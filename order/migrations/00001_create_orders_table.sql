-- +goose Up
-- +goose StatementBegin
create table orders (
    id serial primary key,
    orderUuid text not null,
    userUuid text not null,
    partUuids text[] not null,
    totalPrice REAL not null,
    transactionUuid text,
    paymentMethod text,
    status text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd
