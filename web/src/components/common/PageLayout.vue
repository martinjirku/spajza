<template>
  <page-container>
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
              :to="link.link.href.value"
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

.q-drawer.q-drawer {
  background: none;
}
.q-drawer.q-drawer--mobile {
  background: linear-gradient(
    113.04deg,
    #e68a8a -0.64%,
    #e6a15c 37.48%,
    #e08053 68.8%,
    #ebeb8f 99.61%
  );
}
a,
a:hover {
  text-decoration: none;
  color: inherit;
}
</style>

<script lang="ts" setup>
import { useQuasar } from "quasar";
import { ref } from "vue";
import { useLink } from "vue-router";
import PageContainer from "./PageContainer.vue";
// https://github.com/quasarframework/quasar/issues/13154
// temporal workaroud because v-ripple broke the page
const $q = useQuasar();
defineExpose({
  $q,
});

const isLeftOpen = ref(false);

const toggleLeftDrawer = () => (isLeftOpen.value = !isLeftOpen.value);
const links = [
  { icon: "home", text: "Prehľad", link: useLink({ to: "/" }) },
  { icon: "storage", text: "Špajza", link: useLink({ to: "/spajza" }) },
  { icon: "menu_book", text: "Recepty", link: useLink({ to: "/recepty" }) },
];
</script>
