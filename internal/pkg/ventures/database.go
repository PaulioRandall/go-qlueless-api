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

	err = _create_insert_venture_Trigger(db)
	if err != nil {
		return err
	}

	return nil
}

// _create_venture_Table creates the Venture table within the supplied database.
func _create_venture_Table(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE venture (
		id INTEGER NOT NULL,
		vid INTEGER,
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_alive BOOL DEFAULT TRUE,
		extra TEXT DEFAULT NULL,
		PRIMARY KEY(id, vid)
	);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

// _create_ql_venture_Table creates the query layer Venture table within the
// supplied database.
func _create_ql_venture_Table(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE ql_venture (
		id INTEGER NOT NULL PRIMARY KEY,
		vid INTEGER,
		description TEXT NOT NULL,
		order_ids TEXT NOT NULL,
		state TEXT NOT NULL,
		is_alive BOOL DEFAULT TRUE,
		extra TEXT DEFAULT NULL
	);`)

	if stmt != nil {
		defer stmt.Close()
	}

	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

// _create_insert_venture_Trigger creates a trigger within the supplied
// database that updates the ql_venture table when ever a new record is inserted
// into the venture table.
func _create_insert_venture_Trigger(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TRIGGER new_venture
		AFTER INSERT ON venture
		FOR EACH ROW
		BEGIN
			INSERT OR REPLACE INTO ql_venture (
				id, vid, description, order_ids, state, is_alive, extra
			) VALUES (
				NEW.id, NEW.vid, NEW.description, NEW.order_ids, NEW.state, NEW.is_alive, NEW.extra
			)
			ON CONFLICT (id)
			DO UPDATE SET 
				vid = NEW.vid,
				description = NEW.description,
				order_ids = NEW.order_ids,
				state = NEW.state,
				is_alive = NEW.is_alive,
				extra = NEW.extra;
		END;`)

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
