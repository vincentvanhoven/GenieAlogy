import { createApp } from "vue";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import Aura from "@primeuix/themes/aura";
import App from "./App.vue";
import "./style.css";

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(PrimeVue, {
    theme: {
        preset: Aura,
        options: {
            prefix: "p",
            cssLayer: {
                name: "primevue",
                order: "theme, base, primevue",
            },
        },
    },
});
app.mount("#app");
