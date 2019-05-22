package ventures

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-qlueless-api/api/database"
)

// CreateTables creates all the Venture tables, views and triggers within the
// supplied database.
func CreateTables() (err error) {
	err = createVentureTable()
	if err != nil {
		return
	}

	err = createQlVentureTable()
	if err != nil {
		return
	}

	err = createInsertOnLivingVentureTrigger()
	if err != nil {
		return
	}

	err = createInsertOnDeadVentureTrigger()
	if err != nil {
		return
	}

	err = createUpdateOnVentureTrigger()
	if err != nil {
		return
	}

	err = createDeleteOnVentureTrigger()
	if err != nil {
		return
	}

	return
}

// createVentureTable creates the Venture table within the supplied database.
func createVentureTable() error {
	return execStmt(`CREATE TABLE venture (
		id INTEGER NOT NULL,
		last_modified INTEGER NOT NULL DEFAULT(CAST(ROUND((julianday('now') - 2440587.5)*86400000) As INTEGER)),
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_dead BOOL NOT NULL DEFAULT FALSE,
		extra TEXT NOT NULL DEFAULT "",
		PRIMARY KEY(id, last_modified)
	);`)
}

// createQlVentureTable creates the query layer Venture table within the
// supplied database.
func createQlVentureTable() error {
	return execStmt(`CREATE TABLE ql_venture (
		id INTEGER NOT NULL PRIMARY KEY,
		last_modified INTEGER NOT NULL,
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_dead BOOL NOT NULL,
		extra TEXT NOT NULL
	);`)
}

// createInsertOnLivingVentureTrigger creates a trigger within the
// supplied database that updates the ql_venture table when ever a new, and
// living, Venture is inserted into the venture table.
func createInsertOnLivingVentureTrigger() error {
	return execStmt(`CREATE TRIGGER insert_on_living_venture
		AFTER INSERT ON venture
		FOR EACH ROW
		WHEN (NEW.is_dead = false)
		BEGIN
			REPLACE INTO ql_venture (
				id, last_modified, description, order_ids, state, is_dead, extra
			) VALUES (
				NEW.id, NEW.last_modified, NEW.description, NEW.order_ids, NEW.state, NEW.is_dead, NEW.extra
			);
		END;`)
}

// createInsertOnDeadVentureTrigger creates a trigger within the supplied
// database that removes from the ql_venture table the dead Venture inserted
// into the venture table.
func createInsertOnDeadVentureTrigger() error {
	return execStmt(`CREATE TRIGGER insert_on_dead_venture
		AFTER INSERT ON venture
		FOR EACH ROW
		WHEN (NEW.is_dead = true)
		BEGIN
			DELETE FROM ql_venture 
			WHERE id = NEW.id;
		END;`)
}

// createUpdateOnVentureTrigger creates a trigger within the supplied
// database that raises an error if an update is attempted.
func createUpdateOnVentureTrigger() error {
	return execStmt(`CREATE TRIGGER update_on_venture
		BEFORE UPDATE ON venture
		BEGIN
			SELECT RAISE(FAIL, "Updates not allowed, insert with the same Venture ID!");
		END;`)
}

// createDeleteOnVentureTrigger creates a trigger within the supplied
// database that raises an error if a delete is attempted.
func createDeleteOnVentureTrigger() error {
	return execStmt(`CREATE TRIGGER delete_on_venture
		BEFORE DELETE ON venture
		BEGIN
			SELECT RAISE(FAIL, "Deletions not allowed!");
		END;`)
}

// execStmt executes a SQL statment ensuring it is closed afterwards
func execStmt(sql string) error {
	stmt, err := database.Get().Prepare(sql)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

// QueryFor queries the database for a single Venture.
func QueryFor(id string) (*Venture, error) {
	ven := Venture{}
	err := database.Get().QueryRow(`SELECT
		id,
		last_modified,
		description,
		order_ids,
		state,
		extra
	FROM ql_venture
	WHERE id = ?`, id).Scan(&ven.ID,
		&ven.LastModified,
		&ven.Description,
		&ven.Orders,
		&ven.State,
		&ven.Extra)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case cookies.LogIfErr(err):
		return nil, err
	}

	return &ven, nil
}

// QueryMany queries the database for all specified Ventures.
func QueryMany(ids []interface{}) ([]Venture, error) {
	posParams := strings.Repeat(",?", len(ids))[1:]
	sql := fmt.Sprintf(`SELECT
			id,
			last_modified,
			description,
			order_ids,
			state,
			extra
		FROM ql_venture
		WHERE id IN (%s)`, posParams)

	rows, err := database.Get().Query(sql, ids...)

	if rows != nil {
		defer rows.Close()
	}

	if cookies.LogIfErr(err) {
		return nil, err
	}

	return mapRows(rows)
}

// QueryAll queries the database for all Ventures.
func QueryAll() ([]Venture, error) {
	rows, err := database.Get().Query(`SELECT
		id,
		last_modified,
		description,
		order_ids,
		state,
		extra
	FROM ql_venture`)

	if rows != nil {
		defer rows.Close()
	}

	if cookies.LogIfErr(err) {
		return nil, err
	}

	return mapRows(rows)
}

// mapRows is a file private function that maps rows from a database query into
// a slice of Ventures.
func mapRows(rows *sql.Rows) ([]Venture, error) {
	vens := []Venture{}

	for rows.Next() {
		ven, err := mapRow(rows)
		if err != nil {
			return nil, err
		}
		vens = append(vens, *ven)
	}

	return vens, nil
}

// mapRow is a file private function that maps a single row from a database
// query into a Venture.
func mapRow(rows *sql.Rows) (*Venture, error) {
	ven := Venture{}
	err := rows.Scan(&ven.ID,
		&ven.LastModified,
		&ven.Description,
		&ven.Orders,
		&ven.State,
		&ven.Extra)

	if err != nil {
		return nil, err
	}
	return &ven, err
}
