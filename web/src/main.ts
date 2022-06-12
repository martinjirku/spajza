import { createApp } from "vue";
import App from "./App.vue";
import { quasarLang } from "./quasar";
import { createPinia } from "pinia";
import { Quasar } from "quasar";
import router from "@/router";
import "./localization";

const app = createApp(App);
app.use(createPinia());
app.use(Quasar, {
  plugins: {},
  lang: quasarLang,
});
app.use(router);
app.mount("#app");
