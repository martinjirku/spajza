import { defineStore } from "pinia";

export const useAuthenticationStore = defineStore("auth", {
  state: () => ({
    loggedIn: false,
    returnUrl: "/",
  }),
  actions: {
    async login(username: string, password: string) {
      const response = await fetch("/api/user/login", {
        method: "POST",
        body: JSON.stringify({ username, password }),
        headers: { "Content-Type": "application/json" },
      });
      console.log(response);
      if (response.ok) {
        this.loggedIn = true;
        this.returnUrl = "/";
      }
      return response;
    },
  },
});
