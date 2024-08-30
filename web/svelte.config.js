import adapter from "@sveltejs/adapter-static";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
    // Consult https://kit.svelte.dev/docs/integrations#preprocessors
    // for more information about preprocessors
    preprocess: vitePreprocess(),

    kit: {
        adapter: adapter({
            pages: "build",
            assets: "build",
            fallback: null,
        }),
        prerender: {
            entries: ["*"], // Prerenders all pages by default
        },
        // Optional: If your app uses client-side routing
        paths: {
            base: "",
            assets: "",
        },
        appDir: "internal", // default is '_app', change if needed
    },
};

export default config;
