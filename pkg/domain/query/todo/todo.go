package todo

// Login : query for login auth
var Login = `SELECT * FROM auth a WHERE a.username = $1;`

// AddTodo : query for create a new todo
var AddTodo = `
	INSERT INTO TODO (
					title, 
					description, 
					owner_id, 
					created_at
					) VALUES ($1, $2, $3, $4) RETURNING *;`

// UpdateTodo : query for update a todo
var UpdateTodo = `
	UPDATE todo SET title=$1, description=$2, updated_at=$3 WHERE id = $4 RETURNING *;
`
// GetTodo : query for select a todo
var GetTodo = `SELECT * FROM todo WHERE id = $1;`

// ListTodo : query for list all todo by owner_id
var ListTodo = `SELECT * FROM todo WHERE owner_id = $1;`

// DeleteTodo : query for delete a todo
var DeleteTodo = `DELETE FROM todo WHERE id = $1;`

// IsExist : query for checking if the selected todo is exist
var IsExist = `SELECT COUNT(*) FROM todo WHERE id = $1;`

// IsOwnerExist : query for checking if the selected owner is exist
var IsOwnerExist = `SELECT COUNT(*) FROM auth WHERE id = $1;`