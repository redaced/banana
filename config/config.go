package config
import(
	"os"
	"fmt"
	// "log"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
    Port string
	Db_name string
	Db_user string
	Db_password string
}

func SetConfig() Config{
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
	 var jsontype Config  
	 json.Unmarshal(file, &jsontype)
  	 return jsontype
}