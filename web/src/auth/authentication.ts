import { defineStore } from "pinia";

export const useAuthenticationStore = defineStore("auth", {
  state: () => ({
    loggedIn: true,
  }),
});
