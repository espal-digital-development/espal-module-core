import {h} from 'vue';
import TheLogo from './TheLogo';
import {authService, routerService} from '../services';
import {userModule} from '../modules';

const Logo = TheLogo(50, 50, () => routerService.goToRoute('Home'));

export default h('header', [
    Logo,
    h('a', {onclick: () => authService.goToLoginPage()}, ['Login']),
    h('a', {onclick: () => userModule.goToOverview()}, ['Users'])
]);
