CREATE TABLE IF NOT EXISTS priorities(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

INSERT INTO priorities(name) VALUES
    ('low'),
    ('medium'),
    ('high');

CREATE TABLE IF NOT EXISTS categories(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

INSERT INTO categories(name)
VALUES ('default');

CREATE TABLE IF NOT EXISTS todos(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    priority_id INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL DEFAULT 0,

    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(priority_id) REFERENCES priorities(id)
)
