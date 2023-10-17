package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/domain"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
)

type TaskController struct {
	TaskRepository        repository.TaskRepository
	ProjectRepository     repository.ProjectRepository
	EnvironmentRepository repository.EnvironmentRepository
	RespositoryRepository repository.RepositoryRepository
	Logger                zerolog.Logger
}

func (hc *TaskController) Create(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Projects list"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	if c.Method() == "POST" {
		var task domain.BusTaskResponse
		task.Name = c.FormValue("name")
		task.Description = c.FormValue("description")
		task.ProjectId = c.Params("id")
		task.RepositoryId = c.FormValue("repository_id")
		task.RepositoryPath = c.FormValue("repository_path")
		task.Enabled = c.FormValue("enabled") == "on"

		task.TaskType = c.FormValue("task_type")
		if task.TaskType == "cron" {
			task.Cron = c.FormValue("cron")
		}

		task.ExecType = c.FormValue("exec_type")
		if task.ExecType == "ansible" {
			ansible := domain.AnsibleTask{}
			ansible.DryRun = c.FormValue("ansible_dry_run") == "on"
			ansible.Debug = c.FormValue("ansible_debug") == "on"
			ansible.Verbose = c.FormValue("ansible_verbose") == "on"
			ansible.PlaybookPath = c.FormValue("ansible_playbook_path")
			ansible.InventoryPath = c.FormValue("ansible_inventory_path")
			task.AnsibleTask = ansible
		} else if task.ExecType == "terraform" {
			terraform := domain.TerraformTask{}
			terraform.PlanOnly = c.FormValue("terraform_plan_only") == "on"
			terraform.Workspace = c.FormValue("terraform_workspace")
			task.TerraformTask = terraform
		}

		err := hc.TaskRepository.CreateTask(task)
		if err != nil {
			hc.Logger.Error().Err(err).Str("event", "frontend.task.create").Msg("failed to create task")
			d["Error"] = "Internal server error - failed to create task"
			return c.Render("templates/task_new_update", d, "templates/layouts/main")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "success-flash",
			Value: "Task created successfully",
		})

		return c.Redirect("/app/project/" + task.ProjectId + "/view")
	}

	project, err := hc.ProjectRepository.GetProjectById(c.Params("id"))
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.task.create").Msg("failed to get project")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get project",
		})
		return c.Redirect("/app/project")
	}

	repositories, err := hc.RespositoryRepository.GetAllRepositories()
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.task.create").Msg("failed to get repositories")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get repositories",
		})
		return c.Redirect("/app/project")
	}

	environments, err := hc.EnvironmentRepository.GetAllEnvironmentsByProjectId(c.Params("id"))
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.task.create").Msg("failed to get environments")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get environments",
		})
	}

	d["Project"] = project
	d["Repositories"] = repositories
	d["Environments"] = environments

	return c.Render("templates/task_new_update", d, "templates/layouts/main")
}
