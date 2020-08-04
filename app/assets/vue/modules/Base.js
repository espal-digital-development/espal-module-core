import {storeService, routerService} from '../services';
import {RouterView} from 'vue-router';
import {h} from 'vue';

export class BaseModule {
    constructor(apiEndpoint) {
        this._storeService = storeService;
        this._routerService = routerService;
        this._apiEndpoint = apiEndpoint;
        this._storeService.createAndRegisterModule(this._apiEndpoint);
    }

    init() {
        this._routerService.addRoute(
            this._routerService.createBaseRoutes({
                baseComponent: this.basePage,
                overviewComponent: this.overviewPage,
                createComponent: this.createPage,
                editComponent: this.editPage,
                showComponent: this.showPage,
                endpoint: this._apiEndpoint
            })
        );
    }

    get basePage() {
        return {
            // this looks AWESOME! but it works
            setup: () => () => h(RouterView)
            // mounted: () => this.read()
        };
    }

    get overviewPage() {
        console.warn(`Overview page not explicitly set in Module ${this._apiEndpoint}`);
        return false;
    }

    get createPage() {
        console.warn(`Create page not explicitly set in Module ${this._apiEndpoint}`);
        return false;
    }

    get showPage() {
        console.warn(`Show page not explicitly set in Module ${this._apiEndpoint}`);
        return false;
    }

    get editPage() {
        console.warn(`Edit page not explicitly set in Module ${this._apiEndpoint}`);
        return false;
    }

    goToOverview() {
        this._routerService.goToRoute(this._apiEndpoint + this._routerService.overviewPageNamePart);
    }
}
