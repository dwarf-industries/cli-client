CREATE TABLE IF NOT EXISTS Keys(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    encryption_certificate TEXT NOT NULL,
    identity_certificate TEXT NOT NULL,
    encryption_key TEXT NOT NULL,
    priv TEXT NOT NULL,
    order_sercret TEXT NOT NULL,
    user_id INTEGER UNIQUE
);

CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    identity_contract TEXT,
    encryption_certificate TEXT,
    order_secret TEXT,
    created_at TEXT DEFAULT (datetime('now','localtime'))
);

CREATE TABLE IF NOT EXISTS Accounts (
    key TEXT,
    password TEXT
);

CREATE TABLE IF NOT EXISTS Nodes (
    name TEXT NOT NULL
);
