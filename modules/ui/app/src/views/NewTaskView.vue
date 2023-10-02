<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const message = ref('');
    const repositories = ref([]);
    const project = ref({});
    const environments = ref([]);
    const taskType = ref('');
    const task = {
      name: '',
      description: '',
      project_id: '',
      environment_id: '',
      repository_id: '',
      repository_path: '',
      enabled: true,
      cron: '',
      task_type: '',
      ansible_task: {
        dry_run: false,
        debug: false,
        verbose: false,
        playbook_path: '',
        inventory_path: '',
      },
      terraform_task: {
        workspace: '',
        plan_only: false,
      },
    }

    return {
      task,
      repositories,
      environments,
      project,
      taskType,
      message
    };
  },
  mounted() {
    this.getAllRepositories();
    this.getProject();
    this.getAllEnvironments();
  },
  methods: {
    async getAllRepositories() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/repository/all`);
      this.repositories = response.repositories;
    },
    async onSubmit(task) {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/task/create`, task);
      if (response.message === 'Task created') {
        this.$router.push('/project/'+this.$route.params.id);
      }
    },
    async getProject() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/${this.$route.params.id}`);
      this.project = response.project;
      this.task.project_id = this.project.id;
    },
    async getAllEnvironments() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/environment/project/${this.$route.params.id}`);
      this.environments = response.environments;
    },
    onChangeType() {
      this.taskType = this.task.task_type;
    }
  },
};
</script>

<template>
  <div class="page">
    <NavbarView />
    <div class="page-wrapper">
      <div class="page-header d-print-none">
        <div class="container-xl">
          <div class="row g-2 align-items-center">
            <div class="col">
              <h2 class="page-title">Create new task</h2>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="row row-deck row-cards">
            <div class="col-lg-12">
              <div class="card">
                <div class="card-body">
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">Name</label>
                    <div class="col">
                      <input type="text" class="form-control" name="name" v-model="task.name" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Description</label>
                    <div class="col">
                      <input type="text" class="form-control" name="description" v-model="task.description" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Project name</label>
                    <div class="col">
                      <input type="text" class="form-control" name="project name" v-model="project.name" disabled />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Environment</label>
                    <div class="col">
                      <select class="form-select" v-model="task.environment_id">
                        <option v-for="environment in environments" v-bind:value="environment.id">
                          {{ environment.name }}
                        </option>
                      </select>
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">Cron</label>
                    <div class="col">
                      <input type="text" class="form-control" name="cron" placeholder="like */2 * * * *" v-model="task.cron" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Enabled</label>
                    <div class="col">
                      <input type="checkbox" class="form-check-input" name="enabled" v-model="task.enabled" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">Repository</label>
                    <div class="col">
                      <select class="form-select" v-model="task.repository_id">
                        <option v-for="repository in repositories" v-bind:value="repository.id">
                          {{ repository.name }}
                        </option>
                      </select>
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">Repository path</label>
                    <div class="col">
                      <input type="text" class="form-control" name="repository_path" v-model="task.repository_path" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Task type</label>
                    <div class="col">
                      <select class="form-select" v-model="task.task_type" @change="onChangeType">
                        <option value="ansible">Ansible</option>
                        <option value="terraform">Terraform</option>
                      </select>
                    </div>
                  </div>
                  <div v-if="taskType === 'ansible'">
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Dry run</label>
                      <div class="col">
                        <input type="checkbox" class="form-check-input" name="dry_run" v-model="task.ansible_task.dry_run" />
                      </div>
                    </div>
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Debug</label>
                      <div class="col">
                        <input type="checkbox" class="form-check-input" name="debug" v-model="task.ansible_task.debug" />
                      </div>
                    </div>
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Verbose</label>
                      <div class="col">
                        <input type="checkbox" class="form-check-input" name="verbose" v-model="task.ansible_task.verbose" />
                      </div>
                    </div>
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Playbook file</label>
                      <div class="col">
                        <input type="text" class="form-control" name="playbook_path" v-model="task.ansible_task.playbook_path" />
                      </div>
                    </div>
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Inventory file</label>
                      <div class="col">
                        <input type="text" class="form-control" name="inventory_path" v-model="task.ansible_task.inventory_path" />
                      </div>
                    </div>
                  </div>
                  <div v-if="taskType === 'terraform'">
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Workspace</label>
                      <div class="col">
                        <input type="text" class="form-control" name="workspace" v-model="task.terraform_task.workspace" />
                      </div>
                    </div>
                    <div class="mb-3 row">
                      <label class="col-3 col-form-label">Plan only</label>
                      <div class="col">
                        <input type="checkbox" class="form-check-input" name="plan_only" v-model="task.terraform_task.plan_only" />
                      </div>
                    </div>
                  </div>
                </div>
                <div class="card-footer text-end">
                  <button class="btn btn-primary" @click="onSubmit(task)">Save</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>