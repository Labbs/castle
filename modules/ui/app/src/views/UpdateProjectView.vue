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
    const environments = ref([]);
    const variables = ref([]);
    const variableName = ref('');
    const variableValue = ref('');
    const variableType = ref('text');
    const message = ref('');
    const showEditVariable = ref(false);

    return {
      project,
      environments,
      variables,
      variableName,
      variableValue,
      variableType,
      message,
      showEditVariable,
    };
  },
  mounted() {
    this.getProject();
    this.getEnvironments();
    this.getVariables();
  },
  methods: {
    async onSubmit(project) {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/project/update/${this.$route.params.id}`, project);
      if (response.message === 'Project updated') {
        this.$router.push(`/project/${this.$route.params.id}`);
      }
    },
    async getProject() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/${this.$route.params.id}`);
      this.project = response.project;
    },
    async getEnvironments() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/environment/project/${this.$route.params.id}`);
      this.environments = response.environments;
    },
    async getVariables() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/variable/project/${this.$route.params.id}`);
      this.variables = response.variables;
    },
    async createVariable() {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/variable/create`, {
        name: this.variableName,
        value: this.variableValue,
        project_id: this.$route.params.id,
      });
      if (response.message === 'Variable created') {
        this.getVariables();
      }
    },
    async deleteVariable(id) {
      const response = await fetchWrapper.delete(`${import.meta.env.VITE_API_URL}/variable/delete/${id}`);
      if (response.message === 'Variable deleted') {
        this.getVariables();
      }
    },
  },
};
</script>

<template>
  <NavbarView />
  <div class="page-wrapper">
    <div class="page-header d-print-none">
      <div class="container-xl">
        <div class="row g-2 align-items-center">
          <div class="col">
            <h2 class="page-title">Update project: {{ project.name }}</h2>
          </div>
        </div>
      </div>
    </div>
    <div class="page-body">
      <div class="container-xl">
        <div class="row row-deck row-cards">
          <div class="col-lg-4">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Project config</h3>
              </div>
              <div class="card-body">
                <div class="mb-3 row">
                  <label class="col-3 col-form-label required">Name</label>
                  <div class="col">
                    <input type="text" class="form-control" name="name" v-model="project.name" />
                  </div>
                </div>
                <div class="mb-3 row">
                  <label class="col-3 col-form-label">Description</label>
                  <div class="col">
                    <input type="text" class="form-control" name="description" v-model="project.description" />
                  </div>
                </div>
              </div>
              <div class="card-footer">
                <button class="btn btn-primary" @click="onSubmit(project)">Update</button>
              </div>
            </div>
          </div>
          <div class="col-lg-4">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Project variable(s)</h3>
                <div class="card-actions">
                  <a href="#" class="btn btn-primary btn-icon" data-bs-toggle="modal" data-bs-target="#modal-variable">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M12 5l0 14"></path><path d="M5 12l14 0"></path></svg>
                  </a>
                </div>
              </div>
              <div class="card-table table-responsive" v-if="variables.length > 0">
                <table class="table card-table table-vcenter text-nowrap datatable">
                  <thead>
                    <tr>
                      <th>Name</th>
                      <th>Value</th>
                      <th></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="variable in variables">
                      <td>{{ variable.name }}</td>
                      <td>{{ variable.value }}</td>
                      <td>
                        <router-link :to="{ name: 'variable-update', params: { id: variable.id} }">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-pencil" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 20h4l10.5 -10.5a2.828 2.828 0 1 0 -4 -4l-10.5 10.5v4"></path><path d="M13.5 6.5l4 4"></path></svg>
                        </router-link>
                        <a @click="deleteVariable(variable.id)" style="margin-right: 5px;">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-trash" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M4 7l16 0"></path><path d="M10 11l0 6"></path><path d="M14 11l0 6"></path><path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12"></path><path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3"></path></svg>
                        </a>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div class="card-body" v-if="variables.length === 0">
                <div class="card-stamp">
                  <div class="card-stamp-icon bg-yellow">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                  </div>
                </div>
                <div class="row align-items-center">
                  <div class="col-12">
                    <div class="markdown text-secondary">No variable(s) available.</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-lg-4">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Project environment(s)</h3>
                <div class="card-actions">
                  <a href="#" class="btn btn-primary btn-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M12 5l0 14"></path><path d="M5 12l14 0"></path></svg>
                  </a>
                </div>
              </div>
              <div class="card-table table-responsive" v-if="environments.length > 0">
                <table class="table card-table table-vcenter text-nowrap datatable">
                  <thead>
                    <tr>
                      <th>Name</th>
                      <th></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="environment in environments">
                      <td>{{ environment.name }}</td>
                      <td>
                        
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div class="card-body" v-if="environments.length === 0">
                <div class="card-stamp">
                  <div class="card-stamp-icon bg-yellow">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                  </div>
                </div>
                <div class="row align-items-center">
                  <div class="col-12">
                    <div class="markdown text-secondary">No environment(s) available.</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="modal modal-blur fade" id="modal-variable" tabindex="-1" role="dialog" aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">New variable</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label class="form-label">Name</label>
              <input type="text" class="form-control" name="name" placeholder="Variable name" v-model="variableName" />
            </div>
            <div class="mb-3">
              <label class="form-label">Value</label>
              <textarea rows="5" class="form-control" name="value" v-model="variableValue"></textarea>
            </div>
            <div class="mb-3">
              <label class="form-label">Type</label>
              <select class="form-select" v-model="variableType">
                <option>text</option>
                <option>secret</option>
              </select>
            </div>
          </div>
          <div class="modal-footer">
            <button @click="createVariable" class="btn btn-primary ms-auto" data-bs-dismiss="modal">
              <!-- Download SVG icon from http://tabler-icons.io/i/plus -->
              <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
              Create new variable
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>