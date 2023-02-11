package storage

var (
	// INSERT INTO main_storage (userID, data, dataType, meta)
	// VALUES (($1, $2, $3, $4) RETURNING main_storage.id;
	insertNewDataSQL = "INSERT INTO main_storage (userID, data, dataType, meta) " +
		"VALUES ($1, $2, $3, $4) RETURNING main_storage.id"

	// SELECT data, dataType, meta FROM main_storage WHERE userID=$1 AND id=$2;
	getDataSQL = "SELECT data, dataType, meta FROM main_storage WHERE userID=$1 AND id=$2"

	// UPDATE main_storage SET (data, meta) = ($3, $4) WHERE userID=$1 AND id=$2 RETURNING main_storage.dataType;
	setDataSQL = "UPDATE main_storage SET (data, meta) = " +
		"($3, $4) WHERE userID=$1 AND id=$2 RETURNING main_storage.dataType"

	// DELETE FROM main_storage WHERE userID=$1 AND id=$2;
	deleteSQL = "DELETE FROM main_storage WHERE userID=$1 AND id=$2"

	// SELECT id, data, dataType, meta FROM main_storage WHERE userID=$1 LIMIT $2 OFFSET $3;
	getBatchSQL = "SELECT id, data, dataType, meta FROM main_storage WHERE userID=$1 LIMIT $2 OFFSET $3"

	createTableSQL = "CREATE TABLE IF NOT EXISTS main_storage (" +
		"userID text, " +
		"id SERIAL, " +
		"data bytea, " +
		"dataType smallint," +
		"meta bytea " +
		");"
)
