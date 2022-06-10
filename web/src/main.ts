import { createApp } from "vue";
import App from "./App.vue";
import { createPinia } from "pinia";
import router from "@/router";
import vuetify from "./vuetify";
import "vuetify/styles";
import "./localization";

const app = createApp(App);
app.use(createPinia());
app.use(vuetify);
app.use(router);
app.mount("#app");
