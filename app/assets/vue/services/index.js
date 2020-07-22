import { HTTPService } from './http';
import { AuthService } from './auth';
import { StorageService } from './storage';
import { RouterService } from './router';

export const httpService = new HTTPService();
export const storageService = new StorageService();
export const routerService = new RouterService();
export const authService = new AuthService(httpService, storageService, routerService);
