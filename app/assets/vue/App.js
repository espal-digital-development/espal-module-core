import { h } from 'vue';

import TheLogo from './components/TheLogo';
import TheNavbar from './components/TheNavbar.vue'
import { RouterView } from 'vue-router';

export default {
    setup() {
        return () => [
            h(TheNavbar),
            TheLogo,
            h(RouterView)
        ];
    }
};
