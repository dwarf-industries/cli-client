CREATE TABLE IF NOT EXISTS Keys(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    data TEXT NOT NULL,
    identity INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    identity TEXT,
    encryptionCertificate TEXT,
    orderSecret TEXT,
    created_at TEXT DEFAULT (datetime('now','localtime'))
);

CREATE TABLE IF NOT EXISTS Accounts (
    key TEXT,
    password TEXT
);

CREATE TABLE IF NOT EXISTS Nodes (
    name TEXT NOT NULL
);
