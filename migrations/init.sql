create schema product;

create table product.warehouse (
   id uuid primary key not null,
   name text not null
       constraint name_length check (length(name) >= 2 and length(name) <= 100),
   availability boolean not null
);

create table product.product (
    id uuid primary key not null,
    name text not null
        constraint name_length check (length(name) >= 2 and length(name) <= 100),
    count integer not null
        constraint count_size check (count <= 1000),
    length real not null
        constraint positive_length check (length > 0),
    width real not null
         constraint positive_width check (width > 0),
    height real not null
         constraint positive_height check (height > 0),
    unit text not null,
    warehouse_id uuid not null references product.warehouse(id),
    is_reserved boolean not null default false
);