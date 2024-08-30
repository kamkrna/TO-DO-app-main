import type { CapacitorConfig } from "@capacitor/cli";

const config: CapacitorConfig = {
    appId: "com.todo.app",
    appName: "todoapp",
    webDir: "build",
    // server: {
    //     allowNavigation: ["http://127.0.0.1:8000"],
    // },

    plugins: {
        CapacitorHttp: {
            enabled: true,
        },
    },

    // plugins: {
    //     "CapacitorHttp": {
    //       "enabled": true
    //     }
    //   }
};

export default config;
