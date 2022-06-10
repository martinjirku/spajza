import { createRouter, createWebHistory } from "vue-router";
import Login from "@/pages/Login.vue";
import Home from "@/pages/Home.vue";
import { useAuthenticationStore } from "@/auth/authentication";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Home,
    },
    {
      path: "/login",
      component: Login,
    },
  ],
});

router.beforeEach(async (to) => {
  const publicPages = ["/login"];
  const authRequired = !publicPages.includes(to.path);
  const auth = useAuthenticationStore();

  if (authRequired && !auth.loggedIn) {
    auth.returnUrl = to.fullPath;
    return "/login";
  }
});

export default router;
