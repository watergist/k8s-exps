package wide

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func DrawPodProperties(propertiesPath string) (map[string]string, error) {
	fileBytes, err := os.Open(propertiesPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cfg := map[string]string{}
	fileReader := bufio.NewReader(fileBytes)
	for true {
		var line string
		line, err = fileReader.ReadString('\n')
		err2 := updateKeyValue(&cfg, &line)
		if err2 != nil {
			return nil, err2
		}
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return cfg, err
			}
			break
		}
	}
	return cfg, nil
}

func updateKeyValue(hm *map[string]string, kvLine *string) error {
	if kvSlice := strings.Split(*kvLine, "="); len(kvSlice) != 2 {
		err := errors.New(*kvLine + " is not decodable")
		log.Println(err)
		return err
	} else if key := strings.TrimSpace(kvSlice[0]); len(key) == 0 {
		err := errors.New(*kvLine + " is not decodable")
		log.Println(err)
		return err
	} else {
		value := strings.TrimSpace(kvSlice[1])
		if len(value) < 2 {
			value = "<empty>"
		}
		(*hm)[key] = strings.Trim(value, "\"")
	}
	return nil
}
