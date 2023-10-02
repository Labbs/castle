<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const project = ref({});
    const tasks = ref([]);
    const histories = ref([]);
    const message = ref('');

    return {
      project,
      tasks,
      histories,
      message,
    };
  },
  mounted() {
    this.getProject();
    this.getTasks();
  },

  methods: {
    async getProject() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/${this.$route.params.id}`);
      this.project = response.project;
    },

    async getTasks() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/task/project/${this.$route.params.id}`);
      this.tasks = response.tasks;
    },

    async createTask() {

    },

    async getTaskHistories() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/${this.$route.params.id}/tasks`);
      this.histories = response.histories;
    },
  }
}
</script>

<template>
  <div class="page">
    <NavbarView />
    <div class="page-wrapper">
      <div class="page-header d-print-none">
        <div class="container-xl">
          <div class="row g-2 align-items-center">
            <div class="col">
              <h2 class="page-title">Project: {{ project.name }}</h2>
            </div>
            <div class="col-12 col-md-auto ms-auto d-print-none">
              <div class="btn-list">
                <a :href="`/project/edit/${project.id}`" class="btn btn-primary d-none d-sm-inline-block">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-pencil" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 20h4l10.5 -10.5a2.828 2.828 0 1 0 -4 -4l-10.5 10.5v4"></path><path d="M13.5 6.5l4 4"></path></svg>
                  Update project
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="row row-deck row-cards">
            <div class="col-lg-8">
              <div class="card">
                <div class="card-header">
                  <h3 class="card-title">Tasks</h3>
                  <div class="card-actions">
                    <a :href="`/project/${project.id}/new/task`" class="btn btn-primary">
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M12 5l0 14"></path><path d="M5 12l14 0"></path></svg>
                      Add task
                    </a>
                  </div>
                </div>
                <div class="card-table table-responsive" v-if="tasks.length > 0">
                  <table class="table card-table table-vcenter text-nowrap datatable">
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Enabled</th>
                        <th>Environment</th>
                        <th>Type</th>
                        <th>Action</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="task in tasks" :key="task.id">
                        <td>{{ task.name }}</td>
                        <td>{{ task.enabled }}</td>
                        <td>{{ task.environment }}</td>
                        <td>{{ task.task_type }}</td>
                        <td>
                          <router-link :to="{ name: 'repository-update', params: { id: repository.id } }">
                            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-pencil" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 20h4l10.5 -10.5a2.828 2.828 0 1 0 -4 -4l-10.5 10.5v4"></path><path d="M13.5 6.5l4 4"></path></svg>
                          </router-link>
                          <a @click="deleteRepository(repository.id)">
                            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-trash" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><line x1="4" y1="7" x2="20" y2="7"></line><line x1="10" y1="11" x2="10" y2="17"></line><line x1="14" y1="11" x2="14" y2="17"></line><path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12"></path><path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3"></path></svg> 
                          </a>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div class="card-body" v-if="tasks.length === 0">
                  <div class="card-stamp">
                    <div class="card-stamp-icon bg-yellow">
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                    </div>
                  </div>
                  <div class="row align-items-center">
                    <div class="col-12">
                      <div class="markdown text-secondary">No tasks available.</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-lg-4">
              <div class="card">
                <div class="card-header">Tasks history</div>
                <div class="card-table table-responsive" v-if="histories.length > 0">
                  <table class="table card-table table-vcenter datatable">
                    <thead>
                      <tr>
                        <th>Date</th>
                        <th>Name</th>
                        <th>Status</th>
                        <th></th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="history in histories" :key="history.id">
                        <td>{{ history.created_at }}</td>
                        <td>{{ history.task.name }}</td>
                        <td>{{ history.status }}</td>
                        <td></td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div class="card-body" v-if="histories.length === 0">
                  <div class="card-stamp">
                    <div class="card-stamp-icon bg-yellow">
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                    </div>
                  </div>
                  <div class="row align-items-center">
                    <div class="col-12">
                      <div class="markdown text-secondary">No tasks histories available.</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>