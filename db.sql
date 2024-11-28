CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime'))
);


CREATE TABLE IF NOT EXISTS User_Keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    key_data TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    key_type INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (key_type) REFERENCES Key_Types(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS User_Certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    certificate_data TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    certificate_type INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (certificate_type) REFERENCES Certificate_Types(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS Key_Types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
)


CREATE TABLE IF NOT EXISTS Certificate_Types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
)
