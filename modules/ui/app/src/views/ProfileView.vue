<script>
import NavbarView from '@/views/NavbarView.vue';
import { ref } from 'vue';
import { fetchWrapper,darkMode } from "@/helpers";
import { useDropzone } from "vue3-dropzone";

export default {
  components: {
    NavbarView,
  },
  setup() {
    const profile = ref({});
    const newUsername = ref("");
    const newPassword = ref("");
    const message = ref({});

    function onDrop(acceptFiles, rejectReasons) {
      if (rejectReasons.length) {
        console.log(rejectReasons);
        return;
      }
      const reader = new FileReader();
      reader.readAsDataURL(acceptFiles[0]);
      reader.onload = () => {
        profile.value.avatar = reader.result;
        updateAvatar();
      };
    }

    const { getRootProps, getInputProps, ...rest } = useDropzone({ onDrop });

    return {
      profile,
      newUsername,
      newPassword,
      message,
      getRootProps,
      getInputProps,
      ...rest,
    }
  },

  mounted() {
    this.getProfile();
  },

  methods: {
    async getProfile() {
      const response = await fetchWrapper.get(`${import.meta.env.VITE_API_URL}/user`);
      this.profile = response.user;
      this.newUsername = this.profile.username;
    },

    async updateDarkMode() {
      const response = await fetchWrapper.put(`${import.meta.env.VITE_API_URL}/user/dark_mode`, {
        dark_mode: this.profile.dark_mode,
      });
      this.message = response;
      darkMode(this.profile.dark_mode);
    },

    async updateAvatar() {
      const response = await fetchWrapper.put(`${import.meta.env.VITE_API_URL}/user/avatar`, {
        avatar: this.profile.avatar,
      });
      this.message = response;
    },

    async updateUsername() {
      const response = await fetchWrapper.put(`${import.meta.env.VITE_API_URL}/user/username`, {
        current_username: this.profile.username,
        new_username: this.newUsername,
      });
      this.message = response;
    },
  },
}
</script>

<template>
  <div class="page">
    <NavbarView />
    <div class="page-wrapper">
      <div class="page-header d-print-none">
        <div class="container-xl">
          <div class="alert alert-dismissible" role="alert" v-if="message.status" :class="message.status == 'success' ? 'alert-success': 'alert-danger'">
            <div>{{ message.message }}</div>
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
          </div>
          <div class="row g-2 align-items-center">
            <div class="col">
              <h2 class="page-title">Account Settings</h2>
            </div>
          </div>
        </div>
      </div>
      <div class="page-body">
        <div class="container-xl">
          <div class="card">
            <div class="row g-0">
              <div class="col-3 d-none d-md-block border-end">
                <div class="card-body">
                  <div class="list-group list-group-transparent">
                    <a href="#" class="list-group-item list-group-item-action d-flex align-items-center active">My Account</a>
                    <a href="#" class="list-group-item list-group-item-action d-flex align-items-center">API Token</a>
                  </div>
                </div>
              </div>
              <div class="col d-flex flex-column">
                <div class="card-body">
                  <h2 class="mb-4">My Account</h2>
                  <h3 class="card-title">Profile Details</h3>
                  <div class="row align-items-center">
                    <div class="col-auto"><span class="avatar avatar-xl" style="background-image: url({{ profile.avatar }})"></span>
                    </div>
                    <div class="col-auto"><a href="#" class="btn" data-bs-toggle="modal" data-bs-target="#modal-change-avatar">Change avatar</a></div>
                    <div class="col-auto"><a href="#" class="btn btn-ghost-danger">Delete avatar</a></div>
                  </div>

                  <h3 class="card-title mt-4">Username</h3>
                  <p class="card-subtitle">This contact will be shown to others publicly, so choose it carefully.</p>
                  <div>
                    <div class="row g-2">
                      <div class="col-auto">
                        <input type="text" class="form-control w-auto" v-model="newUsername" disabled>
                      </div>
                    </div>
                  </div>

                  <h3 class="card-title mt-4">Password</h3>
                  <p class="card-subtitle">You can set a permanent password if you don't want to use temporary login codes.</p>
                  <div>
                    <a href="#" class="btn">
                      Set new password
                    </a>
                  </div>
                  <h3 class="card-title mt-4">Dark mode</h3>
                  <p class="card-subtitle">Configure your theme.</p>
                  <div class="col-6">
                    <select class="form-select" v-model="profile.dark_mode" @change="updateDarkMode()">
                      <option value="light" :value="{ darkModeValue: 'light' }">Light</option>
                      <option value="dark" :value="{ darkModeValue: 'dark' }">Dark</option>
                      <option value="auto" :value="{ darkModeValue: 'auto' }">Auto</option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="modal modal-blur fade" id="modal-change-avatar" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Change avatar</h5>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <div v-bind="getRootProps()" class="dropzone">
                <input v-bind="getInputProps()" />
                <p v-if="isDragActive">Drop the files here ...</p>
                <p v-else>Drag 'n' drop some files here, or click to select files</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="modal modal-blur fade" id="modal-create-token" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <form action="/user/token/create" method="post">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title">Create new token</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div class="modal-body">
                <div class="mb-3">
                  <label class="form-label">Name</label>
                  <input type="text" class="form-control" name="name" required />
                </div>
                <div class="mb-3">
                  <label class="form-label">Scope(s)</label>
                  <label class="form-check form-switch">
                    <input class="form-check-input" type="checkbox" name="bookmark">
                    <span class="form-check-label">Bookmark</span>
                  </label>
                </div>
              </div>
              <div class="modal-footer">
                <button type="submit" class="btn btn-primary ms-auto" data-bs-dismiss="modal">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
                  Create
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
      <div class="modal modal-blur fade" id="modal-password" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-lg" role="document">
          <form action="/user/password/edit" method="post">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title">Set new password</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div class="modal-body">
                <div class="mb-3">
                  <label class="form-label">Current password</label>
                  <input type="password" class="form-control" name="current-password" required />
                </div>
                <div class="mb-3">
                  <label class="form-label">New password</label>
                  <input type="password" class="form-control" name="new-password" required />
                </div>
                <div class="mb-3">
                  <label class="form-label">Confirm new password</label>
                  <input type="password" class="form-control" name="confirm-password" required/>
                </div>
              </div>
              <div class="modal-footer">
                <button type="submit" class="btn btn-primary ms-auto" data-bs-dismiss="modal">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><line x1="12" y1="5" x2="12" y2="19" /><line x1="5" y1="12" x2="19" y2="12" /></svg>
                  Update
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.dropzone {
  border: 1px dashed #e6e7e9;
  border-radius: 5px;
  min-height: 150px;
  padding: 1rem;
}
</style>