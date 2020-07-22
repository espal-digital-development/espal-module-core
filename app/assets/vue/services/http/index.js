/**
 * @typedef {(options: RequestInit) => void} RequestMiddleware
 * @typedef {(response: Response) => void} ResponseMiddleware
 */

export class HTTPService {
    constructor() {
        // TODO :: make this dynamic/a setting
        this._baseAPIUrl = 'https://localhost:8443/API/V1/';

        /** @type RequestMiddleware[] */
        this._requestMiddleware = [];
        /** @type ResponseMiddleware[] */
        this._responseMiddleware = [];
    }

    /**
     * send a request to the given endpoint
     * @param {String} endpoint
     * @param {RequestInit} options
     */
    async request(endpoint, options) {
        for (const middleware of this._requestMiddleware) middleware(options);

        // fetch always returns a Response, even it it's not 20*
        const response = await fetch(this._baseAPIUrl + endpoint, options);

        for (const middleware of this._responseMiddleware) middleware(response);

        // TODO :: check response content type header to see what to convert to
        return response.text();
    }

    /**
     * send a get request to the given endpoint
     * @param {String} endpoint
     * @param {RequestInit} [options]
     */
    get(endpoint, options) {
        if (!options) options = {};
        options.method = 'GET';

        return this.request(endpoint, options);
    }

    /**
     * send a post request to the given endpoint with the given data
     * @param {String} endpoint
     * @param {Object.<string,*>} data
     * @param {RequestInit} [options]
     */
    post(endpoint, data, options) {
        if (!options) options = {};
        options.method = 'POST';

        options.body = new FormData();
        for (const key in data) {
            addToFormData(options.body, key, data[key]);
        }

        return this.request(endpoint, options);
    }

    /**
     * Register request middleware
     * @param {RequestMiddleware} middlewareFunc
     */
    registerRequestMiddleware(middlewareFunc) {
        this._requestMiddleware.push(middlewareFunc);
    }

    /**
     * Register Response middleware
     * @param {ResponseMiddleware} middlewareFunc
     */
    registerResponseMiddleware(middlewareFunc) {
        this._responseMiddleware.push(middlewareFunc);
    }
}

/**
 * Add information to the form data
 * it will check what type it is and will fill the form data accordingly
 *
 * @param {FormData} formData
 * @param {String} key
 * @param {*} value
 */
const addToFormData = (formData, key, value) => {
    if (Array.isArray(value)) {
        for (const newValue of value) {
            addToFormData(formData, `${key}[]`, newValue);
        }
        return;
    }

    if (value === Object(value)) {
        for (const newKey in value) {
            addToFormData(formData, `${key}[${newKey}]`, value[newKey]);
        }
        return;
    }

    formData.append(key, value);
};

// shorter verions, but sending more data
// const addToFormData = (formData, key, value) => {
//     if (data === Object(data) || Array.isArray(data)) {
//         for (var i in data) {
//             createFormData(formData, key + '[' + i + ']', data[i]);
//         }
//     } else {
//         formData.append(key, data);
//     }
// }
