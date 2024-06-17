package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Address struct {
	City    string
	Pincode string
	State   string
}

type User struct {
	Name    string
	Age     int64
	Contact string
	Company string
	Address Address
}

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Info(string, ...interface{})
		Warn(string, ...interface{})
		Debug(string, ...interface{})
	}

	DatabaseDriver struct {
		mutex sync.Mutex
		//map of mutextes add later
		mutexes map[string]*sync.Mutex
		log     Logger
		dir     string
	}

	Options struct {
		log Logger
	}
)

func NewDatabase(dir string, option *Options) (*DatabaseDriver, error) {
	dir = filepath.Clean(dir)

	//to plug in option for later logging purposes

	opts := Options{}
	if option != nil {
		opts = *option
	}
	//create the options logger

	driver := DatabaseDriver{
		dir:     dir,
		log:     opts.log,
		mutexes: make(map[string]*sync.Mutex),
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return &driver, err
		} else {
			opts.log.Debug("Database already exists", err)
		}
	}

	return &driver, nil
}

/* func (d *DatabaseDriver) Read() error {
	// these are the struct menthods that will be called using driver
} */

func (d *DatabaseDriver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("collection is required")
	}
	if resource == "" {
		return fmt.Errorf("resource is required")
	}
	if v == nil {
		return fmt.Errorf("object is required")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)

	fnlPath := filepath.Join(dir, resource+".json")

	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	b = append(b, byte('\n'))
	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return nil

}

func (d *DatabaseDriver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

func (d *DatabaseDriver) ReadAll(collection string) ([]string, error) {

	//handle if collection is empty
	if collection == "" {
		return nil, fmt.Errorf("collections is empty")
	}

	//read the collection

	dir := filepath.Join(d.dir, collection)

	files, err := os.ReadDir(dir)

	if err != nil {
		return nil, fmt.Errorf("unable to read the directory: %v", err)
	}

	//iterate over the files from the collection

	var records []string
	for _, file := range files {

		if file.IsDir() {
			continue
		}

		//read the file
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))

		if err != nil {
			return nil, fmt.Errorf("unable to read the file: %v", err)
		}

		//unmarshall the file into User
		var user User
		if err := json.Unmarshal(b, &user); err != nil {
			return nil, fmt.Errorf("unable to unmarshall the file: %v", err)
		}

		records = append(records, string(b))

	}

	return records, nil

}

func (d *DatabaseDriver) Delete(collection, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := os.Stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named ")
	case os.IsNotExist(err):
		return nil
	case err != nil:
		return err
	default:
		if err := os.Remove(dir); err != nil {
			return err
		}
	}

	return nil

}

func main() {
	fmt.Println("Starting go documentg database...")
	dir := "./"

	db, err := NewDatabase(dir, nil)
	if err != nil {
		fmt.Println("Error while connecting to database")
		return
	}

	employees := []User{
		{"Kedar", 20, "123333", "Alphabets", Address{"Pune", "123", "Mh"}},
		{"John", 40, "367", "Tesla", Address{"June", "908", "Mh"}},
		{"Cena", 30, "334", "Text", Address{"Dist", "900", "Mh"}},
		{"BigMan", 10, "100", "Anc", Address{"Bhn", "190", "Mh"}},
		{"Wold", 90, "23", "abc", Address{"aji", "123", "Mh"}},
		{"Would", 40, "222", "eee", Address{"ddd", "123", "Mh"}},
	}

	//write

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	//readAll

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error while reading from database")
		return
	}
	fmt.Println(records) //prints records in json format

	//unmrshall to get the records in golang struct format

	allusers := []User{}

	for _, f := range records {
		var employeeFound User
		err := json.Unmarshal([]byte(f), &employeeFound)
		if err != nil {
			fmt.Println("Error while unmarshalling")
			return
		}
		allusers = append(allusers, employeeFound)
	}

	//print all users
	fmt.Println(allusers)

	//dlelete the records
	/* if err := db.Delete("users", "John"); err != nil {
		fmt.Println("Error", err)
	} */

}
