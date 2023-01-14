import "@quasar/extras/roboto-font-latin-ext/roboto-font-latin-ext.css";
import "@quasar/extras/material-icons/material-icons.css";
import "@quasar/extras/material-symbols-outlined/material-symbols-outlined.css";
import materialIcons from "quasar/icon-set/material-icons";

import "quasar/src/css/index.sass";
import quasarSkLang from "quasar/lang/sk";
export const quasarLang = quasarSkLang;
export const quasarConfig = {
  extras: ["material-icons", "material-symbols-outlined"],
};
export const iconSet = materialIcons;
