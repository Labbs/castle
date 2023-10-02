<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper } from "@/helpers";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const variable = ref({});

    return {
      variable
    };
  },
  mounted() {
    this.getVariable();
  },
  methods: {
    async onSubmit(variable) {
      const response = await fetchWrapper.post(`${import.meta.env.VITE_API_URL}/variable/update/${this.$route.params.id}`, variable);
      if (response.message === 'Variable edited') {
        this.$router.push('/variable');
      }
    },
    async getVariable() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/variable/${this.$route.params.id}`);
      this.variable = response.variable;
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
              <h2 class="page-title">Update variable</h2>
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
                      <input type="text" class="form-control" name="name" v-model="variable.name" />
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label required">Value</label>
                    <div class="col">
                      <textarea rows="5" class="form-control" name="value" v-model="variable.value"></textarea>
                    </div>
                  </div>
                  <div class="mb-3 row">
                    <label class="col-3 col-form-label">Type</label>
                    <div class="col">
                      <select class="form-select" v-model="variable.type" disabled>
                        <option>text</option>
                        <option>secret</option>
                      </select>
                    </div>
                  </div>
                </div>
                <div class="card-footer d-flex justify-content-end">
                  <button class="btn btn-primary ms-auto" @click="onSubmit(variable)">Update variable</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>