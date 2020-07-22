const keepALiveKey = 'keepALive';
/** setting keepALive here so we don't have to Parse it each time we get it */
let keepALive = JSON.parse(localStorage.getItem(keepALiveKey));

export class StorageService {
    /** @param {Boolean} value */
    set keepALive(value) {
        localStorage.setItem(keepALiveKey, value);
        keepALive = value;
    }

    // prettier-ignore
    /** @returns {Boolean} */
    get keepALive() { return keepALive; }

    /**
     * Set item in storage, value will be converted to String if it's not a string yet
     * 
     * @param {String} key the key under which to store the value
     * @param {*} value the value to store
     */
    setItem(key, value) {
        if (!this.keepALive) return;
        if (typeof value !== 'string') value = JSON.stringify(value);
        localStorage.setItem(key, value);
    }

    /**
     * Get item from storage
     * 
     * @param {String} key the key for which value to recieve
     */
    getItem(key) {
        if (!this.keepALive) return null;
        return localStorage.getItem(key);
    }

    /** empty the storage */
    clear() {
        localStorage.clear();
    }
}
