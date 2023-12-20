package config

const (
	InsertFasilities     = `INSERT INTO facilities (name, quantity) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	SelectFasilitiesList = `SELECT id, name, quantity, created_at, updated_at FROM facilities ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectFasilitiesById = `SELECT id, name, quantity, created_at, updated_at FROM facilities WHERE id = $1`
	UpdateFasilities     = `UPDATE facilities SET name = $1, quantity = $2, updated_at = $3 WHERE id = $4 RETURNING id, created_at`
	TotalRowsFasilities  = `SELECT COUNT(*) FROM facilities`
)
