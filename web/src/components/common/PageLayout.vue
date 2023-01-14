<template>
  <page-container>
    <q-inner-loading
      :showing="isLoggingOut"
      label="Odhlasujem Vás..."
      label-class="text-teal"
      label-style="font-size: 1.1em"
    />
    <q-layout view="hHh lpr fFf">
      <q-header elevated>
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
                <q-item>
                  <q-item-section>
                    <div>
                      Prihlásený ako <strong>{{ auth.username }}</strong>
                    </div>
                  </q-item-section>
                </q-item>
                <q-separator />
                <q-item to="/typ-poloziek" active-class="is-active">
                  <q-item-section> Druhy položiek </q-item-section>
                </q-item>
                <q-item to="/miesta-ulozenia" active-class="is-active">
                  <q-item-section> Miesta uloženia </q-item-section>
                </q-item>
                <q-separator />
                <q-item clickable @click="logout">
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
        <q-page :style-fn="calculateHeight">
          <slot></slot>
        </q-page>
      </q-page-container>
    </q-layout>
  </page-container>
</template>
<style>
.is-active {
  font-weight: 600;
  position: relative;
}
.is-active-user-menu {
  cursor: default !important;
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
  transition: none;
}

.q-header {
  background: var(--bg-gradient);
}

a,
a:hover {
  text-decoration: none;
  color: inherit;
}

.q-drawer.q-drawer.q-drawer--mobile {
  background: var(--bg-gradient);
}
.user-dropdown {
  background: var(--bg-gradient);
  color: #fff;
}
.q-drawer.q-drawer {
  background: none;
}

body.desktop .q-focus-helper.q-focus-helper {
  transition: background-color 0.2s cubic-bezier(0.25, 0.8, 0.5, 1),
    opacity 0.1s cubic-bezier(0.25, 0.8, 0.5, 1);
}
body.desktop .q-focus-helper.q-focus-helper::before,
body.desktop .q-focus-helper.q-focus-helper::after {
  transition: background-color 0.2s cubic-bezier(0.25, 0.8, 0.5, 1),
    opacity 0.1s cubic-bezier(0.25, 0.8, 0.5, 1);
}
</style>
<style lang="scss" scoped></style>

<script lang="ts" setup>
import { useAuthenticationStore } from "@auth/authentication";
import { useQuasar } from "quasar";
import { ref } from "vue";
import { useRouter } from "vue-router";
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
  { icon: "storage", text: "Špajza", to: "/spajza" },
  { icon: "menu_book", text: "Recepty", to: "/recepty" },
  { icon: "shopping_bag", text: "Nákup", to: "/nakup" },
];

const logout = () => {
  isLoggingOut.value = false;
  auth
    .logout()
    .then(() => {
      router.push({ name: "login" });
    })
    .finally(() => {
      isLoggingOut.value = true;
    });
};

const calculateHeight = (offset: number) => ({
  height: offset ? `calc(100vh - ${offset}px)` : "100vh",
});
</script>
