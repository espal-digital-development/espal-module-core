/**
 * @typedef {import('../http/index').HTTPService} HTTPService
 */
import {BaseStoreModule} from './storeModule';

export class StoreService {
    /**
     * @param {HTTPService} httpService the service that makes the requests
     */
    constructor(httpService) {
        /** @type {Object<string,BaseStoreModule>} */
        this._store = {};
        this._httpService = httpService;
    }

    /**
     * register a store module in the store
     * @param {String} endpoint
     * @param {BaseStoreModule} storeModule
     */
    registerModule(endpoint, storeModule) {
        this._store[endpoint] = storeModule;
    }

    /**
     * Create a Store Module
     * @param {String} endpoint
     */
    createModule(endpoint) {
        return new BaseStoreModule(endpoint, this._httpService);
    }

    /**
     * Create and a set a default Store Module
     * @param {String} endpoint
     */
    createAndRegisterModule(endpoint) {
        this.registerModule(endpoint, this.createModule(endpoint));
    }
}
