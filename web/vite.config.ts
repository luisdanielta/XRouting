import path from "path"
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'


// Vite configuration
export default defineConfig({
  plugins: [
    tailwindcss(), react(),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"), // Resolve '@' to the 'src' directory
    },
  },
  build: {
    outDir: "dist", // Output directory for production build
    sourcemap: true, // Enable source maps for debugging in production
  },
  define: {
    "process.env": {}, // Ensure compatibility for libraries expecting 'process.env'
  },
})