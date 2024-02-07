import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// const outDir = resolve(__dirname, "dist");
import { resolve } from "path";
const root = resolve(__dirname, "src");

export default defineConfig({
  plugins: [
    react(),
    {
      name: "markdown-loader",
      transform(code, id) {
        if (id.slice(-3) === ".md") {
          return `export default ${JSON.stringify(code)};`;
        }
      },
    },
  ],
  build: {
    manifest: "manifest.json",
    rollupOptions: {
      input: {
        main: resolve(root, "index.html"),
      },
    },
  },
  server: {
    port: 5173,
  },
});
