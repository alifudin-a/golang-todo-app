package todo

var Login = `SELECT * FROM auth a WHERE a.username = $1;`