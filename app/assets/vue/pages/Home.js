import {h, ref} from 'vue';
import {authService} from '../services';

export default {
    setup() {
        const jwtValidString = ref('');

        const horizontalRule = h('hr');
        return () =>
            h('div', [
                h('h1', ['Hello!']),
                horizontalRule,
                h('p', [authService.jwt]),
                horizontalRule,
                h('p', [jwtValidString.value]),
                h(
                    'button',
                    {
                        onclick: async () => (jwtValidString.value = await authService.checkJWT())
                    },
                    ['Check JWT']
                )
            ]);
    }
};
