package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Task struct {
	// Beego ORM default identify Id field as auto increment
	Id        int
	Content   string
	ProjectId int
	CreatedAt time.Time
	CreatedBy int
}

type User struct {
	Id        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	RoleId    int
}

type Project struct {
	Id        int
	Name      string
	CreatedAt time.Time
	CreatedBy int
}

type Role struct {
	Id   int
	Name string
}

type TaskList struct {
	TaskList      []*Task
	UserID        int
	NumberOfTasks int
}

type RoleCode int
const (
	ADMIN = 0
	USER = 1
)

func init() {
	// Register database connection
	err := orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://hovkhypvxvkvpz:ca4753784710fb030779e419d864058876679046a4f78379a87e2d8857e2fee2@ec2-34-206-8-52.compute-1.amazonaws.com:5432/d930hg51licte7")
	if err != nil {
		fmt.Println("Failed to connect to datasource!")
	}
}

// ReadTaskById Read a task by its id
func ReadTaskById(id int) Task {
	o := orm.NewOrm()
	var task Task
	err := o.Raw("SELECT id, content, project_id, created_at, created_by " +
		"FROM public.tasks WHERE id = ?", id).QueryRow(&task)
	if err != nil {
		log.Fatalln("Error in models.ReadTaskById()")
	}
	return task
}

// ReadAllTasks Read all exists tasks in database from oldest to newest
func ReadAllTasks() []Task {
	o := orm.NewOrm()
	var arrayTasks []Task
	// Build query
	_, err := o.Raw("SELECT id, content, project_id, created_at, created_by " +
		"FROM public.tasks " +
		"ORDER BY created_at ASC").QueryRows(&arrayTasks)
	if err != nil {
		log.Fatal("Error in models.ReadAllTasks()")
	}
	return arrayTasks
}

// ReadTasksFromUser Read all tasks from a specific user
func ReadTasksFromUser(u User) []Task {
	o := orm.NewOrm()
	var arrayTasks []Task
	// Build query
	_, err := o.Raw("SELECT id, content, project_id, created_at, created_by " +
		"FROM public.tasks WHERE created_by = ?" +
		"ORDER BY created_at ASC", u.Id).QueryRows(&arrayTasks)
	if err != nil {
		log.Fatal("Error in models.ReadAllTasks()")
	}
	return arrayTasks
}

// CreateTask Create new task into database
func CreateTask(t Task) bool {
	o := orm.NewOrm()
	// Build query
	_, err := o.Raw("INSERT INTO public.tasks (content, project_id, created_at, created_by) "+
		"VALUES (?, ?, ?, ?)", t.Content, t.ProjectId, time.RFC3339, t.CreatedBy).Exec()
	if err != nil {
		log.Fatalln("Error in models.CreateTask()")
		return false
	}
	return true
}

// UpdateTask Update an existing task
func UpdateTask(t Task) bool {
	o := orm.NewOrm()
	// Build query
	_, err := o.Raw("UPDATE public.tasks "+
		"SET content = ?, project_id = ?, created_at = ?, created_by = ? "+
		"WHERE id = ?",
		t.Content, t.ProjectId, t.CreatedAt, t.CreatedBy, t.Id).Exec()
	if err != nil {
		log.Fatalln("Error in models.UpdateTask()")
		return false
	}
	return true
}

// DeleteTask Delete an existing task
func DeleteTask(t Task) bool {
	o := orm.NewOrm()
	// Build query
	_, err := o.Raw("DELETE FROM public.tasks WHERE id = ?", t.Id).Exec()
	if err != nil {
		log.Fatalln("Error in models.DeleteTask()")
		return false
	}
	return true
}

// ReadAllProjects Read all projects
func ReadAllProjects() []Project {
	o := orm.NewOrm()
	var arrayProjects []Project
	_, err := o.Raw("SELECT * FROM public.projects").QueryRows(&arrayProjects)
	if err != nil {
		log.Fatalln("Error in models.ReadAllProjects()")
	}
	return arrayProjects
}

// ReadProjectsByUserID Read all the projects owned by a specific user
func ReadProjectsByUserID(uid int) []Project {
	o := orm.NewOrm()
	var arrayProjects []Project
	_, err := o.Raw("SELECT * FROM public.projects WHERE created_by = ?", uid).QueryRows(&arrayProjects)
	if err != nil {
		log.Fatalln("Error in models.ReadAllProjects()")
	}
	return arrayProjects
}

// ReadProjectById Read a project with its specific id
func ReadProjectById(pid int) Project {
	o := orm.NewOrm()
	var project Project
	err := o.Raw("SELECT * FROM public.projects WHERE id = ?", pid).QueryRow(&project)
	if err != nil {
		log.Fatalln("Error in models.ReadProjectById()")
	}
	return project
}

// UpdateProject Update a project with its specific id
func UpdateProject(p Project) bool {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE public.projects SET name = ?, created_at = ?, created_by = ? WHERE id = ?",
		p.Name, p.CreatedAt, p.CreatedBy, p.Id).Exec()
	if err != nil {
		log.Fatalln("Error in models.UpdateProject()")
		return false
	}
	return true
}

// DeleteProjectById Delete a specific project from its id
func DeleteProjectById(id int) bool {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM public.projects WHERE id = ?", id).Exec()
	if err != nil {
		log.Fatalln("Error in models.DeleteProjectById()")
		return false
	}
	return true
}

// CreateProject Create a new project
func CreateProject(p Project) bool {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO public.projects (name, created_at, created_by) VALUES (?, ?, ?)",
		p.Name, p.CreatedAt, p.CreatedBy).Exec()
	if err != nil {
		log.Fatalln("Error in models.CreateProject()")
		return false
	}
	return true
}