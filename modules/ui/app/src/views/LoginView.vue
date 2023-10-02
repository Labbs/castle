<script setup>
  import { Form, Field } from 'vee-validate';
  import * as Yup from 'yup';

  import { useAuthStore } from '@/stores';

  const schema = Yup.object().shape({
    username: Yup.string().required('Username is required'),
    password: Yup.string().required('Password is required')
  });

  function onSubmit(values, { setErrors }) {
    const authStore = useAuthStore();
    const { username, password } = values;

    return authStore.login(username, password)
        .catch(error => setErrors({ apiError: error }));
  }
</script>

<template>
  <div class="page page-center">
    <div class="container-tight py-4">
      <div class="card card-md">
        <div class="card-body">
          <h2 class="card-title text-center mb-4">Login to your account</h2>
          <Form @submit="onSubmit" :validation-schema="schema" v-slot="{ errors, isSubmitting }">
            <div class="mb-3">
              <label class="form-label">Username</label>
              <Field name="username" type="username" class="form-control" placeholder="Enter username" :class="{ 'is-invalid': errors.username }" />
            </div>
            <div class="mb-2">
              <label class="form-label">Password</label>
              <Field name="password" type="password" class="form-control" placeholder="Enter password" :class="{ 'is-invalid': errors.password }" />
            </div>
            <div class="form-footer">
              <button class="btn btn-primary w-100" :disabled="isSubmitting">Sign in</button>
            </div>
          </Form>
        </div>
      </div>
    </div>
  </div>
</template>