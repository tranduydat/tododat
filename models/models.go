package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Task struct {
	ID        int      `orm:"auto"`
	Content   string   `orm:"size(255)"`
	ProjectId *Project `orm:"rel(one)"`
	CreatedAt time.Time
	CreatedBy *User `orm:"rel(fk)"`
}

type TaskList struct {
	TaskList      []*Task
	UserID        int
	NumberOfTasks int
}

type User struct {
	ID        int    `orm:"auto"`
	Username  string `orm:"size(20)"`
	Password  string `orm:"size(50)"`
	Email     string
	CreatedAt time.Time
	RoleId    *Role `orm:"rel(fk)"`
}

type Project struct {
	ID        int    `orm:"auto"`
	Name      string `orm:"size(255)"`
	CreatedAt time.Time
	CreatedBy *User `orm:"rel(fk)"`
}

type Role struct {
	ID   int
	Name string
}

// TableName Renaming tables
func (u *User) TableName() string {
	return "users"
}
func (t *Task) TableName() string {
	return "tasks"
}
func (p *Project) TableName() string {
	return "projects"
}
func (r *Role) TableName() string {
	return "roles"
}

func init() {
	// Register models
	orm.RegisterModel(new(Task), new(User), new(Project), new(Role))

	// Register database connection
	err := orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://hovkhypvxvkvpz:ca4753784710fb030779e419d864058876679046a4f78379a87e2d8857e2fee2@ec2-34-206-8-52.compute-1.amazonaws.com:5432/d930hg51licte7")
	if err != nil {
		fmt.Println("Failed to connect to datasource!")
	}
}

func AddTask(t Task, p Project, u User) {
	o := orm.NewOrm()
	za, err := o.InsertOrUpdate(&u)
	if err != nil {
		log.Fatal(za)
		//log.Fatal("Error when insert/update User in models.AddTask")
	}
	_, err = o.InsertOrUpdate(&p)
	if err != nil {
		log.Fatal("Error when insert/update Project in models.AddTask")
	}
	_, err = o.Insert(&t)
	if err != nil {
		log.Fatal("Error when insert Task in models.AddTask")
	}
}

func NewUser(id int, username string, password string, email string, createdAt time.Time) *User {
	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func NewTaskList() *TaskList {
	return &TaskList{}
}
