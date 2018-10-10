package src

import (
	"encoding/json"
	"strconv"
	"time"
)

func Create(cmd Cmd) Respond {
	result := Respond{}
	cmdJson, _ := json.Marshal(cmd)
	body := HttpJson("POST", "/cmdp/create", cmdJson)
	json.Unmarshal(body, &result)
	return result
}

func Search(keyword string, page int, size int) CmdsRespond {
	result := CmdsRespond{}
	body := Http("POST", "/cmdp/search", "keyword="+keyword+"&page="+strconv.Itoa(page)+"&size="+strconv.Itoa(size))
	json.Unmarshal(body, &result)
	return result
}

func GetCmd(keyword string) CmdRespond {
	result := CmdRespond{}
	body := Http("POST", "/cmdp/getCmd", "keyword="+keyword)
	json.Unmarshal(body, &result)
	return result
}

func Delete(id string) Respond {
	result := Respond{}
	body := Http("GET", "/cmdp/delete/"+id, "")
	json.Unmarshal(body, &result)
	return result
}

// file

func CreateFile(file *File) Respond {
	result := Respond{}
	fileJson, _ := json.Marshal(file)
	body := HttpJson("POST", "/file/create", fileJson)
	json.Unmarshal(body, &result)
	return result
}

func SearchFile(keyword string, page int, size int) FilesRespond {
	result := FilesRespond{}
	body := Http("POST", "/file/search", "keyword="+keyword+"&page="+strconv.Itoa(page)+"&size="+strconv.Itoa(size))
	json.Unmarshal(body, &result)
	return result
}

func DownloadFile(keyword string) FilesDownloadRespond {
	result := FilesDownloadRespond{}
	body := Http("POST", "/file/download", "keyword="+keyword)
	json.Unmarshal(body, &result)
	return result
}

func DeleteFile(id string) Respond {
	result := Respond{}
	body := Http("GET", "/file/delete/"+id, "")
	json.Unmarshal(body, &result)
	return result
}

// user

func Login(username string, password string) Respond {
	result := Respond{}
	body := Http("POST", "/user/login", "Username="+username+"&Password="+password)
	json.Unmarshal(body, &result)
	return result
}

func Register(username string, password string) Respond {
	result := Respond{}
	body := Http("POST", "/user/register", "Username="+username+"&Password="+password)
	json.Unmarshal(body, &result)
	return result
}

func ResetPassword(password string) Respond {
	result := Respond{}
	body := Http("POST", "/user/reset", "password="+password)
	json.Unmarshal(body, &result)
	return result
}

func UpdateInfo(info string) Respond {
	result := Respond{}
	body := Http("POST", "/user/info", "info="+info)
	json.Unmarshal(body, &result)
	return result
}

// star

type StarUser struct {
	StarId    int
	UserId    int
	Username  string    `orm:"size(255)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	User      User
}

func CreateStar(username string) Respond {
	result := Respond{}
	body := Http("POST", "/star/create", "Username="+username)
	json.Unmarshal(body, &result)
	return result
}

func DeleteStar(id int) CmdsRespond {
	result := CmdsRespond{}
	body := Http("GET", "/star/delete/"+strconv.Itoa(id), "")
	json.Unmarshal(body, &result)
	return result
}

func SearchStar(page int, size int) StarsRespond {
	result := StarsRespond{}
	body := Http("POST", "/star/search", "page="+strconv.Itoa(page)+"&size="+strconv.Itoa(size))
	json.Unmarshal(body, &result)
	return result
}
