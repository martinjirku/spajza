import { createVuetify, ThemeDefinition } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { mdi, aliases } from "vuetify/iconsets/mdi";
import "@mdi/font/css/materialdesignicons.css";

const zasobarTheme: ThemeDefinition = {
  dark: false,
  variables: {
    "theme-surface": "255, 255, 255, 0.85",
  },
  colors: {
    primary: "#409EFF",
    secondary: "#03DAC6",
    error: "#B00020",
    info: "#2196F3",
    success: "#4CAF50",
    warning: "#FB8C00",
  },
};

export default createVuetify({
  components,
  directives,
  locale: {
    defaultLocale: "sk-SK",
  },
  defaults: {
    VTextField: {
      variant: "contained",
      bgColor: "white",
    },
    global: {},
  },
  theme: {
    defaultTheme: "zasobarTheme",
    themes: {
      zasobarTheme,
    },
  },
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: { mdi },
  },
});
