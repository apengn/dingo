package model

const posts = `
CREATE TABLE IF NOT EXISTS
posts (
  id                 integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title              varchar(150) NOT NULL,
  slug               varchar(150) NOT NULL,
  markdown           text,
  html               text,
  image              text,
  featured           BOOLEAN,
  page               BOOLEAN,
  allow_comment      BOOLEAN,
  published          BOOLEAN,
  comment_num        integer NOT NULL DEFAULT '0',
  language           varchar(6) NOT NULL DEFAULT 'en_US',
  meta_title         varchar(150),
  meta_description   varchar(200),
  created_at         datetime NOT NULL,
  created_by         integer NOT NULL,
  updated_at         datetime,
  updated_by         integer,
  published_at       datetime,
  published_by       integer
);`

const tokens = `
CREATE TABLE IF NOT EXISTS
tokens (
  id          integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  value       varchar(40) NOT NULL,
  user_id     integer UNIQUE,
  created_at  datetime,
  expired_at  datetime
);`
const users = `
CREATE TABLE IF NOT EXISTS
users (
  id               integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name             varchar(150) NOT NULL,
  slug             varchar(150) NOT NULL,
  password         varchar(60) NOT NULL,
  email            varchar(254) NOT NULL,
  image            text,
  cover            text,
  bio              varchar(200),
  website          text,
  location         text,
  accessibility    text,
  status           varchar(150) NOT NULL DEFAULT 'active',
  language         varchar(6) NOT NULL DEFAULT 'en_US',
  last_login       datetime,
  created_at       datetime NOT NULL,
  created_by       integer NOT NULL,
  updated_at       datetime,
  updated_by       integer
);`
const categories = `
CREATE TABLE IF NOT EXISTS
categories (
  id                integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name              varchar(150) NOT NULL,
  slug              varchar(150) NOT NULL,
  description       varchar(200),
  parent_id         integer,
  meta_title        varchar(150),
  meta_description  varchar(200),
  created_at        datetime NOT NULL,
  created_by        integer NOT NULL,
  updated_at        datetime,
  updated_by        integer
);
`
const tags = `
CREATE TABLE IF NOT EXISTS
tags (
  id                integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name              varchar(150) NOT NULL,
  slug              varchar(150) NOT NULL,
  hidden            boolean NOT NULL DEFAULT 0,
  created_at        datetime NOT NULL,
  created_by        integer NOT NULL,
  updated_at        datetime,
  updated_by        integer
);
`
const comments = `
CREATE TABLE IF NOT EXISTS
comments (
  id            integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  post_id       varchar(150) NOT NULL,
  author        varchar(150) NOT NULL,
  author_email  varchar(150) NOT NULL,
  author_avatar varchar(50)  NOT NULL,
  author_url    varchar(200) NOT NULL,
  author_ip     varchar(100) NOT NULL,
  created_at    datetime     NOT NULL,
  content       text         NOT NULL,
  approved      tinyint      NOT NULL DEFAULT '0',
  agent         varchar(255) NOT NULL,
  type          varchar(20),
  parent        integer,
  user_id       integer
);
`
const posts_tags = `
CREATE TABLE IF NOT EXISTS
posts_tags (
  id       integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  post_id  integer NOT NULL,
  tag_id   integer NOT NULL
);
`
const posts_categories = `
CREATE TABLE IF NOT EXISTS
posts_categories (
  id           integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  post_id      integer NOT NULL,
  category_id  integer NOT NULL
);
`
const settings = `
CREATE TABLE IF NOT EXISTS
settings (
  id          integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  key_word         varchar(150) NOT NULL,
  value       varchar(20) NOT NULL,
  type        varchar(150) NOT NULL DEFAULT 'core',
  created_at  datetime NOT NULL,
  created_by  integer NOT NULL,
  updated_at  datetime,
  updated_by  integer
);
`
const roles = `
CREATE TABLE IF NOT EXISTS
roles (
  id           integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name         varchar(150) NOT NULL,
  description  varchar(200),
  created_at   datetime NOT NULL,
  created_by   integer NOT NULL,
  updated_at   datetime,
  updated_by   integer
);
`
const messages = `
CREATE TABLE IF NOT EXISTS
messages (
  id           integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  type         varchar(20) NOT NULL,
  data         text NOT NULL,
  is_read      boolean NOT NULL default 0,
  created_at   datetime NOT NULL
);
`

var CreateTables createTables

func init() {
	createTableStr := []string{posts, tokens, users, categories, tags,
		comments, posts_tags, posts_categories, settings, roles, messages}
	CreateTables = createTables{createTableStr: createTableStr}
}

type createTables struct {
	createTableStr []string
}
