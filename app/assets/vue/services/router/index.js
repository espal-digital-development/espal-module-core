/**
 * @typedef {import('vue').Component} Component
 * @typedef {import('vue-router').RouteRecordRaw} RouteRecordRaw
 *
 * @typedef {Object} RouteSettings
 * @property {String} endpoint
 * @property {Component} baseComponent
 * @property {Component|boolean} overviewComponent
 * @property {Component|boolean} createComponent
 * @property {Component|boolean} showComponent
 * @property {Component|boolean} editComponent
 */

import {createRouter, createWebHistory} from 'vue-router';
import Home from '../../pages/Home';

export class RouterService {
    constructor() {
        this._router = createRouter({
            history: createWebHistory(),
            routes: [{name: 'Home', path: '/', component: Home}]
        });
    }

    // prettier-ignore
    /**
     * Add a route to the router
     * @param {RouteRecordRaw} route the route to add
     */
    addRoute(route) { this._router.addRoute(route); }

    // prettier-ignore
    /**
     * Go to the route based on the given name
     * @param {String} name the name of the route to go to
     */
    goToRoute(name) { this._router.push({ name }); }

    /**
     * create basic routes for the given settings
     *
     * @param {RouteSettings} settings the settings on which the routes are based
     */
    createBaseRoutes(settings) {
        const base = this.createBase(settings.endpoint, settings.baseComponent);

        // TODO :: put all strings in getters
        if (settings.overviewComponent) {
            base.children.push(this.createRecord('', settings.endpoint + '.overview', settings.overviewComponent));
        }

        if (settings.createComponent) {
            base.children.push(this.createRecord('/create', settings.endpoint + '.create', settings.createComponent));
        }

        if (settings.showComponent) {
            base.children.push(this.createRecord('/:id', settings.endpoint + '.overview', settings.showComponent));
        }

        if (settings.editComponent) {
            base.children.push(this.createRecord('/:id/edit', settings.endpoint + '.overview', settings.editComponent));
        }

        return base;
    }

    /**
     * Create the base for the routes based on the settings and add the children to it
     *
     * @param {basePath} string
     * @param {Component} baseComponent
     *
     * @returns {RouteRecordRaw}
     */
    createBase(basePath, baseComponent) {
        // TODO :: create createPath method to always add a slash to the given path
        return {
            path: '/' + basePath,
            component: baseComponent,
            children: []
        };
    }

    /**
     * Create a standard route record
     *
     * @param {String} path the name of the path for the route config
     * @param {String} name the name of the route
     * @param {*} component the component to render for this route
     *
     * @returns {RouteRecordRaw}
     */
    createRecord(path, name, component) {
        return {path, name, component};
    }
}
