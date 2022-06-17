<template>
  <page-container>
    <q-inner-loading
      :showing="isLoggingOut"
      label="Odhlasujem Vás..."
      label-class="text-teal"
      label-style="font-size: 1.1em"
    />
    <q-layout>
      <q-header elevated class="bg-transparent">
        <q-toolbar>
          <q-btn
            flat
            dense
            round
            @click="toggleLeftDrawer"
            aria-label="Menu"
            icon="menu"
          />
          <h1 class="q-my-sm q-ml-lg text-h4">Špajza</h1>
          <q-space />
          <q-btn round flat>
            <q-avatar size="26px">
              <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
            </q-avatar>

            <q-menu class="user-dropdown" auto-close max-width="350px">
              <q-list dense>
                <q-item class="GL__menu-link-signed-in">
                  <q-item-section>
                    <div>
                      Prihlásený ako <strong>{{ auth.username }}</strong>
                    </div>
                  </q-item-section>
                </q-item>
                <q-separator />
                <q-item clickable @click="logout" class="GL__menu-link">
                  <q-item-section>Odhlásiť</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </q-toolbar>
      </q-header>
      <q-drawer
        v-model="isLeftOpen"
        dark
        show-if-above
        class="bg-transparent yellow-1"
        :width="250"
      >
        <q-scroll-area class="fit">
          <q-list padding>
            <q-item
              v-for="link in links"
              :to="link.to"
              :key="link.text"
              active-class="is-active"
              v-ripple
              clickable
            >
              <q-item-section avatar>
                <q-icon color="yellow-1" :name="link.icon" />
              </q-item-section>
              <q-item-section>
                <q-item-label>{{ link.text }}</q-item-label>
              </q-item-section>
            </q-item>

            <!-- <q-separator class="q-my-md bg-yellow-1" /> -->
          </q-list>
        </q-scroll-area>
      </q-drawer>

      <q-page-container>
        <slot></slot>
      </q-page-container>
    </q-layout>
  </page-container>
</template>

<style lang="scss">
.is-active {
  font-weight: 600;
  position: relative;
}
.is-active::before {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  opacity: 0.1;
  content: "";
  background-color: #fff;
}

.user-dropdown {
  background: var(--bg-gradient);
}

.q-drawer.q-drawer {
  background: none;
}
.q-drawer.q-drawer--mobile {
  background: var(--bg-gradient);
}
a,
a:hover {
  text-decoration: none;
  color: inherit;
}
</style>

<script lang="ts" setup>
import { useAuthenticationStore } from "@/auth/authentication";
import { useQuasar } from "quasar";
import { ref } from "vue";
import { useLink, useRouter } from "vue-router";
import PageContainer from "./PageContainer.vue";
// https://github.com/quasarframework/quasar/issues/13154
// temporal workaroud because v-ripple broke the page
const $q = useQuasar();
defineExpose({
  $q,
});

const isLeftOpen = ref(false);
const isLoggingOut = ref(false);
const auth = useAuthenticationStore();
const router = useRouter();
const toggleLeftDrawer = () => (isLeftOpen.value = !isLeftOpen.value);
const links = [
  { icon: "home", text: "Prehľad", to: "/" },
  { icon: "storage", text: "Špajza", to: "/spajza" },
  { icon: "menu_book", text: "Recepty", to: "/recepty" },
  { icon: "shopping_bag", text: "Nákup", to: "/nakup" },
];

const logout = () => {
  console.log("logout");
  isLoggingOut.value = false;
  auth
    .logout()
    .then(() => {
      router.push("/prihlasenie");
    })
    .finally(() => {
      isLoggingOut.value = true;
    });
};
</script>
