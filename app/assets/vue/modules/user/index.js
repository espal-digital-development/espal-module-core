import {BaseModule} from '../Base';
import {h} from 'vue';

const API_ENDOINT = 'users';

export class UserModule extends BaseModule {
    constructor() {
        super(API_ENDOINT);
    }

    get overviewPage() {
        return h('div', ['USERS OVERVIEW']);
    }

    get createPage() {
        return false;
    }

    get showPage() {
        return false;
    }

    get editPage() {
        return false;
    }
}
