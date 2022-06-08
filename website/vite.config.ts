import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
    build: {
        rollupOptions: {
            external: "wasm_exec.js",
            input: {
                main: path.resolve(__dirname, 'index.html'),
            },
        },
        outDir: "../docs",
        commonjsOptions: {
            include: [/node_modules/]
        }
    },
    server: {
        port: 8080
    }
});