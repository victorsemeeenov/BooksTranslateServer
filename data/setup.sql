drop table if exists access_tokens CASCADE;
drop table if exists refresh_tokens CASCADE;
drop table if exists users CASCADE;
drop table if exists words CASCADE;
drop table if exists translations CASCADE;
drop table if exists sentences CASCADE;
drop table if exists words_sentences CASCADE;
drop table if exists chapters CASCADE;
drop table if exists books CASCADE;
drop table if exists authors CASCADE;
drop table if exists books_authors CASCADE;
drop table if exists languages CASCADE;
drop table if exists words_translations CASCADE;
drop table if exists translations CASCADE;
drop table if exists sentence_translations CASCADE;
drop table if exists synonims CASCADE;
drop table if exists words_synonims CASCADE;
drop table if exists book_categories CASCADE;
drop table if exists users_words CASCADE;
drop index if exists sentence_index;
drop index if exists chapter_index;
drop index if exists word_index;


create table users (
    id         serial primary key,
    name       varchar(64) not null unique,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    avatar_url varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table access_tokens (
    id         serial primary key,
    value      varchar(500) not null unique,
    user_id    integer references users(id),
    expired_in timestamp not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table refresh_tokens (
    id         serial primary key,
    value      varchar(500) not null unique,
    user_id    integer references users(id),
    expired_in timestamp not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table languages (
    id    serial primary key,
    value varchar(64) not null unique,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table words (
    id             serial primary key,
    value          varchar(255) not null unique,
    transcription  varchar(255),
    part_of_speech varchar(64),
    language_id    integer references languages(id),
    created_at     timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);

create table translations (
    id             serial primary key,
    value          varchar(255) not null unique,
    part_of_speech varchar(64),
    gender         varchar(32),
    created_at     timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);

create table words_translations (
    id             serial primary key,
    word_id        integer references words(id),
    translation_id integer references translations(id),
    created_at     timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);

create index word_index on words((lower(value)));

create table synonims (
    id                serial primary key,
    translation_value varchar(255) not null unique,
    part_of_speech    varchar(64),
    gender            varchar(32),
    created_at        timestamp,
    updated_at        timestamp,
    deleted_at        timestamp
);

create table words_synonims (
    id         serial primary key,
    word_id    integer references words(id),
    synonim_id integer references synonims(id),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table book_categories (
    id         serial primary key,
    name       varchar(255) not null unique,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table books (
    id               serial primary key,
    name             varchar (255) not null,
    number_of_pages  integer,
    year             integer,
    url              varchar (255) not null,
    book_category_id integer references book_categories(id),
    language_id      integer references languages(id),
    created_at       timestamp,
    updated_at       timestamp,
    deleted_at       timestamp
);

create table chapters (
    id           serial primary key,
    title        varchar (255),
    order_number integer,
    order_value  varchar (255),
    book_id      integer references books(id),
    created_at   timestamp,
    updated_at   timestamp,
    deleted_at   timestamp
);

create index chapter_index on chapters(order_number);

create table sentences (
    id           serial primary key,
    value        text unique,
    order_number integer,
    chapter_id   integer references chapters(id),
    language_id  integer references languages(id),
    created_at   timestamp,
    updated_at   timestamp,
    deleted_at   timestamp
);

create index sentence_index on sentences(order_number);

create table words_sentences (
    id          serial primary key,
    word_id     integer references words(id),
    sentence_id integer references sentences(id),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

create table sentence_translations (
    id          serial primary key,
    sentence_id integer references sentences(id),
    value       varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

create table authors (
    id               serial primary key,
    name             varchar (255),
    created_at       timestamp,
    updated_at       timestamp,
    deleted_at       timestamp
);

create table books_authors (
    id         serial primary key,
    book_id    integer references books(id),
    author_id  integer references authors(id),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table users_words (
    id      serial primary key,
    user_id integer references users(id),
    word_id integer references words(id)
);
