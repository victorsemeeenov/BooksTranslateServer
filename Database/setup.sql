drop table if exists sessions CASCADE;
drop table if exists users CASCADE;
drop table if exists words CASCADE;
drop table if exists translations CASCADE;
drop table if exists sentences CASCADE;
drop table if exists word_sentences CASCADE;
drop table if exists chapters CASCADE;
drop table if exists books CASCADE;
drop table if exists authors CASCADE;
drop table if exists books_and_authors CASCADE;

create table users (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    name       varchar(255),
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp not null   
);

create table sessions (
    id         serial primary key,
    uuid       varchar(64) not null unique,
    email      varchar(255),
    user_id    integer references users(id),
    created_at timestamp not null   
);

create table words (
    id    serial primary key,
    value varchar(255)
);

create table translations (
    id          serial primary key,
    word_id     integer references words(id),
    translation varchar (255)
);

create table books (
    id              serial primary key,
    name            varchar (255),
    number_of_pages integer,
    year            integer
);

create table chapters (
    id          serial primary key,
    title       varchar (255),
    index       integer,
    index_value varchar (255),
    book_id     integer references books(id)
);

create table sentences (
    id         serial primary key,
    value      text,
    index      integer,
    chapter_id integer references chapters(id)  
);

create table word_sentences (
    id           serial primary key,
    word_id      integer references words(id),
    sentences_id integer references sentences(id),
    index        integer
);

create table authors (
    id           serial primary key,
    name varchar (255)
);

create table books_and_authors (
    id        serial primary key,
    book_id   integer references books(id),
    author_id integer references authors(id)
);

