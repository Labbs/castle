<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const projects = ref([]);
    const projectName = ref('');
    const message = ref('');
    return {
      projects,
      projectName,
      message,
    };
  },
  mounted() {
    this.getProjects();
  },
  methods: {
    async getProjects() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/all`);
      this.projects = response.projects;
    },
    async createProject() {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/project/create`, {
        name: this.projectName,
      });
      if (response.message === 'Project created') {
        this.getProjects();
      }
    },
    async deleteProject(id) {
      const response = await fetchWrapper.delete(`${import.meta.env.VITE_API_URL}/project/${id}`);
      if (response.message === 'Project deleted') {
        this.getProjects();
      }
    },
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
              <h2 class="page-title">Project List</h2>
            </div>
            <div class="col-12 col-md-auto ms-auto d-print-none">
              <div class="btn-list">
                <a href="#" class="btn btn-primary d-none d-sm-inline-block" data-bs-toggle="modal" data-bs-target="#modal-project">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
                  Create new project
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="card">
            <div class="card-table table-responsive">
              <table class="table card-table table-vcenter text-nowrap datatable">
                <thead>
                  <tr>
                    <th>Name</th>
                    <th></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="project in projects">
                    <td>{{ project.name }}</td>
                    <td>
                      <router-link :to="{name: 'project', params: { id: project.id }}" class="btn btn-primary" style="margin-right: 5px;">
                        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-file-search" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><circle cx="15" cy="15" r="4"></circle><path d="M18.5 18.5l2.5 2.5"></path><path d="M4 6h11"></path><path d="M6 18v-6a2 2 0 0 1 2 -2h6"></path></svg>
                        View
                      </router-link>
                      <a :href="`/project/edit/${project.id}`" class="btn btn-info" style="margin-right: 5px;">
                        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-pencil" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 20h4l10.5 -10.5a2.828 2.828 0 1 0 -4 -4l-10.5 10.5v4"></path><path d="M13.5 6.5l4 4"></path></svg>
                        Edit
                      </a>
                      <a @click="deleteProject(project.id)" class="btn btn-danger" style="margin-right: 5px;">
                        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-trash" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 7l16 0"></path><path d="M10 11l0 6"></path><path d="M14 11l0 6"></path><path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12"></path><path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3"></path></svg>
                        Delete
                      </a>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="modal modal-blur fade" id="modal-project" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-lg" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">New project</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <div class="mb-3">
            <label class="form-label">Name</label>
            <input type="text" class="form-control" name="name" placeholder="Your project name" v-model="projectName" />
          </div>
        </div>
        <div class="modal-footer">
          <button @click="createProject" class="btn btn-primary ms-auto" data-bs-dismiss="modal">
            <!-- Download SVG icon from http://tabler-icons.io/i/plus -->
            <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
            Create new project
          </button>
        </div>
      </div>
    </div>
  </div>
</template>