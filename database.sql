create table customers
(
    id       serial
        primary key,
    name     varchar(255),
    hp       varchar(20),
    address  text,
    username varchar(255),
    password varchar(255)
);

alter table customers
    owner to postgres;

create table customers
(
    id       serial
        primary key,
    name     varchar(255),
    hp       varchar(20),
    address  text,
    username varchar(255),
    password varchar(255)
);

alter table customers
    owner to postgres;

create table bank
(
    id           serial
        primary key,
    bank_name    varchar(100),
    bank_account integer,
    id_cust      integer
);

alter table bank
    owner to postgres;

create table pocket
(
    id      serial
        primary key,
    id_cust integer,
    saldo   integer
);

alter table pocket
    owner to postgres;

create table term
(
    id                      serial
        primary key,
    principal_deposit       integer,
    deposit_interest_profit integer,
    deposit_interest_tax    integer,
    total_investment        integer,
    id_cust                 integer
);

alter table term
    owner to postgres;


INSERT INTO public.bank (id, bank_name, bank_account, id_cust) VALUES (1, 'SUPERBANK', 3847832, 4);
INSERT INTO public.bank (id, bank_name, bank_account, id_cust) VALUES (2, 'BNI', 48499372, 5);

INSERT INTO public.customers (id, name, hp, address, username, password) VALUES (4, 'tester', '62849485933', 'Jl. Imam Bonjol no.1', 'tester', '$2a$10$Kn7gN56ub8ZY/rOeL4RzoerLaCpu1XIULljnWG3ecE5vYmHPp8vmm');
INSERT INTO public.customers (id, name, hp, address, username, password) VALUES (5, 'pentest', '6237378444', 'Jl. Ir. H. Juanda no.1', 'pentest', '$2a$10$Kn7gN56ub8ZY/rOeL4RzoerLaCpu1XIULljnWG3ecE5vYmHPp8vmm');

INSERT INTO public.pocket (id, id_cust, saldo) VALUES (1, 4, 10000000);
INSERT INTO public.pocket (id, id_cust, saldo) VALUES (2, 5, 20000000);

INSERT INTO public.term (id, principal_deposit, deposit_interest_profit, deposit_interest_tax, total_investment, id_cust) VALUES (2, 5000000, 125000, 25000, 5100000, 4);
INSERT INTO public.term (id, principal_deposit, deposit_interest_profit, deposit_interest_tax, total_investment, id_cust) VALUES (3, 10000000, 250000, 50000, 10300000, 5);
