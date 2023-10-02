<script>
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";
import NavbarView from '@/views/NavbarView.vue';

export default {
  components: {
    NavbarView,
  },
  setup() {
    const projects = ref([]);
    const histories = ref([]);
    return {
      projects,
      histories,
    };
  },
  mounted() {
    this.getProjects();
    this.getLatestTaskHistories();
  },
  methods: {
    async getProjects() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/project/all`);
      this.projects = response.projects;
    },

    async getLatestTaskHistories() {

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
              <div class="page-pretitle"> Overview </div>
              <h2 class="page-title">Home Castle</h2>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="row row-deck row-cards">
            <div class="col-lg-6">
              <div class="card">
                <div class="card-header">Projects</div>
                <div class="card-table table-responsive" v-if="projects.length > 0">
                  <table class="table table-vcenter">
                    <thead>
                      <tr>
                        <th>Name</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="project in projects" :key="project.id">
                        <td><router-link class="nav-link" :to="{name: 'project', params: { id: project.id }}">{{ project.name }}</router-link></td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <div class="card-body" v-if="projects.length === 0">
                  <div class="row align-items-center">
                    <div class="col-12">
                      <div class="card-stamp">
                        <div class="card-stamp-icon bg-yellow">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                        </div>
                      </div>
                      <div class="markdown text-secondary">No projects available.</div>
                      <div class="mt-3"><a href="#" class="btn btn-primary" target="_blank" rel="noopener">Create project</a></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="card">
                <div class="card-header">Latest tasks</div>
                <div class="card-table table-responsive" v-if="histories.length > 0">
                </div>
                <div class="card-body" v-if="histories.length === 0">
                  <div class="row align-items-center">
                    <div class="col-12">
                      <div class="card-stamp">
                        <div class="card-stamp-icon bg-yellow">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                        </div>
                      </div>
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