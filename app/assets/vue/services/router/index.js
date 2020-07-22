import { createRouter, createWebHistory } from 'vue-router';
import Home from '../../pages/Home';

export class RouterService {
    constructor() {
        this._router = createRouter({ history: createWebHistory(), routes: [{ name: 'Home', path: '/', component: Home }] });
    }

    // prettier-ignore
    /**
     * Add a route to the router
     * @param {import('vue-router').RouteRecordRaw} route the route to add
     */
    addRoute(route) { this._router.addRoute(route); }

    // prettier-ignore
    /**
     * Go to the route based on the given name
     * @param {String} name the name of the route to go to
     */
    goToRoute(name) { this._router.push({ name }); }
}
