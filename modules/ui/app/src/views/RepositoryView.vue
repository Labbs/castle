<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const repositories = ref([]);
    const message = ref('');
    return {
      repositories,
      message,
    };
  },
  mounted() {
    this.getRepositories();
  },
  methods: {
    async getRepositories() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/repository/all`);
      this.repositories = response.repositories;
    },

    async deleteRepository(id) {
      const response = await fetchWrapper.delete(`${import.meta.env.VITE_API_URL}/repository/delete/${id}`);
      if (response.message === 'Repository deleted') {
        this.getRepositories();
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
              <h2 class="page-title">Repository</h2>
            </div>
            <div class="col-12 col-md-auto ms-auto d-print-none">
              <div class="btn-list">
                <router-link class="btn btn-primary d-none d-sm-inline-block" to="/repository/new">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
                  Create new repository
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="row row-deck row-cards">
            <div class="col-lg-12">
              <div class="card">
                <div class="card-table table-responsive" v-if="repositories.length > 0">
                  <table class="table card-table table-vcenter text-nowrap datatable">
                    <thead>
                      <tr>
                        <th>Name</th>
                        <th>Action</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="repository in repositories" :key="repository.id">
                        <td>{{ repository.name }}</td>
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
                <div class="card-body" v-if="repositories.length === 0">
                  <div class="card-stamp">
                    <div class="card-stamp-icon bg-yellow">
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M10 5a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path><path d="M9 17v1a3 3 0 0 0 6 0v-1"></path></svg>
                    </div>
                  </div>
                  <div class="row align-items-center">
                    <div class="col-12">
                      <div class="markdown text-secondary">No repository available.</div>
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