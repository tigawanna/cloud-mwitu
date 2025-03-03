import { defineConfig } from "vite";
import preact from "@preact/preset-vite";
import tailwindcss from "@tailwindcss/vite";
import tsconfigPaths from "vite-tsconfig-paths";
import analyze from "rollup-plugin-analyzer";
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    preact(),
    
    tailwindcss(),
    tsconfigPaths(),
    analyze({
      // highlight the modules with size > 40kb
      filter(moduleObject) {
        return moduleObject.size > 40000;
      },
    }),

    
  ],
    resolve: {
    // add this line
    dedupe: ["preact"],
    alias: [
      // { find: "react", replacement: "preact/compat" },
      // { find: "react-dom/test-utils", replacement: "preact/test-utils" },
      // { find: "react-dom", replacement: "preact/compat" },
      // { find: "react/jsx-runtime", replacement: "preact/jsx-runtime" },
      // add this line
      { find: "@preact/signals-react", replacement: "@preact/signals" },
    ],
  },
});
