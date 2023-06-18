import { createRouter, createWebHistory } from "vue-router";
import Login from "@pages/Login.vue";
import StorageRoom from "@pages/Storage/Storage.vue";
import Recipies from "@pages/Recipies.vue";
import Shopping from "@pages/Shopping.vue";
import CategoryItems from "@pages/CategoryItems.vue";
import Settings from "@pages/Settings.vue";
import StoragePlaces from "@pages/StoragePlaces.vue";

import { useAuthenticationStore } from "@/auth/authentication";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      redirect: "/spajza",
    },
    {
      name: "login",
      path: "/prihlasenie",
      component: Login,
    },
    {
      name: "storageRoom",
      path: "/spajza",
      component: StorageRoom,
    },
    {
      name: "recipies",
      path: "/recepty",
      component: Recipies,
    },
    {
      name: "shopping",
      path: "/nakup",
      component: Shopping,
    },
    {
      name: "item-type",
      path: "/typ-poloziek",
      component: CategoryItems,
    },
    {
      name: "settings",
      path: "/nastavenia",
      component: Settings,
    },
    {
      name: "storage-location",
      path: "/miesta-ulozenia",
      component: StoragePlaces,
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
