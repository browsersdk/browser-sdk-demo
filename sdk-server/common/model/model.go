package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
)

type Array []int

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Array) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		slog.Error("json.Unmarshal failed", "err", err)
	}
	return err
}

// Value return json value, implement driver.Valuer interface
func (j Array) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

type Map map[string]any

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Map) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		slog.Error("json.Unmarshal failed", "err", err)
	}
	return err
}

// Value return json value, implement driver.Valuer interface
func (j Map) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}
