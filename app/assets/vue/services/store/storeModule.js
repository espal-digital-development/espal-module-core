/**
 * @typedef {import('../http/index').HTTPService} HTTPService
 */
import {ref, computed} from 'vue';

export class BaseStoreModule {
    /**
     * @param {String} endpoint the endpoint
     * @param {HTTPService} httpService the service that makes the requests
     */
    constructor(endpoint, httpService) {
        this._endpoint = endpoint;
        this._httpService = httpService;

        const stored = localStorage.getItem(endpoint);
        this._state = ref(stored ? JSON.parse(stored) : []);
    }

    get all() {
        return computed(() => this._state.value);
    }

    get allById() {
        return computed(() =>
            this._state.value.reduce((acc, item) => {
                acc[item.id] = item.name;
                return acc;
            }, {})
        );
    }

    setAll(data) {
        this._state.value = data;
        localStorage.setItem(this._endpoint, JSON.stringify(data));
    }

    async loadAll() {
        const {data} = await this._httpService.get(this._endpoint);
        this.setAll(data);
    }
}
