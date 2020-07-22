/**
 * @typedef {import('../http').HTTPService} HTTPService
 * @typedef {import('../storage').StorageService} StorageService
 *
 * @typedef {Object} Credentials
 * @property {String} email
 * @property {String} password
 * @property {Boolean} rememberMe
 */

import { ref } from 'vue';

const JWT_STORAGE_KEY = 'Epal-JWT';

export class AuthService {
    /**
     * @param {HTTPService} httpService
     * @param {StorageService} storageService
     */
    constructor(httpService, storageService) {
        this._httpService = httpService;
        this._storageService = storageService;

        const storedJWT = this._storageService.getItem(JWT_STORAGE_KEY);
        this._jwt = ref(storedJWT || '');
    }

    // prettier-ignore
    get jwt() { return this._jwt.value; }

    set jwt(value) {
        this._jwt.value = value;
        this._storageService.setItem(JWT_STORAGE_KEY, value);
    }

    /**
     * @param {Credentials} credentials
     */
    async login(credentials) {
        this._httpService.post('Login', credentials).then((result) => (this.jwt = result));
    }

    /** Checks if jwt is valid */
    checkJWT() {
        return this._httpService.get('Account', {
            headers: { Authorization: `Bearer ${this.jwt}` }
        });
    }
}
