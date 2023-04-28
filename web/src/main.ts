import { createApp } from "vue";
import App from "./App.vue";
import { quasarLang, quasarConfig, iconSet } from "./quasar";
import { createPinia } from "pinia";
import { VueQueryPlugin } from "@tanstack/vue-query";
import { Quasar } from "quasar";
// this need to be called before the Router
import "./localization";
import router from "@/router";

const app = createApp(App);
app.use(createPinia());
app.use(Quasar, {
  plugins: {},
  config: quasarConfig,
  lang: quasarLang,
  iconSet: iconSet,
});
app.use(router);
app.use(VueQueryPlugin);
app.mount("#app");
