/**
 * @typedef {Object} StoreModule
 * @property {*} state
 * @property {*} getters
 * @property {*} mutations
 * @property {*} actions
 */

export class StoreService {
    constructor() {
        /** @type {Object<string,StoreModule>} */
        this._store = {};
    }

    /**
     * register a store module in the store
     * @param {String} name
     * @param {StoreModule} storeModule
     */
    registerModule(name, storeModule) {
        this._store[name] = storeModule;
    }
}
