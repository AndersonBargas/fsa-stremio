package utils

import (
	"fmt"
	"net"
	"strconv"
)

func ValidarPortaTCP(portaStr string) bool {
	porta, err := strconv.Atoi(portaStr)
	if err != nil {
		return false // Não é um número inteiro
	}

	if porta < 0 || porta > 65535 {
		return false // Fora do intervalo válido
	}

	// Tentar conectar à porta para validar se ela está disponível
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", porta))
	if err != nil {
		return false // Porta não disponível ou outro erro de conexão
	}
	defer conn.Close()
	return true // Porta válida e disponível
}
