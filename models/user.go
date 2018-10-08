package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DefaultUserList *UserList

type User struct {
	Id int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:secret@/TestGO", 30)
}

// TaskManager manages a list of tasks in memory.
type UserList struct {
	users  []*User
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewUserList() *UserList {
	return &UserList{}
}


// All returns the list of all the Tasks in the TaskManager.
func (m *UserList) All() []*User {
	return m.users
}


//func main (){
//	o := orm.NewOrm()
//	user := User{Id: 1}
//
//	err := o.Read(&user)
//
//	if err == orm.ErrNoRows {
//		fmt.Println("No result found.")
//	} else if err == orm.ErrMissPK {
//		fmt.Println("No primary key found.")
//	} else {
//		fmt.Println(user.Id, user.Name)
//	}
//}

