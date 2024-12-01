package repositories

import (
	"fmt"
	"strings"

	"client/interfaces"
)

type NodesRepository struct {
	storage interfaces.Storage
}

func (n *NodesRepository) Selected() (*[]string, error) {
	sql := `
		SELECT name FROM Nodes
	`

	query, err := n.storage.Query(&sql, &[]interface{}{})
	if err != nil {
		return nil, err
	}

	var nodes []string

	for query.Next() {
		var node string
		err := query.Scan(&node)

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return &nodes, nil
}

func (n *NodesRepository) Select(node *string) {
	sql := `
		INSERT INTO Nodes (name) VALUES ($1)
	`

	n.storage.Exec(&sql, &[]interface{}{
		&node,
	})
}

func (n *NodesRepository) RemoveAllExcept(nodes []string) error {
	if len(nodes) == 0 {
		sql := `
			DELETE FROM Nodes
		`
		err := n.storage.Exec(&sql, &[]interface{}{})
		return err
	}

	placeholders := make([]string, len(nodes))
	values := make([]interface{}, len(nodes))
	for i, node := range nodes {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		values[i] = node
	}

	sql := fmt.Sprintf(`
		DELETE FROM Nodes
		WHERE name NOT IN (%s)
	`, strings.Join(placeholders, ", "))

	err := n.storage.Exec(&sql, &values)
	return err
}

func (n *NodesRepository) Init(storage interfaces.Storage) {
	n.storage = storage
}
