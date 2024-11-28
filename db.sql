CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    certificate TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime'))
);


CREATE TABLE IF NOT EXISTS User_Keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    key_data TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS User_Certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    certificate_data TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
)

