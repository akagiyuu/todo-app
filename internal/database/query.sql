-- name: CreateCategory :one
INSERT INTO categories(name)
VALUES (?)
RETURNING id;

-- name: GetCategories :many
SELECT name FROM categories;

-- name: GetCategory :one
SELECT id FROM categories WHERE name = ?;

-- name: GetPriorities :many
SELECT name FROM priorities;

-- name: GetPriority :one
SELECT id FROM priorities WHERE name = ?;

-- name: CreateTodo :exec
INSERT INTO todos(title, description, priority_id, category_id)
VALUES (?, ?, ?, ?);

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;

-- name: GetTodo :many
SELECT title, description, priorities.name as priority, categories.name as category
FROM todos
INNER JOIN priorities ON priorities.id = priority_id
INNER JOIN categories ON categories.id = category_id
WHERE
    (priorities.name = @priority OR @priority = '') AND
    (categories.name = @category OR @category = '');
