CREATE TABLE orders
(
    id          varchar(255) NOT NULL,
    price       float        NOT NULL,
    tax         float        NOT NULL,
    final_price float        NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO orders (id, price, tax, final_price) values (1, 100, 10, 110);
INSERT INTO orders (id, price, tax, final_price) values (2, 200, 20, 220);
INSERT INTO orders (id, price, tax, final_price) values (3, 300, 30, 330);
INSERT INTO orders (id, price, tax, final_price) values (4, 300, 30, 330);
INSERT INTO orders (id, price, tax, final_price) values (5, 300, 30, 330);