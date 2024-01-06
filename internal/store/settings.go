package store

import "context"

func (db *Store) Settings(ctx context.Context) (map[string]string, error) {
	rows, errRows := db.QueryBuilder(ctx, db.sb.Select("setting", "value").From("settings"))
	if errRows != nil {
		return nil, errRows
	}

	defer rows.Close()

	settings := make(map[string]string)

	for rows.Next() {
		var (
			key   string
			value string
		)

		if errScan := rows.Scan(&key, &value); errScan != nil {
			return nil, Err(errScan)
		}

		settings[key] = value
	}

	return settings, nil
}

func (db *Store) SettingsSave(ctx context.Context, settings map[string]string) error {
	return nil
}
