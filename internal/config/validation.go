package config

import (
	"fmt"
	"strconv"
	"strings"
)

func validatePort(port string) error {
	port = strings.TrimSpace(port)

	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("port must be a number: %w", err)
	}

	if p < 1 || p > 65535 {
		return fmt.Errorf("port must be between 1 and 65535, got %d", p)
	}

	return nil
}
