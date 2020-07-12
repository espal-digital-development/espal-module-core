/**
 * @typedef {import('./httpService').HTTPService} HTTPService
 *
 * @typedef {Object} Credentials
 * @property {String} email
 * @property {String} password
 * @property {Boolean} rememberMe
 */

import { ref } from 'vue';

export class AuthService {
    /**
     *
     * @param {HTTPService} httpService
     */
    constructor(httpService) {
        this._httpService = httpService;

        this._jwt = ref('');
    }

    // prettier-ignore
    get jwt() { return this._jwt.value; }

    // prettier-ignore
    set jwt(value) { this._jwt.value = value; }

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
