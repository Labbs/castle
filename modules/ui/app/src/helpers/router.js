import { createRouter, createWebHistory } from 'vue-router';

import { useAuthStore } from '@/stores';
import { LoginView,
  HomeView,
  ProfileView,
  ProjectListView,
  ProjectView,
  NewTaskView,
  RepositoryView,
  NewRepositoryView,
  UpdateRepositoryView,
  UpdateProjectView,
  UpdateVariableView
} from '@/views';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: HomeView, name: 'home' },
    { path: '/login', component: LoginView, name: 'login' },
    { path: '/profile', component: ProfileView, name: 'profile' },
    { path: '/project/list', component: ProjectListView, name: 'project-list' },
    { path: '/project/:id', component: ProjectView, name: 'project' },
    { path: '/project/:id/new/task', component: NewTaskView, name: 'project-new-task' },
    { path: '/project/edit/:id', component: UpdateProjectView, name: 'project-update'},
    { path: '/repository', component: RepositoryView, name: 'repository' },
    { path: '/repository/new', component: NewRepositoryView, name: 'repository-new' },
    { path: '/repository/:id', component: UpdateRepositoryView, name: 'repository-update' },
    { path: '/variable/:id', component: UpdateVariableView, name: 'variable-update'},
    { path: '/task/edit/:id', component: NewTaskView, name: 'task-update'}
  ]
});

router.beforeEach(async (to) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const auth = useAuthStore();

  if (authRequired && !auth.user) {
      auth.returnUrl = to.fullPath;
      return '/login';
  }
});