import {h} from 'vue';

import TheLogo from './components/TheLogo';
import TheNavbar from './components/TheNavbar';
import {RouterView} from 'vue-router';

export default {
    setup() {
        return () => [h(TheNavbar), TheLogo(), h(RouterView)];
    }
};
