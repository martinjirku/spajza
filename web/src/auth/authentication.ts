import { defineStore } from "pinia";
type UserResponse = {
  username: string;
};

type Store = {
  authStatus: "checking" | "loaded";
  loggedIn: boolean;
  username?: string;
  returnUrl?: string;
};

const DEFAULT_RETURN_URL = "/";

export const useAuthenticationStore = defineStore("auth", {
  state: (): Store => ({
    authStatus: "checking",
    loggedIn: false,
    username: undefined as string | undefined,
    returnUrl: DEFAULT_RETURN_URL,
  }),
  actions: {
    resetReturnUrl() {
      this.returnUrl = DEFAULT_RETURN_URL;
    },
    async checkAuthentification() {
      const response = await fetch("/api/user/me", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      })
        .then((a) => a.json())
        .then((d) => {
          if (d.username) {
            this.loggedIn = true;
            this.username = d.username;
          }
        })
        .catch(() => {
          this.loggedIn = false;
          this.username = undefined;
        })
        .finally(() => {
          this.authStatus = "loaded";
        });

      return response;
    },
    async login(username: string, password: string) {
      const response = await fetch("/api/user/login", {
        method: "POST",
        body: JSON.stringify({ username, password }),
        headers: { "Content-Type": "application/json" },
      });
      await this.checkAuthentification();
      if (response.ok) {
        this.loggedIn = true;
      }
      return response;
    },
    async logout() {
      const response = await fetch("/api/user/logout", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });
      if (response.ok) {
        this.loggedIn = false;
      }
      return response;
    },
  },
});
