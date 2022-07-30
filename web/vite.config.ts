import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { fileURLToPath, URL } from "url";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "@assets": fileURLToPath(new URL("./src/assets", import.meta.url)),
      "@categories": fileURLToPath(
        new URL("./src/categories", import.meta.url)
      ),
      "@storage": fileURLToPath(new URL("./src/storage", import.meta.url)),
      "@units": fileURLToPath(new URL("./src/units", import.meta.url)),
      "@api": fileURLToPath(new URL("./src/api", import.meta.url)),
      "@components": fileURLToPath(
        new URL("./src/components", import.meta.url)
      ),
      "@pages": fileURLToPath(new URL("./src/pages", import.meta.url)),
      "@auth": fileURLToPath(new URL("./src/auth", import.meta.url)),
    },
  },
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    quasar({
      runMode: "web-client",
      sassVariables: "./src/quasar-variables.sass",
    }),
  ],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8000",
        changeOrigin: true,
        secure: false,
        ws: true,
      },
    },
  },
});
