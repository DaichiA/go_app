CREATE DATABASE go_app;

-- データベースへの接続
\c go_app;

-- テーブル作成
CREATE TABLE breads (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
