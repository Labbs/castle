import { defineStore } from "pinia";

import { fetchWrapper, router } from "@/helpers";

const baseUrl = `${import.meta.env.VITE_API_URL}`;

export const useAuthStore = defineStore({
  id: "auth",
  state: () => ({
    // initialize state from local storage to enable user to stay logged in
    user: JSON.parse(sessionStorage.getItem("user")),
    returnUrl: null,
  }),
  actions: {
    async login(username, password) {
      const user = await fetchWrapper.post(`${baseUrl}/auth/login`, {
        username,
        password,
      });

      // update pinia state
      this.user = user;

      // store user details and jwt in local storage to keep user logged in between page refreshes
      sessionStorage.setItem("user", JSON.stringify(user));

      // redirect to previous url or default to home page
      router.push(this.returnUrl || "/");
    },
    logout() {
      this.user = null;
      sessionStorage.removeItem("user");
      router.push("/login");
    },
  },
});
