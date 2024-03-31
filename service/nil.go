package service

// nil hanya bisa digunakan untuk beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
