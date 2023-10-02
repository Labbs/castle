<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const repository = {
      name: '',
      description: '',
      url: '',
      ssh_key: '',
      ssh_key_pass: ''
    }

    return {
      repository
    };
  },
  mounted() {
  },
  methods: {
    async onSubmit(repository) {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/repository/create`, repository);
      if (response.message === 'Repository created') {
        this.$router.push('/repository');
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
              <h2 class="page-title">New repository</h2>
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
                      <input type="text" class="form-control" name="name" v-model="repository.name" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Description</label>
                    <div class="col">
                      <input type="text" class="form-control" name="description" v-model="repository.description" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">URL</label>
                    <div class="col">
                      <input type="text" class="form-control" name="url" v-model="repository.url" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">SSH Key</label>
                    <div class="col">
                      <textarea rows="5" class="form-control" name="ssh_key" v-model="repository.ssh_key"></textarea>
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">SSH Key Password</label>
                    <div class="col">
                      <input type="password" class="form-control" name="ssh_key_pass" v-model="repository.ssh_key_pass" />
                    </div>
                  </div>
                </div>
                <div class="card-footer text-end">
                  <button class="btn btn-primary" @click="onSubmit(repository)">Save</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>