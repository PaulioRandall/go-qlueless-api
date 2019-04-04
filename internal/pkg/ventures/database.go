package ventures

import (
	"database/sql"
)

// CreateTables creates all the Venture tables, views and triggers within the
// supplied database.
//
// @UNTESTED
func CreateTables(db *sql.DB) error {
	err := _create_venture_Table(db)
	if err != nil {
		return err
	}

	err = _create_ql_venture_Table(db)
	if err != nil {
		return err
	}

	err = _create_insert_on_living_venture_Trigger(db)
	if err != nil {
		return err
	}

	err = _create_insert_on_dead_venture_Trigger(db)
	if err != nil {
		return err
	}

	err = _create_update_on_venture_Trigger(db)
	if err != nil {
		return err
	}

	err = _create_delete_on_venture_Trigger(db)
	if err != nil {
		return err
	}

	return nil
}

// _create_venture_Table creates the Venture table within the supplied database.
func _create_venture_Table(db *sql.DB) error {
	return _execStmt(db, `CREATE TABLE venture (
		id INTEGER NOT NULL,
		vid INTEGER NOT NULL DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_alive BOOL NOT NULL DEFAULT TRUE,
		extra TEXT NOT NULL DEFAULT "",
		PRIMARY KEY(id, vid)
	);`)
}

// _create_ql_venture_Table creates the query layer Venture table within the
// supplied database.
func _create_ql_venture_Table(db *sql.DB) error {
	return _execStmt(db, `CREATE TABLE ql_venture (
		id INTEGER NOT NULL PRIMARY KEY,
		vid INTEGER NOT NULL,
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_alive BOOL NOT NULL,
		extra TEXT NOT NULL
	);`)
}

// _create_insert_on_living_venture_Trigger creates a trigger within the
// supplied database that updates the ql_venture table when ever a new, and
// living, Venture is inserted into the venture table.
func _create_insert_on_living_venture_Trigger(db *sql.DB) error {
	return _execStmt(db, `CREATE TRIGGER insert_on_living_venture
		AFTER INSERT ON venture
		FOR EACH ROW
		WHEN (NEW.is_alive = true)
		BEGIN
			REPLACE INTO ql_venture (
				id, vid, description, order_ids, state, is_alive, extra
			) VALUES (
				NEW.id, NEW.vid, NEW.description, NEW.order_ids, NEW.state, NEW.is_alive, NEW.extra
			);
		END;`)
}

// _create_insert_on_dead_venture_Trigger creates a trigger within the supplied
// database that removes from the ql_venture table the dead Venture inserted
// into the venture table.
func _create_insert_on_dead_venture_Trigger(db *sql.DB) error {
	return _execStmt(db, `CREATE TRIGGER insert_on_dead_venture
		AFTER INSERT ON venture
		FOR EACH ROW
		WHEN (NEW.is_alive = false)
		BEGIN
			DELETE FROM ql_venture 
			WHERE id = NEW.id;
		END;`)
}

// _create_update_on_venture_Trigger creates a trigger within the supplied
// database that raises an error if an update is attempted.
func _create_update_on_venture_Trigger(db *sql.DB) error {
	return _execStmt(db, `CREATE TRIGGER update_on_venture
		BEFORE UPDATE ON venture
		BEGIN
			SELECT RAISE(FAIL, "Updates not allowed, insert with the same Venture ID!");
		END;`)
}

// _create_delete_on_venture_Trigger creates a trigger within the supplied
// database that raises an error if a delete is attempted.
func _create_delete_on_venture_Trigger(db *sql.DB) error {
	return _execStmt(db, `CREATE TRIGGER delete_on_venture
		BEFORE DELETE ON venture
		BEGIN
			SELECT RAISE(FAIL, "Deletions not allowed!");
		END;`)
}

// _execStmt executes a SQL statment ensuring it is closed afterwards
func _execStmt(db *sql.DB, sql string) error {
	stmt, err := db.Prepare(sql)

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
//
// @UNTESTED
func QueryFor(db *sql.DB, id string) (*Venture, error) {
	ven := Venture{}
	err := db.QueryRow(`SELECT
		id,
		description,
		order_ids,
		state,
		is_alive,
		extra
	FROM ql_venture
	WHERE id = ?`, id).Scan(&ven.ID,
		&ven.Description,
		&ven.OrderIDs,
		&ven.State,
		&ven.IsAlive,
		&ven.Extra)

	if err != nil {
		return nil, err
	}

	return &ven, nil
}

// QueryAll queries the database for all Ventures.
//
// @UNTESTED
func QueryAll(db *sql.DB) ([]Venture, error) {
	rows, err := db.Query(`SELECT
		id,
		description,
		order_ids,
		state,
		is_alive,
		extra
	FROM ql_venture`)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return nil, err
	}

	return _mapRows(rows)
}

// _mapRows is a file private function that maps rows from a database query into
// a slice of Ventures.
func _mapRows(rows *sql.Rows) ([]Venture, error) {
	vens := []Venture{}

	for rows.Next() {
		ven, err := _mapRow(rows)
		if err != nil {
			return nil, err
		}
		vens = append(vens, *ven)
	}

	return vens, nil
}

// _mapRow is a file private function that maps a single row from a database
// query into a Venture.
func _mapRow(rows *sql.Rows) (*Venture, error) {
	ven := Venture{}
	err := rows.Scan(&ven.ID,
		&ven.Description,
		&ven.OrderIDs,
		&ven.State,
		&ven.IsAlive,
		&ven.Extra)

	if err != nil {
		return nil, err
	}
	return &ven, err
}
