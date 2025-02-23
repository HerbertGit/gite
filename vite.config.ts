import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

export default defineConfig({
  plugins: [react()],
  // server: {
  //   cors: {
  //     // the origin you will be accessing via browser
  //     origin: "http://my-backend.example.com",
  //   },
  // },
  build: {
    // generate .vite/manifest.json in outDir
    manifest: true,
    rollupOptions: {
      // overwrite default .html entry
      input: "/src/main.tsx",
    },
    // rollupOptions: {
    //   // overwrite default .html entry
    //   input: "/path/to/main.js",
    // },
  },
});
