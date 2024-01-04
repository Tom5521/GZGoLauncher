package files

import "os"

func Copy(src, dest string) error {
	stats, err := os.Stat(src)
	if err != nil {
		return err
	}
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, stats.Mode().Perm())
	if err != nil {
		return err
	}
	return nil
}

func Move(src, dest string) error {
	err := Copy(src, dest)
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}
