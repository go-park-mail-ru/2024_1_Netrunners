package constants

import "os"

func InitUploads() error {
	storagePath := "./uploads/"
	_, err := os.Stat(storagePath)
	if err != nil {
		err = os.Mkdir(storagePath, 0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(storagePath+"users/", 0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(storagePath+"films/", 0755)
		if err != nil {
			return err
		}
	} else {
		storagePath = "./uploads/users/"
		_, err = os.Stat(storagePath)
		if err != nil {
			err = os.Mkdir(storagePath, 0755)
			if err != nil {
				return err
			}
		}

		storagePath := "./uploads/films/"
		_, err = os.Stat(storagePath)
		if err != nil {
			err = os.Mkdir(storagePath, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
