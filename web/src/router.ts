import { createRouter, createWebHistory } from "vue-router";
import Login from "@/pages/Login.vue";
import Home from "@/pages/Home.vue";
import StorageRoom from "@/pages/Storage.vue";
import Recipies from "@/pages/Recipies.vue";
import Shopping from "@/pages/Shopping.vue";

import { useAuthenticationStore } from "@/auth/authentication";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Home,
    },
    {
      path: "/prihlasenie",
      component: Login,
    },
    {
      path: "/spajza",
      component: StorageRoom,
    },
    {
      path: "/recepty",
      component: Recipies,
    },
    {
      path: "/nakup",
      component: Shopping,
    },
  ],
});

router.beforeEach(async (to) => {
  const publicPages = ["/prihlasenie"];
  const authRequired = !publicPages.includes(to.path);
  const auth = useAuthenticationStore();
  if (auth.authStatus === "checking") {
    await auth.checkAuthentification();
  }
  if (to.path === "/prihlasenie") {
    if (auth.loggedIn) {
      auth.resetReturnUrl();
      return auth.returnUrl;
    }
    return;
  }
  if (!authRequired) return;

  if (auth.loggedIn) {
    return;
  } else {
    auth.returnUrl = to.fullPath;
    return "/prihlasenie";
  }
});

export default router;
