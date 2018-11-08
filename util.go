package gedis

import (
	"fmt"
	"bufio"
	"strconv"
	"errors"
)

func send(writer *bufio.Writer, messages ...string) error {
	cmd := fmt.Sprintf("*%d\r\n", len(messages))
	for _, message := range messages {
		cmd += fmt.Sprintf("$%d\r\n", len(message))
		cmd += fmt.Sprintf("%s\r\n", message)
	}

	_, err := writer.Write([]byte(cmd))
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func recv(reader *bufio.Reader) (string, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	message := string(line)

	switch message[0] {
	case '+':
		return message[1:], nil
	case '-':
		return fmt.Sprintf("(error) %s", message[1:]), nil
	case ':':
		return fmt.Sprintf("(integer) %s", message[1:]), nil
	case '$':
		line, _, err = reader.ReadLine()
		if err != nil {
			return "", err
		}
		message = string(line)
		return fmt.Sprintf("\"%s\"", message), nil
	case '*':
		t, err := strconv.Atoi(message[1:])
		if err != nil {
			return "", err
		}

		index, cmd := 0, ""
		for index < t {
			message, err = recv(reader)
			if err != nil {
				return "", nil
			}
			cmd += fmt.Sprintf("%d) %s\n", index + 1, message)
			index += 1
		}

		return cmd, nil
	default:
		return "", errors.New(fmt.Sprintf("%b is unknown", message[0]))
	}
}